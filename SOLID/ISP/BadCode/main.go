package main

import "bad"

func main() {
    doc := &bad.Document{}
    
    // This works fine
    multiMachine := bad.NewMultiPurposeMachine("HP OfficeJet")
    multiMachine.Print(doc)
    multiMachine.Scan(doc)
    multiMachine.Copy(doc)
    
    // This will panic!
    simplePrinter := bad.NewSimplePrinter("HP DeskJet")
    simplePrinter.Print(doc)  // Works
    simplePrinter.Scan(doc)   // Panics!
    simplePrinter.Copy(doc)   // Panics!
}
