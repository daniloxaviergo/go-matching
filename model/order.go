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

/*
func (os *OrderLevel) Add(order Order) (value int, err error) {
	// logica
	return 3, nil
}

func call() {
	ol := OrOrderLevel
	v, err := ol.Add(order)
	if (err == nil) {
		// catch exception
	}

}
*/