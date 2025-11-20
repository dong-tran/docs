# Microservices Architecture
## Building Scalable Distributed Systems

**Presenter:** Dong Tran  
**Date:** November 20, 2025  
**Time:** 02:57 UTC

---

## Slide 1: Title Slide

# Microservices Architecture
### From Monolith to Distributed Systems

**Presenter:** Dong Tran  
**Date:** November 20, 2025

*"Microservices are a distributed system architecture that emphasizes small, autonomous services that work together."*  
— Sam Newman

---

## Slide 2: Agenda

### Today's Journey Through Microservices

1. **The Evolution: Why Microservices?**
2. **Core Concepts & Principles**
3. **Microservices vs Monolith**
4. **Service Design & Boundaries**
5. **Communication Patterns**
6. **Data Management Strategies**
7. **Implementation in Go**
8. **Service Discovery & API Gateway**
9. **Distributed Transactions & Saga Pattern**
10. **Event-Driven Architecture**
11. **Monitoring & Observability**
12. **Testing Strategies**
13. **Deployment & DevOps**
14. **Real-World Case Study: E-Commerce Platform**
15. **Best Practices & Anti-Patterns**
16. **Q&A Session**

---

## Slide 3: The Problem Space

### Why Do We Need Microservices?

**Monolith Challenges:**
- 🔴 **Scaling Issues** - Must scale entire application
- 🔴 **Long Release Cycles** - Any change requires full deployment
- 🔴 **Technology Lock-in** - Single tech stack for everything
- 🔴 **Team Dependencies** - Teams step on each other's toes
- 🔴 **Single Point of Failure** - One bug can bring down everything

**Real-World Pain:**
```
Netflix 2008: 3-day database corruption
→ Led to complete microservices transformation
→ Now: 1000+ microservices, 98.8% uptime

Amazon 2001: 2-pizza teams
→ Service-oriented architecture
→ Deployment every 11.7 seconds
```

**The Promise of Microservices:**
✅ Independent deployment  
✅ Technology diversity  
✅ Fault isolation  
✅ Scalability  
✅ Team autonomy  

---

## Slide 4: What Are Microservices?

### Microservices Definition

**Characteristics (by Martin Fowler):**

1. **Componentization via Services** - Not libraries
2. **Organized around Business Capabilities** - Not technical layers
3. **Products not Projects** - You build it, you run it
4. **Smart endpoints, dumb pipes** - No ESB
5. **Decentralized Governance** - Choose right tool
6. **Decentralized Data Management** - Each service owns its data
7. **Design for Failure** - Expect failures
8. **Evolutionary Design** - Can evolve independently

**Visual Comparison:**
```
Monolith                    Microservices
┌──────────────┐           ┌─────┐ ┌─────┐ ┌─────┐
│              │           │Order│ │User │ │Cart │
│   All Code   │    →      └──┬──┘ └──┬──┘ └──┬──┘
│              │              │       │       │
└──────┬───────┘           ┌──┴──┐ ┌──┴──┐ ┌──┴──┐
       │                   │DB-1 │ │DB-2 │ │DB-3 │
   ┌───┴───┐               └─────┘ └─────┘ └─────┘
   │  DB   │
   └───────┘
```

---

## Slide 5: Microservices vs Monolith

### Architectural Comparison

| Aspect | Monolith | Microservices |
|--------|----------|---------------|
| **Deployment** | All or nothing | Independent services |
| **Scaling** | Vertical/Horizontal (entire app) | Horizontal (per service) |
| **Technology** | Single stack | Polyglot |
| **Team Structure** | Functional teams | Cross-functional teams |
| **Data Management** | Single database | Database per service |
| **Complexity** | Simple deployment | Complex orchestration |
| **Network** | In-process calls | Network calls |
| **Testing** | Easier integration tests | Complex distributed testing |
| **Debugging** | Simple stack traces | Distributed tracing needed |
| **Initial Cost** | Lower | Higher |

**When to Choose What:**
```
Choose Monolith:              Choose Microservices:
- Small team (<5)             - Large team (>20)
- Simple domain              - Complex domain
- MVP/Prototype              - Proven product
- Low scale                  - High scale needed
- Startup phase              - Scale-up phase
```

---

## Slide 6: Service Boundaries

### Defining Service Boundaries

**Domain-Driven Design Approach:**

```
Bounded Context = Microservice Boundary

E-Commerce System:
┌─────────────────────────────────────────────┐
│                                             │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  │
│  │ Catalog  │  │   Cart   │  │  Order   │  │
│  │ Service  │  │  Service │  │ Service  │  │
│  └──────────┘  └──────────┘  └──────────┘  │
│                                             │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  │
│  │   User   │  │ Payment  │  │ Shipping │  │
│  │ Service  │  │  Service │  │ Service  │  │
│  └──────────┘  └──────────┘  └──────────┘  │
│                                             │
└─────────────────────────────────────────────┘
```

**Service Sizing Guidelines:**
- **Too Small**: Network overhead, operational complexity
- **Too Large**: Loses microservices benefits
- **Just Right**: Single business capability, 2-pizza team

**Example Service Definition:**
```go
// User Service - Clear boundary
type UserService struct {
    // Owns user data
    Users        []User
    Profiles     []Profile
    Preferences  []Preference
    
    // Does NOT own
    // - Orders (Order Service)
    // - Payments (Payment Service)
    // - Products (Catalog Service)
}
```

---

## Slide 7: Communication Patterns

### Inter-Service Communication

**1. Synchronous Communication:**

```go
// REST API Call
type OrderService struct {
    userClient UserServiceClient
}

func (s *OrderService) CreateOrder(ctx context.Context, req CreateOrderRequest) error {
    // Synchronous call to User Service
    user, err := s.userClient.GetUser(ctx, req.UserID)
    if err != nil {
        return fmt.Errorf("failed to get user: %w", err)
    }
    
    if !user.IsActive {
        return ErrInactiveUser
    }
    
    // Create order...
}

// gRPC Implementation
type UserServiceClient struct {
    conn *grpc.ClientConn
}

func (c *UserServiceClient) GetUser(ctx context.Context, userID string) (*User, error) {
    client := pb.NewUserServiceClient(c.conn)
    
    resp, err := client.GetUser(ctx, &pb.GetUserRequest{
        UserId: userID,
    })
    
    if err != nil {
        return nil, err
    }
    
    return &User{
        ID:     resp.Id,
        Name:   resp.Name,
        Email:  resp.Email,
        Active: resp.Active,
    }, nil
}
```

