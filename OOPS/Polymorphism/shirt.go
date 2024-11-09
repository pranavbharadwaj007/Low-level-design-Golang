package main

type Shirt struct {
	ProductDetails
	Size  string
	Color string
}

func (s Shirt) calculatePrice() int64 {
	clotingDiscount := float64(s.Price) * .20
	return s.Price - int64(clotingDiscount)
}
