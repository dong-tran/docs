package service

import (
"errors"
"github.com/dong-tran/docs/ddd-example/domain/model"
)

// PricingService is a domain service for pricing logic
type PricingService struct{}

func NewPricingService() *PricingService {
	return &PricingService{}
}

// ApplyDiscount applies a discount to a product
func (s *PricingService) ApplyDiscount(product *model.Product, discountPercent float64) error {
	if discountPercent < 0 || discountPercent > 100 {
		return errors.New("discount must be between 0 and 100")
	}

	currentPrice := product.Price()
	discountAmount := currentPrice.Amount() * (discountPercent / 100)
	newAmount := currentPrice.Amount() - discountAmount

	newPrice, err := model.NewMoney(newAmount, currentPrice.Currency())
	if err != nil {
		return err
	}

	return product.ChangePrice(newPrice)
}
