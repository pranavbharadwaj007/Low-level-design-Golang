package main

type Monitor struct {
	ProductDetails
	Size       string
	Resolution string
}

func (m Monitor) calculatePrice() int64 {
	electronicTax := float64(m.Price) * .20
	return int64(m.Price) + int64(electronicTax)
}
