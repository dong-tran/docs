# Relationships & Integration
## Connecting Architecture Patterns, Design Principles, and Development Practices

**Presenter:** Dong Tran  
**Date:** November 20, 2025  
**Time:** 03:28 UTC

---

## Slide 1: Title Slide

# Relationships & Integration
### Building Cohesive Software Systems Through Pattern Synergy

**Presenter:** Dong Tran  
**Date:** November 20, 2025

*"In software architecture, the whole is greater than the sum of its parts. The true power emerges when patterns, principles, and practices work in harmony."*

---

## Slide 2: Agenda

### Today's Integration Journey

1. **The Integration Challenge**
2. **Architecture Patterns Relationships**
3. **Design Principles Harmony**
4. **Pattern + Principle Integration**
5. **Layered Architecture Integration**
6. **Microservices + DDD + Clean Architecture**
7. **SOLID + GoF Patterns Synergy**
8. **Event-Driven Integration Patterns**
9. **Data Integration Strategies**
10. **API Integration Patterns**
11. **Testing Integration Strategies**
12. **DevOps Integration**
13. **Real-World Case Study: E-Commerce Platform**
14. **Integration Anti-Patterns**
15. **Best Practices & Guidelines**
16. **Q&A Session**

---

## Slide 3: The Integration Challenge

### Why Integration Matters

**The Reality of Modern Systems:**
```
Typical Enterprise System:
├── Multiple Architecture Patterns
├── Various Design Principles  
├── Different Technology Stacks
├── Legacy Systems
├── Third-Party Services
└── Cross-Team Dependencies
```

**Integration Challenges:**
- 🔴 **Pattern Conflicts** - Different patterns may contradict
- 🔴 **Principle Violations** - Integration can break principles
- 🔴 **Complexity Explosion** - More connections = more complexity
- 🔴 **Consistency Issues** - Maintaining coherent design
- 🔴 **Team Boundaries** - Conway's Law effects

**The Goal:**
```
Isolated Patterns          Integrated System
    ┌───┐                    ┌───────────┐
    │ A │                    │     A     │
    └───┘                    │  ┌───┐    │
    ┌───┐         →          │  │ B │    │
    │ B │                    │  └─┬─┘    │
    └───┘                    │    │ C    │
    ┌───┐                    └────┴──────┘
    │ C │                   Synergistic Whole
    └───┘
```

---

## Slide 4: Architecture Patterns Relationships

### How Major Patterns Connect

```go
// Relationship Matrix
type ArchitectureRelationship struct {
    Pattern1     string
    Pattern2     string
    Relationship string
    Integration  string
}

var relationships = []ArchitectureRelationship{
    {
        Pattern1:     "Clean Architecture",
        Pattern2:     "Domain-Driven Design",
        Relationship: "Complementary",
        Integration:  "DDD Entities in Clean Architecture Core",
    },
    {
        Pattern1:     "Microservices",
        Pattern2:     "Domain-Driven Design", 
        Relationship: "Natural Fit",
        Integration:  "Bounded Contexts as Service Boundaries",
    },
    {
        Pattern1:     "Clean Architecture",
        Pattern2:     "Hexagonal Architecture",
        Relationship: "Similar Philosophy",
        Integration:  "Ports & Adapters = Interface Adapters",
    },
    {
        Pattern1:     "Event Sourcing",
        Pattern2:     "CQRS",
        Relationship: "Synergistic",
        Integration:  "Event Store feeds Read Models",
    },
}
```

**Visual Integration Map:**
```
┌─────────────────────────────────────────────────┐
│                 Microservices                   │
│  ┌──────────────┐  ┌──────────────┐            │
│  │   Service A  │  │   Service B  │            │
│  │ ┌──────────┐ │  │ ┌──────────┐ │            │
│  │ │  Clean   │ │  │ │   DDD    │ │            │
│  │ │   Arch   │ │  │ │ Bounded  │ │            │
│  │ │          │ │  │ │ Context  │ │            │
│  │ └──────────┘ │  │ └──────────┘ │            │
│  └──────┬───────┘  └───────┬──────┘            │
│         │                   │                    │
│  ┌──────▼───────────────────▼──────┐            │
│  │      Event Streaming Bus        │            │
│  │         (Kafka/RabbitMQ)        │            │
│  └──────────────────────────────────┘            │
└─────────────────────────────────────────────────┘
```

---

## Slide 5: Clean Architecture + DDD Integration

### Combining Clean Architecture with Domain-Driven Design