**2. Asynchronous Communication:**

```go
// Event-Driven with Message Queue
type OrderService struct {
    publisher EventPublisher
}

func (s *OrderService) CreateOrder(ctx context.Context, order Order) error {
    // Save order
    if err := s.repository.Save(ctx, order); err != nil {
        return err
    }
    
    // Publish event asynchronously
    event := OrderCreatedEvent{
        OrderID:    order.ID,
        UserID:     order.UserID,
        TotalAmount: order.Total,
        CreatedAt:  time.Now(),
    }
    
    return s.publisher.Publish(ctx, "order.created", event)
}

// Consumer in Inventory Service
type InventoryService struct {
    consumer EventConsumer
}

func (s *InventoryService) Start(ctx context.Context) {
    s.consumer.Subscribe("order.created", func(event OrderCreatedEvent) error {
        // React to order creation
        return s.reserveInventory(event.OrderID, event.Items)
    })
}
```

---

## Slide 8: API Gateway Pattern

### API Gateway Implementation

```go
// api-gateway/main.go
package main

import (
    "context"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/sony/gobreaker"
)

type APIGateway struct {
    userService    ServiceClient
    orderService   ServiceClient
    catalogService ServiceClient
    circuitBreaker *gobreaker.CircuitBreaker
}

func NewAPIGateway() *APIGateway {
    settings := gobreaker.Settings{
        Name:        "API-Gateway",
        MaxRequests: 3,
        Interval:    5 * time.Second,
        Timeout:     10 * time.Second,
        OnStateChange: func(name string, from, to gobreaker.State) {
            log.Printf("Circuit breaker %s: %s -> %s", name, from, to)
        },
    }
    
    return &APIGateway{
        userService:    NewServiceClient("http://user-service:8080"),
        orderService:   NewServiceClient("http://order-service:8080"),
        catalogService: NewServiceClient("http://catalog-service:8080"),
        circuitBreaker: gobreaker.NewCircuitBreaker(settings),
    }
}

// Aggregate data from multiple services
func (g *APIGateway) GetUserDashboard(c *gin.Context) {
    userID := c.Param("userID")
    
    // Parallel calls to multiple services
    var wg sync.WaitGroup
    var user *User
    var orders []Order
    var recommendations []Product
    var errors []error
    
    // Get user info
    wg.Add(1)
    go func() {
        defer wg.Done()
        result, err := g.circuitBreaker.Execute(func() (interface{}, error) {
            return g.userService.GetUser(c.Request.Context(), userID)
        })
        if err != nil {
            errors = append(errors, err)
            return
        }
        user = result.(*User)
    }()
    
    // Get recent orders
    wg.Add(1)
    go func() {
        defer wg.Done()
        result, err := g.orderService.GetUserOrders(c.Request.Context(), userID)
        if err != nil {
            // Graceful degradation
            orders = []Order{} // Empty list on failure
            return
        }
        orders = result
    }()
    
    // Get recommendations
    wg.Add(1)
    go func() {
        defer wg.Done()
        result, err := g.catalogService.GetRecommendations(c.Request.Context(), userID)
        if err != nil {
            // Optional data, ignore error
            recommendations = []Product{}
            return
        }
        recommendations = result
    }()
    
    wg.Wait()
    
    // Check for critical errors
    if user == nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "error": "User service unavailable",
        })
        return
    }
    
    // Compose response
    c.JSON(http.StatusOK, gin.H{
        "user":            user,
        "recent_orders":   orders,
        "recommendations": recommendations,
    })
}

// Rate limiting
func (g *APIGateway) RateLimitMiddleware() gin.HandlerFunc {
    limiter := NewRateLimiter(100, 1*time.Minute) // 100 requests per minute
    
    return func(c *gin.Context) {
        userID := c.GetHeader("X-User-ID")
        if !limiter.Allow(userID) {
            c.JSON(http.StatusTooManyRequests, gin.H{
                "error": "Rate limit exceeded",
            })
            c.Abort()
            return
        }
        c.Next()
    }
}
```

---

## Slide 9: Service Discovery

### Service Discovery Patterns

**1. Client-Side Discovery (Consul)**

```go
// service-registry/consul.go
package registry

import (
    "fmt"
    consulapi "github.com/hashicorp/consul/api"
)

type ConsulRegistry struct {
    client *consulapi.Client
}

// Register service with Consul
func (r *ConsulRegistry) Register(service ServiceInfo) error {
    registration := &consulapi.AgentServiceRegistration{
        ID:      service.ID,
        Name:    service.Name,
        Address: service.Address,
        Port:    service.Port,
        Check: &consulapi.AgentServiceCheck{
            HTTP:                           fmt.Sprintf("http://%s:%d/health", service.Address, service.Port),
            Interval:                       "10s",
            Timeout:                        "2s",
            DeregisterCriticalServiceAfter: "30s",
        },
        Tags: service.Tags,
    }
    
    return r.client.Agent().ServiceRegister(registration)
}

// Discover healthy services
func (r *ConsulRegistry) Discover(serviceName string) ([]ServiceInstance, error) {
    services, _, err := r.client.Health().Service(serviceName, "", true, nil)
    if err != nil {
        return nil, err
    }
    
    instances := make([]ServiceInstance, 0, len(services))
    for _, service := range services {
        instances = append(instances, ServiceInstance{
            ID:      service.Service.ID,
            Address: service.Service.Address,
            Port:    service.Service.Port,
        })
    }
    
    return instances, nil
}

// Service client with discovery
type DiscoverableClient struct {
    registry     *ConsulRegistry
    serviceName  string
    loadBalancer LoadBalancer
}

func (c *DiscoverableClient) Call(ctx context.Context, path string, payload interface{}) (interface{}, error) {
    // Discover available instances
    instances, err := c.registry.Discover(c.serviceName)
    if err != nil {
        return nil, fmt.Errorf("service discovery failed: %w", err)
    }
    
    if len(instances) == 0 {
        return nil, fmt.Errorf("no healthy instances of %s found", c.serviceName)
    }
    
    // Load balance between instances
    instance := c.loadBalancer.Choose(instances)
    
    // Make the call
    url := fmt.Sprintf("http://%s:%d%s", instance.Address, instance.Port, path)
    return c.makeHTTPCall(ctx, url, payload)
}
```

