package main

import (

	"github.com/gin-gonic/gin"
	"go-matching/controllers"
	"go-matching/model"
	"go-matching/test"
)

func main() {

	// BRUNO E ELTON
	router := gin.Default()

	router.GET("/:market/order-create", controllers.AddOrder)
	router.GET("/:market/order-cancel", controllers.RemoveOrder)
	router.GET("/:market/orderbook", controllers.OrderBook)

	router.Run("localhost:8080")
	// BRUNO E ELTON


	fmt.Println(c.Name)

	orderLevel := model.OrderLevel{
		Price:  0,
		Volume: 0,
		Orders: make([]model.Order, 0),
	}

	order := model.Order{
		Id:     1,
		Price:  0.5,
		Volume: 10,
	}

	orderLevel.Add(order)
	// fmt.Println(orderLevel.Orders)

	fmt.Printf("%v", orderLevel.Orders)
	fmt.Printf("%v", orderLevel.Volume)
}
