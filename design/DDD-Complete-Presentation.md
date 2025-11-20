# Domain-Driven Design (DDD)
## Building Software That Reflects Business Reality

**Presenter:** Dong Tran  
**Date:** November 19, 2025  
**Time:** 09:18 UTC

---

## Slide 1: Title Slide

# Domain-Driven Design (DDD)
### Building Software That Reflects Business Reality

Presenter: Dong Tran  
Date: November 19, 2025

---

## Slide 2: Agenda

### Agenda

1. What is DDD and Why It Matters
2. Strategic Design: The Big Picture
3. Tactical Patterns: Building Blocks
4. Implementing DDD in Go
5. Real-World Use Case: E-Commerce Platform
6. Best Practices & Common Pitfalls
7. When to Use DDD
8. Q&A Session

---

## Slide 3: The Problem DDD Solves

### The Problem DDD Solves

**Common Software Development Challenges:**
- ❌ Technical jargon vs Business language
- ❌ Complex business logic scattered everywhere
- ❌ Database-driven design
- ❌ Anemic domain models
- ❌ Lost business knowledge in code

**DDD Solution:**
✅ Align software with business needs through collaboration

---

## Slide 4: What is Domain-Driven Design?

### What is Domain-Driven Design?

**Definition:**
An approach to software development that:
- Centers on the core business domain
- Bases complex designs on domain models
- Initiates creative collaboration between technical and domain experts

**Created by Eric Evans (2003)**
"Blue Book" - Domain-Driven Design: Tackling Complexity in the Heart of Software

---

## Slide 5: Core Concepts

### Core Concepts

**Three Pillars of DDD:**

1. **Ubiquitous Language**
   - Shared vocabulary between developers and domain experts
   
2. **Bounded Context**
   - Clear boundaries where a model applies
   
3. **Model-Driven Design**
   - Code reflects domain model directly

---

## Slide 6: Strategic Design - Bounded Contexts

### Strategic Design - Bounded Contexts

**Example: E-Commerce System**

```
┌─────────────────┐  ┌─────────────────┐
│   Inventory     │  │    Shopping     │
│    Context      │  │     Context     │
├─────────────────┤  ├─────────────────┤
│ - Product       │  │ - Cart          │
│ - Stock         │  │ - Product View  │
│ - Warehouse     │  │ - Customer      │
└─────────────────┘  └─────────────────┘
         ↓                    ↓
┌─────────────────┐  ┌─────────────────┐
│    Billing      │  │   Shipping      │
│    Context      │  │    Context      │
├─────────────────┤  ├─────────────────┤
│ - Invoice       │  │ - Delivery      │
│ - Payment       │  │ - Address       │
│ - Tax           │  │ - Tracking      │
└─────────────────┘  └─────────────────┘
```

---

## Slide 7: Context Mapping Patterns

### Context Mapping Patterns

**How Bounded Contexts Relate:**

- **Shared Kernel**: Shared subset of domain model
- **Customer/Supplier**: Upstream/downstream relationship
- **Conformist**: Downstream conforms to upstream
- **Anti-Corruption Layer**: Translation layer between contexts
- **Open Host Service**: Well-defined protocol for access
- **Published Language**: Well-documented shared language

---

## Slide 8: Tactical Patterns - Building Blocks

### Tactical Patterns - Building Blocks

**Domain Model Components:**

1. **Entity**: Has identity (User, Order)
2. **Value Object**: No identity (Money, Address)
3. **Aggregate**: Consistency boundary
4. **Domain Service**: Business operations
5. **Repository**: Persistence abstraction
6. **Factory**: Complex creation logic
7. **Domain Event**: Something that happened

---

## Slide 9: Entity vs Value Object in Go

### Entity vs Value Object in Go

**Entity Example:**
```go
type UserID uuid.UUID

type User struct {
    ID        UserID
    Email     Email
    Profile   UserProfile
    CreatedAt time.Time
}

func (u *User) Equals(other *User) bool {
    return u.ID == other.ID // Identity comparison
}
```

