package payment

type UPI struct{}

func NewUPI() *UPI {
    return &UPI{}
}

func (u *UPI) Pay(amount float64) error {
    println("Making payment via UPI:", amount)
    return nil
}
