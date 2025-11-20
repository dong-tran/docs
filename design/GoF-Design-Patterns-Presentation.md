# Gang of Four Design Patterns
## 23 Classic Solutions to Recurring Design Problems

**Presenter:** Dong Tran  
**Date:** November 20, 2025  
**Time:** 03:12 UTC

---

## Slide 1: Title Slide

# Gang of Four Design Patterns
### Timeless Solutions for Object-Oriented Design

**Presenter:** Dong Tran  
**Date:** November 20, 2025

*"Design patterns are typical solutions to common problems in software design. Each pattern is like a blueprint that you can customize to solve a particular design problem in your code."*  
— Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides

---

## Slide 2: Agenda

### Our Journey Through Design Patterns

1. **Introduction to Design Patterns**
2. **History and the Gang of Four**
3. **Pattern Categories Overview**
4. **Creational Patterns** (5 patterns)
   - Singleton, Factory Method, Abstract Factory, Builder, Prototype
5. **Structural Patterns** (7 patterns)
   - Adapter, Bridge, Composite, Decorator, Facade, Flyweight, Proxy
6. **Behavioral Patterns** (11 patterns)
   - Chain of Responsibility, Command, Iterator, Mediator, Memento
   - Observer, State, Strategy, Template Method, Visitor, Interpreter
7. **Patterns in Modern Go**
8. **Real-World Case Study**
9. **Pattern Selection Guide**
10. **Anti-Patterns to Avoid**
11. **Q&A Session**

---

## Slide 3: What Are Design Patterns?

### Understanding Design Patterns

**Definition:**
Reusable solutions to commonly occurring problems in software design

**What Patterns Are:**
- ✅ Templates for solving problems
- ✅ Communication tools between developers
- ✅ Best practices distilled from experience
- ✅ Language-agnostic concepts

**What Patterns Are NOT:**
- ❌ Copy-paste code solutions
- ❌ Silver bullets for all problems
- ❌ Mandatory for good design
- ❌ Substitutes for good architecture

**Pattern Structure:**
```
┌─────────────────────────────┐
│      Pattern Name           │ ← Intent
├─────────────────────────────┤
│      Problem Context        │ ← When to use
├─────────────────────────────┤
│      Solution Structure     │ ← How it works
├─────────────────────────────┤
│      Consequences          │ ← Trade-offs
└─────────────────────────────┘
```

---

## Slide 4: The Gang of Four

### History and Background

**The Book (1994):**
"Design Patterns: Elements of Reusable Object-Oriented Software"

**The Authors (Gang of Four):**
- **Erich Gamma** - JUnit, Eclipse
- **Richard Helm** - IBM Research
- **Ralph Johnson** - University of Illinois
- **John Vlissides** - IBM Research (1961-2005)

**Pattern Categories:**
```
         Design Patterns (23)
                │
    ┌───────────┼───────────┐
    │           │           │
Creational  Structural  Behavioral
    (5)         (7)         (11)
    │           │           │
  Object     Object      Object
 Creation   Composition Interaction
```

**Impact:**
- Foundation of modern OOP design
- Common vocabulary for developers
- Influenced all modern languages
- Basis for enterprise patterns

---

## Slide 5: Pattern Categories Overview

### Three Categories of GoF Patterns

```go
// 1. CREATIONAL PATTERNS - Object Creation
// "How objects are created"
type CreationalPattern interface {
    // Controls object instantiation process
    // Examples: Singleton, Factory, Builder
    Create() interface{}
}

// 2. STRUCTURAL PATTERNS - Object Composition
// "How objects are composed"
type StructuralPattern interface {
    // Defines relationships between objects
    // Examples: Adapter, Decorator, Proxy
    Structure() interface{}
}

// 3. BEHAVIORAL PATTERNS - Object Collaboration
// "How objects interact"
type BehavioralPattern interface {
    // Defines communication between objects
    // Examples: Observer, Strategy, Command
    Behave() interface{}
}
```

**Quick Reference:**
| Category | Focus | Question | Patterns Count |
|----------|-------|----------|----------------|
| Creational | Object creation | "How to create?" | 5 |
| Structural | Object composition | "How to structure?" | 7 |
| Behavioral | Object collaboration | "How to behave?" | 11 |

---

## Slide 6: Singleton Pattern

### Creational Pattern #1: Singleton

**Intent:** Ensure a class has only one instance and provide global access to it

**When to Use:**
- Exactly one instance needed (database connection, logger)
- Global access point required
- Lazy initialization desired

**Implementation in Go:**
```go
package singleton

import (
    "sync"
    "database/sql"
    "fmt"
)

// Classic Singleton with sync.Once (Thread-safe)
type Database struct {
    connection *sql.DB
    name       string
}

var (
    instance *Database
    once     sync.Once
)

// GetInstance returns the singleton instance
func GetInstance() *Database {
    once.Do(func() {
        fmt.Println("Creating singleton instance")
        conn, _ := sql.Open("postgres", "connection_string")
        instance = &Database{
            connection: conn,
            name:       "MainDB",
        }
    })
    return instance
}

func (db *Database) Query(sql string) ([]interface{}, error) {
    // Execute query
    fmt.Printf("Executing on %s: %s\n", db.name, sql)
    return nil, nil
}

// Better approach in Go - using package initialization
package config

var (
    Config *AppConfig
)

type AppConfig struct {
    DatabaseURL string
    APIKey      string
    MaxRetries  int
    mu          sync.RWMutex
}

func init() {
    Config = &AppConfig{
        DatabaseURL: getEnv("DATABASE_URL", "localhost:5432"),
        APIKey:      getEnv("API_KEY", ""),
        MaxRetries:  3,
    }
}

func (c *AppConfig) Get(key string) string {
    c.mu.RLock()
    defer c.mu.RUnlock()
    // Return config value
    return ""
}

func (c *AppConfig) Set(key, value string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    // Set config value
}

// Usage
func main() {
    db1 := GetInstance()
    db2 := GetInstance()
    
    fmt.Println(db1 == db2) // true - same instance
    
    // Using config singleton
    apiKey := config.Config.Get("APIKey")
}
```

**Pros & Cons:**
- ✅ Controlled access to single instance
- ✅ Lazy initialization
- ✅ Global access point
- ❌ Difficult to test (global state)
- ❌ Violates Single Responsibility Principle
- ❌ Hidden dependencies

---

## Slide 7: Factory Method Pattern

### Creational Pattern #2: Factory Method

**Intent:** Define interface for creating objects, let subclasses decide which class to instantiate

**When to Use:**
- Class can't anticipate the class of objects it must create
- Class wants subclasses to specify objects it creates
- Classes delegate responsibility to helper subclasses

