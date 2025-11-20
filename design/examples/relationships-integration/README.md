# Relationships & Integration Example
## E-Commerce Order Management System

This comprehensive example demonstrates how **Clean Architecture**, **DDD**, **SOLID Principles**, **Design Patterns**, and **Microservices concepts** work together in a real-world application.

## 🎯 What This Example Demonstrates

### Complete Integration of All Concepts

```
┌─────────────────────────────────────────────────────────────┐
│                    HTTP Layer (Echo)                         │
│                   (Frameworks & Drivers)                     │
└──────────────────────┬──────────────────────────────────────┘
                       │ depends on ↓
┌──────────────────────────────────────────────────────────────┐
│                   Handler Layer                              │
│              (Interface Adapters)                            │
└──────────────────────┬──────────────────────────────────────┘
                       │ depends on ↓
┌──────────────────────────────────────────────────────────────┐
│                 Use Case Layer                               │
│          (Application Business Rules)                        │
│    • Uses Strategy Pattern for payments                     │
│    • Uses Observer Pattern for events                       │
│    • Uses Factory Pattern for object creation               │
└──────────────────────┬──────────────────────────────────────┘
                       │ depends on ↓
┌──────────────────────────────────────────────────────────────┐
│                   Domain Layer (DDD)                         │
│         (Enterprise Business Rules)                          │
│    • Order (Aggregate Root)                                 │
│    • Money, OrderID (Value Objects)                         │
│    • OrderItem (Entity)                                     │
│    • Domain Events                                          │
│    • Repository Interfaces (DIP)                            │
└──────────────────────────────────────────────────────────────┘
```

## 📁 Project Structure

```
relationships-integration/
├── cmd/
│   └── main.go                    # Application entry point
├── shared/
│   └── patterns/                  # Design Patterns
│       ├── observer.go            # Observer Pattern
│       ├── strategy.go            # Strategy Pattern
│       └── factory.go             # Factory Pattern
├── domain/
│   └── order/                     # DDD Bounded Context
│       ├── order.go               # Aggregate Root + Value Objects
│       ├── repository.go          # Repository Interface (DIP)
│       └── events.go              # Domain Events
├── usecase/
│   └── order_usecase.go           # Application Services (Clean Architecture)
├── repository/
│   └── order_repository_impl.go   # Repository Implementation (Infrastructure)
├── infrastructure/
│   ├── database.go                # Database setup
│   └── event_handlers.go          # Event handlers (Observer)
└── handler/
    └── order_handler.go           # HTTP handlers (Presentation)
```

## 🔗 How Patterns Work Together

### 1. Clean Architecture Integration

**Layers**:
- **Domain Layer**: Pure business logic (DDD entities, value objects)
- **Use Case Layer**: Application-specific business rules
- **Infrastructure Layer**: External concerns (database, HTTP)
- **Presentation Layer**: HTTP handlers

**Dependency Rule**: All dependencies point inward toward the domain.

### 2. DDD (Domain-Driven Design) Integration

**Tactical Patterns**:
- **Aggregate Root**: `Order` manages consistency boundary
- **Value Objects**: `Money`, `OrderID`, `CustomerID` (immutable)
- **Entities**: `OrderItem` (has identity within aggregate)
- **Repository**: Interface in domain, implementation in infrastructure
- **Domain Events**: `OrderCreatedEvent`, `OrderPaidEvent`, `OrderShippedEvent`
- **Domain Services**: Business logic that doesn't belong to single entity

**Business Rules**:
```go
// Only pending orders can be paid
func (o *Order) MarkAsPaid() error {
    if o.status != OrderStatusPending {
        return errors.New("only pending orders can be marked as paid")
    }
    o.status = OrderStatusPaid
    return nil
}
```

### 3. SOLID Principles Integration

**Single Responsibility Principle (SRP)**:
- Each class has one reason to change
- `Order` handles order business logic
- `OrderRepository` handles persistence
- `OrderHandler` handles HTTP requests

**Open/Closed Principle (OCP)**:
- Strategy Pattern allows new payment methods without modifying existing code
- Observer Pattern allows new event handlers without modifying publisher

**Liskov Substitution Principle (LSP)**:
- All payment strategies can substitute `PaymentStrategy` interface
- All event handlers can substitute `EventObserver` interface

**Interface Segregation Principle (ISP)**:
- Small, focused interfaces (`PaymentStrategy`, `EventObserver`)
- Clients depend only on methods they use

**Dependency Inversion Principle (DIP)**:
- Use case depends on `OrderRepository` interface (abstraction)
- Infrastructure provides concrete implementation
- High-level policy doesn't depend on low-level details

### 4. Design Patterns Integration

**Observer Pattern** (Behavioral):
```go
// Event publisher notifies multiple subscribers
eventPublisher.Publish(Event{
    Type: "OrderCreated",
    Data: OrderCreatedEvent{...},
})

// Multiple handlers react to events
- EmailNotificationHandler
- LoggingHandler
- AnalyticsHandler
```

**Strategy Pattern** (Behavioral):
```go
// Interchangeable payment algorithms
type PaymentStrategy interface {
    ProcessPayment(amount float64, orderID string) error
}

// Different implementations
- CreditCardPayment
- PayPalPayment
- CryptoPayment
```

**Factory Pattern** (Creational):
```go
// Creates payment strategies dynamically
paymentStrategy, err := paymentFactory.CreatePayment("credit_card")
```

### 5. Microservices Concepts

