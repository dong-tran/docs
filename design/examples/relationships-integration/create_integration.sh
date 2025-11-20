#!/bin/bash

BASE="/Users/mrku/dev/docs/design/examples/relationships-integration"
cd "$BASE"

# Create directory structure
mkdir -p {shared/patterns,domain/order,usecase,repository,handler,infrastructure,cmd}

# ==========================================
# SHARED PATTERNS (Design Patterns + SOLID)
# ==========================================

cat > shared/patterns/observer.go << 'EOF'
package patterns

// Observer Pattern - notifies multiple subscribers of events
type Event struct {
	Type string
	Data interface{}
}

type EventObserver interface {
	OnEvent(event Event)
}

type EventPublisher struct {
	observers []EventObserver
}

func NewEventPublisher() *EventPublisher {
	return &EventPublisher{observers: make([]EventObserver, 0)}
}

func (p *EventPublisher) Subscribe(observer EventObserver) {
	p.observers = append(p.observers, observer)
}

func (p *EventPublisher) Publish(event Event) {
	for _, observer := range p.observers {
		observer.OnEvent(event)
	}
}
EOF

cat > shared/patterns/strategy.go << 'EOF'
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
EOF

cat > shared/patterns/factory.go << 'EOF'
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
EOF

# ==========================================
# DOMAIN LAYER (DDD + SOLID SRP)
# ==========================================

cat > domain/order/order.go << 'EOF'
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
EOF

cat > domain/order/repository.go << 'EOF'
package order

// OrderRepository - Repository interface (DDD pattern)
// Defined in domain layer but implemented in infrastructure (DIP)
type OrderRepository interface {
	Save(order *Order) error
	FindByID(id OrderID) (*Order, error)
	FindByCustomerID(customerID CustomerID) ([]*Order, error)
	Update(order *Order) error
}
EOF

cat > domain/order/events.go << 'EOF'
package order

// Domain Events (DDD pattern)
type OrderCreatedEvent struct {
	OrderID    string
	CustomerID string
	Total      float64
}

type OrderPaidEvent struct {
	OrderID       string
	PaymentMethod string
	Amount        float64
}

type OrderShippedEvent struct {
	OrderID        string
	TrackingNumber string
}
EOF

# ==========================================
# USE CASE LAYER (Clean Architecture)
# ==========================================

cat > usecase/order_usecase.go << 'EOF'
package usecase

import (
"github.com/dong-tran/docs/integration-example/domain/order"
"github.com/dong-tran/docs/integration-example/shared/patterns"
)

// OrderUseCase - Application Service (Clean Architecture Use Case Layer)
// Orchestrates domain objects and publishes events
type OrderUseCase struct {
	orderRepo      order.OrderRepository
	paymentFactory *patterns.PaymentFactory
	eventPublisher *patterns.EventPublisher
}

func NewOrderUseCase(
orderRepo order.OrderRepository,
paymentFactory *patterns.PaymentFactory,
eventPublisher *patterns.EventPublisher,
) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:      orderRepo,
		paymentFactory: paymentFactory,
		eventPublisher: eventPublisher,
	}
}

// CreateOrderDTO - Input DTO
type CreateOrderDTO struct {
	CustomerID string
	Items      []OrderItemDTO
}

type OrderItemDTO struct {
	ProductID   string
	ProductName string
	Quantity    int
	Price       float64
	Currency    string
}

// CreateOrder - Use case method
func (uc *OrderUseCase) CreateOrder(dto CreateOrderDTO) (*order.Order, error) {
	// Convert DTOs to domain objects
	customerID := order.NewCustomerID(dto.CustomerID)
	
	items := make([]order.OrderItem, 0, len(dto.Items))
	for _, itemDTO := range dto.Items {
		price, err := order.NewMoney(itemDTO.Price, itemDTO.Currency)
		if err != nil {
			return nil, err
		}
		
		item, err := order.NewOrderItem(
itemDTO.ProductID,
itemDTO.ProductName,
itemDTO.Quantity,
price,
)
		if err != nil {
			return nil, err
		}
		items = append(items, *item)
	}

	// Create order using domain logic
	newOrder, err := order.NewOrder(customerID, items)
	if err != nil {
		return nil, err
	}

	// Persist
	if err := uc.orderRepo.Save(newOrder); err != nil {
		return nil, err
	}

	// Publish domain event
	uc.eventPublisher.Publish(patterns.Event{
Type: "OrderCreated",
Data: order.OrderCreatedEvent{
OrderID:    newOrder.ID().String(),
			CustomerID: newOrder.CustomerID().String(),
			Total:      newOrder.TotalAmount().Amount(),
		},
	})

	return newOrder, nil
}