**2. Server-Side Discovery (Kubernetes)**

```yaml
# kubernetes/user-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: user-service
  labels:
    app: user-service
spec:
  selector:
    app: user-service
  ports:
    - port: 8080
      targetPort: 8080
  type: ClusterIP
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: user-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: user-service
  template:
    metadata:
      labels:
        app: user-service
    spec:
      containers:
      - name: user-service
        image: user-service:v1.0.0
        ports:
        - containerPort: 8080
        env:
        - name: DB_HOST
          valueFrom:
            secretKeyRef:
              name: db-secret
              key: host
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
```

---

## Slide 10: Data Management

### Database per Service Pattern

**The Challenge:**
```
Traditional:                 Microservices:
┌──────────┐                ┌──────────┐
│  App 1   │                │Service 1 │
├──────────┤                └────┬─────┘
│  App 2   │ → Shared DB         │ Private DB
├──────────┤                ┌────┴─────┐
│  App 3   │                │   DB 1   │
└────┬─────┘                └──────────┘
     │
┌────┴─────┐                Each service owns its data!
│    DB    │
└──────────┘
```

**Implementation Strategies:**

```go
// 1. Shared Data via API
type UserService struct {
    db *sql.DB
}

// Other services call this API
func (s *UserService) GetUser(ctx context.Context, userID string) (*User, error) {
    var user User
    err := s.db.QueryRowContext(ctx, 
        "SELECT id, name, email FROM users WHERE id = $1", 
        userID,
    ).Scan(&user.ID, &user.Name, &user.Email)
    
    return &user, err
}

// 2. Event Sourcing for Data Sync
type OrderService struct {
    db        *sql.DB
    publisher EventPublisher
}

func (s *OrderService) CreateOrder(ctx context.Context, order Order) error {
    tx, err := s.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    // Save order
    _, err = tx.ExecContext(ctx,
        "INSERT INTO orders (id, user_id, total) VALUES ($1, $2, $3)",
        order.ID, order.UserID, order.Total,
    )
    
    if err != nil {
        return err
    }
    
    // Publish event for other services
    event := OrderCreatedEvent{
        OrderID:  order.ID,
        UserID:   order.UserID,
        Total:    order.Total,
        Items:    order.Items,
        Timestamp: time.Now(),
    }
    
    if err := s.publisher.Publish(ctx, event); err != nil {
        return err
    }
    
    return tx.Commit()
}

// 3. CQRS Pattern
type CommandService struct {
    writeDB *sql.DB
}

type QueryService struct {
    readDB *sql.DB // Read replica or different database
}

func (s *CommandService) CreateProduct(ctx context.Context, product Product) error {
    // Write to master database
    _, err := s.writeDB.ExecContext(ctx,
        "INSERT INTO products (id, name, price) VALUES ($1, $2, $3)",
        product.ID, product.Name, product.Price,
    )
    return err
}

func (s *QueryService) SearchProducts(ctx context.Context, query string) ([]Product, error) {
    // Read from optimized read database
    rows, err := s.readDB.QueryContext(ctx,
        "SELECT id, name, price FROM product_search WHERE name ILIKE $1",
        "%"+query+"%",
    )
    // ...
}
```

---

## Slide 11: Distributed Transactions - Saga Pattern

### Implementing Saga Pattern

```go
// saga/orchestrator.go
package saga

import (
    "context"
    "fmt"
)

type SagaStep struct {
    Name        string
    Transaction func(ctx context.Context, data interface{}) error
    Compensation func(ctx context.Context, data interface{}) error
}

type SagaOrchestrator struct {
    steps []SagaStep
}

func (s *SagaOrchestrator) Execute(ctx context.Context, data interface{}) error {
    completedSteps := []SagaStep{}
    
    for _, step := range s.steps {
        if err := step.Transaction(ctx, data); err != nil {
            // Rollback completed steps in reverse order
            for i := len(completedSteps) - 1; i >= 0; i-- {
                compensationErr := completedSteps[i].Compensation(ctx, data)
                if compensationErr != nil {
                    // Log compensation failure
                    fmt.Printf("Compensation failed for %s: %v\n", 
                        completedSteps[i].Name, compensationErr)
                }
            }
            return fmt.Errorf("saga failed at step %s: %w", step.Name, err)
        }
        completedSteps = append(completedSteps, step)
    }
    
    return nil
}

// Order Processing Saga Example
type OrderSaga struct {
    orchestrator    *SagaOrchestrator
    orderService    OrderService
    paymentService  PaymentService
    inventoryService InventoryService
    shippingService ShippingService
}

func NewOrderSaga() *OrderSaga {
    saga := &OrderSaga{}
    
    saga.orchestrator = &SagaOrchestrator{
        steps: []SagaStep{
            {
                Name: "Create Order",
                Transaction: func(ctx context.Context, data interface{}) error {
                    order := data.(*OrderData)
                    return saga.orderService.CreateOrder(ctx, order)
                },
                Compensation: func(ctx context.Context, data interface{}) error {
                    order := data.(*OrderData)
                    return saga.orderService.CancelOrder(ctx, order.ID)
                },
            },
            {
                Name: "Reserve Inventory",
                Transaction: func(ctx context.Context, data interface{}) error {
                    order := data.(*OrderData)
                    return saga.inventoryService.ReserveItems(ctx, order.Items)
                },
                Compensation: func(ctx context.Context, data interface{}) error {
                    order := data.(*OrderData)
                    return saga.inventoryService.ReleaseItems(ctx, order.Items)
                },
            },
            {
                Name: "Process Payment",
                Transaction: func(ctx context.Context, data interface{}) error {
                    order := data.(*OrderData)
                    return saga.paymentService.ChargePayment(ctx, order.PaymentInfo)
                },
                Compensation: func(ctx context.Context, data interface{}) error {
                    order := data.(*OrderData)
                    return saga.paymentService.RefundPayment(ctx, order.PaymentInfo)
                },
            },
            {
                Name: "Create Shipment",
                Transaction: func(ctx context.Context, data interface{}) error {
                    order := data.(*OrderData)
                    return saga.shippingService.CreateShipment(ctx, order)
                },
                Compensation: func(ctx context.Context, data interface{}) error {
                    order := data.(*OrderData)
                    return saga.shippingService.CancelShipment(ctx, order.ID)
                },
            },
        },
    }
    
    return saga
}

func (s *OrderSaga) ProcessOrder(ctx context.Context, order *OrderData) error {
    return s.orchestrator.Execute(ctx, order)
}

// Choreography-based Saga (Event-driven)
type PaymentService struct {
    eventBus EventBus
}

func (s *PaymentService) HandleOrderCreated(event OrderCreatedEvent) error {
    // Process payment
    payment := Payment{
        OrderID: event.OrderID,
        Amount:  event.Total,
    }
    
    if err := s.processPayment(payment); err != nil {
        // Publish compensation event
        return s.eventBus.Publish(PaymentFailedEvent{
            OrderID: event.OrderID,
            Reason:  err.Error(),
        })
    }
    
    // Publish success event
    return s.eventBus.Publish(PaymentCompletedEvent{
        OrderID:   event.OrderID,
        PaymentID: payment.ID,
    })
}
```