**Value Object Example:**
```go
type Money struct {
    Amount   decimal.Decimal
    Currency string
}

func (m Money) Equals(other Money) bool {
    return m.Amount.Equal(other.Amount) && 
           m.Currency == other.Currency // Value comparison
}
```

---

## Slide 10: Aggregate Pattern in Go

### Aggregate Pattern in Go

**Order Aggregate Example:**
```go
// Aggregate Root
type Order struct {
    ID         OrderID
    CustomerID CustomerID
    Items      []OrderItem
    Total      Money
    Status     OrderStatus
    CreatedAt  time.Time
}

// Aggregate member
type OrderItem struct {
    ProductID ProductID
    Quantity  int
    Price     Money
}

// Business invariants protected
func (o *Order) AddItem(product Product, quantity int) error {
    if o.Status != OrderStatusDraft {
        return ErrOrderNotEditable
    }
    
    if quantity <= 0 {
        return ErrInvalidQuantity
    }
    
    o.Items = append(o.Items, OrderItem{
        ProductID: product.ID,
        Quantity:  quantity,
        Price:     product.Price,
    })
    
    o.recalculateTotal()
    return nil
}
```

---

## Slide 11: Repository Pattern in Go

### Repository Pattern in Go

**Interface Definition:**
```go
// Domain layer - defines the contract
type OrderRepository interface {
    Save(ctx context.Context, order *Order) error
    FindByID(ctx context.Context, id OrderID) (*Order, error)
    FindByCustomer(ctx context.Context, customerID CustomerID) ([]*Order, error)
}

// Infrastructure layer - implements the contract
type PostgresOrderRepository struct {
    db *sql.DB
}

func (r *PostgresOrderRepository) Save(ctx context.Context, order *Order) error {
    query := `
        INSERT INTO orders (id, customer_id, total, status, created_at)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (id) DO UPDATE SET
            total = $3, status = $4
    `
    _, err := r.db.ExecContext(ctx, query,
        order.ID, order.CustomerID, order.Total.Amount, 
        order.Status, order.CreatedAt)
    return err
}
```

---

## Slide 12: Domain Service in Go

### Domain Service in Go

**When Business Logic Doesn't Fit in Entity:**
```go
type PricingService struct {
    discountRepo DiscountRepository
    taxCalc      TaxCalculator
}

func (s *PricingService) CalculateFinalPrice(
    order *Order,
    customer *Customer,
) (Money, error) {
    
    subtotal := order.CalculateSubtotal()
    
    // Apply customer discount
    discount := s.discountRepo.GetCustomerDiscount(customer.Tier)
    afterDiscount := subtotal.ApplyPercentage(discount)
    
    // Calculate tax
    tax := s.taxCalc.Calculate(afterDiscount, customer.Address)
    
    return afterDiscount.Add(tax), nil
}
```

---

## Slide 13: Domain Events in Go

### Domain Events in Go

**Capturing What Happened:**
```go
// Domain Event
type OrderPlaced struct {
    OrderID    OrderID
    CustomerID CustomerID
    Total      Money
    OccurredAt time.Time
}

// Event Publisher
type EventPublisher interface {
    Publish(ctx context.Context, events ...DomainEvent) error
}

// Aggregate with events
type Order struct {
    // ... fields
    events []DomainEvent
}

func (o *Order) PlaceOrder() error {
    if err := o.validate(); err != nil {
        return err
    }
    
    o.Status = OrderStatusPlaced
    
    // Record domain event
    o.events = append(o.events, OrderPlaced{
        OrderID:    o.ID,
        CustomerID: o.CustomerID,
        Total:      o.Total,
        OccurredAt: time.Now(),
    })
    
    return nil
}
```

---

## Slide 14: Project Structure in Go

### Project Structure in Go

