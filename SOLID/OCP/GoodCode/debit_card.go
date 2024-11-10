package payment

type DebitCard struct{}

func NewDebitCard() *DebitCard {
    return &DebitCard{}
}

func (d *DebitCard) Pay(amount float64) error {
    println("Making payment via Debit Card:", amount)
    return nil
}