**Implementation in Go:**
```go
package factory

import "fmt"

// Product interface
type Vehicle interface {
    Drive() string
    GetType() string
}

// Concrete Products
type Car struct {
    model string
}

func (c *Car) Drive() string {
    return fmt.Sprintf("Driving car: %s", c.model)
}

func (c *Car) GetType() string {
    return "Car"
}

type Motorcycle struct {
    model string
}

func (m *Motorcycle) Drive() string {
    return fmt.Sprintf("Riding motorcycle: %s", m.model)
}

func (m *Motorcycle) GetType() string {
    return "Motorcycle"
}

type Truck struct {
    model string
    capacity int
}

func (t *Truck) Drive() string {
    return fmt.Sprintf("Driving truck: %s with capacity %d tons", t.model, t.capacity)
}

func (t *Truck) GetType() string {
    return "Truck"
}

// Factory Method
type VehicleFactory interface {
    CreateVehicle(model string) Vehicle
}

// Concrete Factories
type CarFactory struct{}

func (f *CarFactory) CreateVehicle(model string) Vehicle {
    return &Car{model: model}
}

type MotorcycleFactory struct{}

func (f *MotorcycleFactory) CreateVehicle(model string) Vehicle {
    return &Motorcycle{model: model}
}

type TruckFactory struct{}

func (f *TruckFactory) CreateVehicle(model string) Vehicle {
    return &Truck{model: model, capacity: 10}
}

// Factory Registry Pattern (Go idiomatic)
type VehicleType string

const (
    CarType        VehicleType = "car"
    MotorcycleType VehicleType = "motorcycle"
    TruckType      VehicleType = "truck"
)

var factories = map[VehicleType]VehicleFactory{
    CarType:        &CarFactory{},
    MotorcycleType: &MotorcycleFactory{},
    TruckType:      &TruckFactory{},
}

func CreateVehicle(vehicleType VehicleType, model string) (Vehicle, error) {
    factory, exists := factories[vehicleType]
    if !exists {
        return nil, fmt.Errorf("unknown vehicle type: %s", vehicleType)
    }
    return factory.CreateVehicle(model), nil
}

// Real-world example: Database drivers
type DatabaseConnection interface {
    Connect() error
    Query(sql string) ([]interface{}, error)
    Close() error
}

type PostgresConnection struct {
    connectionString string
}

func (p *PostgresConnection) Connect() error {
    fmt.Printf("Connecting to PostgreSQL: %s\n", p.connectionString)
    return nil
}

func (p *PostgresConnection) Query(sql string) ([]interface{}, error) {
    return nil, nil
}

func (p *PostgresConnection) Close() error {
    return nil
}

type MySQLConnection struct {
    connectionString string
}

func (m *MySQLConnection) Connect() error {
    fmt.Printf("Connecting to MySQL: %s\n", m.connectionString)
    return nil
}

func (m *MySQLConnection) Query(sql string) ([]interface{}, error) {
    return nil, nil
}

func (m *MySQLConnection) Close() error {
    return nil
}

// Database Factory
func NewDatabaseConnection(dbType, connectionString string) (DatabaseConnection, error) {
    switch dbType {
    case "postgres":
        return &PostgresConnection{connectionString: connectionString}, nil
    case "mysql":
        return &MySQLConnection{connectionString: connectionString}, nil
    default:
        return nil, fmt.Errorf("unsupported database type: %s", dbType)
    }
}

// Usage
func main() {
    // Using vehicle factory
    car, _ := CreateVehicle(CarType, "Tesla Model 3")
    fmt.Println(car.Drive())
    
    // Using database factory
    db, _ := NewDatabaseConnection("postgres", "localhost:5432/mydb")
    db.Connect()
    defer db.Close()
}
```

---

## Slide 8: Abstract Factory Pattern

### Creational Pattern #3: Abstract Factory

**Intent:** Provide interface for creating families of related objects

**When to Use:**
- System should be independent of product creation
- System should be configured with multiple families of products
- Family of products must be used together

**Implementation in Go:**
```go
package abstractfactory

// Abstract Products
type Button interface {
    Click() string
}

type Checkbox interface {
    Check() string
}

type Window interface {
    Open() string
}

// Concrete Products - Material Design
type MaterialButton struct{}
func (b *MaterialButton) Click() string { return "Material Button clicked" }

type MaterialCheckbox struct{}
func (c *MaterialCheckbox) Check() string { return "Material Checkbox checked" }

type MaterialWindow struct{}
func (w *MaterialWindow) Open() string { return "Material Window opened" }

// Concrete Products - iOS Style
type IOSButton struct{}
func (b *IOSButton) Click() string { return "iOS Button tapped" }

type IOSCheckbox struct{}
func (c *IOSCheckbox) Check() string { return "iOS Switch toggled" }

type IOSWindow struct{}
func (w *IOSWindow) Open() string { return "iOS View presented" }

// Abstract Factory
type UIFactory interface {
    CreateButton() Button
    CreateCheckbox() Checkbox
    CreateWindow() Window
}

// Concrete Factories
type MaterialUIFactory struct{}

func (f *MaterialUIFactory) CreateButton() Button {
    return &MaterialButton{}
}

func (f *MaterialUIFactory) CreateCheckbox() Checkbox {
    return &MaterialCheckbox{}
}

func (f *MaterialUIFactory) CreateWindow() Window {
    return &MaterialWindow{}
}

type IOSUIFactory struct{}

func (f *IOSUIFactory) CreateButton() Button {
    return &IOSButton{}
}

func (f *IOSUIFactory) CreateCheckbox() Checkbox {
    return &IOSCheckbox{}
}

func (f *IOSUIFactory) CreateWindow() Window {
    return &IOSWindow{}
}

// Client code
type Application struct {
    factory UIFactory
}

func NewApplication(factory UIFactory) *Application {
    return &Application{factory: factory}
}

func (app *Application) CreateUI() {
    button := app.factory.CreateButton()
    checkbox := app.factory.CreateCheckbox()
    window := app.factory.CreateWindow()
    
    println(button.Click())
    println(checkbox.Check())
    println(window.Open())
}

// Real-world example: Cloud Provider Services
type CloudProvider interface {
    CreateCompute() ComputeInstance
    CreateStorage() StorageService
    CreateNetwork() NetworkService
}

type ComputeInstance interface {
    Start() error
    Stop() error
    GetStatus() string
}

type StorageService interface {
    Upload(file []byte) error
    Download(key string) ([]byte, error)
}

type NetworkService interface {
    CreateVPC(cidr string) error
    CreateSubnet(vpc, cidr string) error
}

// AWS Implementation
type AWSProvider struct{}

func (a *AWSProvider) CreateCompute() ComputeInstance {
    return &EC2Instance{}
}

func (a *AWSProvider) CreateStorage() StorageService {
    return &S3Bucket{}
}

func (a *AWSProvider) CreateNetwork() NetworkService {
    return &VPCService{}
}

type EC2Instance struct {
    instanceID string
}

func (e *EC2Instance) Start() error {
    println("Starting EC2 instance")
    return nil
}

func (e *EC2Instance) Stop() error {
    println("Stopping EC2 instance")
    return nil
}

func (e *EC2Instance) GetStatus() string {
    return "running"
}

// Azure Implementation
type AzureProvider struct{}

func (a *AzureProvider) CreateCompute() ComputeInstance {
    return &AzureVM{}
}

func (a *AzureProvider) CreateStorage() StorageService {
    return &BlobStorage{}
}

func (a *AzureProvider) CreateNetwork() NetworkService {
    return &VNetService{}
}

// Usage
func DeployApplication(provider CloudProvider) {
    compute := provider.CreateCompute()
    storage := provider.CreateStorage()
    network := provider.CreateNetwork()
    
    network.CreateVPC("10.0.0.0/16")
    compute.Start()
    storage.Upload([]byte("application data"))
}
```