**Clean Architecture with DDD:**
```
project/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── domain/
│   │   ├── order/
│   │   │   ├── order.go         # Aggregate
│   │   │   ├── repository.go    # Repository interface
│   │   │   ├── service.go       # Domain service
│   │   │   └── events.go        # Domain events
│   │   ├── customer/
│   │   ├── product/
│   │   └── shared/
│   │       ├── money.go          # Shared value objects
│   │       └── address.go
│   ├── application/
│   │   ├── command/              # Use cases (write)
│   │   │   └── place_order.go
│   │   └── query/                # Use cases (read)
│   │       └── get_order.go
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   └── messaging/
│   └── interfaces/
│       ├── http/
│       └── grpc/
└── go.mod
```

---

## Slide 15: Real Use Case - Place Order

### Real Use Case - Place Order

**Complete Flow Example:**
```go
// Application Service (Use Case)
type PlaceOrderCommand struct {
    CustomerID CustomerID
    Items      []OrderItemRequest
}

type PlaceOrderHandler struct {
    orderRepo    OrderRepository
    productRepo  ProductRepository
    customerRepo CustomerRepository
    eventBus     EventPublisher
}

func (h *PlaceOrderHandler) Handle(
    ctx context.Context, 
    cmd PlaceOrderCommand,
) error {
    // 1. Load customer
    customer, err := h.customerRepo.FindByID(ctx, cmd.CustomerID)
    if err != nil {
        return err
    }
    
    // 2. Create order aggregate
    order := NewOrder(cmd.CustomerID)
    
    // 3. Add items (domain logic)
    for _, item := range cmd.Items {
        product, err := h.productRepo.FindByID(ctx, item.ProductID)
        if err != nil {
            return err
        }
        
        if err := order.AddItem(product, item.Quantity); err != nil {
            return err
        }
    }
    
    // 4. Place order (more domain logic)
    if err := order.PlaceOrder(); err != nil {
        return err
    }
    
    // 5. Persist
    if err := h.orderRepo.Save(ctx, order); err != nil {
        return err
    }
    
    // 6. Publish events
    return h.eventBus.Publish(ctx, order.GetEvents()...)
}
```

---

## Slide 16: Specification Pattern

### Specification Pattern

**Complex Business Rules:**
```go
// Specification interface
type Specification interface {
    IsSatisfiedBy(candidate interface{}) bool
    And(other Specification) Specification
    Or(other Specification) Specification
    Not() Specification
}

// Concrete specification
type PremiumCustomerSpec struct{}

func (s PremiumCustomerSpec) IsSatisfiedBy(candidate interface{}) bool {
    customer, ok := candidate.(*Customer)
    if !ok {
        return false
    }
    return customer.Tier == TierPremium
}

// Usage
type FreeShippingSpec struct {
    minAmount Money
}

func (s FreeShippingSpec) IsSatisfiedBy(candidate interface{}) bool {
    order, ok := candidate.(*Order)
    if !ok {
        return false
    }
    return order.Total.GreaterThan(s.minAmount)
}

// Combine specifications
premiumOrLargeOrder := PremiumCustomerSpec{}.
    Or(FreeShippingSpec{minAmount: NewMoney(100, "USD")})
```

---

## Slide 17: CQRS with DDD

### CQRS with DDD

**Command Query Responsibility Segregation:**

```go
// Write Model (Command Side)
type Order struct {
    ID     OrderID
    Items  []OrderItem
    Status OrderStatus
    // Rich domain model with behavior
}

// Read Model (Query Side)
type OrderView struct {
    ID           string    `json:"id"`
    CustomerName string    `json:"customer_name"`
    TotalAmount  float64   `json:"total_amount"`
    StatusText   string    `json:"status"`
    OrderedAt    time.Time `json:"ordered_at"`
    // Optimized for display
}

// Query Handler
type GetOrdersHandler struct {
    readDB *sql.DB
}

func (h *GetOrdersHandler) Handle(
    ctx context.Context, 
    customerID string,
) ([]OrderView, error) {
    query := `
        SELECT o.id, c.name, o.total, o.status, o.created_at
        FROM order_views o
        JOIN customers c ON o.customer_id = c.id
        WHERE o.customer_id = $1
        ORDER BY o.created_at DESC
    `
    // Direct query to read-optimized view
    rows, err := h.readDB.QueryContext(ctx, query, customerID)
    // ...
}
```