```go
// project structure showing integration
// project/
// ├── domain/           (DDD + Clean Architecture Entities)
// ├── application/      (Clean Architecture Use Cases)
// ├── infrastructure/   (Clean Architecture Frameworks)
// └── interfaces/       (Clean Architecture Interface Adapters)

// Domain Layer (DDD Entities + Clean Architecture Entities)
package domain

import (
    "errors"
    "time"
)

// DDD Value Object
type Money struct {
    Amount   decimal.Decimal
    Currency string
}

func NewMoney(amount decimal.Decimal, currency string) (Money, error) {
    if amount.LessThan(decimal.Zero) {
        return Money{}, errors.New("amount cannot be negative")
    }
    if currency == "" {
        return Money{}, errors.New("currency is required")
    }
    return Money{Amount: amount, Currency: currency}, nil
}

func (m Money) Add(other Money) (Money, error) {
    if m.Currency != other.Currency {
        return Money{}, errors.New("cannot add different currencies")
    }
    return Money{
        Amount:   m.Amount.Add(other.Amount),
        Currency: m.Currency,
    }, nil
}

// DDD Aggregate Root (Clean Architecture Entity)
type Order struct {
    id         OrderID
    customerID CustomerID
    items      []OrderItem
    total      Money
    status     OrderStatus
    placedAt   time.Time
    events     []DomainEvent // DDD Domain Events
}

// DDD Domain Service
type PricingDomainService struct {
    // No dependencies on infrastructure
}

func (s *PricingDomainService) CalculateDiscount(
    order *Order,
    customer *Customer,
) (Money, error) {
    // Pure business logic
    if customer.IsVIP() {
        return order.total.Multiply(0.20), nil
    }
    if order.total.Amount.GreaterThan(decimal.NewFromFloat(100)) {
        return order.total.Multiply(0.10), nil
    }
    return Money{Amount: decimal.Zero, Currency: order.total.Currency}, nil
}

// Application Layer (Clean Architecture Use Cases)
package application

import (
    "context"
    "project/domain"
)

// Use Case with DDD Repository
type PlaceOrderUseCase struct {
    orderRepo    domain.OrderRepository    // DDD Repository Interface
    customerRepo domain.CustomerRepository
    eventBus     domain.EventPublisher     // Domain Events Publisher
    pricing      *domain.PricingDomainService
}

func (uc *PlaceOrderUseCase) Execute(
    ctx context.Context,
    input PlaceOrderInput,
) (*PlaceOrderOutput, error) {
    // Load aggregate
    customer, err := uc.customerRepo.FindByID(ctx, input.CustomerID)
    if err != nil {
        return nil, err
    }
    
    // Create order using DDD patterns
    order, err := domain.NewOrder(customer.ID, input.Items)
    if err != nil {
        return nil, err
    }
    
    // Apply domain service
    discount, err := uc.pricing.CalculateDiscount(order, customer)
    if err != nil {
        return nil, err
    }
    
    order.ApplyDiscount(discount)
    
    // Domain logic
    if err := order.Place(); err != nil {
        return nil, err
    }
    
    // Persist aggregate
    if err := uc.orderRepo.Save(ctx, order); err != nil {
        return nil, err
    }
    
    // Publish domain events
    for _, event := range order.Events() {
        uc.eventBus.Publish(ctx, event)
    }
    
    return &PlaceOrderOutput{
        OrderID: order.ID().String(),
        Total:   order.Total().Amount.Float64(),
    }, nil
}

// Infrastructure Layer (Clean Architecture + DDD)
package infrastructure

import (
    "context"
    "database/sql"
    "project/domain"
)

// DDD Repository Implementation
type PostgresOrderRepository struct {
    db *sql.DB
}

func (r *PostgresOrderRepository) Save(
    ctx context.Context,
    order *domain.Order,
) error {
    // Map domain model to database
    tx, err := r.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    // Save aggregate root
    _, err = tx.ExecContext(ctx, `
        INSERT INTO orders (id, customer_id, total_amount, total_currency, status, placed_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `, order.ID(), order.CustomerID(), order.Total().Amount, 
       order.Total().Currency, order.Status(), order.PlacedAt())
    
    if err != nil {
        return err
    }
    
    // Save aggregate members
    for _, item := range order.Items() {
        _, err = tx.ExecContext(ctx, `
            INSERT INTO order_items (order_id, product_id, quantity, price)
            VALUES ($1, $2, $3, $4)
        `, order.ID(), item.ProductID, item.Quantity, item.Price.Amount)
        
        if err != nil {
            return err
        }
    }
    
    return tx.Commit()
}

func (r *PostgresOrderRepository) FindByID(
    ctx context.Context,
    id domain.OrderID,
) (*domain.Order, error) {
    // Reconstruct aggregate from database
    // ...
    return nil, nil
}
```

---

## Slide 6: Microservices + Event-Driven Architecture

### Integration Through Events

```go
// Event-Driven Microservices Integration

// Shared Event Definitions (Schema Registry)
package events

import (
    "time"
    "encoding/json"
)

// Event Interface
type Event interface {
    EventType() string
    AggregateID() string
    Timestamp() time.Time
    Version() string
}

// Order Events
type OrderPlaced struct {
    OrderID    string    `json:"order_id"`
    CustomerID string    `json:"customer_id"`
    Items      []Item    `json:"items"`
    Total      float64   `json:"total"`
    PlacedAt   time.Time `json:"placed_at"`
}

func (e OrderPlaced) EventType() string   { return "order.placed" }
func (e OrderPlaced) AggregateID() string { return e.OrderID }
func (e OrderPlaced) Timestamp() time.Time { return e.PlacedAt }
func (e OrderPlaced) Version() string     { return "v1" }

// Service A: Order Service
package orderservice

import (
    "context"
    "events"
)

type OrderService struct {
    repo      OrderRepository
    publisher EventPublisher
}

func (s *OrderService) CreateOrder(ctx context.Context, req CreateOrderRequest) error {
    // Business logic
    order := &Order{
        ID:         generateID(),
        CustomerID: req.CustomerID,
        Items:      req.Items,
        Status:     "pending",
    }
    
    // Save to database
    if err := s.repo.Save(ctx, order); err != nil {
        return err
    }
    
    // Publish event for other services
    event := events.OrderPlaced{
        OrderID:    order.ID,
        CustomerID: order.CustomerID,
        Items:      order.Items,
        Total:      order.CalculateTotal(),
        PlacedAt:   time.Now(),
    }
    
    return s.publisher.Publish(ctx, event)
}

// Service B: Inventory Service
package inventoryservice

import (
    "context"
    "events"
)

type InventoryService struct {
    repo      InventoryRepository
    consumer  EventConsumer
    publisher EventPublisher
}

func (s *InventoryService) Start(ctx context.Context) {
    s.consumer.Subscribe("order.placed", s.handleOrderPlaced)
    s.consumer.Subscribe("payment.completed", s.handlePaymentCompleted)
}

func (s *InventoryService) handleOrderPlaced(ctx context.Context, event events.Event) error {
    orderPlaced := event.(events.OrderPlaced)
    
    // Reserve inventory
    for _, item := range orderPlaced.Items {
        if err := s.reserveStock(ctx, item.ProductID, item.Quantity); err != nil {
            // Publish compensation event
            return s.publisher.Publish(ctx, events.InventoryReservationFailed{
                OrderID: orderPlaced.OrderID,
                Reason:  err.Error(),
            })
        }
    }
    
    // Publish success event
    return s.publisher.Publish(ctx, events.InventoryReserved{
        OrderID:    orderPlaced.OrderID,
        ReservedAt: time.Now(),
    })
}

// Service C: Notification Service
package notificationservice

type NotificationService struct {
    consumer     EventConsumer
    emailSender  EmailSender
    smsSender    SMSSender
    pushNotifier PushNotifier
}

func (s *NotificationService) Start(ctx context.Context) {
    // Listen to multiple events from different services
    s.consumer.Subscribe("order.placed", s.notifyOrderPlaced)
    s.consumer.Subscribe("order.shipped", s.notifyOrderShipped)
    s.consumer.Subscribe("payment.completed", s.notifyPaymentCompleted)
    s.consumer.Subscribe("inventory.low", s.notifyLowInventory)
}

// Event Choreography vs Orchestration
// Choreography - Services react to events independently
type ChoreographyExample struct{}

func (c *ChoreographyExample) OrderFlow() {
    // Each service listens and publishes events
    // No central coordinator
    /*
    Order Service    → publishes → OrderPlaced
                                      ↓
    Payment Service  ← subscribes ← OrderPlaced
                     → publishes → PaymentProcessed
                                      ↓
    Inventory Service ← subscribes ← PaymentProcessed
                      → publishes → ItemsShipped
    */
}

// Orchestration - Central coordinator
type OrderSagaOrchestrator struct {
    orderService     OrderService
    paymentService   PaymentService
    inventoryService InventoryService
    shippingService  ShippingService
}

func (o *OrderSagaOrchestrator) ProcessOrder(ctx context.Context, order Order) error {
    // Step 1: Create Order
    if err := o.orderService.Create(ctx, order); err != nil {
        return err
    }
    
    // Step 2: Process Payment
    payment, err := o.paymentService.Process(ctx, order.Payment)
    if err != nil {
        o.orderService.Cancel(ctx, order.ID) // Compensation
        return err
    }
    
    // Step 3: Reserve Inventory
    reservation, err := o.inventoryService.Reserve(ctx, order.Items)
    if err != nil {
        o.paymentService.Refund(ctx, payment.ID) // Compensation
        o.orderService.Cancel(ctx, order.ID)     // Compensation
        return err
    }
    
    // Step 4: Create Shipment
    if err := o.shippingService.Ship(ctx, order); err != nil {
        o.inventoryService.Release(ctx, reservation.ID) // Compensation
        o.paymentService.Refund(ctx, payment.ID)       // Compensation
        o.orderService.Cancel(ctx, order.ID)           // Compensation
        return err
    }
    
    return nil
}
```

---

## Slide 7: SOLID Principles + GoF Patterns

### How Patterns Implement SOLID

```go
// SOLID Principles demonstrated through GoF Patterns

// 1. Strategy Pattern implements Open/Closed Principle
type PaymentStrategy interface {
    ProcessPayment(amount float64) error
}

// Closed for modification - interface doesn't change
// Open for extension - add new payment methods
type CreditCardPayment struct {
    cardNumber string
}

func (c *CreditCardPayment) ProcessPayment(amount float64) error {
    fmt.Printf("Processing $%.2f via Credit Card\n", amount)
    return nil
}

type PayPalPayment struct {
    email string
}

func (p *PayPalPayment) ProcessPayment(amount float64) error {
    fmt.Printf("Processing $%.2f via PayPal\n", amount)
    return nil
}

// New payment method - no existing code changed (OCP)
type CryptoPayment struct {
    walletAddress string
}

func (c *CryptoPayment) ProcessPayment(amount float64) error {
    fmt.Printf("Processing $%.2f via Crypto\n", amount)
    return nil
}

// 2. Decorator Pattern implements Single Responsibility Principle
// Each decorator has single responsibility

// Core responsibility - basic logging
type Logger interface {
    Log(message string)
}

type BasicLogger struct{}

func (l *BasicLogger) Log(message string) {
    fmt.Println(message)
}

// Single responsibility - add timestamp
type TimestampDecorator struct {
    logger Logger
}

func (t *TimestampDecorator) Log(message string) {
    timestamped := fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), message)
    t.logger.Log(timestamped)
}

// Single responsibility - add log level
type LevelDecorator struct {
    logger Logger
    level  string
}

func (l *LevelDecorator) Log(message string) {
    leveled := fmt.Sprintf("[%s] %s", l.level, message)
    l.logger.Log(leveled)
}

// 3. Factory Pattern implements Dependency Inversion Principle
// High-level modules depend on abstractions, not concrete implementations

// Abstraction (high-level)
type Database interface {
    Connect() error
    Query(sql string) ([]Row, error)
}

// Factory depends on abstraction
type DatabaseFactory struct{}

func (f *DatabaseFactory) Create(dbType string) Database {
    switch dbType {
    case "mysql":
        return &MySQLDatabase{} // Low-level module
    case "postgres":
        return &PostgresDatabase{} // Low-level module
    default:
        return nil
    }
}

// Business logic depends on abstraction (DIP)
type UserService struct {
    db Database // Depends on interface, not concrete type
}

func NewUserService(db Database) *UserService {
    return &UserService{db: db}
}

// 4. Observer Pattern implements Interface Segregation Principle
// Observers only implement what they need

type EventListener interface {
    OnEvent(event Event)
}

// Segregated interfaces for different event types
type OrderEventListener interface {
    OnOrderCreated(order Order)
}

type PaymentEventListener interface {
    OnPaymentReceived(payment Payment)
}

type ShippingEventListener interface {
    OnItemShipped(shipment Shipment)
}

// Concrete observer implements only needed interfaces
type InventoryManager struct{}

func (i *InventoryManager) OnOrderCreated(order Order) {
    fmt.Printf("Inventory: Reserving items for order %s\n", order.ID)
}

func (i *InventoryManager) OnItemShipped(shipment Shipment) {
    fmt.Printf("Inventory: Updating stock for shipment %s\n", shipment.ID)
}

// Doesn't need to implement OnPaymentReceived (ISP)

// 5. Template Method implements Liskov Substitution Principle
// Subclasses can be substituted without breaking behavior

type DataProcessor interface {
    Process(data []byte) ([]byte, error)
}

// Template
type BaseProcessor struct{}

func (b *BaseProcessor) Process(data []byte) ([]byte, error) {
    // 1. Validate
    if err := b.validate(data); err != nil {
        return nil, err
    }
    
    // 2. Transform
    transformed := b.transform(data)
    
    // 3. Save
    return b.save(transformed)
}

func (b *BaseProcessor) validate(data []byte) error {
    if len(data) == 0 {
        return errors.New("empty data")
    }
    return nil
}

func (b *BaseProcessor) transform(data []byte) []byte {
    return data // Default: no transformation
}

func (b *BaseProcessor) save(data []byte) ([]byte, error) {
    return data, nil
}

// Concrete processor - substitutable (LSP)
type JSONProcessor struct {
    BaseProcessor
}

func (j *JSONProcessor) transform(data []byte) []byte {
    // JSON-specific transformation
    var obj map[string]interface{}
    json.Unmarshal(data, &obj)
    formatted, _ := json.MarshalIndent(obj, "", "  ")
    return formatted
}

// Another concrete processor - also substitutable (LSP)
type XMLProcessor struct {
    BaseProcessor
}

func (x *XMLProcessor) transform(data []byte) []byte {
    // XML-specific transformation
    return data
}

// Client code works with any processor (LSP)
func ProcessData(processor DataProcessor, data []byte) ([]byte, error) {
    return processor.Process(data)
}
```

---

## Slide 8: API Gateway Integration Pattern

### Unified Entry Point for Microservices

```go
// API Gateway integrating multiple patterns

package gateway

import (
    "context"
    "time"
    "sync"
)

// API Gateway with multiple responsibilities
type APIGateway struct {
    // Service clients
    userService      ServiceClient
    orderService     ServiceClient
    inventoryService ServiceClient
    
    // Cross-cutting concerns
    rateLimiter  RateLimiter
    cache        Cache
    auth         Authenticator
    circuitBreaker CircuitBreaker
    logger       Logger
}

// Request aggregation pattern
func (g *APIGateway) GetUserDashboard(ctx context.Context, userID string) (*Dashboard, error) {
    // Authentication
    if err := g.auth.Validate(ctx); err != nil {
        return nil, err
    }
    
    // Rate limiting
    if !g.rateLimiter.Allow(userID) {
        return nil, ErrRateLimited
    }
    
    // Check cache
    if cached, found := g.cache.Get("dashboard:" + userID); found {
        return cached.(*Dashboard), nil
    }
    
    // Parallel service calls with circuit breaker
    var (
        wg        sync.WaitGroup
        mu        sync.Mutex
        dashboard = &Dashboard{}
        errors    []error
    )
    
    // User info
    wg.Add(1)
    go func() {
        defer wg.Done()
        
        result, err := g.circuitBreaker.Execute(func() (interface{}, error) {
            return g.userService.GetUser(ctx, userID)
        })
        
        mu.Lock()
        defer mu.Unlock()
        
        if err != nil {
            errors = append(errors, err)
            // Fallback to cached or default
            dashboard.User = g.getUserFallback(userID)
        } else {
            dashboard.User = result.(*User)
        }
    }()
    
    // Recent orders
    wg.Add(1)
    go func() {
        defer wg.Done()
        
        result, err := g.orderService.GetUserOrders(ctx, userID)
        
        mu.Lock()
        defer mu.Unlock()
        
        if err != nil {
            g.logger.Warn("Failed to fetch orders", "error", err)
            dashboard.Orders = []Order{} // Empty list as fallback
        } else {
            dashboard.Orders = result
        }
    }()
    
    // Recommendations (non-critical)
    wg.Add(1)
    go func() {
        defer wg.Done()
        
        ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
        defer cancel()
        
        recommendations, _ := g.inventoryService.GetRecommendations(ctx, userID)
        
        mu.Lock()
        dashboard.Recommendations = recommendations
        mu.Unlock()
    }()
    
    wg.Wait()
    
    // Cache result
    g.cache.Set("dashboard:"+userID, dashboard, 5*time.Minute)
    
    return dashboard, nil
}

// Backend for Frontend (BFF) Pattern
type MobileBFF struct {
    gateway *APIGateway
}

func (m *MobileBFF) GetMobileUserProfile(ctx context.Context, userID string) (*MobileProfile, error) {
    // Get full dashboard
    dashboard, err := m.gateway.GetUserDashboard(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    // Transform to mobile-specific format
    return &MobileProfile{
        UserID:      dashboard.User.ID,
        Name:        dashboard.User.Name,
        Avatar:      m.getOptimizedAvatar(dashboard.User.Avatar),
        OrderCount:  len(dashboard.Orders),
        LastOrder:   m.getLastOrderSummary(dashboard.Orders),
        // Reduced data for mobile
    }, nil
}

type WebBFF struct {
    gateway *APIGateway
}

func (w *WebBFF) GetWebUserProfile(ctx context.Context, userID string) (*WebProfile, error) {
    dashboard, err := w.gateway.GetUserDashboard(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    // Transform to web-specific format with more details
    return &WebProfile{
        User:             dashboard.User,
        Orders:           dashboard.Orders,
        Recommendations:  dashboard.Recommendations,
        Analytics:        w.calculateAnalytics(dashboard),
        // Full data for web
    }, nil
}

// GraphQL Gateway Integration
type GraphQLGateway struct {
    gateway *APIGateway
}

func (g *GraphQLGateway) ResolveUser(ctx context.Context, userID string) (*User, error) {
    // Dataloader pattern for batching
    return g.gateway.userService.GetUser(ctx, userID)
}

func (g *GraphQLGateway) ResolveOrders(ctx context.Context, userID string) ([]Order, error) {
    // Can be called independently or as part of user resolution
    return g.gateway.orderService.GetUserOrders(ctx, userID)
}

// Protocol Translation
type ProtocolAdapter struct {
    gateway *APIGateway
}

func (p *ProtocolAdapter) HandleHTTP(w http.ResponseWriter, r *http.Request) {
    // HTTP to internal protocol
    ctx := r.Context()
    userID := r.URL.Query().Get("user_id")
    
    dashboard, err := p.gateway.GetUserDashboard(ctx, userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(dashboard)
}

func (p *ProtocolAdapter) HandleGRPC(ctx context.Context, req *pb.GetDashboardRequest) (*pb.Dashboard, error) {
    // gRPC to internal protocol
    dashboard, err := p.gateway.GetUserDashboard(ctx, req.UserId)
    if err != nil {
        return nil, err
    }
    
    // Convert to protobuf
    return &pb.Dashboard{
        User:   p.userToProto(dashboard.User),
        Orders: p.ordersToProto(dashboard.Orders),
    }, nil
}

func (p *ProtocolAdapter) HandleWebSocket(conn *websocket.Conn) {
    // WebSocket for real-time updates
    userID := p.authenticateWebSocket(conn)
    
    // Subscribe to events
    eventChan := p.gateway.SubscribeToUserEvents(userID)
    
    for event := range eventChan {
        if err := conn.WriteJSON(event); err != nil {
            break
        }
    }
}
```

---

## Slide 9: Data Integration Patterns

### Integrating Data Across Services

```go
// Data Integration Patterns

// 1. Shared Database Anti-Pattern (Avoid)
type SharedDatabaseAntiPattern struct {
    // Multiple services accessing same database
    // ❌ Tight coupling
    // ❌ No clear ownership
    // ❌ Schema changes affect all services
}

// 2. Database per Service with Data Synchronization
type DataSynchronizationPattern struct {
    sourceDB EventStore
    targetDB ReadModel
    syncer   EventProcessor
}

func (d *DataSynchronizationPattern) Synchronize(ctx context.Context) {
    // Subscribe to change events
    events := d.sourceDB.Subscribe(ctx, "user.updated", "order.created")
    
    for event := range events {
        // Transform and update read model
        switch e := event.(type) {
        case UserUpdatedEvent:
            d.targetDB.UpdateUser(e.UserID, e.Changes)
        case OrderCreatedEvent:
            d.targetDB.AddOrder(e.Order)
        }
        
        // Mark event as processed
        d.sourceDB.Acknowledge(event.ID)
    }
}

// 3. Change Data Capture (CDC) Pattern
type CDCIntegration struct {
    debezium DebeziumConnector
    kafka    KafkaProducer
}

func (c *CDCIntegration) StreamChanges(ctx context.Context) {
    // Capture database changes
    changes := c.debezium.CaptureChanges("postgres.public.orders")
    
    for change := range changes {
        // Publish to event stream
        event := c.transformToEvent(change)
        c.kafka.Publish(ctx, "order-events", event)
    }
}

// 4. CQRS Pattern for Read/Write Separation
type CQRSDataIntegration struct {
    commandDB WriteModel
    queryDB   ReadModel
    projector EventProjector
}

// Command side
func (c *CQRSDataIntegration) CreateOrder(ctx context.Context, cmd CreateOrderCommand) error {
    // Write to command database
    order := &Order{
        ID:     generateID(),
        Items:  cmd.Items,
        Status: "pending",
    }
    
    if err := c.commandDB.Save(ctx, order); err != nil {
        return err
    }
    
    // Publish event for projection
    event := OrderCreatedEvent{
        OrderID:   order.ID,
        Items:     order.Items,
        CreatedAt: time.Now(),
    }
    
    return c.projector.Project(ctx, event)
}

// Query side
func (c *CQRSDataIntegration) GetOrderView(ctx context.Context, orderID string) (*OrderView, error) {
    // Read from optimized read model
    return c.queryDB.GetOrderView(ctx, orderID)
}

// 5. Saga Pattern for Distributed Transactions
type DistributedDataSaga struct {
    coordinator SagaCoordinator
    steps       []SagaStep
}

type SagaStep struct {
    Name         string
    Execute      func(ctx context.Context, data interface{}) error
    Compensate   func(ctx context.Context, data interface{}) error
}

func (s *DistributedDataSaga) ExecuteTransaction(ctx context.Context, data interface{}) error {
    completedSteps := []SagaStep{}
    
    for _, step := range s.steps {
        if err := step.Execute(ctx, data); err != nil {
            // Rollback in reverse order
            for i := len(completedSteps) - 1; i >= 0; i-- {
                if compensateErr := completedSteps[i].Compensate(ctx, data); compensateErr != nil {
                    // Log compensation failure
                    s.coordinator.LogCompensationFailure(completedSteps[i].Name, compensateErr)
                }
            }
            return fmt.Errorf("saga failed at step %s: %w", step.Name, err)
        }
        completedSteps = append(completedSteps, step)
    }
    
    return nil
}

// 6. Data Lake Integration
type DataLakeIntegration struct {
    s3Client     S3Client
    glueClient   GlueClient
    athenaClient AthenaClient
}

func (d *DataLakeIntegration) StoreRawData(ctx context.Context, service string, data []byte) error {
    // Store raw data in S3
    key := fmt.Sprintf("raw/%s/%s/%s.json",
        service,
        time.Now().Format("2006/01/02"),
        uuid.New().String())
    
    return d.s3Client.PutObject(ctx, "data-lake-bucket", key, data)
}

func (d *DataLakeIntegration) QueryAcrossServices(ctx context.Context, sql string) ([]map[string]interface{}, error) {
    // Query using Athena
    queryID, err := d.athenaClient.StartQuery(ctx, sql)
    if err != nil {
        return nil, err
    }
    
    // Wait for completion
    result, err := d.athenaClient.GetQueryResults(ctx, queryID)
    if err != nil {
        return nil, err
    }
    
    return d.parseAthenaResults(result), nil
}

// 7. API Composition for Data Integration
type APIComposer struct {
    services map[string]ServiceClient
    cache    Cache
}

func (a *APIComposer) GetCompositeData(ctx context.Context, entityID string) (*CompositeEntity, error) {
    // Check cache first
    if cached, found := a.cache.Get("composite:" + entityID); found {
        return cached.(*CompositeEntity), nil
    }
    
    // Parallel API calls
    type result struct {
        name string
        data interface{}
        err  error
    }
    
    results := make(chan result, len(a.services))
    
    for name, service := range a.services {
        go func(n string, s ServiceClient) {
            data, err := s.GetData(ctx, entityID)
            results <- result{name: n, data: data, err: err}
        }(name, service)
    }
    
    // Aggregate results
    composite := &CompositeEntity{
        ID:   entityID,
        Data: make(map[string]interface{}),
    }
    
    for i := 0; i < len(a.services); i++ {
        r := <-results
        if r.err != nil {
            // Log error but continue with partial data
            log.Printf("Failed to get data from %s: %v", r.name, r.err)
            continue
        }
        composite.Data[r.name] = r.data
    }
    
    // Cache composite result
    a.cache.Set("composite:"+entityID, composite, 5*time.Minute)
    
    return composite, nil
}
```

---

## Slide 10: Testing Integration Strategies

### Testing Across Patterns and Boundaries

```go
// Integrated Testing Strategies

// 1. Contract Testing Between Services
type ContractTest struct {
    provider string
    consumer string
    pact     PactFramework
}

func (c *ContractTest) TestOrderServiceContract(t *testing.T) {
    // Consumer defines expected interaction
    c.pact.
        AddInteraction().
        Given("Order 123 exists").
        UponReceiving("A request for order 123").
        WithRequest(Request{
            Method: "GET",
            Path:   "/orders/123",
        }).
        WillRespondWith(Response{
            Status: 200,
            Body: Map{
                "id":     "123",
                "status": "shipped",
                "total":  99.99,
            },
        })
    
    // Test consumer code against mock
    err := c.pact.Verify(func() error {
        client := NewOrderServiceClient(c.pact.MockServerURL())
        order, err := client.GetOrder("123")
        
        assert.NoError(t, err)
        assert.Equal(t, "123", order.ID)
        assert.Equal(t, "shipped", order.Status)
        
        return nil
    })
    
    assert.NoError(t, err)
}

// 2. Integration Testing with TestContainers
type IntegrationTestSuite struct {
    postgres     testcontainers.Container
    redis        testcontainers.Container
    kafka        testcontainers.Container
    services     map[string]*Service
}

func (s *IntegrationTestSuite) SetupSuite() {
    ctx := context.Background()
    
    // Start PostgreSQL
    postgresReq := testcontainers.ContainerRequest{
        Image:        "postgres:14",
        ExposedPorts: []string{"5432/tcp"},
        Env: map[string]string{
            "POSTGRES_PASSWORD": "test",
        },
        WaitingFor: wait.ForSQL("5432/tcp", "postgres", 
            func(port nat.Port) string {
                return fmt.Sprintf("postgres://postgres:test@localhost:%s/postgres?sslmode=disable", port.Port())
            }).WithStartupTimeout(time.Second * 5),
    }
    
    s.postgres, _ = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
        ContainerRequest: postgresReq,
        Started:          true,
    })
    
    // Start Redis
    redisReq := testcontainers.ContainerRequest{
        Image:        "redis:7",
        ExposedPorts: []string{"6379/tcp"},
        WaitingFor:   wait.ForListeningPort("6379/tcp"),
    }
    
    s.redis, _ = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
        ContainerRequest: redisReq,
        Started:          true,
    })
    
    // Start Kafka
    kafkaReq := testcontainers.ContainerRequest{
        Image:        "confluentinc/cp-kafka:latest",
        ExposedPorts: []string{"9092/tcp"},
        Env: map[string]string{
            "KAFKA_ZOOKEEPER_CONNECT": "zookeeper:2181",
        },
    }
    
    s.kafka, _ = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
        ContainerRequest: kafkaReq,
        Started:          true,
    })
    
    // Initialize services with test containers
    s.initializeServices()
}

func (s *IntegrationTestSuite) TestEndToEndOrderFlow() {
    ctx := context.Background()
    
    // Create user
    user := &User{
        Name:  "Test User",
        Email: "test@example.com",
    }
    
    createdUser, err := s.services["user"].CreateUser(ctx, user)
    s.NoError(err)
    
    // Create order
    order := &Order{
        UserID: createdUser.ID,
        Items: []Item{
            {ProductID: "PROD-1", Quantity: 2},
        },
    }
    
    createdOrder, err := s.services["order"].CreateOrder(ctx, order)
    s.NoError(err)
    
    // Wait for async processing
    time.Sleep(2 * time.Second)
    
    // Verify inventory was updated
    inventory, err := s.services["inventory"].GetInventory(ctx, "PROD-1")
    s.NoError(err)
    s.Equal(98, inventory.Available) // Started with 100
    
    // Verify notification was sent
    notifications, err := s.services["notification"].GetUserNotifications(ctx, createdUser.ID)
    s.NoError(err)
    s.Len(notifications, 1)
    s.Contains(notifications[0].Message, "Order confirmed")
}

// 3. Component Testing with Mocks
type ComponentTest struct {
    service      *OrderService
    mockRepo     *MockOrderRepository
    mockEventBus *MockEventBus
    mockCache    *MockCache
}

func (c *ComponentTest) TestOrderServiceWithMocks(t *testing.T) {
    // Setup mocks
    c.mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
    c.mockEventBus.On("Publish", mock.Anything, mock.Anything).Return(nil)
    c.mockCache.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)
    
    // Test service
    order := &Order{
        UserID: "user-123",
        Items:  []Item{{ProductID: "prod-1", Quantity: 1}},
    }
    
    err := c.service.CreateOrder(context.Background(), order)
    
    assert.NoError(t, err)
    c.mockRepo.AssertExpectations(t)
    c.mockEventBus.AssertExpectations(t)
    c.mockCache.AssertExpectations(t)
}

// 4. Chaos Testing for Resilience
type ChaosTest struct {
    chaosMonkey ChaosMonkey
    service     Service
}

func (c *ChaosTest) TestServiceResilience(t *testing.T) {
    // Inject failures
    c.chaosMonkey.Configure(ChaosConfig{
        NetworkLatency:     500 * time.Millisecond,
        ErrorRate:          0.1, // 10% error rate
        ServiceFailureRate: 0.05, // 5% complete failure
    })
    
    // Run tests under chaos
    results := make([]bool, 100)
    
    for i := 0; i < 100; i++ {
        ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
        _, err := c.service.ProcessRequest(ctx, &Request{ID: i})
        results[i] = err == nil
        cancel()
    }
    
    // Verify resilience metrics
    successRate := c.calculateSuccessRate(results)
    assert.Greater(t, successRate, 0.85, "Service should maintain >85% success rate under chaos")
}

// 5. Performance Testing Integration
type PerformanceTest struct {
    loadGenerator LoadGenerator
    metrics       MetricsCollector
}

func (p *PerformanceTest) TestSystemPerformance(t *testing.T) {
    // Configure load pattern
    load := LoadPattern{
        RampUpTime:   30 * time.Second,
        SteadyTime:   5 * time.Minute,
        RampDownTime: 30 * time.Second,
        MaxUsers:     1000,
        RequestRate:  100, // requests per second
    }
    
    // Start monitoring
    p.metrics.Start()
    
    // Generate load
    results := p.loadGenerator.Generate(load, func() error {
        // Simulate user journey
        user := p.createUser()
        order := p.createOrder(user)
        p.processPayment(order)
        p.checkOrderStatus(order)
        return nil
    })
    
    // Analyze results
    stats := p.metrics.GetStatistics()
    
    assert.Less(t, stats.P95Latency, 500*time.Millisecond)
    assert.Less(t, stats.P99Latency, 1*time.Second)
    assert.Greater(t, stats.SuccessRate, 0.99)
    assert.Less(t, stats.ErrorRate, 0.01)
}
```

---

## Slide 11: DevOps Integration

### CI/CD Pipeline Integration

```go
// DevOps Integration Patterns

// 1. GitOps Integration
type GitOpsController struct {
    git        GitClient
    kubernetes KubernetesClient
    argo       ArgoCD
}

func (g *GitOpsController) DeployService(service ServiceManifest) error {
    // Update Git repository
    if err := g.git.UpdateManifest(service); err != nil {
        return err
    }
    
    // Trigger ArgoCD sync
    return g.argo.SyncApplication(service.Name)
}

// 2. Blue-Green Deployment
type BlueGreenDeployer struct {
    k8s           KubernetesClient
    loadBalancer  LoadBalancer
    healthChecker HealthChecker
}

func (b *BlueGreenDeployer) Deploy(newVersion string) error {
    // Deploy to green environment
    green := &Deployment{
        Name:    "app-green",
        Version: newVersion,
        Replicas: 3,
    }
    
    if err := b.k8s.CreateDeployment(green); err != nil {
        return err
    }
    
    // Health check green environment
    if err := b.healthChecker.WaitForHealthy(green, 5*time.Minute); err != nil {
        b.k8s.DeleteDeployment(green)
        return err
    }
    
    // Switch traffic
    if err := b.loadBalancer.SwitchToGreen(); err != nil {
        return err
    }
    
    // Clean up blue environment
    time.Sleep(10 * time.Minute) // Grace period
    return b.k8s.DeleteDeployment(&Deployment{Name: "app-blue"})
}

// 3. Canary Deployment
type CanaryDeployer struct {
    k8s     KubernetesClient
    istio   IstioClient
    metrics MetricsClient
}

func (c *CanaryDeployer) DeployCanary(version string, percentage int) error {
    // Deploy canary version
    canary := &Deployment{
        Name:     "app-canary",
        Version:  version,
        Replicas: 1,
    }
    
    if err := c.k8s.CreateDeployment(canary); err != nil {
        return err
    }
    
    // Configure traffic split
    virtualService := &VirtualService{
        Name: "app",
        Routes: []Route{
            {Destination: "app-stable", Weight: 100 - percentage},
            {Destination: "app-canary", Weight: percentage},
        },
    }
    
    if err := c.istio.UpdateVirtualService(virtualService); err != nil {
        return err
    }
    
    // Monitor metrics
    return c.monitorCanary(version, 30*time.Minute)
}

func (c *CanaryDeployer) monitorCanary(version string, duration time.Duration) error {
    ticker := time.NewTicker(1 * time.Minute)
    defer ticker.Stop()
    
    timeout := time.After(duration)
    
    for {
        select {
        case <-ticker.C:
            metrics := c.metrics.GetCanaryMetrics(version)
            
            if metrics.ErrorRate > 0.05 { // 5% error threshold
                return c.rollbackCanary()
            }
            
            if metrics.Latency > 500*time.Millisecond {
                return c.rollbackCanary()
            }
            
        case <-timeout:
            return c.promoteCanary()
        }
    }
}

// 4. Infrastructure as Code Integration
type InfrastructureManager struct {
    terraform TerraformClient
    ansible   AnsibleClient
    vault     VaultClient
}

func (i *InfrastructureManager) ProvisionEnvironment(env Environment) error {
    // Create infrastructure with Terraform
    tfVars := map[string]interface{}{
        "environment": env.Name,
        "region":      env.Region,
        "instance_type": env.InstanceType,
    }
    
    if err := i.terraform.Apply(tfVars); err != nil {
        return err
    }
    
    // Get outputs
    outputs := i.terraform.GetOutputs()
    
    // Configure with Ansible
    inventory := i.createAnsibleInventory(outputs)
    
    // Get secrets from Vault
    secrets, err := i.vault.GetSecrets(env.Name)
    if err != nil {
        return err
    }
    
    // Run Ansible playbook
    return i.ansible.RunPlaybook("configure.yml", inventory, secrets)
}

// 5. Monitoring & Observability Integration
type ObservabilityStack struct {
    prometheus   PrometheusClient
    grafana      GrafanaClient
    elastic      ElasticClient
    jaeger       JaegerClient
}

func (o *ObservabilityStack) SetupServiceMonitoring(service Service) error {
    // Configure Prometheus scraping
    promConfig := &PrometheusConfig{
        JobName: service.Name,
        Targets: service.Endpoints,
        Metrics: []string{
            "http_requests_total",
            "http_request_duration_seconds",
            "error_rate",
        },
    }
    
    if err := o.prometheus.AddScrapeConfig(promConfig); err != nil {
        return err
    }
    
    // Create Grafana dashboard
    dashboard := o.createDashboard(service)
    if err := o.grafana.ImportDashboard(dashboard); err != nil {
        return err
    }
    
    // Configure log aggregation
    logConfig := &LogstashConfig{
        Input: fmt.Sprintf("kubernetes.labels.app=%s", service.Name),
        Filter: []Filter{
            {Type: "grok", Pattern: "%{COMMONAPACHELOG}"},
            {Type: "date", Match: "dd/MMM/yyyy:HH:mm:ss Z"},
        },
        Output: "elasticsearch",
    }
    
    if err := o.elastic.ConfigureLogstash(logConfig); err != nil {
        return err
    }
    
    // Enable distributed tracing
    return o.jaeger.EnableTracing(service.Name)
}

// 6. Security Integration
type SecurityPipeline struct {
    sonarqube   SonarQubeClient
    snyk        SnykClient
    vault       VaultClient
    trivy       TrivyClient
}

func (s *SecurityPipeline) RunSecurityChecks(code CodeRepository) (*SecurityReport, error) {
    report := &SecurityReport{}
    
    // Static code analysis
    sonarResults, err := s.sonarqube.Analyze(code)
    if err != nil {
        return nil, err
    }
    report.CodeVulnerabilities = sonarResults.Issues
    
    // Dependency scanning
    snykResults, err := s.snyk.TestDependencies(code)
    if err != nil {
        return nil, err
    }
    report.DependencyVulnerabilities = snykResults.Vulnerabilities
    
    // Container image scanning
    if code.HasDockerfile() {
        image := s.buildImage(code)
        trivyResults, err := s.trivy.ScanImage(image)
        if err != nil {
            return nil, err
        }
        report.ImageVulnerabilities = trivyResults.Vulnerabilities
    }
    
    // Secret scanning
    secrets := s.scanForSecrets(code)
    if len(secrets) > 0 {
        report.ExposedSecrets = secrets
        
        // Rotate exposed secrets
        for _, secret := range secrets {
            s.vault.RotateSecret(secret.Path)
        }
    }
    
    return report, nil
}
```

---

## Slide 12: Real-World Case Study

### E-Commerce Platform Integration

```go
// Complete E-Commerce Platform Integration Example

// System Architecture Overview
type ECommercePlatform struct {
    // Microservices
    userService      *UserService
    catalogService   *CatalogService
    cartService      *CartService
    orderService     *OrderService
    paymentService   *PaymentService
    inventoryService *InventoryService
    shippingService  *ShippingService
    
    // Infrastructure
    apiGateway       *APIGateway
    eventBus         *EventBus
    cache            *DistributedCache
    
    // Patterns Applied
    sagaOrchestrator *SagaOrchestrator
    cqrsHandler      *CQRSHandler
}

// 1. User Journey Integration
func (e *ECommercePlatform) ProcessUserPurchase(ctx context.Context, userID string, cartID string) error {
    // Clean Architecture: Use Case Layer
    useCase := &PurchaseUseCase{
        platform: e,
    }
    
    return useCase.Execute(ctx, userID, cartID)
}

type PurchaseUseCase struct {
    platform *ECommercePlatform
}

func (p *PurchaseUseCase) Execute(ctx context.Context, userID, cartID string) error {
    // Step 1: Validate user (DDD - Aggregate)
    user, err := p.platform.userService.GetUser(ctx, userID)
    if err != nil {
        return err
    }
    
    if !user.CanMakePurchase() {
        return ErrUserRestricted
    }
    
    // Step 2: Get cart items (CQRS - Read Model)
    cart, err := p.platform.cqrsHandler.Query(ctx, GetCartQuery{CartID: cartID})
    if err != nil {
        return err
    }
    
    // Step 3: Check inventory (Event Sourcing)
    availabilityEvents := p.platform.inventoryService.CheckAvailability(ctx, cart.Items)
    
    for _, event := range availabilityEvents {
        if event.Type == "OutOfStock" {
            return ErrItemOutOfStock
        }
    }
    
    // Step 4: Calculate pricing (Domain Service)
    pricing := &PricingService{
        catalogService: p.platform.catalogService,
    }
    
    total, err := pricing.CalculateTotal(ctx, cart, user)
    if err != nil {
        return err
    }
    
    // Step 5: Process payment (Saga Pattern)
    saga := p.platform.createPurchaseSaga(user, cart, total)
    
    if err := saga.Execute(ctx); err != nil {
        // Saga handles compensations automatically
        return err
    }
    
    // Step 6: Publish success event (Event-Driven)
    event := PurchaseCompletedEvent{
        UserID:    userID,
        OrderID:   saga.OrderID,
        Total:     total,
        Timestamp: time.Now(),
    }
    
    return p.platform.eventBus.Publish(ctx, event)
}

func (e *ECommercePlatform) createPurchaseSaga(user *User, cart *Cart, total Money) *Saga {
    return &Saga{
        Steps: []SagaStep{
            {
                Name: "Reserve Inventory",
                Execute: func(ctx context.Context) error {
                    return e.inventoryService.Reserve(ctx, cart.Items)
                },
                Compensate: func(ctx context.Context) error {
                    return e.inventoryService.Release(ctx, cart.Items)
                },
            },
            {
                Name: "Process Payment",
                Execute: func(ctx context.Context) error {
                    return e.paymentService.Charge(ctx, user.PaymentMethod, total)
                },
                Compensate: func(ctx context.Context) error {
                    return e.paymentService.Refund(ctx, user.PaymentMethod, total)
                },
            },
            {
                Name: "Create Order",
                Execute: func(ctx context.Context) error {
                    order := &Order{
                        UserID: user.ID,
                        Items:  cart.Items,
                        Total:  total,
                        Status: "confirmed",
                    }
                    return e.orderService.Create(ctx, order)
                },
                Compensate: func(ctx context.Context) error {
                    return e.orderService.Cancel(ctx, saga.OrderID)
                },
            },
            {
                Name: "Create Shipment",
                Execute: func(ctx context.Context) error {
                    return e.shippingService.CreateShipment(ctx, saga.OrderID)
                },
                Compensate: func(ctx context.Context) error {
                    return e.shippingService.CancelShipment(ctx, saga.OrderID)
                },
            },
        },
    }
}

// 2. Real-time Updates Integration
type RealtimeIntegration struct {
    websocket *WebSocketServer
    sse       *SSEServer
    eventBus  *EventBus
}

func (r *RealtimeIntegration) StreamOrderUpdates(userID string) {
    // Subscribe to relevant events
    events := r.eventBus.Subscribe(
        fmt.Sprintf("order.*.user.%s", userID),
        fmt.Sprintf("shipment.*.user.%s", userID),
    )
    
    for event := range events {
        // Transform to client-friendly format
        update := r.transformToUpdate(event)
        
        // Send via WebSocket
        r.websocket.Broadcast(userID, update)
        
        // Also send via SSE for fallback
        r.sse.Send(userID, update)
    }
}

// 3. Analytics Integration
type AnalyticsIntegration struct {
    clickhouse  *ClickHouseClient
    redshift    *RedshiftClient
    bigquery    *BigQueryClient
    transformer *DataTransformer
}

func (a *AnalyticsIntegration) ProcessAnalyticsEvent(event AnalyticsEvent) error {
    // Transform event
    transformed := a.transformer.Transform(event)
    
    // Write to multiple analytics stores
    var wg sync.WaitGroup
    errors := make(chan error, 3)
    
    wg.Add(3)
    
    // ClickHouse for real-time analytics
    go func() {
        defer wg.Done()
        errors <- a.clickhouse.Insert(transformed)
    }()
    
    // Redshift for business intelligence
    go func() {
        defer wg.Done()
        errors <- a.redshift.Load(transformed)
    }()
    
    // BigQuery for machine learning
    go func() {
        defer wg.Done()
        errors <- a.bigquery.Stream(transformed)
    }()
    
    wg.Wait()
    close(errors)
    
    // Check for errors
    for err := range errors {
        if err != nil {
            return err
        }
    }
    
    return nil
}

// 4. Search Integration
type SearchIntegration struct {
    elastic     *ElasticsearchClient
    algolia     *AlgoliaClient
    indexer     *SearchIndexer
}

func (s *SearchIntegration) IndexProduct(product Product) error {
    // Create search document
    doc := s.indexer.CreateDocument(product)
    
    // Index in Elasticsearch for internal search
    if err := s.elastic.Index("products", doc); err != nil {
        return err
    }
    
    // Index in Algolia for customer-facing search
    return s.algolia.SaveObject("products", doc)
}

func (s *SearchIntegration) Search(query SearchQuery) (*SearchResults, error) {
    // Determine search backend based on context
    if query.IsInternal {
        return s.elastic.Search(query)
    }
    
    // Customer search with facets and filters
    return s.algolia.Search(query)
}

// 5. Machine Learning Integration
type MLIntegration struct {
    tensorflow  *TensorFlowServing
    sagemaker   *SageMakerClient
    mlflow      *MLFlowClient
}

func (m *MLIntegration) GetRecommendations(userID string) ([]Product, error) {
    // Get user features
    features := m.getUserFeatures(userID)
    
    // Call ML model
    predictions, err := m.tensorflow.Predict("recommendation-model", features)
    if err != nil {
        // Fallback to rule-based recommendations
        return m.getRuleBasedRecommendations(userID)
    }
    
    // Track model performance
    m.mlflow.LogPrediction("recommendation-model", userID, predictions)
    
    return m.convertToProducts(predictions)
}

// 6. Performance Metrics
func (e *ECommercePlatform) GetSystemMetrics() *SystemMetrics {
    return &SystemMetrics{
        Architecture: "Microservices + Event-Driven + DDD",
        Services:     12,
        Metrics: PerformanceMetrics{
            RequestsPerSecond:  15000,
            AverageLatency:     45 * time.Millisecond,
            P99Latency:         200 * time.Millisecond,
            ErrorRate:          0.001, // 0.1%
            Availability:       0.9999, // 99.99%
        },
        Scale: ScaleMetrics{
            DailyActiveUsers:   1000000,
            PeakConcurrent:     50000,
            TransactionsPerDay: 5000000,
            DataProcessed:      "10TB/day",
        },
        Integration: IntegrationMetrics{
            EventsProcessedPerDay:  50000000,
            APICallsPerDay:        100000000,
            CacheHitRate:          0.95, // 95%
            DatabaseConnections:   500,
        },
    }
}
```

---

## Slide 13: Integration Anti-Patterns

### Common Pitfalls to Avoid

```go
// Anti-Patterns in Integration

// 1. Distributed Monolith Anti-Pattern
type DistributedMonolith struct {
    // ❌ Microservices that are tightly coupled
    // ❌ Synchronous calls everywhere
    // ❌ Shared database
    // ❌ Coordinated deployments required
}

// Problem:
func (d *DistributedMonolith) GetUserProfile(userID string) (*UserProfile, error) {
    // Synchronous cascade of calls
    user := d.callUserService(userID)           // Blocked waiting
    orders := d.callOrderService(userID)        // Blocked waiting
    payments := d.callPaymentService(userID)    // Blocked waiting
    preferences := d.callPrefService(userID)    // Blocked waiting
    
    // If any service is down, entire request fails
    // Latency adds up: 50ms + 50ms + 50ms + 50ms = 200ms
    
    return &UserProfile{user, orders, payments, preferences}, nil
}

// Solution: Async communication, data duplication, eventual consistency
type ProperMicroservices struct {
    cache    Cache
    eventBus EventBus
}

func (p *ProperMicroservices) GetUserProfile(userID string) (*UserProfile, error) {
    // Use cached/replicated data
    profile := p.cache.Get("profile:" + userID)
    
    if profile == nil {
        // Async event to rebuild cache
        p.eventBus.Publish(ProfileRebuildNeeded{UserID: userID})
        return p.getPartialProfile(userID), nil
    }
    
    return profile, nil
}

// 2. Chatty Service Anti-Pattern
type ChattyService struct {
    // ❌ Too many small requests
    // ❌ N+1 query problem across services
}

// Problem:
func (c *ChattyService) GetOrdersWithDetails(userID string) ([]OrderDetails, error) {
    orders := c.getOrders(userID) // 1 call
    
    details := make([]OrderDetails, len(orders))
    for i, order := range orders {
        // N additional calls!
        details[i].Products = c.getProducts(order.ProductIDs)      // N calls
        details[i].Shipping = c.getShipping(order.ShippingID)     // N calls
        details[i].Payment = c.getPayment(order.PaymentID)        // N calls
    }
    
    return details, nil
}

// Solution: Batch APIs, GraphQL, or Data duplication
type EfficientService struct {
    batchAPI BatchAPI
}

func (e *EfficientService) GetOrdersWithDetails(userID string) ([]OrderDetails, error) {
    orders := e.getOrders(userID) // 1 call
    
    // Collect all IDs
    var productIDs, shippingIDs, paymentIDs []string
    for _, order := range orders {
        productIDs = append(productIDs, order.ProductIDs...)
        shippingIDs = append(shippingIDs, order.ShippingID)
        paymentIDs = append(paymentIDs, order.PaymentID)
    }
    
    // Batch calls
    products := e.batchAPI.GetProducts(productIDs)     // 1 call
    shippings := e.batchAPI.GetShippings(shippingIDs)  // 1 call
    payments := e.batchAPI.GetPayments(paymentIDs)     // 1 call
    
    // Total: 4 calls instead of 1 + 3N
    
    return e.assembleDetails(orders, products, shippings, payments), nil
}

// 3. God Service Anti-Pattern
type GodService struct {
    // ❌ One service does everything
    // ❌ Single point of failure
    // ❌ Impossible to scale independently
}

// Problem:
type EverythingService struct{}

func (e *EverythingService) DoEverything(req Request) error {
    e.validateUser()
    e.checkInventory()
    e.processPayment()
    e.sendEmail()
    e.updateAnalytics()
    e.generateReport()
    e.notifySlack()
    // ... and 50 more responsibilities
    return nil
}

// Solution: Proper service boundaries based on business capabilities
type WellDesignedServices struct {
    userService      UserService
    inventoryService InventoryService
    paymentService   PaymentService
    notificationSvc  NotificationService
    analyticsService AnalyticsService
}

// 4. Shared Database Anti-Pattern
type SharedDatabase struct {
    // ❌ Multiple services access same database
    // ❌ Schema coupling
    // ❌ Can't use different database types
}

// Problem:
type ServiceA struct {
    db Database // shared DB
}

func (s *ServiceA) UpdateUser(user User) error {
    // Direct database access
    return s.db.Exec("UPDATE users SET name = ? WHERE id = ?", user.Name, user.ID)
}

type ServiceB struct {
    db Database // same shared DB
}

func (s *ServiceB) GetUser(id string) (*User, error) {
    // Depends on ServiceA's schema
    return s.db.Query("SELECT * FROM users WHERE id = ?", id)
}

// Solution: Database per service + Events
type ServiceAFixed struct {
    db       Database // own database
    eventBus EventBus
}

func (s *ServiceAFixed) UpdateUser(user User) error {
    // Update own database
    if err := s.db.Exec("UPDATE users SET name = ? WHERE id = ?", user.Name, user.ID); err != nil {
        return err
    }
    
    // Publish event for others
    return s.eventBus.Publish(UserUpdatedEvent{
        UserID: user.ID,
        Name:   user.Name,
    })
}

type ServiceBFixed struct {
    db       Database // own database with replicated data
    eventBus EventBus
}

func (s *ServiceBFixed) Start() {
    // Subscribe to events
    s.eventBus.Subscribe("user.updated", s.handleUserUpdated)
}

func (s *ServiceBFixed) handleUserUpdated(event UserUpdatedEvent) error {
    // Update local replica
    return s.db.Exec("UPDATE user_cache SET name = ? WHERE id = ?", 
        event.Name, event.UserID)
}

// 5. Point-to-Point Integration Spaghetti
type PointToPointMess struct {
    // ❌ Every service calls every other service directly
    // ❌ N*(N-1) integration points
}

// Problem:
/*
Service A → Service B
Service A → Service C
Service A → Service D
Service B → Service C
Service B → Service D
Service C → Service D
... = N*(N-1)/2 connections!
*/

// Solution: Event Bus or API Gateway
type EventDrivenIntegration struct {
    eventBus EventBus
}

func (e *EventDrivenIntegration) ServiceA() {
    // Publish event - don't know who listens
    e.eventBus.Publish(OrderCreatedEvent{})
}

func (e *EventDrivenIntegration) ServiceB() {
    // Subscribe to relevant events
    e.eventBus.Subscribe("order.created", e.handleOrderCreated)
}

// Now: N services = N connections to event bus

// 6. Anemic Domain Model Anti-Pattern
type AnemicDomainModel struct {
    // ❌ Domain objects with no behavior
    // ❌ All logic in services
}

// Problem:
type Order struct {
    ID     string
    Items  []Item
    Total  float64
    Status string
    // No methods - just data
}

type OrderService struct {
    // All logic here instead of domain model
}

func (s *OrderService) CalculateTotal(order *Order) float64 {
    var total float64
    for _, item := range order.Items {
        total += item.Price * float64(item.Quantity)
    }
    return total
}

func (s *OrderService) CanCancel(order *Order) bool {
    return order.Status == "pending" || order.Status == "confirmed"
}

// Solution: Rich Domain Model
type OrderFixed struct {
    id     string
    items  []Item
    total  Money
    status OrderStatus
}

// Business logic in domain model
func (o *OrderFixed) CalculateTotal() Money {
    var total decimal.Decimal
    for _, item := range o.items {
        total = total.Add(item.Price.Multiply(item.Quantity))
    }
    return Money{Amount: total, Currency: o.items[0].Price.Currency}
}

func (o *OrderFixed) CanCancel() bool {
    return o.status == StatusPending || o.status == StatusConfirmed
}

func (o *OrderFixed) Cancel() error {
    if !o.CanCancel() {
        return ErrCannotCancel
    }
    o.status = StatusCancelled
    o.addEvent(OrderCancelledEvent{OrderID: o.id})
    return nil
}
```

---

## Slide 14: Best Practices & Guidelines

### Integration Excellence

```go
// Integration Best Practices

// 1. Design for Failure
type ResilientService struct {
    circuitBreaker CircuitBreaker
    retry          RetryPolicy
    timeout        TimeoutPolicy
    fallback       FallbackHandler
}

func (r *ResilientService) CallExternalService(ctx context.Context, req Request) (*Response, error) {
    // Apply timeout
    ctx, cancel := context.WithTimeout(ctx, r.timeout.Duration)
    defer cancel()
    
    // Circuit breaker
    result, err := r.circuitBreaker.Execute(func() (interface{}, error) {
        // Retry policy
        return r.retry.Do(func() (interface{}, error) {
            return r.doActualCall(ctx, req)
        })
    })
    
    if err != nil {
        // Fallback
        return r.fallback.Handle(ctx, req, err)
    }
    
    return result.(*Response), nil
}

// 2. Observability First
type ObservableService struct {
    logger  Logger
    metrics MetricsCollector
    tracer  Tracer
}

func (o *ObservableService) ProcessRequest(ctx context.Context, req Request) error {
    // Start trace
    span, ctx := o.tracer.Start(ctx, "process-request")
    defer span.End()
    
    // Log with context
    o.logger.Info("Processing request",
        "request_id", req.ID,
        "user_id", req.UserID,
        "trace_id", span.TraceID(),
    )
    
    // Metrics
    timer := o.metrics.StartTimer("request_duration")
    defer timer.Observe()
    
    o.metrics.Inc("requests_total", map[string]string{
        "method": req.Method,
        "path":   req.Path,
    })
    
    // Business logic
    err := o.doProcess(ctx, req)
    
    if err != nil {
        o.metrics.Inc("request_errors_total")
        o.logger.Error("Request failed", "error", err)
        span.RecordError(err)
        return err
    }
    
    o.metrics.Inc("requests_success_total")
    return nil
}

// 3. API Versioning Strategy
type VersionedAPI struct {
    v1 *APIv1
    v2 *APIv2
}

func (v *VersionedAPI) Handle(w http.ResponseWriter, r *http.Request) {
    // Version from header
    version := r.Header.Get("API-Version")
    
    switch version {
    case "2024-11-01", "v2":
        v.v2.Handle(w, r)
    case "2023-06-01", "v1", "":
        v.v1.Handle(w, r)
    default:
        http.Error(w, "Unsupported API version", http.StatusBadRequest)
    }
}

// 4. Backwards Compatibility
type BackwardsCompatibleAPI struct{}

// Evolve schema without breaking clients
type UserV1 struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

type UserV2 struct {
    // New fields with defaults
    ID        string `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
    
    // Deprecated but still supported
    Name string `json:"name,omitempty"` // Computed from FirstName + LastName
}

