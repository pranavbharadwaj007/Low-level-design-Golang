package payment

type CreditCard struct{}

func NewCreditCard() *CreditCard {
    return &CreditCard{}
}

func (c *CreditCard) Pay(amount float64) error {
    println("Making payment via Credit Card:", amount)
    return nil
}
