package main

import (
	"fmt"

	"go-matching/model"
	"go-matching/test"
)

func main() {
	fmt.Println("Hello World!!")
	c := test.Creature{
		Name: "Sammy the Shark",
	}

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
