package model

type Order struct {
	Id     int64
	Volume float64
	Price  float64
}

type OrderLevel struct {
	Price  float64
	Volume float64
	Orders []Order
}

type Market struct {
	Ask OrderBook
	Bid OrderBook
}

type OrderBook struct {
	Levels []OrderLevel
}