---

## Slide 12: Event-Driven Architecture

### Event Streaming with Kafka

```go
// events/producer.go
package events

import (
    "context"
    "encoding/json"
    "github.com/segmentio/kafka-go"
)

type EventProducer struct {
    writer *kafka.Writer
}

func NewEventProducer(brokers []string, topic string) *EventProducer {
    return &EventProducer{
        writer: kafka.NewWriter(kafka.WriterConfig{
            Brokers:  brokers,
            Topic:    topic,
            Balancer: &kafka.LeastBytes{},
            Async:    true,
        }),
    }
}

func (p *EventProducer) PublishEvent(ctx context.Context, event Event) error {
    data, err := json.Marshal(event)
    if err != nil {
        return err
    }
    
    message := kafka.Message{
        Key:   []byte(event.AggregateID()),
        Value: data,
        Headers: []kafka.Header{
            {Key: "event_type", Value: []byte(event.EventType())},
            {Key: "timestamp", Value: []byte(event.Timestamp().String())},
        },
    }
    
    return p.writer.WriteMessages(ctx, message)
}

// events/consumer.go
type EventConsumer struct {
    reader   *kafka.Reader
    handlers map[string]EventHandler
}

type EventHandler func(ctx context.Context, event []byte) error

func NewEventConsumer(brokers []string, topic, groupID string) *EventConsumer {
    return &EventConsumer{
        reader: kafka.NewReader(kafka.ReaderConfig{
            Brokers:  brokers,
            Topic:    topic,
            GroupID:  groupID,
            MinBytes: 10e3, // 10KB
            MaxBytes: 10e6, // 10MB
        }),
        handlers: make(map[string]EventHandler),
    }
}

func (c *EventConsumer) RegisterHandler(eventType string, handler EventHandler) {
    c.handlers[eventType] = handler
}

func (c *EventConsumer) Start(ctx context.Context) error {
    for {
        message, err := c.reader.FetchMessage(ctx)
        if err != nil {
            return err
        }
        
        eventType := c.getEventType(message.Headers)
        handler, exists := c.handlers[eventType]
        
        if exists {
            if err := handler(ctx, message.Value); err != nil {
                // Handle error (retry, DLQ, etc.)
                continue
            }
        }
        
        if err := c.reader.CommitMessages(ctx, message); err != nil {
            return err
        }
    }
}

// Event Store Implementation
type EventStore struct {
    db       *sql.DB
    producer *EventProducer
}

func (s *EventStore) SaveEvent(ctx context.Context, event Event) error {
    tx, err := s.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    // Store event in database
    _, err = tx.ExecContext(ctx,
        `INSERT INTO events (aggregate_id, event_type, event_data, created_at)
         VALUES ($1, $2, $3, $4)`,
        event.AggregateID(),
        event.EventType(),
        event.Data(),
        event.Timestamp(),
    )
    
    if err != nil {
        return err
    }
    
    // Publish to Kafka
    if err := s.producer.PublishEvent(ctx, event); err != nil {
        return err
    }
    
    return tx.Commit()
}
```

---

## Slide 13: Resilience Patterns

### Circuit Breaker & Retry

