package dip

// BAD: Violates DIP - high-level depends on low-level
type StripePayment struct{}

func (s *StripePayment) ProcessPayment(amount float64) error {
// Stripe-specific implementation
return nil
}

type OrderServiceBad struct {
stripe *StripePayment // Depends on concrete implementation
}

func (o *OrderServiceBad) ProcessOrder(amount float64) error {
return o.stripe.ProcessPayment(amount)
// Locked to Stripe! Can't switch to PayPal without changing code
}

// GOOD: Follows DIP - both depend on abstraction
type PaymentProcessor interface {
	Process(amount float64) error
}

type OrderService struct {
	processor PaymentProcessor // Depends on abstraction
}

func (o *OrderService) ProcessOrder(amount float64) error {
	return o.processor.Process(amount)
}

// Low-level modules implement the abstraction
type StripeProcessor struct{}

func (s *StripeProcessor) Process(amount float64) error {
	// Stripe implementation
	return nil
}

type PayPalProcessor struct{}

func (p *PayPalProcessor) Process(amount float64) error {
	// PayPal implementation
	return nil
}

type CryptoProcessor struct{}

func (c *CryptoProcessor) Process(amount float64) error {
	// Cryptocurrency implementation
	return nil
}

// Can inject any processor without changing OrderService
func NewOrderService(processor PaymentProcessor) *OrderService {
	return &OrderService{processor: processor}
}
