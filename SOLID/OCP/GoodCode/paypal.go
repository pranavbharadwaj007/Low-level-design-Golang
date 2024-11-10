package payment

type PayPal struct{}

func NewPayPal() *PayPal {
    return &PayPal{}
}

func (p *PayPal) Pay(amount float64) error {
    println("Making payment via PayPal:", amount)
    return nil
}