func (u UserV2) MarshalJSON() ([]byte, error) {
    // Ensure backwards compatibility
    type Alias UserV2
    
    // Compute deprecated field
    if u.Name == "" {
        u.Name = u.FirstName + " " + u.LastName
    }
    
    return json.Marshal(Alias(u))
}

// 5. Idempotency
type IdempotentHandler struct {
    cache Cache
}

func (i *IdempotentHandler) HandleRequest(ctx context.Context, req Request) (*Response, error) {
    // Check idempotency key
    if req.IdempotencyKey != "" {
        // Check if already processed
        if cached, found := i.cache.Get("idem:" + req.IdempotencyKey); found {
            return cached.(*Response), nil
        }
    }
    
    // Process request
    resp, err := i.process(ctx, req)
    if err != nil {
        return nil, err
    }
    
    // Cache result
    if req.IdempotencyKey != "" {
        i.cache.Set("idem:"+req.IdempotencyKey, resp, 24*time.Hour)
    }
    
    return resp, nil
}

// 6. Graceful Degradation
type GracefulService struct {
    primary   ServiceClient
    secondary ServiceClient
    cache     Cache
}

func (g *GracefulService) GetData(ctx context.Context, id string) (*Data, error) {
    // Try primary source
    data, err := g.primary.Get(ctx, id)
    if err == nil {
        // Cache for fallback
        g.cache.Set("data:"+id, data, 10*time.Minute)
        return data, nil
    }
    
    // Try secondary source
    data, err = g.secondary.Get(ctx, id)
    if err == nil {
        return data, nil
    }
    
    // Fall back to cache
    if cached, found := g.cache.Get("data:" + id); found {
        log.Warn("Using stale cached data", "id", id)
        return cached.(*Data), nil
    }
    
    // Graceful degradation - return partial data
    return &Data{ID: id, Partial: true}, nil
}

