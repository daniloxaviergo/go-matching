package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-matching/model"
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

	if id == "" || price == "" || volume == "" || market == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Faltam parametros"})
		return
	}

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
