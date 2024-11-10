package main

import (
    "fmt"
    "payment"
)

func main() {
    // Create payment processor
    processor := payment.NewPaymentProcessor()

    // Create different payment methods
    creditCard := payment.NewCreditCard()
    debitCard := payment.NewDebitCard()
    paypal := payment.NewPayPal()
    upi := payment.NewUPI()

    // Process payments using different methods
    _ = processor.ProcessPayment(creditCard, 100.0)
    _ = processor.ProcessPayment(debitCard, 200.0)
    _ = processor.ProcessPayment(paypal, 150.0)
    _ = processor.ProcessPayment(upi, 175.0)
}
