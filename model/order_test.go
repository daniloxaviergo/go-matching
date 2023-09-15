package model_test

import (
	"go-matching/model"
	"reflect"
	"testing"
)

func AddOrderIntoOrderLevel(t *testing.T) {
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

	err := orderLevel.Add(order)
	if (err != nil) {
		t.Fatal("Error when adding Order")
	}

	if !reflect.DeepEqual(orderLevel.Orders[0], order) {
		t.Fatal("structs are not equal")
	}
}