```go
// resilience/circuit_breaker.go
package resilience

import (
    "context"
    "fmt"
    "time"
    "github.com/sony/gobreaker"
)

type ResilientClient struct {
    circuitBreaker *gobreaker.CircuitBreaker
    retryPolicy    RetryPolicy
}

type RetryPolicy struct {
    MaxAttempts int
    InitialDelay time.Duration
    MaxDelay     time.Duration
    Multiplier   float64
}

func NewResilientClient(name string) *ResilientClient {
    settings := gobreaker.Settings{
        Name:        name,
        MaxRequests: 3,
        Interval:    10 * time.Second,
        Timeout:     60 * time.Second,
        ReadyToTrip: func(counts gobreaker.Counts) bool {
            failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
            return counts.Requests >= 3 && failureRatio >= 0.6
        },
        OnStateChange: func(name string, from, to gobreaker.State) {
            fmt.Printf("Circuit Breaker %s: %s -> %s\n", name, from, to)
        },
    }
    
    return &ResilientClient{
        circuitBreaker: gobreaker.NewCircuitBreaker(settings),
        retryPolicy: RetryPolicy{
            MaxAttempts:  3,
            InitialDelay: 100 * time.Millisecond,
            MaxDelay:     2 * time.Second,
            Multiplier:   2.0,
        },
    }
}

func (c *ResilientClient) Execute(ctx context.Context, fn func() (interface{}, error)) (interface{}, error) {
    return c.circuitBreaker.Execute(func() (interface{}, error) {
        return c.executeWithRetry(ctx, fn)
    })
}

func (c *ResilientClient) executeWithRetry(ctx context.Context, fn func() (interface{}, error)) (interface{}, error) {
    delay := c.retryPolicy.InitialDelay
    
    for attempt := 1; attempt <= c.retryPolicy.MaxAttempts; attempt++ {
        result, err := fn()
        if err == nil {
            return result, nil
        }
        
        if attempt == c.retryPolicy.MaxAttempts {
            return nil, fmt.Errorf("max retries exceeded: %w", err)
        }
        
        // Exponential backoff
        select {
        case <-time.After(delay):
            delay = time.Duration(float64(delay) * c.retryPolicy.Multiplier)
            if delay > c.retryPolicy.MaxDelay {
                delay = c.retryPolicy.MaxDelay
            }
        case <-ctx.Done():
            return nil, ctx.Err()
        }
    }
    
    return nil, fmt.Errorf("retry failed")
}

// Bulkhead Pattern
type Bulkhead struct {
    semaphore chan struct{}
}

func NewBulkhead(maxConcurrent int) *Bulkhead {
    return &Bulkhead{
        semaphore: make(chan struct{}, maxConcurrent),
    }
}

func (b *Bulkhead) Execute(ctx context.Context, fn func() error) error {
    select {
    case b.semaphore <- struct{}{}:
        defer func() { <-b.semaphore }()
        return fn()
    case <-ctx.Done():
        return ctx.Err()
    }
}

// Timeout Pattern
func WithTimeout(timeout time.Duration, fn func(context.Context) error) error {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()
    
    done := make(chan error, 1)
    go func() {
        done <- fn(ctx)
    }()
    
    select {
    case err := <-done:
        return err
    case <-ctx.Done():
        return fmt.Errorf("operation timed out after %v", timeout)
    }
}
```

---

## Slide 14: Monitoring & Observability

### Distributed Tracing with OpenTelemetry

```go
// observability/tracing.go
package observability

import (
    "context"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/propagation"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    "go.opentelemetry.io/otel/trace"
)

func InitTracer(serviceName, jaegerEndpoint string) func() {
    exporter, err := jaeger.New(
        jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerEndpoint)),
    )
    if err != nil {
        panic(err)
    }
    
    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
        sdktrace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String(serviceName),
        )),
    )
    
    otel.SetTracerProvider(tp)
    otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
        propagation.TraceContext{},
        propagation.Baggage{},
    ))
    
    return func() {
        _ = tp.Shutdown(context.Background())
    }
}

// Tracing Middleware
func TracingMiddleware(serviceName string) gin.HandlerFunc {
    tracer := otel.Tracer(serviceName)
    
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        
        // Extract parent span from headers
        ctx = otel.GetTextMapPropagator().Extract(ctx, 
            propagation.HeaderCarrier(c.Request.Header))
        
        // Start new span
        ctx, span := tracer.Start(ctx, c.Request.URL.Path,
            trace.WithSpanKind(trace.SpanKindServer),
            trace.WithAttributes(
                attribute.String("http.method", c.Request.Method),
                attribute.String("http.url", c.Request.URL.String()),
                attribute.String("http.user_agent", c.Request.UserAgent()),
            ),
        )
        defer span.End()
        
        // Pass context down
        c.Request = c.Request.WithContext(ctx)
        c.Next()
        
        // Record response
        span.SetAttributes(
            attribute.Int("http.status_code", c.Writer.Status()),
        )
        
        if c.Writer.Status() >= 400 {
            span.RecordError(fmt.Errorf("HTTP %d", c.Writer.Status()))
        }
    }
}

// Metrics with Prometheus
type Metrics struct {
    requestDuration *prometheus.HistogramVec
    requestCount    *prometheus.CounterVec
    errorCount      *prometheus.CounterVec
}

func NewMetrics(namespace, service string) *Metrics {
    m := &Metrics{
        requestDuration: prometheus.NewHistogramVec(
            prometheus.HistogramOpts{
                Namespace: namespace,
                Subsystem: service,
                Name:      "request_duration_seconds",
                Help:      "Request duration in seconds",
                Buckets:   prometheus.DefBuckets,
            },
            []string{"method", "endpoint", "status"},
        ),
        requestCount: prometheus.NewCounterVec(
            prometheus.CounterOpts{
                Namespace: namespace,
                Subsystem: service,
                Name:      "request_total",
                Help:      "Total number of requests",
            },
            []string{"method", "endpoint", "status"},
        ),
        errorCount: prometheus.NewCounterVec(
            prometheus.CounterOpts{
                Namespace: namespace,
                Subsystem: service,
                Name:      "error_total",
                Help:      "Total number of errors",
            },
            []string{"method", "endpoint", "error_type"},
        ),
    }
    
    prometheus.MustRegister(m.requestDuration, m.requestCount, m.errorCount)
    return m
}

// Centralized Logging
type StructuredLogger struct {
    logger *zap.Logger
}

func NewStructuredLogger(serviceName string) *StructuredLogger {
    config := zap.NewProductionConfig()
    config.InitialFields = map[string]interface{}{
        "service": serviceName,
    }
    
    logger, _ := config.Build()
    return &StructuredLogger{logger: logger}
}

func (l *StructuredLogger) LogRequest(ctx context.Context, method, path string, duration time.Duration, status int) {
    span := trace.SpanFromContext(ctx)
    
    l.logger.Info("request",
        zap.String("trace_id", span.SpanContext().TraceID().String()),
        zap.String("span_id", span.SpanContext().SpanID().String()),
        zap.String("method", method),
        zap.String("path", path),
        zap.Duration("duration", duration),
        zap.Int("status", status),
    )
}
```

---

## Slide 15: Testing Microservices

### Testing Strategy

