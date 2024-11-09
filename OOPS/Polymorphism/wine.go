package main

type Wine struct {
	ProductDetails
	Year string
	Kind string
}

func (w Wine) calculatePrice() int64 {
	liquorTax := float64(w.Price) * .28
	return w.Price + int64(liquorTax)
}
