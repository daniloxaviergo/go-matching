package orderbook_test

import (
  "fmt"
  "testing"
	"go-matching/orderbook"
  "go-matching/order"
)

func TestCreateNewLevel(t *testing.T) {
  book := orderbook.OrderBook{}
  level := book.FindOrCreate(10.0)

  fmt.Printf("---------%v", level)

  if level.Price != 10.0 {
    t.Errorf("Price: got %f, want %f", level.Price, 10.0)
  }

  if level.Volume != 0.0 {
    t.Errorf("Volume: got %f, want %f", level.Volume, 0.0)
  }

  if len(level.Orders) != 0 {
    t.Errorf("Orders: got %d, want %f", len(level.Orders), 0.0)
  }

  if len(book.Levels) != 1 {
    t.Errorf("Levels: got %d, want %f", len(book.Levels), 0.0)
  }
}

func TestCreateAppendLevel(t *testing.T) {
  book := orderbook.OrderBook{}
  book.FindOrCreate(20.0)
  book.FindOrCreate(10.0)
  level := book.FindOrCreate(10.0)

  if level.Price != 10.0 {
    t.Errorf("Price: got %f, want %f", level.Price, 10.0)
  }

  if level.Volume != 0.0 {
    t.Errorf("Volume: got %f, want %f", level.Volume, 0.0)
  }

  if len(level.Orders) != 0 {
    t.Errorf("Orders: got %d, want %d", len(level.Orders), 0)
  }

  if len(book.Levels) != 2 {
    t.Errorf("Levels: got %d, want %d", len(book.Levels), 2)
  }
}

func TestRemoveByVolume(t *testing.T) {
  book := orderbook.OrderBook{}
  orderLevel1 := book.FindOrCreate(10.0)
  orderLevel2 := book.FindOrCreate(20.0)
  book.FindOrCreate(30.0)

  order1 := order.Order {
    Id: 1,
    Price: 10.0,
    Volume: 0.5,
  }

  order3 := order.Order {
    Id: 3,
    Price: 10.0,
    Volume: 0.2,
  }

  order2 := order.Order {
    Id: 2,
    Price: 20.0,
    Volume: 0.3,
  }

  orderLevel1.Add(order1)
  orderLevel1.Add(order3)
  orderLevel2.Add(order2)

  fmt.Printf("orderLevel1 %v\n", orderLevel1)
  fmt.Printf("orderLevel1.Volume %f\n", orderLevel1.Volume)
  fmt.Printf("orderLevel2 %v\n", orderLevel2)
  fmt.Printf("orderLevel2.Volume %f\n", orderLevel2.Volume)
  fmt.Printf("---------%v %v %v\n", book.Levels[0], book.Levels[1], book.Levels[2])
  fmt.Println("---------------------------------------")

  orderLevel1.Remove(order1)
  orderLevel2.Remove(order2)

  book.RemoveByVolume(order1.Price)
  book.RemoveByVolume(order2.Price)

  fmt.Printf("orderLevel1 %v\n", orderLevel1)
  fmt.Printf("orderLevel1.Volume %f\n", orderLevel1.Volume)
  fmt.Printf("orderLevel2 %v\n", orderLevel2)
  fmt.Printf("orderLevel2.Volume %f\n", orderLevel2.Volume)
  fmt.Printf("---------%v %v\n", book.Levels[0], book.Levels[1])

  if len(book.Levels) != 2 {
    t.Errorf("Levels: got %d, want %d", len(book.Levels), 2)
  }

  if book.Levels[0].Price != 10.0 {
    t.Errorf("first Levels: got %f, want %f", book.Levels[0].Price, 10.0)
  }

  if len(book.Levels[0].Orders) != 1 {
    t.Errorf("Levels: got %d, want %d", len(book.Levels[0].Orders), 1)
  }

  if book.Levels[1].Price != 30.0 {
    t.Errorf("first Levels: got %f, want %f", book.Levels[0].Price, 30.0)
  }

  if len(book.Levels[1].Orders) != 0 {
    t.Errorf("Levels: got %d, want %d", len(book.Levels[1].Orders), 0)
  }
}
