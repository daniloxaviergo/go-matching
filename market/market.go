package market

import (
  // "fmt"
  "go-matching/order"
  "go-matching/orderbook"
)

type Market struct {
  Name string
  Ask  *orderbook.OrderBook
  Bid  *orderbook.OrderBook
}

type Markets struct {
  Markets []*Market
}

func (market *Market) AddOrder(side string, order order.Order) (err error, success bool) {
  var side_book *orderbook.OrderBook

  if side == "Ask" || side == "ask" {
    side_book = market.Ask
  } else {
    side_book = market.Bid
  }

  // fmt.Printf("market_side: %v", market.ValueOf("Ask"))
  order_level := side_book.FindOrCreate(order.Price)
  order_level.Add(order)

  // fmt.Printf("market_side: %v\n\n\n", side_book)
  // fmt.Printf("order_level: %v", order_level)
  // orderLevel = market[side].FindOrCreate(order.Price)

  return nil, true
}

func (market *Market) RemoveOrder(side string, order order.Order) (err error, success bool) {
  var side_book *orderbook.OrderBook

  if side == "Ask" || side == "ask" {
    side_book = market.Ask
  } else {
    side_book = market.Bid
  }

  // fmt.Printf("market_side: %v", market.ValueOf("Ask"))
  order_level := side_book.FindOrCreate(order.Price)

  order_level.Remove(order)
  side_book.RemoveByVolume(order.Price)

  // fmt.Printf("market_side: %v\n\n\n", side_book)
  // fmt.Printf("order_level: %v", order_level)
  // orderLevel = market[side].FindOrCreate(order.Price)

  return nil, true
}

func (markets *Markets) FindOrCreate(market_name string) *Market {
  index := -1
  var mmmarket *Market

  for idx, mmarket := range markets.Markets {
    if markets.Markets[idx].Name == market_name {
      index = idx

      mmmarket = mmarket
    }
  }

  if index > -1 {
    return mmmarket
  }

  new_market := Market{
    Name: market_name,
    Ask:  &orderbook.OrderBook{},
    Bid:  &orderbook.OrderBook{},
  }

  markets.Markets = append(markets.Markets, &new_market)
  mmmarket = &new_market
  return mmmarket
}
