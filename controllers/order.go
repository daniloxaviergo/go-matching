package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"go-matching/model"
	"strconv"
)

type OrderBookResponse struct {
	Market string    `json:"market"`
	Asks   []float64 `json:"asks"`
	Bids   []float64 `json:"bids"`
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

var markets []model.Market

func InitMarket() {
	markets = make([]model.Market, 0)
}

func AddOrder(c *gin.Context) {
	market := c.Param("market")
	id := c.Query("id")
	price := c.Query("price")
	volume := c.Query("volume")
	side := c.Query("side")

	if id == "" || price == "" || volume == "" || market == "" || side == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Faltam parametros"})
		return
	}

	create_market := true
	var mkt model.Market
	var idx int

	for i, o := range markets {
		if o.Name == market{
			create_market = false
			mkt = o
			idx = i
		}		
	}

	if create_market {
		fmt.Println("CRIANDO MERCADO")

		orderBookAsk := model.OrderBook {
			Levels: make([]model.OrderLevel, 0),
		}
	
		orderBookBid := model.OrderBook {
			Levels: make([]model.OrderLevel, 0),
		}
	
		book := model.Market{
			Name: market,
			Ask: orderBookAsk,
			Bid: orderBookBid,
		}

		markets = append(markets, book)

		mkt = book
	}

	for i, o := range markets {
		if o.Name == market{			
			mkt = o
			idx = i
		}		
	}

	newId, _ := strconv.ParseInt(id, 64, 0)
	newPrice, _ := strconv.ParseFloat(price, 32)
	newVolume, _ := strconv.ParseFloat(volume, 32)

	order := model.Order{
		Id:     newId,
		Price:  newPrice,
		Volume: newVolume,
	}	

	orderLevel := model.OrderLevel{
		Price:  newPrice,
		Volume: 0.0,
		Orders: make([]model.Order, 0),
	}
	_, _ = orderLevel.Add(order)	

	if side == "ask" {		
		markets[idx].Ask.Levels = append(markets[idx].Ask.Levels, orderLevel)		
	}else{
		markets[idx].Bid.Levels = append(markets[idx].Bid.Levels, orderLevel)
	}	

	// markets = append(markets, mkt)

	fmt.Printf("markets: %v\n", markets)

	fmt.Printf("%v", mkt)

	var response OrderCreateResponse
	response.Id = id
	response.Price = price
	response.Volume = volume
	response.Status = "ok"
	response.Market = market

	c.JSON(http.StatusOK, response)
}

func RemoveOrder(c *gin.Context) {
	market := c.Param("market")
	id := c.Query("id")

	if id == "" || market == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Faltam parametros"})
		return
	}

	var response OrderRemoveResponse
	response.Id = id
	response.Status = "cancel"
	response.Market = market

	c.JSON(http.StatusOK, response)
}

func OrderBook(c *gin.Context) {
	market := c.Param("market")

	var response OrderBookResponse
	response.Market = market

	c.JSON(http.StatusOK, response)
}
