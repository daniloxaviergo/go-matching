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

	o := model.Order{
		Side: "213.0",
	}

	btcbrl := model.OrderBook

	fmt.Println(c.Name)
	fmt.Println(o.Side)
}
