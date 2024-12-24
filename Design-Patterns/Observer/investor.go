package observer

import "fmt"

// InvestorA implements the Observer interface
type InvestorA struct{}

func (i *InvestorA) Update(symbol string, price float64) {
	fmt.Printf("Investor A notified: Stock %s has a new price: $%.2f\n", symbol, price)
}

// InvestorB implements the Observer interface
type InvestorB struct{}

func (i *InvestorB) Update(symbol string, price float64) {
	fmt.Printf("Investor B notified: Stock %s has a new price: $%.2f\n", symbol, price)
}
