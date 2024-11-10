package main

import (
    "invoice"
    "email"
    "repository"
)

func main() {
    inv := invoice.NewInvoice(100.0)
    emailService := email.NewEmailService()
    repo := repository.NewInvoiceRepository()
    inv.GenerateInvoice()
    emailService.SendEmailNotification()
    repo.SaveToDatabase()
}
