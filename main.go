package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go-matching/controllers"
	"go-matching/model"
)

func main() {

	// BRUNO E ELTON
	router := gin.Default()

	controllers.InitMarket()
	router.GET("/:market/order-create", controllers.AddOrder)
	router.GET("/:market/order-cancel", controllers.RemoveOrder)
	router.GET("/:market/orderbook", controllers.OrderBook)

	// router.Run("127.0.0.1:8080")
	// BRUNO E ELTON
	orderBookAsk := model.OrderBook {
		Levels: make([]model.OrderLevel, 0),
	}

	orderBookBid := model.OrderBook {
		Levels: make([]model.OrderLevel, 0),
	}

	btcbrl := model.Market{
		Ask: orderBookAsk,
		Bid: orderBookBid,
	}

	orderAskBtc := model.Order{
		Id:     1,
		Price:  0.5,
		Volume: 10.1,
	}

	orderLevel0 := model.OrderLevel{
		Price:  0.5,
		Volume: 0.0,
		Orders: make([]model.Order, 0),
	}

	orderBookAsk.Levels = append(orderBookAsk.Levels, orderLevel0)
	btcbrl.Ask = orderBookAsk

	orderBookAsk.Levels[0].Add(orderAskBtc)
	fmt.Printf("len: %d\n", len(orderBookAsk.Levels))
	fmt.Printf("btcbrl: %v\n", btcbrl)

	// btcbrl
	// var err0, success0 = btcbrl.Ask.Levels.AddOrder(orderBidBtc)
	// if (!success0) {
	// 	fmt.Println("error - btcbrl")
	// 	fmt.Println(err0)
	// }

	order := model.Order{
		Id:     1,
		Price:  0.5,
		Volume: 10.1,
	}

	order2 := model.Order{
		Id:     1,
		Price:  0.5,
		Volume: 1.0,
	}

	orderLevel := model.OrderLevel{
		Price:  0.5,
		Volume: 0.0,
		Orders: make([]model.Order, 0),
	}

	var err, success = orderLevel.Add(order)
	// fmt.Println(orderLevel.Orders)
	if (!success) {
		fmt.Println("error")
		fmt.Println(err)
	}

	var err2, success2 = orderLevel.Add(order2)
	// fmt.Println(orderLevel.Orders)
	if (!success2) {
		fmt.Println("error")
		fmt.Println(err2)
	}

	fmt.Printf("Orders: %v\n", orderLevel.Orders)
	fmt.Printf("Volume: %v\n", orderLevel.Volume)
}
