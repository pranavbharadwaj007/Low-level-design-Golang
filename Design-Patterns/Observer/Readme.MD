# Understanding the Observer Pattern

## Pattern Structure
```mermaid
classDiagram
    class Subject {
        <<interface>>
        +RegisterObserver(Observer)
        +RemoveObserver(Observer)
        +NotifyObservers(string, float64)
    }

    class Observer {
        <<interface>>
        +Update(string, float64)
    }

    class StockMarket {
        -observers: map
        -priceChangeThreshold: float64
        -updateChan: chan StockUpdate
        -mu: sync.RWMutex
        +NewStockMarket(float64) *StockMarket
        +processUpdates()
        +SetStockPrice(string, float64, float64)
    }

    class StockUpdate {
        +Symbol: string
        +NewPrice: float64
        +OldPrice: float64
    }

    class InvestorA {
        +Update(string, float64)
    }

    class InvestorB {
        +Update(string, float64)
    }

    Subject <|.. StockMarket : implements
    Observer <|.. InvestorA : implements
    Observer <|.. InvestorB : implements
    StockMarket ..> StockUpdate : uses
```

## Sequence of Operations
```
flowchart TD
    A[Main Program] -->|Create| B[StockMarket]
    B -->|Start| C((processUpdates Goroutine))
    
    D[SetStockPrice] -->|Send Update| E[(updateChan)]
    E -->|Receive| C
    
    C -->|Threshold Check| F[NotifyObservers]
    F -->|Spawn Goroutines| G[WaitGroup]
    
    G -->|Concurrent Update| H[Observer A Goroutine]
    G -->|Concurrent Update| I[Observer B Goroutine]
    
    H -->|Complete| J[WaitGroup Done]
    I -->|Complete| J
    
    subgraph Synchronization
        K[RWMutex] -.->|Protects| L[Observers Map]
    end
```
