package model

import "errors"

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

func (os *OrderLevel) Add(order Order) (err error, success bool) {
	if os.Price != order.Price {
		return errors.New("price is not matching with Order Level"), false
	}

	os.Orders = append(os.Orders, order)

	var volumes float64 = 0
	for _, o := range os.Orders {
		volumes = volumes + o.Volume
	}

	os.Volume = volumes

	return nil, true
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
