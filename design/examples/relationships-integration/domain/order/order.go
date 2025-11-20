package order

import (
"errors"
"time"

"github.com/google/uuid"
)

// Order - DDD Aggregate Root with business rules
type Order struct {
	id          OrderID
	customerID  CustomerID
	items       []OrderItem
	totalAmount Money
	status      OrderStatus
	createdAt   time.Time
	updatedAt   time.Time
}

// OrderID - Value Object
type OrderID struct {
	value string
}

func NewOrderID() OrderID {
	return OrderID{value: uuid.New().String()}
}

func (id OrderID) String() string {
	return id.value
}

// CustomerID - Value Object
type CustomerID struct {
	value string
}

func NewCustomerID(value string) CustomerID {
	return CustomerID{value: value}
}

func (id CustomerID) String() string {
	return id.value
}

// Money - Value Object (immutable)
type Money struct {
	amount   float64
	currency string
}

func NewMoney(amount float64, currency string) (Money, error) {
	if amount < 0 {
		return Money{}, errors.New("amount cannot be negative")
	}
	if currency == "" {
		currency = "USD"
	}
	return Money{amount: amount, currency: currency}, nil
}

func (m Money) Amount() float64 {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, errors.New("currency mismatch")
	}
	return NewMoney(m.amount+other.amount, m.currency)
}

// OrderStatus - Value Object
type OrderStatus string

const (
OrderStatusPending   OrderStatus = "PENDING"
OrderStatusPaid      OrderStatus = "PAID"
OrderStatusShipped   OrderStatus = "SHIPPED"
OrderStatusDelivered OrderStatus = "DELIVERED"
OrderStatusCancelled OrderStatus = "CANCELLED"
)

// OrderItem - Entity within Order aggregate
type OrderItem struct {
	productID   string
	productName string
	quantity    int
	price       Money
}

func NewOrderItem(productID, productName string, quantity int, price Money) (*OrderItem, error) {
	if quantity <= 0 {
		return nil, errors.New("quantity must be positive")
	}
	return &OrderItem{
		productID:   productID,
		productName: productName,
		quantity:    quantity,
		price:       price,
	}, nil
}

func (i *OrderItem) Total() Money {
	total, _ := NewMoney(i.price.Amount()*float64(i.quantity), i.price.Currency())
	return total
}

// NewOrder - Factory method for creating orders
func NewOrder(customerID CustomerID, items []OrderItem) (*Order, error) {
	if len(items) == 0 {
		return nil, errors.New("order must have at least one item")
	}

	total, err := NewMoney(0, "USD")
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		itemTotal := item.Total()
		total, err = total.Add(itemTotal)
		if err != nil {
			return nil, err
		}
	}

	now := time.Now()
	return &Order{
		id:          NewOrderID(),
		customerID:  customerID,
		items:       items,
		totalAmount: total,
		status:      OrderStatusPending,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

// Domain Methods (Business Logic)

func (o *Order) ID() OrderID {
	return o.id
}

func (o *Order) CustomerID() CustomerID {
	return o.customerID
}

func (o *Order) Items() []OrderItem {
	return o.items
}

func (o *Order) TotalAmount() Money {
	return o.totalAmount
}

func (o *Order) Status() OrderStatus {
	return o.status
}

func (o *Order) CreatedAt() time.Time {
	return o.createdAt
}

func (o *Order) UpdatedAt() time.Time {
	return o.updatedAt
}

// MarkAsPaid - Domain method with business rules
func (o *Order) MarkAsPaid() error {
	if o.status != OrderStatusPending {
		return errors.New("only pending orders can be marked as paid")
	}
	o.status = OrderStatusPaid
	o.updatedAt = time.Now()
	return nil
}

// Ship - Domain method
func (o *Order) Ship() error {
	if o.status != OrderStatusPaid {
		return errors.New("only paid orders can be shipped")
	}
	o.status = OrderStatusShipped
	o.updatedAt = time.Now()
	return nil
}

// Cancel - Domain method
func (o *Order) Cancel() error {
	if o.status == OrderStatusShipped || o.status == OrderStatusDelivered {
		return errors.New("cannot cancel shipped or delivered orders")
	}
	o.status = OrderStatusCancelled
	o.updatedAt = time.Now()
	return nil
}
