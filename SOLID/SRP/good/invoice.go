package invoice

type Invoice struct {
    amount float64
}

func NewInvoice(amount float64) *Invoice {
    return &Invoice{amount: amount}
}

func (i *Invoice) GenerateInvoice() {
    println("Invoice generated & printed for amount", i.amount)
}