```go
// Unit Testing
func TestOrderService_CreateOrder(t *testing.T) {
    // Mock dependencies
    mockRepo := new(MockOrderRepository)
    mockEventBus := new(MockEventBus)
    
    service := NewOrderService(mockRepo, mockEventBus)
    
    order := Order{
        UserID: "user-123",
        Items:  []Item{{ProductID: "prod-456", Quantity: 2}},
    }
    
    mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
    mockEventBus.On("Publish", mock.Anything, mock.Anything).Return(nil)
    
    err := service.CreateOrder(context.Background(), order)
    
    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
    mockEventBus.AssertExpectations(t)
}

// Integration Testing with TestContainers
func TestOrderServiceIntegration(t *testing.T) {
    ctx := context.Background()
    
    // Start PostgreSQL container
    postgres, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
        ContainerRequest: testcontainers.ContainerRequest{
            Image:        "postgres:13",
            ExposedPorts: []string{"5432/tcp"},
            Env: map[string]string{
                "POSTGRES_PASSWORD": "test",
                "POSTGRES_DB":       "testdb",
            },
            WaitingFor: wait.ForListeningPort("5432/tcp"),
        },
        Started: true,
    })
    require.NoError(t, err)
    defer postgres.Terminate(ctx)
    
    // Start Kafka container
    kafka, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
        ContainerRequest: testcontainers.ContainerRequest{
            Image:        "confluentinc/cp-kafka:latest",
            ExposedPorts: []string{"9092/tcp"},
            WaitingFor:   wait.ForListeningPort("9092/tcp"),
        },
        Started: true,
    })
    require.NoError(t, err)
    defer kafka.Terminate(ctx)
    
    // Get connection strings
    postgresHost, _ := postgres.Host(ctx)
    postgresPort, _ := postgres.MappedPort(ctx, "5432")
    
    kafkaHost, _ := kafka.Host(ctx)
    kafkaPort, _ := kafka.MappedPort(ctx, "9092")
    
    // Initialize service with real dependencies
    db, _ := sql.Open("postgres", fmt.Sprintf(
        "postgres://postgres:test@%s:%s/testdb?sslmode=disable",
        postgresHost, postgresPort.Port(),
    ))
    
    eventBus := NewKafkaEventBus([]string{
        fmt.Sprintf("%s:%s", kafkaHost, kafkaPort.Port()),
    })
    
    service := NewOrderService(db, eventBus)
    
    // Run tests
    order := Order{UserID: "user-123"}
    err = service.CreateOrder(ctx, order)
    assert.NoError(t, err)
}

// Contract Testing with Pact
func TestUserServiceContract(t *testing.T) {
    // Create Pact client
    pact := &dsl.Pact{
        Consumer: "OrderService",
        Provider: "UserService",
        Host:     "localhost",
    }
    defer pact.Teardown()
    
    // Define interaction
    pact.AddInteraction().
        Given("User user-123 exists").
        UponReceiving("A request for user user-123").
        WithRequest(dsl.Request{
            Method: "GET",
            Path:   dsl.String("/users/user-123"),
        }).
        WillRespondWith(dsl.Response{
            Status: 200,
            Body: dsl.Like(map[string]interface{}{
                "id":     "user-123",
                "name":   "John Doe",
                "email":  "john@example.com",
                "active": true,
            }),
        })
    
    // Test
    err := pact.Verify(func() error {
        client := NewUserServiceClient(
            fmt.Sprintf("http://localhost:%d", pact.Server.Port),
        )
        
        user, err := client.GetUser(context.Background(), "user-123")
        if err != nil {
            return err
        }
        
        assert.Equal(t, "user-123", user.ID)
        assert.Equal(t, "John Doe", user.Name)
        return nil
    })
    
    assert.NoError(t, err)
}

// End-to-End Testing
func TestE2EOrderFlow(t *testing.T) {
    // Deploy services to test environment
    services := DeployTestStack()
    defer services.Teardown()
    
    // Create test client
    client := NewAPIClient(services.GatewayURL)
    
    // Test complete order flow
    // 1. Create user
    user, err := client.CreateUser(User{
        Name:  "Test User",
        Email: "test@example.com",
    })
    assert.NoError(t, err)
    
    // 2. Add product to catalog
    product, err := client.AddProduct(Product{
        Name:  "Test Product",
        Price: 99.99,
    })
    assert.NoError(t, err)
    
    // 3. Create order
    order, err := client.CreateOrder(Order{
        UserID: user.ID,
        Items: []Item{
            {ProductID: product.ID, Quantity: 1},
        },
    })
    assert.NoError(t, err)
    
    // 4. Verify order status
    time.Sleep(2 * time.Second) // Wait for async processing
    
    status, err := client.GetOrderStatus(order.ID)
    assert.NoError(t, err)
    assert.Equal(t, "CONFIRMED", status)
}
```

---

## Slide 16: Deployment Strategies

### Container Orchestration with Kubernetes

```yaml
# kubernetes/microservices-deployment.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: microservices
---
# ConfigMap for shared configuration
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
  namespace: microservices
data:
  KAFKA_BROKERS: "kafka-0.kafka:9092,kafka-1.kafka:9092"
  JAEGER_ENDPOINT: "http://jaeger-collector:14268/api/traces"
---
# Order Service Deployment
apiVersion: apps/v1
kind: Deployment
metadata:
  name: order-service
  namespace: microservices
spec:
  replicas: 3
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
  selector:
    matchLabels:
      app: order-service
  template:
    metadata:
      labels:
        app: order-service
        version: v1.0.0
      annotations:
        prometheus.io/scrape: "true"
        prometheus.io/port: "8080"
        prometheus.io/path: "/metrics"
    spec:
      containers:
      - name: order-service
        image: order-service:v1.0.0
        ports:
        - containerPort: 8080
          name: http
        - containerPort: 8081
          name: grpc
        env:
        - name: SERVICE_NAME
          value: "order-service"
        - name: DB_HOST
          valueFrom:
            secretKeyRef:
              name: db-credentials
              key: host
        - name: KAFKA_BROKERS
          valueFrom:
            configMapKeyRef:
              name: app-config
              key: KAFKA_BROKERS
        resources:
          requests:
            memory: "128Mi"
            cpu: "100m"
          limits:
            memory: "256Mi"
            cpu: "200m"
        livenessProbe:
          httpGet:
            path: /health/live
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /health/ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
---
# Service Mesh with Istio
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: order-service
  namespace: microservices
spec:
  hosts:
  - order-service
  http:
  - match:
    - headers:
        version:
          exact: v2
    route:
    - destination:
        host: order-service
        subset: v2
      weight: 100
  - route:
    - destination:
        host: order-service
        subset: v1
      weight: 90
    - destination:
        host: order-service
        subset: v2
      weight: 10  # Canary deployment
---
# Horizontal Pod Autoscaler
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: order-service-hpa
  namespace: microservices
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: order-service
  minReplicas: 3
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
  - type: Pods
    pods:
      metric:
        name: custom_metric
      target:
        type: AverageValue
        averageValue: "30"
```