---

## Slide 18: Testing DDD Code

### Testing DDD Code

**Unit Testing Domain Logic:**
```go
func TestOrder_AddItem(t *testing.T) {
    // Given
    order := NewOrder(CustomerID("cust-123"))
    product := Product{
        ID:    ProductID("prod-456"),
        Name:  "Laptop",
        Price: NewMoney(1000, "USD"),
    }
    
    // When
    err := order.AddItem(product, 2)
    
    // Then
    assert.NoError(t, err)
    assert.Len(t, order.Items, 1)
    assert.Equal(t, 2, order.Items[0].Quantity)
    assert.Equal(t, NewMoney(2000, "USD"), order.Total)
}

func TestOrder_CannotModifyAfterPlaced(t *testing.T) {
    // Given
    order := &Order{
        Status: OrderStatusPlaced,
    }
    
    // When
    err := order.AddItem(Product{}, 1)
    
    // Then
    assert.Error(t, err)
    assert.Equal(t, ErrOrderNotEditable, err)
}
```

---

## Slide 19: Best Practices

### Best Practices

**DO's:**
✅ Start with Ubiquitous Language
✅ Focus on Core Domain
✅ Keep Aggregates small
✅ Use Value Objects liberally
✅ Make implicit concepts explicit
✅ Test domain logic thoroughly
✅ Enforce invariants in aggregates

**DON'Ts:**
❌ Don't start with database schema
❌ Don't create anemic domain models
❌ Don't ignore bounded contexts
❌ Don't over-engineer simple domains
❌ Don't violate aggregate boundaries
❌ Don't put business logic in application services

---

## Slide 20: Common Pitfalls & Solutions

### Common Pitfalls & Solutions

**Pitfall 1: Anemic Domain Model**
```go
// ❌ Bad: Anemic model
type Order struct {
    ID     string
    Items  []OrderItem
    Status string
}

// ✅ Good: Rich domain model
type Order struct {
    id     OrderID
    items  []OrderItem
    status OrderStatus
}

func (o *Order) Cancel(reason CancellationReason) error {
    if !o.canBeCancelled() {
        return ErrCannotCancel
    }
    o.status = OrderStatusCancelled
    o.recordEvent(OrderCancelled{...})
    return nil
}
```

---

## Slide 21: When to Use DDD

### When to Use DDD

**Good Fit:**
✅ Complex business logic
✅ Business rules change frequently
✅ Long-term project (>6 months)
✅ Domain experts available
✅ Multiple bounded contexts

**Not a Good Fit:**
❌ Simple CRUD applications
❌ Technical/infrastructure projects
❌ Short-term projects
❌ No domain complexity
❌ No access to domain experts

---

## Slide 22: DDD Maturity Levels

### DDD Maturity Levels

**Level 1: Tactical Patterns**
- Using Entities, Value Objects, Repositories
- Basic separation of concerns

**Level 2: Strategic Design**
- Bounded Contexts identified
- Context Mapping
- Ubiquitous Language established

**Level 3: Full DDD**
- Event Storming sessions
- Domain Events
- Event Sourcing / CQRS
- Continuous refinement with domain experts

---

## Slide 23: Real-World Success Stories

### Real-World Success Stories

**Companies Using DDD:**

1. **Amazon** - Service boundaries align with bounded contexts
2. **Netflix** - Microservices following DDD principles
3. **Alibaba** - Large-scale e-commerce with DDD
4. **ING Bank** - Banking domain complexity managed with DDD

**Benefits Reported:**
- 40% reduction in bugs related to business logic
- 60% faster feature development in complex domains
- Better communication with business stakeholders