// 7. Rate Limiting & Backpressure
type RateLimitedService struct {
    limiter *rate.Limiter
    queue   chan Request
    workers int
}

func (r *RateLimitedService) HandleRequest(ctx context.Context, req Request) error {
    // Check rate limit
    if !r.limiter.Allow() {
        return ErrRateLimitExceeded
    }
    
    // Apply backpressure
    select {
    case r.queue <- req:
        // Request queued
        return nil
    case <-time.After(1 * time.Second):
        // Queue full
        return ErrSystemOverloaded
    case <-ctx.Done():
        return ctx.Err()
    }
}

// 8. Contract Testing
func TestServiceContract(t *testing.T) {
    // Define contract
    contract := Contract{
        Provider: "order-service",
        Consumer: "notification-service",
        Interactions: []Interaction{
            {
                Description: "Get order by ID",
                Request: Request{
                    Method: "GET",
                    Path:   "/orders/123",
                },
                Response: Response{
                    Status: 200,
                    Body: map[string]interface{}{
                        "id":     "123",
                        "status": "shipped",
                    },
                },
            },
        },
    }
    
    // Verify both sides
    contract.VerifyProvider(t)
    contract.VerifyConsumer(t)
}

// 9. Documentation as Code
// OpenAPI/Swagger specification
type APIDocumentation struct {
    OpenAPI string `json:"openapi"`
    Info    struct {
        Title   string `json:"title"`
        Version string `json:"version"`
    } `json:"info"`
    Paths map[string]PathItem `json:"paths"`
}

