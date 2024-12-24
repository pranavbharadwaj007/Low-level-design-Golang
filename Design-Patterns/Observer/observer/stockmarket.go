package observer

import (
	"math"
	"sync"
)

// StockUpdate represents a stock price update
type StockUpdate struct {
	Symbol    string
	NewPrice  float64
	OldPrice  float64
}

// Subject interface defines methods for managing observers
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(symbol string, price float64)
}

// Observer interface defines the update method for receiving notifications
type Observer interface {
	Update(symbol string, price float64)
}

// StockMarket implements the Subject interface
type StockMarket struct {
	observers           map[Observer]struct{}
	priceChangeThreshold float64
	updateChan         chan StockUpdate
	mu                 sync.RWMutex
}

// NewStockMarket creates a new StockMarket instance
func NewStockMarket(priceChangeThreshold float64) *StockMarket {
	sm := &StockMarket{
		observers:           make(map[Observer]struct{}),
		priceChangeThreshold: priceChangeThreshold,
		updateChan:         make(chan StockUpdate, 100),
	}
	go sm.processUpdates()
	return sm
}

func (sm *StockMarket) processUpdates() {
	for update := range sm.updateChan {
		priceChange := math.Abs(update.NewPrice-update.OldPrice) / update.OldPrice * 100
		if priceChange >= sm.priceChangeThreshold {
			sm.NotifyObservers(update.Symbol, update.NewPrice)
		}
	}
}

func (sm *StockMarket) RegisterObserver(observer Observer) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.observers[observer] = struct{}{}
}

func (sm *StockMarket) RemoveObserver(observer Observer) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.observers, observer)
}

func (sm *StockMarket) NotifyObservers(symbol string, price float64) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	var wg sync.WaitGroup
	for observer := range sm.observers {
		wg.Add(1)
		go func(o Observer) {
			defer wg.Done()
			o.Update(symbol, price)
		}(observer)
	}
	wg.Wait()
}

func (sm *StockMarket) SetStockPrice(symbol string, newPrice, oldPrice float64) {
	sm.updateChan <- StockUpdate{
		Symbol:    symbol,
		NewPrice:  newPrice,
		OldPrice:  oldPrice,
	}
}
