package model

import (
"errors"
"time"

"github.com/google/uuid"
)

// Product is an aggregate root
type Product struct {
	id          ProductID
	name        string
	description string
	price       Money
	category    Category
	createdAt   time.Time
	updatedAt   time.Time
}

type ProductID struct {
	value string
}

func NewProductID() ProductID {
	return ProductID{value: uuid.New().String()}
}

func (id ProductID) String() string {
	return id.value
}

// Money is a value object
type Money struct {
	amount   float64
	currency string
}

func NewMoney(amount float64, currency string) (Money, error) {
	if amount < 0 {
		return Money{}, errors.New("money amount cannot be negative")
	}
	return Money{amount: amount, currency: currency}, nil
}

func (m Money) Amount() float64 {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}

// Category is a value object
type Category struct {
	name string
}

func NewCategory(name string) (Category, error) {
	if name == "" {
		return Category{}, errors.New("category name cannot be empty")
	}
	return Category{name: name}, nil
}

func (c Category) Name() string {
	return c.name
}

// NewProduct creates a new product aggregate
func NewProduct(name, description string, price Money, category Category) (*Product, error) {
	if name == "" {
		return nil, errors.New("product name cannot be empty")
	}

	now := time.Now()
	return &Product{
		id:          NewProductID(),
		name:        name,
		description: description,
		price:       price,
		category:    category,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

func (p *Product) ID() ProductID {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Description() string {
	return p.description
}

func (p *Product) Price() Money {
	return p.price
}

func (p *Product) Category() Category {
	return p.category
}

func (p *Product) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Product) UpdatedAt() time.Time {
	return p.updatedAt
}

// ChangePrice is a domain method
func (p *Product) ChangePrice(newPrice Money) error {
	if newPrice.amount <= 0 {
		return errors.New("price must be positive")
	}
	p.price = newPrice
	p.updatedAt = time.Now()
	return nil
}

// UpdateInfo updates product information
func (p *Product) UpdateInfo(name, description string) error {
	if name == "" {
		return errors.New("product name cannot be empty")
	}
	p.name = name
	p.description = description
	p.updatedAt = time.Now()
	return nil
}