---

## Slide 17: Real-World Case Study

### E-Commerce Platform Migration

**Before: Monolithic Architecture**
```
Single Application (500K LOC)
├── User Management
├── Product Catalog  
├── Shopping Cart
├── Order Processing
├── Payment
├── Inventory
├── Shipping
└── Notifications

Problems:
- 2-hour deployment time
- Single database (20TB)
- 15-minute startup time
- Can't scale individual features
```

**After: Microservices Architecture**

```go
// Service Boundaries Definition
services := []Service{
    {Name: "user-service",      Team: "Identity",    Language: "Go"},
    {Name: "catalog-service",   Team: "Catalog",     Language: "Go"},
    {Name: "cart-service",      Team: "Shopping",    Language: "Go"},
    {Name: "order-service",     Team: "Orders",      Language: "Go"},
    {Name: "payment-service",   Team: "Payments",    Language: "Java"},
    {Name: "inventory-service", Team: "Warehouse",   Language: "Go"},
    {Name: "shipping-service",  Team: "Logistics",   Language: "Python"},
    {Name: "notification-service", Team: "Engagement", Language: "Node.js"},
}

// Migration Strategy
type MigrationPhase struct {
    Phase    int
    Duration string
    Services []string
    Strategy string
}

phases := []MigrationPhase{
    {
        Phase:    1,
        Duration: "3 months",
        Services: []string{"user-service"},
        Strategy: "Strangler Fig Pattern",
    },
    {
        Phase:    2,
        Duration: "4 months",
        Services: []string{"catalog-service", "cart-service"},
        Strategy: "Branch by Abstraction",
    },
    {
        Phase:    3,
        Duration: "6 months",
        Services: []string{"order-service", "payment-service"},
        Strategy: "Event Sourcing Migration",
    },
    {
        Phase:    4,
        Duration: "3 months",
        Services: []string{"inventory-service", "shipping-service", "notification-service"},
        Strategy: "Parallel Run",
    },
}
```

**Results After 18 Months:**
- **Deployment**: 2 hours → 10 minutes per service
- **Scaling**: Can scale services independently
- **Availability**: 99.5% → 99.95%
- **Development Speed**: 2x faster feature delivery
- **Team Autonomy**: Teams deploy independently
- **Cost**: 30% reduction in infrastructure costs

---

## Slide 18: Security in Microservices

### Security Patterns

```go
// security/auth.go
package security

import (
    "context"
    "crypto/rsa"
    "github.com/dgrijalva/jwt-go"
    "google.golang.org/grpc"
    "google.golang.org/grpc/metadata"
)

// Service-to-Service Authentication (mTLS)
func NewSecureGRPCClient(certFile, keyFile, caFile, serverAddr string) (*grpc.ClientConn, error) {
    // Load client certificates
    cert, err := tls.LoadX509KeyPair(certFile, keyFile)
    if err != nil {
        return nil, err
    }
    
    // Load CA certificate
    caCert, err := ioutil.ReadFile(caFile)
    if err != nil {
        return nil, err
    }
    
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)
    
    // Create TLS configuration
    tlsConfig := &tls.Config{
        Certificates: []tls.Certificate{cert},
        RootCAs:      caCertPool,
    }
    
    // Create gRPC connection with TLS
    creds := credentials.NewTLS(tlsConfig)
    return grpc.Dial(serverAddr, grpc.WithTransportCredentials(creds))
}

// JWT Token Validation
type JWTValidator struct {
    publicKey *rsa.PublicKey
}

func (v *JWTValidator) ValidateToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return v.publicKey, nil
    })
    
    if err != nil {
        return nil, err
    }
    
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }
    
    return nil, fmt.Errorf("invalid token")
}

// API Gateway Security Middleware
func SecurityMiddleware(validator *JWTValidator) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Extract token from header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(401, gin.H{"error": "missing authorization header"})
            c.Abort()
            return
        }
        
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        
        // Validate token
        claims, err := validator.ValidateToken(tokenString)
        if err != nil {
            c.JSON(401, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }
        
        // Add claims to context
        c.Set("user_id", claims.UserID)
        c.Set("roles", claims.Roles)
        
        // Rate limiting per user
        if !rateLimiter.Allow(claims.UserID) {
            c.JSON(429, gin.H{"error": "rate limit exceeded"})
            c.Abort()
            return
        }
        
        c.Next()
    }
}

// Service Mesh Security Policy (Istio)
apiVersion: security.istio.io/v1beta1
kind: PeerAuthentication
metadata:
  name: default
  namespace: microservices
spec:
  mtls:
    mode: STRICT  # Enforce mTLS

---
apiVersion: security.istio.io/v1beta1
kind: AuthorizationPolicy
metadata:
  name: order-service-authz
  namespace: microservices
spec:
  selector:
    matchLabels:
      app: order-service
  action: ALLOW
  rules:
  - from:
    - source:
        principals: ["cluster.local/ns/microservices/sa/api-gateway"]
    to:
    - operation:
        methods: ["GET", "POST"]
        paths: ["/api/v1/orders/*"]
```

---

## Slide 19: Best Practices

### Microservices Best Practices

**✅ DO's:**

