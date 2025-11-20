package behavioral

// Strategy - Behavioral Pattern
// Defines family of algorithms, encapsulates each one, makes them interchangeable

type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCardStrategy struct {
	cardNumber string
}

func (c *CreditCardStrategy) Pay(amount float64) string {
	return "Paid with credit card"
}

type PayPalStrategy struct {
	email string
}

func (p *PayPalStrategy) Pay(amount float64) string {
	return "Paid with PayPal"
}

type BitcoinStrategy struct {
	walletAddress string
}

func (b *BitcoinStrategy) Pay(amount float64) string {
	return "Paid with Bitcoin"
}

type ShoppingCart struct {
	strategy PaymentStrategy
}

func (s *ShoppingCart) SetStrategy(strategy PaymentStrategy) {
	s.strategy = strategy
}

func (s *ShoppingCart) Checkout(amount float64) string {
	return s.strategy.Pay(amount)
}
