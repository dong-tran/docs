package patterns

// Strategy Pattern - interchangeable algorithms (OCP + DIP)
type PaymentStrategy interface {
	ProcessPayment(amount float64, orderID string) error
	GetName() string
}

type CreditCardPayment struct{}

func (c *CreditCardPayment) ProcessPayment(amount float64, orderID string) error {
	// Process credit card payment
	return nil
}

func (c *CreditCardPayment) GetName() string {
	return "Credit Card"
}

type PayPalPayment struct{}

func (p *PayPalPayment) ProcessPayment(amount float64, orderID string) error {
	// Process PayPal payment
	return nil
}

func (p *PayPalPayment) GetName() string {
	return "PayPal"
}

type CryptoPayment struct{}

func (c *CryptoPayment) ProcessPayment(amount float64, orderID string) error {
	// Process cryptocurrency payment
	return nil
}

func (c *CryptoPayment) GetName() string {
	return "Cryptocurrency"
}
