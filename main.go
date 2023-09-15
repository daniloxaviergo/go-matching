package main

import (

	"github.com/gin-gonic/gin"
	"go-matching/controllers"
)

// func main() {
// 	fmt.Println("Hello World!!")
// 	c := test.Creature{
// 		Name: "Sammy the Shark",
// 	}

// 	fmt.Println(c.Name)
// }


func main() {
	router := gin.Default()

	router.GET("/:market/order-create", controllers.AddOrder)
	router.GET("/:market/order-cancel", controllers.RemoveOrder)
	router.GET("/:market/orderbook", controllers.OrderBook)

	router.Run("localhost:8080")
}