// ProcessPayment - Use case using Strategy pattern
func (uc *OrderUseCase) ProcessPayment(orderID string, paymentMethod string) error {
	// Get order
	ord, err := uc.orderRepo.FindByID(order.OrderID{})
	if err != nil {
		return err
	}

	// Use Factory to create payment strategy (Factory + Strategy patterns)
	paymentStrategy, err := uc.paymentFactory.CreatePayment(paymentMethod)
	if err != nil {
		return err
	}

	// Process payment using strategy
	if err := paymentStrategy.ProcessPayment(ord.TotalAmount().Amount(), ord.ID().String()); err != nil {
		return err
	}

	// Update order status (domain logic)
	if err := ord.MarkAsPaid(); err != nil {
		return err
	}

	// Persist changes
	if err := uc.orderRepo.Update(ord); err != nil {
		return err
	}

	// Publish event
	uc.eventPublisher.Publish(patterns.Event{
Type: "OrderPaid",
Data: order.OrderPaidEvent{
OrderID:       ord.ID().String(),
			PaymentMethod: paymentStrategy.GetName(),
			Amount:        ord.TotalAmount().Amount(),
		},
	})

	return nil
}

// GetOrder - Query use case
func (uc *OrderUseCase) GetOrder(orderID string) (*order.Order, error) {
	return uc.orderRepo.FindByID(order.OrderID{})
}

// GetCustomerOrders - Query use case
func (uc *OrderUseCase) GetCustomerOrders(customerID string) ([]*order.Order, error) {
	return uc.orderRepo.FindByCustomerID(order.NewCustomerID(customerID))
}

// ShipOrder - Use case
func (uc *OrderUseCase) ShipOrder(orderID string, trackingNumber string) error {
	ord, err := uc.orderRepo.FindByID(order.OrderID{})
	if err != nil {
		return err
	}

	if err := ord.Ship(); err != nil {
		return err
	}

	if err := uc.orderRepo.Update(ord); err != nil {
		return err
	}

	uc.eventPublisher.Publish(patterns.Event{
Type: "OrderShipped",
Data: order.OrderShippedEvent{
OrderID:        ord.ID().String(),
			TrackingNumber: trackingNumber,
		},
	})

	return nil
}
EOF

# ==========================================
# INFRASTRUCTURE LAYER
# ==========================================

cat > repository/order_repository_impl.go << 'EOF'
package repository

import (
"database/sql"
"encoding/json"

"github.com/dong-tran/docs/integration-example/domain/order"
"github.com/jmoiron/sqlx"
)

// OrderRepositoryImpl - Infrastructure implementation (Clean Architecture + DIP)
type OrderRepositoryImpl struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) order.OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

type orderDB struct {
	ID          string  `db:"id"`
	CustomerID  string  `db:"customer_id"`
	Items       string  `db:"items"`
	TotalAmount float64 `db:"total_amount"`
	Currency    string  `db:"currency"`
	Status      string  `db:"status"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   string  `db:"updated_at"`
}

func (r *OrderRepositoryImpl) Save(ord *order.Order) error {
	itemsJSON, _ := json.Marshal(ord.Items())
	
	query := `
		INSERT INTO orders (id, customer_id, items, total_amount, currency, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query,
ord.ID().String(),
		ord.CustomerID().String(),
		itemsJSON,
		ord.TotalAmount().Amount(),
		ord.TotalAmount().Currency(),
		string(ord.Status()),
		ord.CreatedAt(),
		ord.UpdatedAt(),
	)
	return err
}

func (r *OrderRepositoryImpl) FindByID(id order.OrderID) (*order.Order, error) {
	// Implementation details...
	return nil, sql.ErrNoRows
}

func (r *OrderRepositoryImpl) FindByCustomerID(customerID order.CustomerID) ([]*order.Order, error) {
	// Implementation details...
	return nil, nil
}