---

## Slide 24: Tools & Resources

### Tools & Resources

**Event Storming Tools:**
- Miro / Mural (Online collaboration)
- Physical sticky notes

**Go Libraries:**
- `github.com/ThreeDotsLabs/wild-workouts-go-ddd-example`
- `github.com/vardius/go-api-boilerplate`
- `github.com/evrone/go-clean-template`

**Books:**
- Domain-Driven Design (Eric Evans)
- Implementing DDD (Vaughn Vernon)
- Domain-Driven Design Distilled (Vaughn Vernon)

**Online Resources:**
- Martin Fowler's DDD articles
- DDD Community on GitHub
- Domain-Driven Design Europe YouTube

---

## Slide 25: Implementation Checklist

### Implementation Checklist

**Starting Your DDD Journey:**

□ Identify Core Domain
□ Map Bounded Contexts
□ Establish Ubiquitous Language
□ Model Key Aggregates
□ Define Repository Interfaces
□ Implement Domain Services
□ Add Domain Events
□ Create Application Services
□ Set up Infrastructure
□ Write Tests
□ Iterate with Domain Experts

---

## Slide 26: Key Takeaways

### Key Takeaways

1. **DDD is about understanding the business domain**
2. **Ubiquitous Language is the foundation**
3. **Bounded Contexts prevent model pollution**
4. **Rich domain models encapsulate business logic**
5. **Start simple, evolve based on complexity**
6. **Not every project needs DDD**

---

## Slide 27: Questions & Discussion

### Questions & Discussion

**Let's Discuss:**
- Which parts of our system could benefit from DDD?
- What are our core domains?
- Who are our domain experts?
- What challenges do you see in implementing DDD?

**Contact:**
Dong Tran
dong.tran@company.com

---

## Slide 28: Thank You!

### Thank You!

**Next Steps:**
- Form study group for DDD book club
- Identify pilot project for DDD implementation
- Schedule Event Storming workshop

**Remember:**
"The goal of DDD is to create better software by focusing on a model of the domain rather than technology."
- Eric Evans

---

## Appendix A: Complete Order Aggregate Implementation

