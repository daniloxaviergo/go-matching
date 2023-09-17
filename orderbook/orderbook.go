package orderbook

import (
	"fmt"
	"go-matching/orderlevel"
	"go-matching/order"
)

type OrderBook struct {
  Levels []*orderlevel.OrderLevel
}

func (orderBook *OrderBook) FindOrCreate(price float64) (*orderlevel.OrderLevel) {
	index := -1
	var lllevel *orderlevel.OrderLevel

	for idx, llevel := range orderBook.Levels {
    if (orderBook.Levels[idx].Price == price) {
    	index = idx

    	lllevel = llevel
    }
  }

  if index > -1 {
  	// return &orderBook.Levels[index]
  	return lllevel
  }

  level := orderlevel.OrderLevel {
  	Price: price,
  	Volume: 0.0,
  	Orders: make([]order.Order, 0),
  }

  orderBook.Levels = append(orderBook.Levels, &level)
  // return &orderBook.Levels[len(orderBook.Levels) - 1]
  lllevel = &level
  return lllevel
}

func (orderBook *OrderBook) RemoveByVolume(price float64) (level []*orderlevel.OrderLevel) {
	index := -1
	for idx, _ := range orderBook.Levels {
    if (orderBook.Levels[idx].Price == price) {
    	index = idx
    }
  }

  if index == -1 {
  	return nil
  }

  fmt.Printf("Price: %f Volume: %f\n\n", orderBook.Levels[index].Price, orderBook.Levels[index].Volume)
  if orderBook.Levels[index].Volume == 0.0 {
  	orderBook.Levels = append(orderBook.Levels[:index], orderBook.Levels[index+1:]...)
  }

  return orderBook.Levels
}