---

## Slide 9: Builder Pattern

### Creational Pattern #4: Builder

**Intent:** Separate construction of complex object from its representation

**When to Use:**
- Creating complex objects with many optional parameters
- Construction process must allow different representations
- Avoid "telescoping constructors"

**Implementation in Go:**
```go
package builder

import (
    "fmt"
    "strings"
)

// Product
type SQLQuery struct {
    selectClause []string
    fromClause   string
    whereClause  []string
    joinClause   []string
    orderBy      []string
    groupBy      []string
    limit        int
}

func (q *SQLQuery) String() string {
    var query strings.Builder
    
    // SELECT
    query.WriteString("SELECT ")
    if len(q.selectClause) > 0 {
        query.WriteString(strings.Join(q.selectClause, ", "))
    } else {
        query.WriteString("*")
    }
    
    // FROM
    query.WriteString(fmt.Sprintf(" FROM %s", q.fromClause))
    
    // JOIN
    for _, join := range q.joinClause {
        query.WriteString(fmt.Sprintf(" %s", join))
    }
    
    // WHERE
    if len(q.whereClause) > 0 {
        query.WriteString(" WHERE ")
        query.WriteString(strings.Join(q.whereClause, " AND "))
    }
    
    // GROUP BY
    if len(q.groupBy) > 0 {
        query.WriteString(" GROUP BY ")
        query.WriteString(strings.Join(q.groupBy, ", "))
    }
    
    // ORDER BY
    if len(q.orderBy) > 0 {
        query.WriteString(" ORDER BY ")
        query.WriteString(strings.Join(q.orderBy, ", "))
    }
    
    // LIMIT
    if q.limit > 0 {
        query.WriteString(fmt.Sprintf(" LIMIT %d", q.limit))
    }
    
    return query.String()
}

// Builder
type QueryBuilder struct {
    query *SQLQuery
}

func NewQueryBuilder() *QueryBuilder {
    return &QueryBuilder{
        query: &SQLQuery{},
    }
}

func (b *QueryBuilder) Select(columns ...string) *QueryBuilder {
    b.query.selectClause = append(b.query.selectClause, columns...)
    return b
}

func (b *QueryBuilder) From(table string) *QueryBuilder {
    b.query.fromClause = table
    return b
}

func (b *QueryBuilder) Where(condition string) *QueryBuilder {
    b.query.whereClause = append(b.query.whereClause, condition)
    return b
}

func (b *QueryBuilder) Join(join string) *QueryBuilder {
    b.query.joinClause = append(b.query.joinClause, join)
    return b
}

func (b *QueryBuilder) OrderBy(column string) *QueryBuilder {
    b.query.orderBy = append(b.query.orderBy, column)
    return b
}

func (b *QueryBuilder) GroupBy(column string) *QueryBuilder {
    b.query.groupBy = append(b.query.groupBy, column)
    return b
}

func (b *QueryBuilder) Limit(limit int) *QueryBuilder {
    b.query.limit = limit
    return b
}

func (b *QueryBuilder) Build() *SQLQuery {
    return b.query
}

// Real-world example: HTTP Request Builder
type HTTPRequest struct {
    method  string
    url     string
    headers map[string]string
    params  map[string]string
    body    []byte
    timeout int
}

type RequestBuilder struct {
    request *HTTPRequest
}

func NewRequestBuilder() *RequestBuilder {
    return &RequestBuilder{
        request: &HTTPRequest{
            headers: make(map[string]string),
            params:  make(map[string]string),
        },
    }
}

func (b *RequestBuilder) Method(method string) *RequestBuilder {
    b.request.method = method
    return b
}

func (b *RequestBuilder) URL(url string) *RequestBuilder {
    b.request.url = url
    return b
}

func (b *RequestBuilder) Header(key, value string) *RequestBuilder {
    b.request.headers[key] = value
    return b
}

func (b *RequestBuilder) Param(key, value string) *RequestBuilder {
    b.request.params[key] = value
    return b
}

func (b *RequestBuilder) Body(body []byte) *RequestBuilder {
    b.request.body = body
    return b
}

func (b *RequestBuilder) Timeout(seconds int) *RequestBuilder {
    b.request.timeout = seconds
    return b
}

func (b *RequestBuilder) Build() (*HTTPRequest, error) {
    if b.request.method == "" {
        b.request.method = "GET"
    }
    if b.request.url == "" {
        return nil, fmt.Errorf("URL is required")
    }
    return b.request, nil
}

// Configuration Builder
type ServerConfig struct {
    host        string
    port        int
    useTLS      bool
    certFile    string
    keyFile     string
    maxConn     int
    timeout     int
    middlewares []string
}

type ServerBuilder struct {
    config *ServerConfig
}

func NewServerBuilder() *ServerBuilder {
    return &ServerBuilder{
        config: &ServerConfig{
            port:    8080,
            maxConn: 100,
            timeout: 30,
        },
    }
}

func (b *ServerBuilder) Host(host string) *ServerBuilder {
    b.config.host = host
    return b
}

func (b *ServerBuilder) Port(port int) *ServerBuilder {
    b.config.port = port
    return b
}

func (b *ServerBuilder) WithTLS(certFile, keyFile string) *ServerBuilder {
    b.config.useTLS = true
    b.config.certFile = certFile
    b.config.keyFile = keyFile
    return b
}

func (b *ServerBuilder) MaxConnections(max int) *ServerBuilder {
    b.config.maxConn = max
    return b
}

func (b *ServerBuilder) Middleware(name string) *ServerBuilder {
    b.config.middlewares = append(b.config.middlewares, name)
    return b
}

func (b *ServerBuilder) Build() *ServerConfig {
    return b.config
}

// Usage
func main() {
    // SQL Query Builder
    query := NewQueryBuilder().
        Select("users.name", "orders.total").
        From("users").
        Join("JOIN orders ON users.id = orders.user_id").
        Where("orders.total > 100").
        Where("users.active = true").
        OrderBy("orders.total DESC").
        Limit(10).
        Build()
    
    fmt.Println(query.String())
    
    // HTTP Request Builder
    request, _ := NewRequestBuilder().
        Method("POST").
        URL("https://api.example.com/users").
        Header("Content-Type", "application/json").
        Header("Authorization", "Bearer token").
        Body([]byte(`{"name":"John"}`)).
        Timeout(10).
        Build()
    
    // Server Config Builder
    config := NewServerBuilder().
        Host("localhost").
        Port(443).
        WithTLS("cert.pem", "key.pem").
        Middleware("auth").
        Middleware("logging").
        Middleware("cors").
        Build()
}
```

