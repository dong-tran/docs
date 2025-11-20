package patterns

import "errors"

// Factory Pattern - creates payment strategies
type PaymentFactory struct{}

func NewPaymentFactory() *PaymentFactory {
	return &PaymentFactory{}
}

func (f *PaymentFactory) CreatePayment(paymentType string) (PaymentStrategy, error) {
	switch paymentType {
	case "credit_card":
		return &CreditCardPayment{}, nil
	case "paypal":
		return &PayPalPayment{}, nil
	case "crypto":
		return &CryptoPayment{}, nil
	default:
		return nil, errors.New("unsupported payment type")
	}
}
