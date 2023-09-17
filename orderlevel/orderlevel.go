package orderlevel

import (
  "errors"
  "go-matching/order"
)

type OrderLevel struct {
  Price  float64
  Volume float64
  Orders []order.Order
}

func (orderLevel *OrderLevel) Add(order order.Order) (err error, success bool) {
  if orderLevel.Price != order.Price {
    return errors.New("price is not matching with Order Level"), false
  }

  orderLevel.Orders = append(orderLevel.Orders, order)

  var volumes float64 = 0
  for _, o := range orderLevel.Orders {
    volumes = volumes + o.Volume
  }

  orderLevel.Volume = volumes

  return nil, true
}

func (orderLevel *OrderLevel) Remove(order order.Order) (err error, success bool) {
  if orderLevel.Price != order.Price {
    return errors.New("price is not matching with Order Level"), false
  }

  index := -1
  for idx, o := range orderLevel.Orders {
    if o.Id == order.Id {
      index = idx
    }
  }

  orderLevel.Orders = append(orderLevel.Orders[:index], orderLevel.Orders[index+1:]...)

  var volumes float64 = 0
  for _, o := range orderLevel.Orders {
    volumes = volumes + o.Volume
  }

  orderLevel.Volume = volumes

  return nil, true
}