---

## Slide 10: Prototype Pattern

### Creational Pattern #5: Prototype

**Intent:** Create new objects by cloning existing instance

**When to Use:**
- Creating objects is expensive
- Need variations of objects
- Avoid subclasses of object creator

**Implementation in Go:**
```go
package prototype

import (
    "bytes"
    "encoding/gob"
    "fmt"
)

// Prototype interface
type Prototype interface {
    Clone() Prototype
}

// Document prototype
type Document struct {
    Title    string
    Content  string
    Author   string
    Comments []string
    Metadata map[string]string
}

// Shallow clone
func (d *Document) Clone() Prototype {
    cloned := *d
    // Need to manually copy slices and maps for deep clone
    cloned.Comments = make([]string, len(d.Comments))
    copy(cloned.Comments, d.Comments)
    
    cloned.Metadata = make(map[string]string)
    for k, v := range d.Metadata {
        cloned.Metadata[k] = v
    }
    
    return &cloned
}

// Deep clone using gob
func (d *Document) DeepClone() (*Document, error) {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    dec := gob.NewDecoder(&buf)
    
    err := enc.Encode(d)
    if err != nil {
        return nil, err
    }
    
    var cloned Document
    err = dec.Decode(&cloned)
    if err != nil {
        return nil, err
    }
    
    return &cloned, nil
}

// Cell prototype for spreadsheet
type Cell struct {
    Value   string
    Formula string
    Style   *CellStyle
}

type CellStyle struct {
    Font      string
    FontSize  int
    Bold      bool
    TextColor string
    BGColor   string
}

func (c *Cell) Clone() *Cell {
    cloned := *c
    if c.Style != nil {
        styleCopy := *c.Style
        cloned.Style = &styleCopy
    }
    return &cloned
}

// Shape prototype for graphics editor
type Shape interface {
    Clone() Shape
    GetInfo() string
}

type Rectangle struct {
    X      int
    Y      int
    Width  int
    Height int
    Color  string
}

func (r *Rectangle) Clone() Shape {
    return &Rectangle{
        X:      r.X,
        Y:      r.Y,
        Width:  r.Width,
        Height: r.Height,
        Color:  r.Color,
    }
}

func (r *Rectangle) GetInfo() string {
    return fmt.Sprintf("Rectangle at (%d,%d) with size %dx%d, color: %s",
        r.X, r.Y, r.Width, r.Height, r.Color)
}

type Circle struct {
    X      int
    Y      int
    Radius int
    Color  string
}

func (c *Circle) Clone() Shape {
    return &Circle{
        X:      c.X,
        Y:      c.Y,
        Radius: c.Radius,
        Color:  c.Color,
    }
}

func (c *Circle) GetInfo() string {
    return fmt.Sprintf("Circle at (%d,%d) with radius %d, color: %s",
        c.X, c.Y, c.Radius, c.Color)
}

// Prototype Registry
type ShapeRegistry struct {
    shapes map[string]Shape
}

func NewShapeRegistry() *ShapeRegistry {
    return &ShapeRegistry{
        shapes: make(map[string]Shape),
    }
}

func (r *ShapeRegistry) Register(name string, shape Shape) {
    r.shapes[name] = shape
}

func (r *ShapeRegistry) Create(name string) (Shape, error) {
    prototype, exists := r.shapes[name]
    if !exists {
        return nil, fmt.Errorf("prototype %s not found", name)
    }
    return prototype.Clone(), nil
}

// Complex object with nested structures
type GameCharacter struct {
    Name       string
    Level      int
    Health     int
    Inventory  []Item
    Skills     map[string]int
    Position   Position
    Equipment  *Equipment
}

type Item struct {
    Name   string
    Weight int
}

type Position struct {
    X, Y, Z float64
}

type Equipment struct {
    Weapon string
    Armor  string
}

func (g *GameCharacter) Clone() *GameCharacter {
    cloned := *g
    
    // Deep copy inventory
    cloned.Inventory = make([]Item, len(g.Inventory))
    copy(cloned.Inventory, g.Inventory)
    
    // Deep copy skills
    cloned.Skills = make(map[string]int)
    for k, v := range g.Skills {
        cloned.Skills[k] = v
    }
    
    // Deep copy equipment
    if g.Equipment != nil {
        equipCopy := *g.Equipment
        cloned.Equipment = &equipCopy
    }
    
    return &cloned
}

// Usage
func main() {
    // Document cloning
    original := &Document{
        Title:    "Design Patterns",
        Content:  "Gang of Four patterns...",
        Author:   "GoF",
        Comments: []string{"Great!", "Useful"},
        Metadata: map[string]string{"version": "1.0"},
    }
    
    cloned := original.Clone().(*Document)
    cloned.Title = "Design Patterns - Copy"
    
    // Shape registry
    registry := NewShapeRegistry()
    registry.Register("red-rectangle", &Rectangle{
        Width: 100, Height: 50, Color: "red",
    })
    registry.Register("blue-circle", &Circle{
        Radius: 30, Color: "blue",
    })
    
    // Create shapes from prototypes
    shape1, _ := registry.Create("red-rectangle")
    shape2, _ := registry.Create("blue-circle")
    
    fmt.Println(shape1.GetInfo())
    fmt.Println(shape2.GetInfo())
    
    // Game character template
    warriorTemplate := &GameCharacter{
        Name:   "Warrior",
        Level:  1,
        Health: 100,
        Skills: map[string]int{"strength": 10, "defense": 8},
        Equipment: &Equipment{
            Weapon: "Sword",
            Armor:  "Plate Mail",
        },
    }
    
    player1 := warriorTemplate.Clone()
    player1.Name = "Conan"
    
    player2 := warriorTemplate.Clone()
    player2.Name = "Aragorn"
}
```

---

## Slide 11: Adapter Pattern

### Structural Pattern #1: Adapter

**Intent:** Convert interface of a class into another interface clients expect

