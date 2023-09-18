package controllers

import (
  "github.com/gin-gonic/gin"
  "go-matching/market"
  "go-matching/order"
  "net/http"
  "strconv"
)

type OrderBookResponse struct {
  Market string       `json:"market"`
  Asks   [][2]float64 `json:"asks"`
  Bids   [][2]float64 `json:"bids"`
}

type OrderRemoveResponse struct {
  Id     string `json:"id"`
  Market string `json:"market"`
  Status string `json:"status"`
}

type OrderCreateResponse struct {
  Id     string `json:"id"`
  Price  string `json:"price"`
  Volume string `json:"volume"`
  Market string `json:"market"`
  Status string `json:"status"`
}

var markets market.Markets

func AddOrder(c *gin.Context) {
  q_market := c.Param("market")
  q_id := c.Query("id")
  q_price := c.Query("price")
  q_volume := c.Query("volume")
  q_side := c.Query("side")

  if q_id == "" || q_price == "" || q_volume == "" || q_market == "" || q_side == "" {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Faltam parametros"})
    return
  }

  new_id, _ := strconv.ParseInt(q_id, 64, 0)
  new_price, _ := strconv.ParseFloat(q_price, 32)
  new_volume, _ := strconv.ParseFloat(q_volume, 32)

  market := markets.FindOrCreate(q_market)
  req_order := order.Order{
    Id:     new_id,
    Price:  new_price,
    Volume: new_volume,
  }

  market.AddOrder(q_side, req_order)

  var response OrderCreateResponse
  response.Id = q_id
  response.Price = q_price
  response.Volume = q_volume
  response.Status = "ok"
  response.Market = q_market

  c.JSON(http.StatusOK, response)
}

func RemoveOrder(c *gin.Context) {
  q_market := c.Param("market")
  q_id := c.Query("id")

  if q_id == "" || q_market == "" {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Faltam parametros"})
    return
  }

  market := markets.FindOrCreate(q_market)
  new_id, _ := strconv.ParseInt(q_id, 64, 0)

  var rm_side string
  var rm_order order.Order

  for _, lv := range market.Ask.Levels {
    for _, or := range lv.Orders {
      if or.Id == new_id {
        rm_side = "Ask"
        rm_order = or

        break
      }
    }
  }

  for _, lv := range market.Bid.Levels {
    if rm_side == "Ask" {
      break
    }

    for _, or := range lv.Orders {
      if or.Id == new_id {
        rm_side = "Bid"
        rm_order = or

        break
      }
    }
  }

  market.RemoveOrder(rm_side, rm_order)

  var response OrderRemoveResponse
  response.Id = q_id
  response.Status = "cancel"
  response.Market = q_market

  c.JSON(http.StatusOK, response)
}

func OrderBook(c *gin.Context) {
  q_market := c.Param("market")

  market := markets.FindOrCreate(q_market)

  var Asks [][2]float64
  var Bids [][2]float64

  for _, lv := range market.Ask.Levels {
    level := [2]float64{
      lv.Price,
      lv.Volume,
    }

    Asks = append(Asks, level)
  }

  for _, lv := range market.Bid.Levels {
    level := [2]float64{
      lv.Price,
      lv.Volume,
    }

    Bids = append(Bids, level)
  }

  var response OrderBookResponse
  response.Market = q_market
  response.Asks = Asks
  response.Bids = Bids

  c.JSON(http.StatusOK, response)
}