```go
package order

import (
    "errors"
    "time"
    "github.com/google/uuid"
    "github.com/shopspring/decimal"
)

// Error definitions
var (
    ErrInvalidQuantity   = errors.New("quantity must be positive")
    ErrOrderNotEditable  = errors.New("order cannot be modified")
    ErrInsufficientStock = errors.New("insufficient stock")
    ErrOrderNotFound     = errors.New("order not found")
    ErrInvalidStatus     = errors.New("invalid order status")
)

// OrderID is a value object representing unique order identifier
type OrderID string

func NewOrderID() OrderID {
    return OrderID(uuid.New().String())
}

// OrderStatus represents the state of an order
type OrderStatus int

const (
    OrderStatusDraft OrderStatus = iota
    OrderStatusPlaced
    OrderStatusPaid
    OrderStatusShipped
    OrderStatusDelivered
    OrderStatusCancelled
)

// Order is the aggregate root for the order bounded context
type Order struct {
    id         OrderID
    customerID CustomerID
    items      []OrderItem
    total      Money
    status     OrderStatus
    placedAt   *time.Time
    events     []DomainEvent
    version    int
}

// NewOrder creates a new order in draft status
func NewOrder(customerID CustomerID) *Order {
    return &Order{
        id:         NewOrderID(),
        customerID: customerID,
        items:      make([]OrderItem, 0),
        status:     OrderStatusDraft,
        events:     make([]DomainEvent, 0),
    }
}

// AddItem adds a product to the order
func (o *Order) AddItem(product Product, quantity int) error {
    if o.status != OrderStatusDraft {
        return ErrOrderNotEditable
    }
    
    if quantity <= 0 {
        return ErrInvalidQuantity
    }
    
    if !product.HasSufficientStock(quantity) {
        return ErrInsufficientStock
    }
    
    // Check if item already exists
    for i, item := range o.items {
        if item.productID == product.ID() {
            o.items[i].quantity += quantity
            o.recalculateTotal()
            return nil
        }
    }
    
    // Add new item
    item := OrderItem{
        productID:   product.ID(),
        productName: product.Name(),
        unitPrice:   product.Price(),
        quantity:    quantity,
    }
    
    o.items = append(o.items, item)
    o.recalculateTotal()
    
    return nil
}

// RemoveItem removes a product from the order
func (o *Order) RemoveItem(productID ProductID) error {
    if o.status != OrderStatusDraft {
        return ErrOrderNotEditable
    }
    
    for i, item := range o.items {
        if item.productID == productID {
            o.items = append(o.items[:i], o.items[i+1:]...)
            o.recalculateTotal()
            return nil
        }
    }
    
    return errors.New("item not found in order")
}

// PlaceOrder transitions the order from draft to placed status
func (o *Order) PlaceOrder() error {
    if o.status != OrderStatusDraft {
        return errors.New("order already placed")
    }
    
    if len(o.items) == 0 {
        return errors.New("order must have at least one item")
    }
    
    now := time.Now()
    o.status = OrderStatusPlaced
    o.placedAt = &now
    
    o.recordEvent(OrderPlacedEvent{
        OrderID:    o.id,
        CustomerID: o.customerID,
        Total:      o.total,
        PlacedAt:   now,
    })
    
    return nil
}

// Cancel cancels the order
func (o *Order) Cancel(reason string) error {
    if o.status == OrderStatusCancelled {
        return errors.New("order already cancelled")
    }
    
    if o.status == OrderStatusDelivered {
        return errors.New("cannot cancel delivered order")
    }
    
    o.status = OrderStatusCancelled
    
    o.recordEvent(OrderCancelledEvent{
        OrderID:     o.id,
        Reason:      reason,
        CancelledAt: time.Now(),
    })
    
    return nil
}

// MarkAsPaid marks the order as paid
func (o *Order) MarkAsPaid(paymentID PaymentID) error {
    if o.status != OrderStatusPlaced {
        return errors.New("only placed orders can be marked as paid")
    }
    
    o.status = OrderStatusPaid
    
    o.recordEvent(OrderPaidEvent{
        OrderID:   o.id,
        PaymentID: paymentID,
        PaidAt:    time.Now(),
    })
    
    return nil
}

// Ship marks the order as shipped
func (o *Order) Ship(trackingNumber string) error {
    if o.status != OrderStatusPaid {
        return errors.New("only paid orders can be shipped")
    }
    
    o.status = OrderStatusShipped
    
    o.recordEvent(OrderShippedEvent{
        OrderID:        o.id,
        TrackingNumber: trackingNumber,
        ShippedAt:      time.Now(),
    })
    
    return nil
}

// Deliver marks the order as delivered
func (o *Order) Deliver() error {
    if o.status != OrderStatusShipped {
        return errors.New("only shipped orders can be delivered")
    }
    
    o.status = OrderStatusDelivered
    
    o.recordEvent(OrderDeliveredEvent{
        OrderID:     o.id,
        DeliveredAt: time.Now(),
    })
    
    return nil
}

// Private methods
func (o *Order) recalculateTotal() {
    total := decimal.Zero
    for _, item := range o.items {
        itemTotal := item.unitPrice.Amount.Mul(
            decimal.NewFromInt(int64(item.quantity)),
        )
        total = total.Add(itemTotal)
    }
    o.total = Money{
        Amount:   total,
        Currency: "USD",
    }
}

func (o *Order) recordEvent(event DomainEvent) {
    o.events = append(o.events, event)
}

// Getters
func (o *Order) ID() OrderID             { return o.id }
func (o *Order) CustomerID() CustomerID  { return o.customerID }
func (o *Order) Status() OrderStatus     { return o.status }
func (o *Order) Total() Money            { return o.total }
func (o *Order) Items() []OrderItem      { return o.items }
func (o *Order) PlacedAt() *time.Time    { return o.placedAt }
func (o *Order) Version() int            { return o.version }

// GetEvents returns all uncommitted events
func (o *Order) GetEvents() []DomainEvent {
    return o.events
}

// ClearEvents clears all events after they've been published
func (o *Order) ClearEvents() {
    o.events = []DomainEvent{}
}

// OrderItem is a value object within the Order aggregate
type OrderItem struct {
    productID   ProductID
    productName string
    unitPrice   Money
    quantity    int
}

// Getters for OrderItem
func (i OrderItem) ProductID() ProductID { return i.productID }
func (i OrderItem) ProductName() string  { return i.productName }
func (i OrderItem) UnitPrice() Money     { return i.unitPrice }
func (i OrderItem) Quantity() int        { return i.quantity }

func (i OrderItem) Subtotal() Money {
    return Money{
        Amount: i.unitPrice.Amount.Mul(
            decimal.NewFromInt(int64(i.quantity)),
        ),
        Currency: i.unitPrice.Currency,
    }
}
```