**When to Use:**
- Want to use existing class with incompatible interface
- Create reusable class that cooperates with unrelated classes
- Need to use several existing subclasses

**Implementation in Go:**
```go
package adapter

import (
    "fmt"
    "strings"
)

// Target interface (what client expects)
type PaymentProcessor interface {
    ProcessPayment(amount float64) error
}

// Adaptee (existing implementation with different interface)
type LegacyPaymentSystem struct {
    merchantID string
}

func (l *LegacyPaymentSystem) MakePayment(merchantID string, amount float64) string {
    return fmt.Sprintf("Payment of $%.2f processed for merchant %s", amount, merchantID)
}

func (l *LegacyPaymentSystem) GetTransactionStatus(transactionID string) string {
    return "SUCCESS"
}

// Adapter
type LegacyPaymentAdapter struct {
    legacySystem *LegacyPaymentSystem
    merchantID   string
}

func NewLegacyPaymentAdapter(merchantID string) *LegacyPaymentAdapter {
    return &LegacyPaymentAdapter{
        legacySystem: &LegacyPaymentSystem{merchantID: merchantID},
        merchantID:   merchantID,
    }
}

func (a *LegacyPaymentAdapter) ProcessPayment(amount float64) error {
    result := a.legacySystem.MakePayment(a.merchantID, amount)
    if strings.Contains(result, "processed") {
        return nil
    }
    return fmt.Errorf("payment failed")
}

// Real-world example: Different data sources
type DataAnalyzer interface {
    Analyze(data []float64) AnalysisResult
}

type AnalysisResult struct {
    Mean   float64
    Median float64
    StdDev float64
}

// External library with different interface
type StatisticalLibrary struct{}

func (s *StatisticalLibrary) CalculateMean(numbers []float64) float64 {
    sum := 0.0
    for _, n := range numbers {
        sum += n
    }
    return sum / float64(len(numbers))
}

func (s *StatisticalLibrary) CalculateMedian(numbers []float64) float64 {
    // Simplified median calculation
    if len(numbers) == 0 {
        return 0
    }
    return numbers[len(numbers)/2]
}

func (s *StatisticalLibrary) CalculateStandardDeviation(numbers []float64) float64 {
    // Simplified std dev calculation
    return 1.5
}

// Adapter for statistical library
type StatisticalAdapter struct {
    library *StatisticalLibrary
}

func NewStatisticalAdapter() *StatisticalAdapter {
    return &StatisticalAdapter{
        library: &StatisticalLibrary{},
    }
}

func (a *StatisticalAdapter) Analyze(data []float64) AnalysisResult {
    return AnalysisResult{
        Mean:   a.library.CalculateMean(data),
        Median: a.library.CalculateMedian(data),
        StdDev: a.library.CalculateStandardDeviation(data),
    }
}

// Multiple adaptees example: Different cloud storage providers
type CloudStorage interface {
    Upload(key string, data []byte) error
    Download(key string) ([]byte, error)
    Delete(key string) error
}

// AWS S3 (Adaptee 1)
type S3Client struct{}

func (s *S3Client) PutObject(bucket, key string, data []byte) error {
    fmt.Printf("S3: Uploading to %s/%s\n", bucket, key)
    return nil
}

func (s *S3Client) GetObject(bucket, key string) ([]byte, error) {
    fmt.Printf("S3: Downloading from %s/%s\n", bucket, key)
    return []byte("s3 data"), nil
}

func (s *S3Client) DeleteObject(bucket, key string) error {
    fmt.Printf("S3: Deleting %s/%s\n", bucket, key)
    return nil
}

// S3 Adapter
type S3Adapter struct {
    client *S3Client
    bucket string
}

func NewS3Adapter(bucket string) *S3Adapter {
    return &S3Adapter{
        client: &S3Client{},
        bucket: bucket,
    }
}

func (a *S3Adapter) Upload(key string, data []byte) error {
    return a.client.PutObject(a.bucket, key, data)
}

func (a *S3Adapter) Download(key string) ([]byte, error) {
    return a.client.GetObject(a.bucket, key)
}

func (a *S3Adapter) Delete(key string) error {
    return a.client.DeleteObject(a.bucket, key)
}

// Azure Blob Storage (Adaptee 2)
type AzureBlobClient struct{}

func (a *AzureBlobClient) UploadBlob(container, blob string, content []byte) error {
    fmt.Printf("Azure: Uploading to %s/%s\n", container, blob)
    return nil
}

func (a *AzureBlobClient) DownloadBlob(container, blob string) ([]byte, error) {
    fmt.Printf("Azure: Downloading from %s/%s\n", container, blob)
    return []byte("azure data"), nil
}

func (a *AzureBlobClient) RemoveBlob(container, blob string) error {
    fmt.Printf("Azure: Removing %s/%s\n", container, blob)
    return nil
}

// Azure Adapter
type AzureAdapter struct {
    client    *AzureBlobClient
    container string
}

func NewAzureAdapter(container string) *AzureAdapter {
    return &AzureAdapter{
        client:    &AzureBlobClient{},
        container: container,
    }
}

func (a *AzureAdapter) Upload(key string, data []byte) error {
    return a.client.UploadBlob(a.container, key, data)
}

func (a *AzureAdapter) Download(key string) ([]byte, error) {
    return a.client.DownloadBlob(a.container, key)
}

func (a *AzureAdapter) Delete(key string) error {
    return a.client.RemoveBlob(a.container, key)
}

// Client code works with any adapter
func ProcessFiles(storage CloudStorage) {
    storage.Upload("file.txt", []byte("content"))
    data, _ := storage.Download("file.txt")
    fmt.Printf("Downloaded: %s\n", data)
    storage.Delete("file.txt")
}

// Usage
func main() {
    // Payment adapter
    payment := NewLegacyPaymentAdapter("MERCHANT123")
    err := payment.ProcessPayment(99.99)
    if err != nil {
        fmt.Printf("Payment failed: %v\n", err)
    }
    
    // Statistical adapter
    analyzer := NewStatisticalAdapter()
    result := analyzer.Analyze([]float64{1, 2, 3, 4, 5})
    fmt.Printf("Analysis: %+v\n", result)
    
    // Cloud storage adapters
    s3Storage := NewS3Adapter("my-bucket")
    ProcessFiles(s3Storage)
    
    azureStorage := NewAzureAdapter("my-container")
    ProcessFiles(azureStorage)
}
```

---

## Slide 12: Bridge Pattern

### Structural Pattern #2: Bridge

**Intent:** Decouple abstraction from implementation so both can vary independently

**When to Use:**
- Want to avoid permanent binding between abstraction and implementation
- Both abstraction and implementation should be extensible
- Changes in implementation shouldn't affect clients

