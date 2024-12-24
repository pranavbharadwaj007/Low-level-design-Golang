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
```mermaid
flowchart TD
    MainProgram[Main Program] -->|1. Creates| StockMarket
    StockMarket -->|2. Starts| ProcessUpdates[processUpdates Goroutine]
    
    SetPrice[SetStockPrice] -->|3. Sends Update| Channel[(updateChan)]
    Channel -->|4. Receives| ProcessUpdates
    
    ProcessUpdates -->|5. Threshold Check| Notify{NotifyObservers}
    Notify -->|6. Spawn Goroutines| WG[WaitGroup]
    
    WG -->|7. Concurrent Update| ObserverA[Observer A Goroutine]
    WG -->|7. Concurrent Update| ObserverB[Observer B Goroutine]
    
    ObserverA -->|8. Complete| Done[WaitGroup Done]
    ObserverB -->|8. Complete| Done
    
    subgraph Synchronization
        Mutex[RWMutex] -.->|Protects| Map[Observers Map]
    end
```

## Implementation Details

### 1. Core Components
- **Subject Interface**: Defines the contract for managing observers
- **Observer Interface**: Defines the update method for receiving notifications
- **StockMarket**: Concrete implementation of the Subject interface
- **InvestorA/B**: Concrete implementations of the Observer interface

### 2. Concurrency Mechanisms
- **Channels**: `updateChan` buffers and processes stock updates asynchronously
- **Goroutines**: Used for concurrent notification of observers
- **WaitGroup**: Synchronizes completion of observer notifications
- **RWMutex**: Ensures thread-safe access to the observers map

### 3. Data Flow
1. Main program creates StockMarket instance
2. Background goroutine starts processing updates
3. Updates are sent through channel
4. Each update is processed and checked against threshold
5. If threshold is met, observers are notified concurrently
6. WaitGroup ensures all notifications complete

### 4. Thread Safety
- Observer registration/removal is protected by RWMutex
- Channel ensures safe communication between goroutines
- WaitGroup coordinates concurrent notifications
