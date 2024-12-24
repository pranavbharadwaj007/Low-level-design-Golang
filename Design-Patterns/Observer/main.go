package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"./observer"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	// Read price change threshold
	fmt.Print("Enter price change threshold: ")
	scanner.Scan()
	threshold, _ := strconv.ParseFloat(scanner.Text(), 64)
	
	stockMarket := observer.NewStockMarket(threshold)
	investorA := &observer.InvestorA{}
	investorB := &observer.InvestorB{}
	
	// Register observers
	stockMarket.RegisterObserver(investorA)
	stockMarket.RegisterObserver(investorB)
	
	// Read number of updates
	fmt.Print("Enter number of updates: ")
	scanner.Scan()
	updates, _ := strconv.Atoi(scanner.Text())
	
	for i := 1; i <= updates; i++ {
		if i == 5 {
			stockMarket.RemoveObserver(investorB)
		}
		
		fmt.Printf("Update %d - Enter stock symbol, new price, old price: ", i)
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		
		symbol := input[0]
		newPrice, _ := strconv.ParseFloat(input[1], 64)
		oldPrice, _ := strconv.ParseFloat(input[2], 64)
		
		stockMarket.SetStockPrice(symbol, newPrice, oldPrice)
	}
}