**Implementation in Go:**
```go
package bridge

import "fmt"

// Implementation interface
type MessageSender interface {
    SendMessage(title, body string) error
}

// Concrete Implementations
type EmailSender struct {
    emailAddress string
}

func (e *EmailSender) SendMessage(title, body string) error {
    fmt.Printf("Sending Email to %s:\nSubject: %s\nBody: %s\n", 
        e.emailAddress, title, body)
    return nil
}

type SMSSender struct {
    phoneNumber string
}

func (s *SMSSender) SendMessage(title, body string) error {
    fmt.Printf("Sending SMS to %s:\n%s: %s\n", 
        s.phoneNumber, title, body)
    return nil
}

type SlackSender struct {
    channel string
}

func (s *SlackSender) SendMessage(title, body string) error {
    fmt.Printf("Sending to Slack channel %s:\n*%s*\n%s\n", 
        s.channel, title, body)
    return nil
}

// Abstraction
type Notification struct {
    sender MessageSender
}

func NewNotification(sender MessageSender) *Notification {
    return &Notification{sender: sender}
}

func (n *Notification) Send(title, body string) error {
    return n.sender.SendMessage(title, body)
}

// Refined Abstractions
type UrgentNotification struct {
    Notification
}

func NewUrgentNotification(sender MessageSender) *UrgentNotification {
    return &UrgentNotification{
        Notification: Notification{sender: sender},
    }
}

func (u *UrgentNotification) Send(title, body string) error {
    urgentTitle := fmt.Sprintf("🚨 URGENT: %s", title)
    return u.sender.SendMessage(urgentTitle, body)
}

type ReminderNotification struct {
    Notification
}

func NewReminderNotification(sender MessageSender) *ReminderNotification {
    return &ReminderNotification{
        Notification: Notification{sender: sender},
    }
}

func (r *ReminderNotification) Send(title, body string) error {
    reminderTitle := fmt.Sprintf("⏰ Reminder: %s", title)
    reminderBody := fmt.Sprintf("%s\n\nThis is a reminder message.", body)
    return r.sender.SendMessage(reminderTitle, reminderBody)
}

// Real-world example: Remote Control and Devices
type Device interface {
    TurnOn()
    TurnOff()
    SetVolume(volume int)
    GetStatus() string
}

type TV struct {
    on     bool
    volume int
    channel int
}

func (t *TV) TurnOn() {
    t.on = true
    fmt.Println("TV is ON")
}

func (t *TV) TurnOff() {
    t.on = false
    fmt.Println("TV is OFF")
}

func (t *TV) SetVolume(volume int) {
    t.volume = volume
    fmt.Printf("TV volume set to %d\n", volume)
}

func (t *TV) GetStatus() string {
    if t.on {
        return fmt.Sprintf("TV is ON, volume: %d, channel: %d", t.volume, t.channel)
    }
    return "TV is OFF"
}

type Radio struct {
    on     bool
    volume int
    frequency float64
}

func (r *Radio) TurnOn() {
    r.on = true
    fmt.Println("Radio is ON")
}

func (r *Radio) TurnOff() {
    r.on = false
    fmt.Println("Radio is OFF")
}

func (r *Radio) SetVolume(volume int) {
    r.volume = volume
    fmt.Printf("Radio volume set to %d\n", volume)
}

func (r *Radio) GetStatus() string {
    if r.on {
        return fmt.Sprintf("Radio is ON, volume: %d, frequency: %.1f FM", 
            r.volume, r.frequency)
    }
    return "Radio is OFF"
}

// Remote Control abstraction
type RemoteControl struct {
    device Device
}

func NewRemoteControl(device Device) *RemoteControl {
    return &RemoteControl{device: device}
}

func (r *RemoteControl) PowerToggle() {
    status := r.device.GetStatus()
    if strings.Contains(status, "OFF") {
        r.device.TurnOn()
    } else {
        r.device.TurnOff()
    }
}

func (r *RemoteControl) VolumeUp() {
    // Get current volume and increase
    r.device.SetVolume(50) // Simplified
}

func (r *RemoteControl) VolumeDown() {
    r.device.SetVolume(30) // Simplified
}

// Advanced Remote with more features
type AdvancedRemoteControl struct {
    RemoteControl
}

func NewAdvancedRemoteControl(device Device) *AdvancedRemoteControl {
    return &AdvancedRemoteControl{
        RemoteControl: RemoteControl{device: device},
    }
}

func (a *AdvancedRemoteControl) Mute() {
    fmt.Println("Muting device")
    a.device.SetVolume(0)
}

// Database example: Different databases and operations
type Database interface {
    Connect(connectionString string) error
    ExecuteQuery(query string) ([]map[string]interface{}, error)
    Close() error
}

type PostgreSQL struct {
    connection string
}

func (p *PostgreSQL) Connect(connectionString string) error {
    p.connection = connectionString
    fmt.Printf("Connected to PostgreSQL: %s\n", connectionString)
    return nil
}

func (p *PostgreSQL) ExecuteQuery(query string) ([]map[string]interface{}, error) {
    fmt.Printf("PostgreSQL executing: %s\n", query)
    return []map[string]interface{}{{"result": "pg_data"}}, nil
}

func (p *PostgreSQL) Close() error {
    fmt.Println("PostgreSQL connection closed")
    return nil
}

type MongoDB struct {
    connection string
}

func (m *MongoDB) Connect(connectionString string) error {
    m.connection = connectionString
    fmt.Printf("Connected to MongoDB: %s\n", connectionString)
    return nil
}

func (m *MongoDB) ExecuteQuery(query string) ([]map[string]interface{}, error) {
    fmt.Printf("MongoDB executing: %s\n", query)
    return []map[string]interface{}{{"result": "mongo_data"}}, nil
}

func (m *MongoDB) Close() error {
    fmt.Println("MongoDB connection closed")
    return nil
}

// Data Access abstraction
type DataAccess struct {
    db Database
}

func NewDataAccess(db Database) *DataAccess {
    return &DataAccess{db: db}
}

func (d *DataAccess) GetUsers() ([]map[string]interface{}, error) {
    d.db.Connect("localhost:5432")
    defer d.db.Close()
    return d.db.ExecuteQuery("SELECT * FROM users")
}

// Cached Data Access
type CachedDataAccess struct {
    DataAccess
    cache map[string][]map[string]interface{}
}

func NewCachedDataAccess(db Database) *CachedDataAccess {
    return &CachedDataAccess{
        DataAccess: DataAccess{db: db},
        cache:      make(map[string][]map[string]interface{}),
    }
}

func (c *CachedDataAccess) GetUsers() ([]map[string]interface{}, error) {
    if data, exists := c.cache["users"]; exists {
        fmt.Println("Returning cached data")
        return data, nil
    }
    
    data, err := c.DataAccess.GetUsers()
    if err == nil {
        c.cache["users"] = data
    }
    return data, err
}
```