---

## Appendix B: Domain Events Implementation

```go
package order

import "time"

// DomainEvent is the base interface for all domain events
type DomainEvent interface {
    EventName() string
    OccurredAt() time.Time
}

// OrderPlacedEvent is raised when an order is placed
type OrderPlacedEvent struct {
    OrderID    OrderID
    CustomerID CustomerID
    Total      Money
    PlacedAt   time.Time
}

func (e OrderPlacedEvent) EventName() string     { return "order.placed" }
func (e OrderPlacedEvent) OccurredAt() time.Time { return e.PlacedAt }

// OrderCancelledEvent is raised when an order is cancelled
type OrderCancelledEvent struct {
    OrderID     OrderID
    Reason      string
    CancelledAt time.Time
}

func (e OrderCancelledEvent) EventName() string     { return "order.cancelled" }
func (e OrderCancelledEvent) OccurredAt() time.Time { return e.CancelledAt }

// OrderPaidEvent is raised when an order is paid
type OrderPaidEvent struct {
    OrderID   OrderID
    PaymentID PaymentID
    PaidAt    time.Time
}

func (e OrderPaidEvent) EventName() string     { return "order.paid" }
func (e OrderPaidEvent) OccurredAt() time.Time { return e.PaidAt }

// OrderShippedEvent is raised when an order is shipped
type OrderShippedEvent struct {
    OrderID        OrderID
    TrackingNumber string
    ShippedAt      time.Time
}

func (e OrderShippedEvent) EventName() string     { return "order.shipped" }
func (e OrderShippedEvent) OccurredAt() time.Time { return e.ShippedAt }

// OrderDeliveredEvent is raised when an order is delivered
type OrderDeliveredEvent struct {
    OrderID     OrderID
    DeliveredAt time.Time
}

func (e OrderDeliveredEvent) EventName() string     { return "order.delivered" }
func (e OrderDeliveredEvent) OccurredAt() time.Time { return e.DeliveredAt }
```

---

## Appendix C: Repository Implementation Example