1. **Start with a Modular Monolith**
   - Identify boundaries first
   - Extract services gradually
   - Learn from the domain

2. **Design for Failure**
```go
// Always implement timeouts
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Always have circuit breakers
result, err := circuitBreaker.Execute(func() (interface{}, error) {
    return client.Call(ctx, request)
})

// Always have fallbacks
if err != nil {
    return getFallbackResponse()
}
```

3. **Implement Comprehensive Monitoring**
   - Distributed tracing (Jaeger/Zipkin)
   - Metrics (Prometheus)
   - Centralized logging (ELK Stack)
   - Health checks

4. **Use API Versioning**
```go
// URL versioning
router.GET("/api/v1/users/:id", getUserV1)
router.GET("/api/v2/users/:id", getUserV2)

// Header versioning
version := c.GetHeader("API-Version")
switch version {
case "2.0":
    return getUserV2(c)
default:
    return getUserV1(c)
}
```

5. **Implement Idempotency**
```go
type IdempotencyMiddleware struct {
    cache Cache
}

func (m *IdempotencyMiddleware) Handle(next HandlerFunc) HandlerFunc {
    return func(c *Context) error {
        idempotencyKey := c.GetHeader("Idempotency-Key")
        if idempotencyKey == "" {
            return next(c)
        }
        
        // Check cache
        if result, exists := m.cache.Get(idempotencyKey); exists {
            return c.JSON(200, result)
        }
        
        // Process request
        err := next(c)
        if err == nil {
            m.cache.Set(idempotencyKey, c.Response, 24*time.Hour)
        }
        
        return err
    }
}
```

---

## Slide 20: Anti-Patterns to Avoid

### Common Microservices Anti-Patterns

**❌ DON'Ts:**

1. **Distributed Monolith**
```go
// ❌ Bad: Services tightly coupled
type OrderService struct {
    userService    *UserService    // Direct dependency
    paymentService *PaymentService // Direct dependency
}

// ✅ Good: Loose coupling through interfaces
type OrderService struct {
    userClient    UserClient    // Interface
    paymentClient PaymentClient // Interface
}
```

2. **Chatty Services**
```go
// ❌ Bad: Multiple calls for single operation
func GetOrderDetails(orderID string) (*OrderDetails, error) {
    order := orderService.GetOrder(orderID)
    user := userService.GetUser(order.UserID)
    
    for _, item := range order.Items {
        product := productService.GetProduct(item.ProductID)
        // N+1 problem!
    }
}

// ✅ Good: Batch operations or data aggregation
func GetOrderDetails(orderID string) (*OrderDetails, error) {
    // Single call with all needed data
    return orderService.GetOrderDetailsWithProducts(orderID)
}
```

3. **Shared Database**
```go
// ❌ Bad: Multiple services sharing database
SELECT * FROM shared_db.users WHERE id = ?
SELECT * FROM shared_db.orders WHERE user_id = ?

// ✅ Good: Each service owns its data
// User Service: user_db.users
// Order Service: order_db.orders
```

4. **Synchronous Communication Everywhere**
```go
// ❌ Bad: Everything is synchronous
result1 := service1.Call()
result2 := service2.Call()  // Blocks
result3 := service3.Call()  // Blocks

// ✅ Good: Async where possible
eventBus.Publish(Event{})  // Fire and forget
go processInBackground()    // Async processing
```

5. **Missing Service Boundaries**
```go
// ❌ Bad: Nano-services
services := []string{
    "user-validation-service",
    "user-creation-service",
    "user-update-service",
    "user-deletion-service",
    // Too granular!
}

// ✅ Good: Cohesive services
services := []string{
    "user-service",  // Handles all user operations
}
```

---

## Slide 21: Performance Optimization

### Performance in Microservices

```go
// performance/optimization.go

// 1. Connection Pooling
type ServicePool struct {
    connections chan *grpc.ClientConn
    factory     func() (*grpc.ClientConn, error)
}

func NewServicePool(size int, factory func() (*grpc.ClientConn, error)) *ServicePool {
    pool := &ServicePool{
        connections: make(chan *grpc.ClientConn, size),
        factory:     factory,
    }
    
    // Pre-warm pool
    for i := 0; i < size; i++ {
        conn, _ := factory()
        pool.connections <- conn
    }
    
    return pool
}

func (p *ServicePool) Get() *grpc.ClientConn {
    return <-p.connections
}

func (p *ServicePool) Put(conn *grpc.ClientConn) {
    p.connections <- conn
}

// 2. Caching Strategy
type CacheAside struct {
    cache Cache
    db    Database
}

func (c *CacheAside) Get(ctx context.Context, key string) (interface{}, error) {
    // Check cache
    if val, found := c.cache.Get(key); found {
        return val, nil
    }
    
    // Load from database
    val, err := c.db.Get(ctx, key)
    if err != nil {
        return nil, err
    }
    
    // Update cache
    c.cache.Set(key, val, 5*time.Minute)
    
    return val, nil
}

// 3. Batch Processing
type BatchProcessor struct {
    batchSize    int
    batchTimeout time.Duration
    processor    func([]Request) []Response
}

func (b *BatchProcessor) Process(requests <-chan Request) <-chan Response {
    responses := make(chan Response)
    
    go func() {
        batch := make([]Request, 0, b.batchSize)
        timer := time.NewTimer(b.batchTimeout)
        
        for {
            select {
            case req := <-requests:
                batch = append(batch, req)
                
                if len(batch) >= b.batchSize {
                    b.processBatch(batch, responses)
                    batch = batch[:0]
                    timer.Reset(b.batchTimeout)
                }
                
            case <-timer.C:
                if len(batch) > 0 {
                    b.processBatch(batch, responses)
                    batch = batch[:0]
                }
                timer.Reset(b.batchTimeout)
            }
        }
    }()
    
    return responses
}

// 4. GraphQL Federation for Reducing Calls
type GraphQLGateway struct {
    services map[string]GraphQLService
}

func (g *GraphQLGateway) ResolveQuery(query string) (interface{}, error) {
    // Parse query and determine required services
    requiredServices :=