---

## Slide 13: Composite Pattern

### Structural Pattern #3: Composite

**Intent:** Compose objects into tree structures to represent part-whole hierarchies

**When to Use:**
- Want to represent part-whole hierarchies
- Want clients to treat individual objects and compositions uniformly
- Structure can have any level of complexity

**Implementation in Go:**
```go
package composite

import (
    "fmt"
    "strings"
)

// Component interface
type FileSystemNode interface {
    GetName() string
    GetSize() int64
    Display(indent string)
    Add(node FileSystemNode) error
    Remove(node FileSystemNode) error
}

// Leaf
type File struct {
    name string
    size int64
}

func NewFile(name string, size int64) *File {
    return &File{name: name, size: size}
}

func (f *File) GetName() string {
    return f.name
}

func (f *File) GetSize() int64 {
    return f.size
}

func (f *File) Display(indent string) {
    fmt.Printf("%s📄 %s (%d bytes)\n", indent, f.name, f.size)
}

func (f *File) Add(node FileSystemNode) error {
    return fmt.Errorf("cannot add to a file")
}

func (f *File) Remove(node FileSystemNode) error {
    return fmt.Errorf("cannot remove from a file")
}

// Composite
type Directory struct {
    name     string
    children []FileSystemNode
}

func NewDirectory(name string) *Directory {
    return &Directory{
        name:     name,
        children: make([]FileSystemNode, 0),
    }
}

func (d *Directory) GetName() string {
    return d.name
}

func (d *Directory) GetSize() int64 {
    var totalSize int64
    for _, child := range d.children {
        totalSize += child.GetSize()
    }
    return totalSize
}

func (d *Directory) Display(indent string) {
    fmt.Printf("%s📁 %s/ (%d bytes)\n", indent, d.name, d.GetSize())
    for _, child := range d.children {
        child.Display(indent + "  ")
    }
}

func (d *Directory) Add(node FileSystemNode) error {
    d.children = append(d.children, node)
    return nil
}

func (d *Directory) Remove(node FileSystemNode) error {
    for i, child := range d.children {
        if child == node {
            d.children = append(d.children[:i], d.children[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("node not found")
}

// Organization hierarchy example
type Employee interface {
    GetName() string
    GetSalary() float64
    GetDetails() string
    Add(Employee) error
    Remove(Employee) error
    GetSubordinates() []Employee
}

// Leaf - Individual Employee
type Developer struct {
    name     string
    position string
    salary   float64
    skills   []string
}

func NewDeveloper(name, position string, salary float64, skills []string) *Developer {
    return &Developer{
        name:     name,
        position: position,
        salary:   salary,
        skills:   skills,
    }
}

func (d *Developer) GetName() string {
    return d.name
}

func (d *Developer) GetSalary() float64 {
    return d.salary
}

func (d *Developer) GetDetails() string {
    return fmt.Sprintf("%s (%s) - $%.2f - Skills: %s", 
        d.name, d.position, d.salary, strings.Join(d.skills, ", "))
}

func (d *Developer) Add(e Employee) error {
    return fmt.Errorf("cannot add subordinate to developer")
}

func (d *Developer) Remove(e Employee) error {
    return fmt.Errorf("cannot remove subordinate from developer")
}

func (d *Developer) GetSubordinates() []Employee {
    return nil
}

// Composite - Manager
type Manager struct {
    name         string
    position     string
    salary       float64
    subordinates []Employee
}

func NewManager(name, position string, salary float64) *Manager {
    return &Manager{
        name:         name,
        position:     position,
        salary:       salary,
        subordinates: make([]Employee, 0),
    }
}

func (m *Manager) GetName() string {
    return m.name
}

func (m *Manager) GetSalary() float64 {
    return m.salary
}

func (m *Manager) GetDetails() string {
    totalSalary := m.salary
    for _, subordinate := range m.subordinates {
        totalSalary += subordinate.GetSalary()
    }
    return fmt.Sprintf("%s (%s) - $%.2f (Team cost: $%.2f)",
        m.name, m.position, m.salary, totalSalary)
}

func (m *Manager) Add(e Employee) error {
    m.subordinates = append(m.subordinates, e)
    return nil
}

func (m *Manager) Remove(e Employee) error {
    for i, subordinate := range m.subordinates {
        if subordinate == e {
            m.subordinates = append(m.subordinates[:i], m.subordinates[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("employee not found")
}

func (m *Manager) GetSubordinates() []Employee {
    return m.subordinates
}

// UI Components example
type UIComponent interface {
    Render() string
    Add(UIComponent) error
    Remove(UIComponent) error
}

// Leaf components
type Button struct {
    text string
}

func NewButton(text string) *Button {
    return &Button{text: text}
}

func (b *Button) Render() string {
    return fmt.Sprintf("<button>%s</button>", b.text)
}

func (b *Button) Add(c UIComponent) error {
    return fmt.Errorf("cannot add to button")
}

func (b *Button) Remove(c UIComponent) error {
    return fmt.Errorf("cannot remove from button")
}

type Label struct {
    text string
}

func NewLabel(text string) *Label {
    return &Label{text: text}
}

func (l *Label) Render() string {
    return fmt.Sprintf("<label>%s</label>", l.text)
}

func (l *Label) Add(c UIComponent) error {
    return fmt.Errorf("cannot add to label")
}

func (l *Label) Remove(c UIComponent) error {
    return fmt.Errorf("cannot remove from label")
}

// Composite components
type Panel struct {
    id       string
    children []UIComponent
}

func NewPanel(id string) *Panel {
    return &Panel{
        id:       id,
        children: make([]UIComponent, 0),
    }
}

func (p *Panel) Render() string {
    var result strings.Builder
    result.WriteString(fmt.Sprintf(`<div id="%s">`, p.id))
    for _, child := range p.children {
        result.WriteString(child.Render())
    }
    result.WriteString("</div>")
    return result.String()
}

func (p *Panel) Add(c UIComponent) error {
    p.children = append(p.children, c)
    return nil
}

func (p *Panel) Remove(c UIComponent) error {
    for i, child := range p.children {
        if child == c {
            p.children = append(p.children[:i], p.children[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("component not found")
}

// Usage
func main() {
    // File system
    root := NewDirectory("root")
    
    docs := NewDirectory("documents")
    docs.Add(NewFile("resume.pdf", 2048))
    docs.Add(NewFile("cover_letter.docx", 1024))
    
    photos := NewDirectory("photos")
    photos.Add(NewFile("vacation.jpg", 4096))
    photos.Add(NewFile("family.png", 3072))
    
    root.Add(docs)
    root.Add(photos)
    root.Add(NewFile("readme.txt", 512))
    
    root.Display("")
    
    // Organization
    ceo := NewManager("John CEO", "Chief Executive Officer", 200000)
    
    cto := NewManager("Jane CTO", "Chief Technology Officer", 180000)
    dev1 := NewDeveloper("Bob", "Senior Developer", 120000, []string{"Go", "Python"})
    dev2 := NewDeveloper("Alice", "Junior Developer", 80000, []string{"JavaScript"})
    
    cto.Add(dev1)
    cto.Add(dev2)
    ceo.Add(cto)
    
    fmt.Println(ceo.GetDetails())
    
    // UI Components
    mainPanel := NewPanel("main")
    headerPanel := NewPanel("header")
    headerPanel.Add(NewLabel("Welcome"))
    headerPanel.Add(NewButton("Login"))
    
    mainPanel.Add(headerPanel)
    mainPanel.Add(NewButton("Submit"))
    
    fmt.Println(mainPanel.Render())
}
```

