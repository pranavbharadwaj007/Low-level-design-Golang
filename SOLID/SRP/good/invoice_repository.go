package repository

type InvoiceRepository struct{}

func NewInvoiceRepository() *InvoiceRepository {
    return &InvoiceRepository{}
}

func (r *InvoiceRepository) SaveToDatabase() {
    println("Saving invoice to Database")
}
