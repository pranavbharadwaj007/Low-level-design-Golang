package main

import (
	"fmt"
)

type Purchaser interface {
	calculatePrice() int64
}

var cart []Purchaser

func addToCart(products ...Purchaser) {
	cart = append(cart, products...)

}
func getCartTotal() int64 {
	var total int64
	for _, product := range cart {
		fmt.Println("Product: ", product, "Price: ", product.calculatePrice())
		total += product.calculatePrice()
	}
	return total
}
func main() {
	fmt.Println("Hello, World!")
	myshirt := Shirt{ProductDetails{Price: 100, Brand: "Nike"}, "M", "blue"}
	mymonitor := Monitor{ProductDetails{Price: 7000, Brand: "Lg"}, "24inch", "130*340"}

	myWinewine := Wine{ProductDetails{Price: 860, Brand: "coco"}, "2017", "Riesling"}
	addToCart(myshirt, mymonitor, myWinewine)
	fmt.Println(getCartTotal())

}