// Generate from code
func GenerateOpenAPISpec() *APIDocumentation {
    spec := &APIDocumentation{
        OpenAPI: "3.0.0",
    }
    spec.Info.Title = "E-Commerce API"
    spec.Info.Version = "1.0.0"
    
    // Automatically generated from handler annotations
    spec.Paths = generatePathsFromHandlers()
    
    return spec
}

// 10. Security Best Practices
type SecureIntegration struct {
    authZ     Authorizer
    authN     Authenticator
    encryptor Encryptor
    auditor   AuditLogger
}

func (s *SecureIntegration) HandleRequest(ctx context.Context, req Request) error {
    // Authentication
    user, err := s.authN.Authenticate(ctx, req.Token)
    if err != nil {
        return ErrUnauthorized
    }
    
    // Authorization
    if !s.authZ.Can(user, req.Action, req.Resource) {
        s.auditor.LogUnauthorizedAttempt(user, req)
        return ErrForbidden
    }
    
    // Encrypt sensitive data
    if req.ContainsSensitiveData() {
        req.Data = s.encryptor.Encrypt(req.Data)
    }
    
    // Process
    err = s.process(ctx, req)
    
    // Audit
    s.auditor.Log(user, req.Action, req.Resource, err)
    
    return err
}
```

---

## Slide 15: Integration Patterns Summary

### Quick Reference Guide

```
┌─────────────────────────────────────────────────────────────────┐
│                   Integration Pattern Matrix                    │
├──────────────────┬──────────────────┬──────────────────────────┤
│ Pattern          │ Use When         │ Combines With            │
├──────────────────┼──────────────────┼──────────────────────────┤
│ Clean + DDD      │ Complex domains  │ Hexagonal, Event         │
│                  │ Clear boundaries │ Sourcing                 │
├──────────────────┼──────────────────┼──────────────────────────┤
│ Microservices +  │ Distributed      │ CQRS, Saga, API          │
│ Event-Driven     │ systems          │ Gateway                  │
├──────────────────┼──────────────────┼──────────────────────────┤
│ CQRS + Event     │ Read/write       │ Event Sourcing,          │
│ Sourcing         │ separation       │ Microservices            │
├──────────────────┼──────────────────┼──────────────────────────┤
│ API Gateway +    │ Multiple clients │ BFF, Circuit Breaker,    │
│ BFF              │ Different needs  │ Rate Limiting            │
├──────────────────┼──────────────────┼──────────────────────────┤
│ Saga Pattern     │ Distributed      │ Event-Driven,            │
│                  │ transactions     │ Compensation             │
├──────────────────┼──────────────────┼──────────────────────────┤
│ SOLID + GoF      │ All scenarios    │ Clean Architecture,      │
│                  │ Foundation       │ DDD patterns             │
└──────────────────┴──────────────────┴──────────────────────────┘
```

**Key Integration Principles:**

1. **Loose Coupling** - Services should be independent
2. **High Cohesion** - Related functionality together
3. **Async Communication** - Prefer events over direct calls
4. **Failure Isolation** - One failure shouldn't cascade
5. **Observable** - Monitor everything
6. **Scalable** - Scale components independently
7. **Testable** - Contract testing, integration tests
8. **Secure** - Defense in depth
9. **Documented** - Clear contracts and interfaces
10. **Evolvable** - Support versioning and migration

---

## Slide 16: Conclusion & Q&A

### Key Takeaways

**Integration is About Balance:**

```
Complexity ←→ Simplicity
Flexibility ←→ Constraints
Autonomy ←→ Governance
Consistency ←→ Availability
Performance ←→ Reliability
```

**Essential Integration Skills:**

1. **Pattern Recognition**
   - Identify when to apply which pattern
   - Understand trade-offs
   - Combine patterns effectively

2. **System Thinking**
   - See the big picture
   - Understand dependencies
   - Plan for evolution

3. **Pragmatism**
   - Start simple
   - Add complexity when needed
   - Refactor continuously

**Success Metrics:**

```go
type SuccessfulIntegration struct {
    // Technical
    LowCoupling         bool
    HighCohesion        bool
    GoodPerformance     bool
    HighAvailability    bool
    
    // Organizational
    TeamAutonomy        bool
    FastDeployments     bool
    EasyOnboarding      bool
    
    // Business
    FastTimeToMarket    bool
    FlexibleArchitecture bool
    LowMaintenanceCost   bool
}
```

**Final Wisdom:**

> *"The best architecture is the one that allows you to defer decisions for as long as possible, while still delivering value."*
> 
> *"Integration is not about perfection—it's about making thoughtful trade-offs that serve your specific context."*
> 
> *"Patterns are guidelines, not laws. Understand the principles, then adapt to your needs."*

---

### Questions to Consider

1. How do your current services integrate?
2. What integration anti-patterns exist in your system?
3. Which patterns would benefit your architecture?
4. How can you improve observability?
5. What's your migration strategy?

---

### Resources

**Books:**
- "Building Microservices" by Sam Newman
- "Domain-Driven Design" by Eric Evans
- "Clean Architecture" by Robert C. Martin
- "Enterprise Integration Patterns" by Gregor Hohpe

**Online:**
- microservices.io - Microservice patterns
- martinfowler.com - Architecture articles
- AWS Well-Architected Framework
- Google Cloud Architecture Center

**Tools:**
- Kubernetes - Container orchestration
- Kafka - Event streaming
- Istio - Service mesh
- Prometheus/Grafana - Monitoring

---

## Thank You!

**Contact:**
- **Presenter:** Dong Tran
- **Email:** dong.tran@example.com
- **GitHub:** github.com/dong-tran
- **LinkedIn:** linkedin.com/in/dong-tran

**Questions?**

*Let's discuss your integration challenges!*

---

**Presentation Complete** ✅

Total Slides: 16  
Duration: 90 minutes  
Difficulty: Advanced  
Topics Covered: Architecture Patterns, Design Principles, Integration Strategies, Real-World Applications

            