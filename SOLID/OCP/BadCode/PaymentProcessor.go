package bad

type PaymentProcessor struct{}

func NewPaymentProcessor() *PaymentProcessor {
    return &PaymentProcessor{}
}

func (p *PaymentProcessor) ProcessPayment(paymentMethod string, amount float64) error {
    switch paymentMethod {
    case "CreditCard":
        println("Making payment via Credit Card:", amount)
    case "DebitCard":
        println("Making payment via Debit Card:", amount)
    case "PayPal":
        println("Making payment via PayPal:", amount)
    default:
        return fmt.Errorf("unsupported payment method: %s", paymentMethod)
    }
    return nil
}
