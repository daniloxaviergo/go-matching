package main

import (
  "github.com/gin-gonic/gin"
  "go-matching/controllers"
)

func main() {
  // curl -v 'http://localhost:8080/btcbrl/order-create?id=1&price=12.34&volume=78.85&side=ask'
  // curl -v 'http://localhost:8080/btcbrl/order-create?id=2&price=12.34&volume=78.85&side=ask'
  // curl -v 'http://localhost:8080/btcbrl/order-cancel?id=1'
  // curl -v 'http://localhost:8080/btcbrl/order-cancel?id=2'
  // curl -v 'http://localhost:8080/btcbrl/orderbook'

  //docker run --rm -p 8080:8080 -v $(pwd):/app -it go-matching bash
  router := gin.Default()

  router.GET("/:market/order-create", controllers.AddOrder)
  router.GET("/:market/order-cancel", controllers.RemoveOrder)
  router.GET("/:market/orderbook", controllers.OrderBook)

  router.Run("0.0.0.0:8080")
}