```go
package infrastructure

import (
    "context"
    "database/sql"
    "encoding/json"
    "fmt"
    
    "github.com/example/project/internal/domain/order"
)

type PostgresOrderRepository struct {
    db *sql.DB
}

func NewPostgresOrderRepository(db *sql.DB) *PostgresOrderRepository {
    return &PostgresOrderRepository{db: db}
}

func (r *PostgresOrderRepository) Save(ctx context.Context, ord *order.Order) error {
    tx, err := r.db.BeginTx(ctx, nil)
    if err != nil {
        return fmt.Errorf("begin transaction: %w", err)
    }
    defer tx.Rollback()
    
    // Save order
    orderQuery := `
        INSERT INTO orders (
            id, customer_id, total_amount, total_currency, 
            status, placed_at, version
        ) VALUES ($1, $2, $3, $4, $5, $6, $7)
        ON CONFLICT (id) DO UPDATE SET
            total_amount = $3,
            total_currency = $4,
            status = $5,
            placed_at = $6,
            version = orders.version + 1
        WHERE orders.version = $7 - 1
    `
    
    result, err := tx.ExecContext(ctx, orderQuery,
        ord.ID(),
        ord.CustomerID(),
        ord.Total().Amount,
        ord.Total().Currency,
        ord.Status(),
        ord.PlacedAt(),
        ord.Version()+1,
    )
    
    if err != nil {
        return fmt.Errorf("save order: %w", err)
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("get rows affected: %w", err)
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("optimistic locking: order was modified")
    }
    
    // Delete existing items
    deleteItemsQuery := `DELETE FROM order_items WHERE order_id = $1`
    if _, err := tx.ExecContext(ctx, deleteItemsQuery, ord.ID()); err != nil {
        return fmt.Errorf("delete order items: %w", err)
    }
    
    // Save order items
    itemQuery := `
        INSERT INTO order_items (
            order_id, product_id, product_name, 
            unit_price_amount, unit_price_currency, quantity
        ) VALUES ($1, $2, $3, $4, $5, $6)
    `
    
    for _, item := range ord.Items() {
        _, err := tx.ExecContext(ctx, itemQuery,
            ord.ID(),
            item.ProductID(),
            item.ProductName(),
            item.UnitPrice().Amount,
            item.UnitPrice().Currency,
            item.Quantity(),
        )
        if err != nil {
            return fmt.Errorf("save order item: %w", err)
        }
    }
    
    // Save events
    if len(ord.GetEvents()) > 0 {
        eventQuery := `
            INSERT INTO domain_events (
                aggregate_id, event_name, event_data, occurred_at
            ) VALUES ($1, $2, $3, $4)
        `
        
        for _, event := range ord.GetEvents() {
            eventData, err := json.Marshal(event)
            if err != nil {
                return fmt.Errorf("marshal event: %w", err)
            }
            
            _, err = tx.ExecContext(ctx, eventQuery,
                ord.ID(),
                event.EventName(),
                eventData,
                event.OccurredAt(),
            )
            if err != nil {
                return fmt.Errorf("save event: %w", err)
            }
        }
    }
    
    if err := tx.Commit(); err != nil {
        return fmt.Errorf("commit transaction: %w", err)
    }
    
    ord.ClearEvents()
    return nil
}

func (r *PostgresOrderRepository) FindByID(
    ctx context.Context, 
    id order.OrderID,
) (*order.Order, error) {
    // Implementation details...
    return nil, nil
}

func (r *PostgresOrderRepository) FindByCustomer(
    ctx context.Context,
    customerID order.CustomerID,
) ([]*order.Order, error) {
    // Implementation details...
    return nil, nil
}
```

---

## Notes for Presenter

### Time Management (60 minutes total):
- Introduction & Problem Statement: 5 minutes
- Core Concepts: 10 minutes
- Strategic Design: 10 minutes
- Tactical Patterns: 15 minutes
- Implementation & Code Examples: 10 minutes
- Best Practices: 5 minutes
- Q&A: 5 minutes

### Key Points to Emphasize:
1. Start with the domain, not the database
2. Collaborate with domain experts
3. Use ubiquitous language consistently
4. Keep aggregates small and focused
5. Test domain logic thoroughly
6. Don't over-engineer - start simple

### Interactive Elements:
- Ask audience to identify bounded contexts in their systems
- Quick exercise on finding aggregates
- Discussion on current pain points that DDD could solve

### Demo Preparation:
- Have IDE ready with the example code
- Prepare to run tests live
- Show refactoring from anemic to rich domain model

---

**Created by:** Dong Tran  
**Date:** November 19, 2025  
**Version:** 1.0  
**License:** MIT