---

## Slide 14: Decorator Pattern

### Structural Pattern #4: Decorator

**Intent:** Attach additional responsibilities to object dynamically

**When to Use:**
- Add responsibilities to objects without modifying their code
- Responsibilities can be withdrawn
- Extension by subclassing is impractical

**Implementation in Go:**
```go
package decorator

import (
    "fmt"
    "strings"
    "time"
)

// Component interface
type Coffee interface {
    GetDescription() string
    GetCost() float64
}

// Concrete Component
type SimpleCoffee struct{}

func (c *SimpleCoffee) GetDescription() string {
    return "Simple Coffee"
}

func (c *SimpleCoffee) GetCost() float64 {
    return 2.00
}

// Base Decorator
type CoffeeDecorator struct {
    coffee Coffee
}

// Concrete Decorators
type MilkDecorator struct {
    CoffeeDecorator
}

func NewMilkDecorator(coffee Coffee) *MilkDecorator {
    return &MilkDecorator{
        CoffeeDecorator: CoffeeDecorator{coffee: coffee},
    }
}

func (m *MilkDecorator) GetDescription() string {
    return m.coffee.GetDescription() + ", Milk"
}

func (m *MilkDecorator) GetCost() float64 {
    return m.coffee.GetCost() + 0.50
}

type SugarDecorator struct {
    CoffeeDecorator
}

func NewSugarDecorator(coffee Coffee) *SugarDecorator {
    return &SugarDecorator{
        CoffeeDecorator: CoffeeDecorator{coffee: coffee},
    }
}

func (s *SugarDecorator) GetDescription() string {
    return s.coffee.GetDescription() + ", Sugar"
}

func (s *SugarDecorator) GetCost() float64 {
    return s.coffee.GetCost() + 0.25
}

type WhipCreamDecorator struct {
    CoffeeDecorator
}

func NewWhipCreamDecorator(coffee Coffee) *WhipCreamDecorator {
    return &WhipCreamDecorator{
        CoffeeDecorator: CoffeeDecorator{coffee: coffee},
    }
}

func (w *WhipCreamDecorator) GetDescription() string {
    return w.coffee.GetDescription() + ", Whip Cream"
}

func (w *WhipCreamDecorator) GetCost() float64 {
    return w.coffee.GetCost() + 0.75
}

// Real-world example: HTTP middleware
type Handler interface {
    ServeHTTP(request string) string
}

// Base handler
type BaseHandler struct{}

func (h *BaseHandler) ServeHTTP(request string) string {
    return fmt.Sprintf("Handling request: %s", request)
}

// Logging decorator
type LoggingHandler struct {
    handler Handler
    logger  func(string)
}

func NewLoggingHandler(handler Handler) *LoggingHandler {
    return &LoggingHandler{
        handler: handler,
        logger:  func(s string) { fmt.Println("[LOG]", s) },
    }
}

func (l *LoggingHandler) ServeHTTP(request string) string {
    l.logger(fmt.Sprintf("Request received: %s at %s", request, time.Now()))
    response := l.handler.ServeHTTP(request)
    l.logger(fmt.Sprintf("Response sent: %s", response))
    return response
}

// Authentication decorator
type AuthHandler struct {
    handler Handler
    token   string
}

func NewAuthHandler(handler Handler, token string) *AuthHandler {
    return &AuthHandler{
        handler: handler,
        token:   token,
    }
}

func (a *AuthHandler) ServeHTTP(request string) string {
    if !strings.Contains(request, a.token) {
        return "401 Unauthorized"
    }
    return a.handler.ServeHTTP(request)
}

// Rate limiting decorator
type RateLimitHandler struct {
    handler      Handler
    requests     int
    maxRequests  int
    resetTime    time.Time
}

func NewRateLimitHandler(handler Handler, maxRequests int) *RateLimitHandler {
    return &RateLimitHandler{
        handler:     handler,
        maxRequests: maxRequests,
        resetTime:   time.Now().Add(time.Minute),
    }
}

func (r *RateLimitHandler) ServeHTTP(request string) string {
    if time.Now().After(r.resetTime) {
        r.requests = 0
        r.resetTime = time.Now().Add(time.Minute)
    }
    
    if r.requests >= r.maxRequests {
        return "429 Too Many Requests"
    }
    
    r.requests++
    return r.handler.ServeHTTP(request)
}

// Text processing example
type TextProcessor interface {
    Process(text string) string
}

type PlainTextProcessor struct{}

func (p *PlainTextProcessor) Process(text string) string {
    return text
}

// Decorators for text processing
type UpperCaseDecorator struct {
    processor TextProcessor
}

func NewUpperCaseDecorator(processor TextProcessor) *UpperCaseDecorator {
    return &UpperCaseDecorator{processor: processor}
}

func (u *UpperCaseDecorator) Process(text string) string {
    return strings.ToUpper(u.processor.Process(text))
}

type TrimDecorator struct {
    processor TextProcessor
}

func NewTrimDecorator(processor TextProcessor) *TrimDecorator {
    return &TrimDecorator{processor: processor}
}

func (t *TrimDecorator) Process(text string) string {
    return strings.TrimSpace(t.processor.Process(text))
}

type ReplaceDecorator struct {
    processor TextProcessor
    old       string
    new       string
}

func NewReplaceDecorator(processor TextProcessor, old, new string) *ReplaceDecorator {
    return &ReplaceDecorator{
        processor: processor,
        old:       old,
        new:       new,
    }
}

func (r *ReplaceDecorator) Process(text string) string {
    return strings.ReplaceAll(r.processor.Process(text), r.old, r.new)
}

// Data stream decorator
type DataStream interface {
    Write(data []byte) error
    Read() ([]byte, error)
}

type FileStream struct {
    filename string
    data     []byte
}

func NewFileStream(filename string) *FileStream {
    return &FileStream{filename: filename}
}

func (f *FileStream) Write(data []byte) error {