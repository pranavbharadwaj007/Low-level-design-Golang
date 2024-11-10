package payment

type PaymentProcessor struct{}

func NewPaymentProcessor() *PaymentProcessor {
    return &PaymentProcessor{}
}

func (p *PaymentProcessor) ProcessPayment(method PaymentMethod, amount float64) error {
    return method.Pay(amount)
}