func (r *OrderRepositoryImpl) Update(ord *order.Order) error {
	query := `
		UPDATE orders
		SET status = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.Exec(query, string(ord.Status()), ord.UpdatedAt(), ord.ID().String())
	return err
}
EOF

cat > infrastructure/database.go << 'EOF'
package infrastructure

import (
"github.com/jmoiron/sqlx"
_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "./orders.db")
	if err != nil {
		return nil, err
	}

	schema := `
	CREATE TABLE IF NOT EXISTS orders (
id TEXT PRIMARY KEY,
customer_id TEXT NOT NULL,
items TEXT NOT NULL,
total_amount REAL NOT NULL,
currency TEXT NOT NULL,
status TEXT NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL
);
	`

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return db, nil
}
EOF

cat > infrastructure/event_handlers.go << 'EOF'
package infrastructure

import (
"fmt"
"github.com/dong-tran/docs/integration-example/shared/patterns"
)

// Event handlers demonstrating Observer pattern

type EmailNotificationHandler struct{}

func (h *EmailNotificationHandler) OnEvent(event patterns.Event) {
	fmt.Printf("📧 Email Handler: %s - %+v\n", event.Type, event.Data)
}

type LoggingHandler struct{}

func (h *LoggingHandler) OnEvent(event patterns.Event) {
	fmt.Printf("📝 Logger: %s - %+v\n", event.Type, event.Data)
}

type AnalyticsHandler struct{}

func (h *AnalyticsHandler) OnEvent(event patterns.Event) {
	fmt.Printf("📊 Analytics: %s - %+v\n", event.Type, event.Data)
}
EOF

# ==========================================
# HTTP HANDLER LAYER
# ==========================================

cat > handler/order_handler.go << 'EOF'
package handler

import (
"net/http"

"github.com/dong-tran/docs/integration-example/usecase"
"github.com/labstack/echo/v4"
)

// OrderHandler - Presentation layer (Clean Architecture)
type OrderHandler struct {
	orderUseCase *usecase.OrderUseCase
}

func NewOrderHandler(orderUseCase *usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{orderUseCase: orderUseCase}
}

type CreateOrderRequest struct {
	CustomerID string               `json:"customer_id"`
	Items      []OrderItemRequest   `json:"items"`
}

type OrderItemRequest struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
}

type ProcessPaymentRequest struct {
	PaymentMethod string `json:"payment_method"`
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	var req CreateOrderRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Convert to DTO
	dto := usecase.CreateOrderDTO{
		CustomerID: req.CustomerID,
		Items:      make([]usecase.OrderItemDTO, len(req.Items)),
	}

	for i, item := range req.Items {
		dto.Items[i] = usecase.OrderItemDTO{
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			Price:       item.Price,
			Currency:    item.Currency,
		}
	}

	order, err := h.orderUseCase.CreateOrder(dto)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
"id":          order.ID().String(),
		"customer_id": order.CustomerID().String(),
		"total":       order.TotalAmount().Amount(),
		"currency":    order.TotalAmount().Currency(),
		"status":      order.Status(),
	})
}

func (h *OrderHandler) ProcessPayment(c echo.Context) error {
	orderID := c.Param("id")
	
	var req ProcessPaymentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := h.orderUseCase.ProcessPayment(orderID, req.PaymentMethod); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "payment processed"})
}

func (h *OrderHandler) GetOrder(c echo.Context) error {
	orderID := c.Param("id")
	
	order, err := h.orderUseCase.GetOrder(orderID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "order not found"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
"id":          order.ID().String(),
		"customer_id": order.CustomerID().String(),
		"total":       order.TotalAmount().Amount(),
		"currency":    order.TotalAmount().Currency(),
		"status":      order.Status(),
	})
}
EOF

# ==========================================
# MAIN APPLICATION
# ==========================================

cat > cmd/main.go << 'EOF'
package main

import (
"log"

"github.com/dong-tran/docs/integration-example/handler"
"github.com/dong-tran/docs/integration-example/infrastructure"
"github.com/dong-tran/docs/integration-example/repository"
"github.com/dong-tran/docs/integration-example/shared/patterns"
"github.com/dong-tran/docs/integration-example/usecase"
"github.com/labstack/echo/v4"
"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize infrastructure
	db, err := infrastructure.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Setup event system (Observer pattern)
	eventPublisher := patterns.NewEventPublisher()
	eventPublisher.Subscribe(&infrastructure.EmailNotificationHandler{})
	eventPublisher.Subscribe(&infrastructure.LoggingHandler{})
	eventPublisher.Subscribe(&infrastructure.AnalyticsHandler{})

	// Setup factories (Factory pattern)
	paymentFactory := patterns.NewPaymentFactory()

	// Dependency injection (DIP)
	orderRepo := repository.NewOrderRepository(db)
	orderUseCase := usecase.NewOrderUseCase(orderRepo, paymentFactory, eventPublisher)
	orderHandler := handler.NewOrderHandler(orderUseCase)

	// Setup Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.POST("/orders", orderHandler.CreateOrder)
	e.GET("/orders/:id", orderHandler.GetOrder)
	e.POST("/orders/:id/payment", orderHandler.ProcessPayment)

	log.Println("🚀 Integration Example Server starting on :8080")
	log.Println("📚 Demonstrates: Clean Architecture + DDD + SOLID + Design Patterns + Microservices concepts")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
EOF

echo "Integration example created successfully!"
