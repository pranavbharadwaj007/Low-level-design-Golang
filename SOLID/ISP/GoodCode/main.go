package main

import "good"

// Generic function that works with any printer
func printDocument(printer good.Printer, doc *good.Document) {
    printer.Print(doc)
}

// Generic function that works with any scanner
func scanDocument(scanner good.Scanner, doc *good.Document) {
    scanner.Scan(doc)
}

func main() {
    doc := &good.Document{}

    // Multi-purpose machine
    multiMachine := good.NewMultiPurposeMachine("HP OfficeJet")
    printDocument(multiMachine, doc)
    scanDocument(multiMachine, doc)
    multiMachine.Copy(doc)

    // Simple printer
    simplePrinter := good.NewSimplePrinter("HP DeskJet")
    printDocument(simplePrinter, doc)
    // Can't call Scan or Copy - not implemented!

    // Scanner only
    scanner := good.NewScannerOnly("Epson Scanner")
    scanDocument(scanner, doc)
    // Can't call Print or Copy - not implemented!
}
