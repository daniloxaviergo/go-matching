package orderlevel_test

import (
  "fmt"
  "testing"
  "go-matching/orderlevel"
  "go-matching/order"
)

func TestAddLevel(t *testing.T) {
  order1 := order.Order {
    Id: 1,
    Price: 10.0,
    Volume: 0.5,
  }

  order2 := order.Order {
    Id: 2,
    Price: 10.0,
    Volume: 0.2,
  }

  orderLevel1 := orderlevel.OrderLevel {
    Price: 10.0,
    Volume: 0.0,
  }

  _, success1 := orderLevel1.Add(order1)
  _, success2 := orderLevel1.Add(order2)

  fmt.Printf("---------%v", orderLevel1)

  if !success1 {
    t.Errorf("add function return false")
  }

  if !success2 {
    t.Errorf("add function return false")
  }

  if orderLevel1.Price != 10.0 {
    t.Errorf("Price: got %f, want %f", orderLevel1.Price, 10.0)
  }

  if orderLevel1.Volume != 0.7 {
    t.Errorf("Volume: got %f, want %f", orderLevel1.Volume, 0.7)
  }

  if len(orderLevel1.Orders) != 2 {
    t.Errorf("Orders: got %d, want %f", len(orderLevel1.Orders), 1.0)
  }

  if orderLevel1.Orders[1].Id != order2.Id {
    t.Errorf("Order.Id: got %d, want %d", orderLevel1.Orders[1].Id, order2.Id)
  }

  if orderLevel1.Orders[1].Price != order2.Price {
    t.Errorf("Order.Price: got %f, want %f", orderLevel1.Orders[1].Price, order2.Price)
  }

  if orderLevel1.Orders[1].Volume != order2.Volume {
    t.Errorf("Order.Volume: got %f, want %f", orderLevel1.Orders[1].Volume, order2.Volume)
  }
}

func TestRemoveLevel(t *testing.T) {
  order1 := order.Order {
    Id: 1,
    Price: 10.0,
    Volume: 0.5,
  }

  order2 := order.Order {
    Id: 2,
    Price: 10.0,
    Volume: 0.2,
  }

  order3 := order.Order {
    Id: 3,
    Price: 10.0,
    Volume: 0.3,
  }

  orderLevel1 := orderlevel.OrderLevel {
    Price: 10.0,
    Volume: 0.0,
  }

  orderLevel1.Add(order1)
  orderLevel1.Add(order2)
  orderLevel1.Add(order3)

  orderLevel1.Remove(order2)

  fmt.Printf("---------%v", orderLevel1)

  if orderLevel1.Price != 10.0 {
    t.Errorf("Price: got %f, want %f", orderLevel1.Price, 10.0)
  }

  if orderLevel1.Volume != 0.8 {
    t.Errorf("Volume: got %f, want %f", orderLevel1.Volume, 0.8)
  }

  if len(orderLevel1.Orders) != 2 {
    t.Errorf("Orders: got %d, want %f", len(orderLevel1.Orders), 1.0)
  }

  if orderLevel1.Orders[1].Id != order3.Id {
    t.Errorf("Order.Id: got %d, want %d", orderLevel1.Orders[1].Id, order3.Id)
  }

  if orderLevel1.Orders[1].Price != order3.Price {
    t.Errorf("Order.Price: got %f, want %f", orderLevel1.Orders[1].Price, order3.Price)
  }

  if orderLevel1.Orders[1].Volume != order3.Volume {
    t.Errorf("Order.Volume: got %f, want %f", orderLevel1.Orders[1].Volume, order3.Volume)
  }
}
