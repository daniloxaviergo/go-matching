package market_test

import (
  "fmt"
  "go-matching/market"
  "go-matching/order"
  "testing"
  // "go-matching/orderbook"
)

func TestCreateNewMarket(t *testing.T) {
  markets := market.Markets{}
  market := markets.FindOrCreate("btcbrl")

  fmt.Printf("---------%v", market)

  if market.Name != "btcbrl" {
    t.Errorf("Price: got %s, want %s", market.Name, "btcbrl")
  }
}

func TestAddOrder(t *testing.T) {
  markets := market.Markets{}
  market := markets.FindOrCreate("btcbrl")

  ask_order := order.Order{
    Id:     1,
    Price:  10.0,
    Volume: 0.5,
  }

  bid_order := order.Order{
    Id:     2,
    Price:  20.0,
    Volume: 0.7,
  }
  market.AddOrder("Ask", ask_order)
  market.AddOrder("Bid", bid_order)

  ask_level := market.Ask.FindOrCreate(ask_order.Price)
  bid_level := market.Bid.FindOrCreate(bid_order.Price)

  ask_level_order := ask_level.Orders[0]
  bid_level_order := bid_level.Orders[0]

  if len(market.Ask.Levels) != 1 {
    t.Errorf("len OrderBook: got %d, want %d", len(market.Ask.Levels), 1)
  }

  if len(market.Bid.Levels) != 1 {
    t.Errorf("len OrderBook: got %d, want %d", len(market.Bid.Levels), 0)
  }

  if ask_level.Price != 10.0 {
    t.Errorf("ask level: got %f, want %f", ask_level.Price, 10.0)
  }

  if bid_level.Price != 20.0 {
    t.Errorf("bid level: got %f, want %f", bid_level.Price, 20.0)
  }

  if ask_level.Volume != 0.5 {
    t.Errorf("ask level: got %f, want %f", ask_level.Volume, 0.5)
  }

  if bid_level.Volume != 0.7 {
    t.Errorf("bid level: got %f, want %f", bid_level.Volume, 0.7)
  }

  if len(ask_level.Orders) != 1 {
    t.Errorf("len OrderBook: got %d, want %d", len(ask_level.Orders), 1)
  }

  if len(bid_level.Orders) != 1 {
    t.Errorf("len OrderBook: got %d, want %d", len(bid_level.Orders), 1)
  }

  if ask_level_order.Id != 1 {
    t.Errorf("ask order id: got %d, want %d", ask_level_order.Id, 1)
  }

  if bid_level_order.Id != 2 {
    t.Errorf("bid order id: got %d, want %d", bid_level_order.Id, 1)
  }
}

func TestRemoveOrder(t *testing.T) {
  markets := market.Markets{}
  market := markets.FindOrCreate("btcbrl")

  ask_order1 := order.Order{
    Id:     1,
    Price:  10.0,
    Volume: 0.3,
  }

  ask_order2 := order.Order{
    Id:     2,
    Price:  10.0,
    Volume: 0.5,
  }

  bid_order1 := order.Order{
    Id:     4,
    Price:  20.0,
    Volume: 0.7,
  }

  bid_order2 := order.Order{
    Id:     5,
    Price:  20.0,
    Volume: 0.2,
  }

  bid_order3 := order.Order{
    Id:     6,
    Price:  30.0,
    Volume: 0.2,
  }

  market.AddOrder("Ask", ask_order1)
  market.AddOrder("Ask", ask_order2)
  market.AddOrder("Bid", bid_order1)
  market.AddOrder("Bid", bid_order2)
  market.AddOrder("Bid", bid_order3)

  market.RemoveOrder("Bid", bid_order3)
  market.RemoveOrder("Ask", ask_order2)
  market.RemoveOrder("Bid", bid_order1)

  ask_level := market.Ask.FindOrCreate(ask_order1.Price)
  bid_level := market.Bid.FindOrCreate(bid_order1.Price)

  ask_level_order := ask_level.Orders[0]
  bid_level_order := bid_level.Orders[0]

  if len(market.Ask.Levels) != 1 {
    t.Errorf("len OrderBook: got %d, want %d", len(market.Ask.Levels), 1)
  }

  if len(market.Bid.Levels) != 1 {
    t.Errorf("len OrderBook: got %d, want %d", len(market.Bid.Levels), 1)
  }

  if ask_level.Price != 10.0 {
    t.Errorf("ask level: got %f, want %f", ask_level.Price, 10.0)
  }

  if bid_level.Price != 20.0 {
    t.Errorf("bid level: got %f, want %f", bid_level.Price, 20.0)
  }

  if ask_level.Volume != 0.3 {
    t.Errorf("ask level: got %f, want %f", ask_level.Volume, 0.3)
  }

  if bid_level.Volume != 0.2 {
    t.Errorf("bid level: got %f, want %f", bid_level.Volume, 0.2)
  }

  if len(ask_level.Orders) != 1 {
    t.Errorf("len OrderBook: got %d, want %d", len(ask_level.Orders), 1)
  }

  if len(bid_level.Orders) != 1 {
    t.Errorf("len OrderBook: got %d, want %d", len(bid_level.Orders), 1)
  }

  if ask_level_order.Id != 1 {
    t.Errorf("ask order id: got %d, want %d", ask_level_order.Id, 1)
  }

  if bid_level_order.Id != 5 {
    t.Errorf("bid order id: got %d, want %d", bid_level_order.Id, 5)
  }
}