Although this is a single service, it demonstrates microservices principles:
- **Bounded Context**: Order management is isolated
- **Event-Driven**: Publishes domain events
- **API-First**: REST API for inter-service communication
- **Independent Deployment**: Self-contained service
- **Database per Service**: Own database schema

## 🚀 Running the Example

### Prerequisites

```bash
cd relationships-integration
go mod download
```

### Start the Server

```bash
go run cmd/main.go
```

Server starts on `http://localhost:8080`

## 📡 API Usage

### Create Order

```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": "customer-123",
    "items": [
      {
        "product_id": "product-1",
        "product_name": "Laptop",
        "quantity": 1,
        "price": 999.99,
        "currency": "USD"
      },
      {
        "product_id": "product-2",
        "product_name": "Mouse",
        "quantity": 2,
        "price": 29.99,
        "currency": "USD"
      }
    ]
  }'
```

**Response**:
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "customer_id": "customer-123",
  "total": 1059.97,
  "currency": "USD",
  "status": "PENDING"
}
```

**Events Triggered** (visible in console):
```
📧 Email Handler: OrderCreated - ...
📝 Logger: OrderCreated - ...
📊 Analytics: OrderCreated - ...
```

### Process Payment

```bash
curl -X POST http://localhost:8080/orders/{order-id}/payment \
  -H "Content-Type: application/json" \
  -d '{
    "payment_method": "credit_card"
  }'
```

**Available Payment Methods**:
- `credit_card`
- `paypal`
- `crypto`

**Events Triggered**:
```
📧 Email Handler: OrderPaid - ...
📝 Logger: OrderPaid - ...
📊 Analytics: OrderPaid - ...
```

### Get Order

```bash
curl http://localhost:8080/orders/{order-id}
```

## 🎓 Learning Points

### See How Everything Connects

1. **HTTP Request** → Handler (`handler/order_handler.go`)
2. **Handler** → Use Case (`usecase/order_usecase.go`)
3. **Use Case** → Domain Logic (`domain/order/order.go`)
4. **Use Case** → Design Patterns:
   - Factory creates payment strategy
   - Strategy processes payment
   - Observer publishes events
5. **Use Case** → Repository Interface (`domain/order/repository.go`)
6. **Repository** → Database (`repository/order_repository_impl.go`)

### Architecture Benefits Demonstrated

✅ **Testability**: Each layer can be tested independently  
✅ **Flexibility**: Easy to swap implementations (different database, payment provider)  
✅ **Maintainability**: Clear separation of concerns  
✅ **Extensibility**: Add new features without modifying existing code  
✅ **Domain Focus**: Business logic is central and protected  
✅ **Event-Driven**: Loose coupling through domain events

## 📊 Pattern Mapping

| Concept | Implementation | Location |
|---------|---------------|----------|
| Clean Architecture | 4-layer structure | Entire project |
| DDD - Aggregate | Order | `domain/order/order.go` |
| DDD - Value Object | Money, OrderID | `domain/order/order.go` |
| DDD - Repository | OrderRepository | `domain/order/repository.go` |
| DDD - Domain Events | OrderCreatedEvent | `domain/order/events.go` |
| SOLID - SRP | Single responsibility classes | All files |
| SOLID - OCP | Strategy pattern | `shared/patterns/strategy.go` |
| SOLID - DIP | Interface-based design | Repository, Use Case |
| Observer Pattern | Event system | `shared/patterns/observer.go` |
| Strategy Pattern | Payment methods | `shared/patterns/strategy.go` |
| Factory Pattern | Payment creation | `shared/patterns/factory.go` |

## 🔍 Code Examples

### Domain Logic (DDD + Business Rules)

```go
// Rich domain model with business rules
func (o *Order) MarkAsPaid() error {
    if o.status != OrderStatusPending {
        return errors.New("only pending orders can be marked as paid")
    }
    o.status = OrderStatusPaid
    o.updatedAt = time.Now()
    return nil
}
```

### Strategy Pattern (OCP)

```go
// Can add new payment methods without changing existing code
type PaymentStrategy interface {
    ProcessPayment(amount float64, orderID string) error
}

// New payment method - just implement interface
type NewPaymentMethod struct{}
func (n *NewPaymentMethod) ProcessPayment(amount float64, orderID string) error {
    // Implementation
    return nil
}
```

### Dependency Injection (DIP)

```go
// Use case depends on abstractions
type OrderUseCase struct {
    orderRepo      order.OrderRepository      // Interface
    paymentFactory *patterns.PaymentFactory   // Factory
    eventPublisher *patterns.EventPublisher   // Observer
}
```

## 🎯 Key Takeaways

1. **Patterns complement each other** - They work better together than alone
2. **Architecture provides structure** - Clean Architecture gives a place for everything
3. **Domain is king** - Business logic is protected and central
4. **Flexibility through abstraction** - Interfaces enable testing and swapping implementations
5. **Events enable loose coupling** - Components communicate without direct dependencies

## 🚀 Next Steps

Try extending this example:
1. Add new payment method (Strategy Pattern)
2. Add new event handler (Observer Pattern)
3. Implement complete repository methods
4. Add validation use cases
5. Create a second bounded context (Product catalog)
6. Implement CQRS pattern
7. Add integration tests

## 📚 Related Examples

- See individual pattern examples in:
  - `/design/examples/clean-architecture`
  - `/design/examples/ddd`
  - `/design/examples/solid-principles`
  - `/design/examples/design-patterns`
  - `/design/examples/microservices`

---

**This example is the culmination of all architectural concepts working in harmony!** 🎉
