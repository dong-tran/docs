# SOLID Principles
## The Foundation of Object-Oriented Design

**Presenter:** Dong Tran  
**Date:** November 20, 2025  
**Time:** 03:03 UTC

---

## Slide 1: Title Slide

# SOLID Principles
### Building Maintainable Object-Oriented Software

**Presenter:** Dong Tran  
**Date:** November 20, 2025

*"The SOLID principles are not rules. They are not laws. They are not perfect truths. They are statements of principle that, when followed, tend to make software more maintainable."*  
— Robert C. Martin (Uncle Bob)

---

## Slide 2: Agenda

### Our Journey Through SOLID

1. **Why SOLID Matters**
2. **Overview of SOLID Principles**
3. **S - Single Responsibility Principle**
4. **O - Open/Closed Principle**
5. **L - Liskov Substitution Principle**
6. **I - Interface Segregation Principle**
7. **D - Dependency Inversion Principle**
8. **SOLID in Go: Practical Examples**
9. **Real-World Case Study**
10. **Common Violations & How to Fix Them**
11. **SOLID vs Other Principles**
12. **Testing & SOLID**
13. **Refactoring to SOLID**
14. **Best Practices & Anti-Patterns**
15. **Q&A Session**

---

## Slide 3: The Problem

### Why Do We Need SOLID?

**Common Software Problems:**
- 🔴 **Rigidity** - Hard to change, every change affects too many parts
- 🔴 **Fragility** - Changes break unexpected parts
- 🔴 **Immobility** - Can't reuse code in other projects
- 🔴 **Viscosity** - Easier to do the wrong thing than the right thing
- 🔴 **Complexity** - Over-engineered, unnecessary complexity

**The Cost of Bad Design:**
```
Technical Debt Impact:
- 42% of developer time spent on technical debt (Stripe, 2018)
- 23% slower feature delivery in poorly designed systems
- 4x more bugs in tightly coupled code
- 60% higher maintenance cost
```

**SOLID Promise:**
✅ Maintainable code  
✅ Flexible architecture  
✅ Testable components  
✅ Reusable modules  
✅ Understandable design  

---

## Slide 4: SOLID Overview

### What is SOLID?

**SOLID** = Five design principles for object-oriented programming

Created by **Robert C. Martin** (Uncle Bob) in 2000

```
S - Single Responsibility Principle
O - Open/Closed Principle  
L - Liskov Substitution Principle
I - Interface Segregation Principle
D - Dependency Inversion Principle
```

**The Goal:**
Transform tightly coupled code into loosely coupled, maintainable systems

**Visual Representation:**
```
        Bad Design                Good Design (SOLID)
    ┌──────────────┐            ┌────┐  ┌────┐  ┌────┐
    │              │            │ S  │  │ R  │  │ P  │
    │   Big Ball   │     →      └──┬─┘  └──┬─┘  └──┬─┘
    │   of Mud     │               │       │       │
    │              │            ┌──▼─┐  ┌──▼─┐  ┌──▼─┐
    └──────────────┘            │    │  │    │  │    │
                                └────┘  └────┘  └────┘
    Tangled mess               Clear responsibilities
```

---

## Slide 5: Single Responsibility Principle (SRP)

### S - Single Responsibility Principle

> **"A class should have only one reason to change"**

**What it means:**
- Each class/module should have ONE job
- Separate concerns into different classes
- High cohesion within classes

**❌ BAD Example - Multiple Responsibilities:**
```go
// Bad: This class has multiple reasons to change
type Employee struct {
    ID     int
    Name   string
    Salary float64
}

// Reason 1: Business logic changes
func (e *Employee) CalculatePay() float64 {
    return e.Salary * 0.8 // After tax
}

// Reason 2: Persistence logic changes
func (e *Employee) SaveToDatabase() error {
    db, err := sql.Open("postgres", "...")
    if err != nil {
        return err
    }
    
    _, err = db.Exec("INSERT INTO employees...", e.ID, e.Name, e.Salary)
    return err
}

// Reason 3: Report format changes
func (e *Employee) GenerateReport() string {
    return fmt.Sprintf("Employee Report\nID: %d\nName: %s\nSalary: $%.2f",
        e.ID, e.Name, e.Salary)
}
```

---

## Slide 6: SRP - Correct Implementation

### Single Responsibility - The Right Way

**✅ GOOD Example - Separated Responsibilities:**
```go
// Good: Each struct has a single responsibility

// 1. Employee entity - represents business data
type Employee struct {
    ID     int
    Name   string
    Salary float64
}

// 2. Payroll service - handles payment calculations
type PayrollService struct {
    taxCalculator TaxCalculator
}

func (p *PayrollService) CalculatePay(employee Employee) float64 {
    grossPay := employee.Salary
    tax := p.taxCalculator.Calculate(grossPay)
    return grossPay - tax
}

// 3. Employee repository - handles persistence
type EmployeeRepository struct {
    db *sql.DB
}

func (r *EmployeeRepository) Save(employee Employee) error {
    _, err := r.db.Exec(
        "INSERT INTO employees (id, name, salary) VALUES ($1, $2, $3)",
        employee.ID, employee.Name, employee.Salary,
    )
    return err
}

func (r *EmployeeRepository) FindByID(id int) (*Employee, error) {
    var emp Employee
    err := r.db.QueryRow(
        "SELECT id, name, salary FROM employees WHERE id = $1", 
        id,
    ).Scan(&emp.ID, &emp.Name, &emp.Salary)
    
    if err != nil {
        return nil, err
    }
    return &emp, nil
}

// 4. Report generator - handles reporting
type EmployeeReportGenerator struct {
    template ReportTemplate
}

func (g *EmployeeReportGenerator) Generate(employee Employee) string {
    return g.template.Render(map[string]interface{}{
        "ID":     employee.ID,
        "Name":   employee.Name,
        "Salary": employee.Salary,
    })
}
```

**Benefits:**
- Each class can change independently
- Easier to test
- More reusable
- Clear responsibilities

---

## Slide 7: Open/Closed Principle (OCP)

### O - Open/Closed Principle

> **"Software entities should be open for extension but closed for modification"**

**What it means:**
- Add new functionality without changing existing code
- Use abstraction and polymorphism
- Extend behavior through inheritance or composition

**❌ BAD Example - Violating OCP:**
```go
// Bad: Need to modify existing code for new shapes

type Rectangle struct {
    Width  float64
    Height float64
}

type Circle struct {
    Radius float64
}

// This function violates OCP - must modify for each new shape
func CalculateArea(shape interface{}) float64 {
    switch s := shape.(type) {
    case Rectangle:
        return s.Width * s.Height
    case Circle:
        return math.Pi * s.Radius * s.Radius
    // Must add new case for each new shape - violates OCP!
    // case Triangle:
    //     return ...
    default:
        return 0
    }
}

// Problem: Adding Triangle requires modifying CalculateArea function
type Triangle struct {
    Base   float64
    Height float64
}
```

---

## Slide 8: OCP - Correct Implementation

### Open/Closed - The Right Way

**✅ GOOD Example - Following OCP:**
```go
// Good: Open for extension, closed for modification

// Define abstraction
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Implementations
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// New shape - no modification needed to existing code!
type Triangle struct {
    Base   float64
    Height float64
    SideA  float64
    SideB  float64
    SideC  float64
}

func (t Triangle) Area() float64 {
    return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
    return t.SideA + t.SideB + t.SideC
}

// Client code - never needs to change
func CalculateTotalArea(shapes []Shape) float64 {
    total := 0.0
    for _, shape := range shapes {
        total += shape.Area()
    }
    return total
}

// Strategy Pattern Example - Another OCP implementation
type PaymentProcessor interface {
    ProcessPayment(amount float64) error
}

type PaymentService struct {
    processor PaymentProcessor
}

func (s *PaymentService) Process(amount float64) error {
    // Validation
    if amount <= 0 {
        return errors.New("invalid amount")
    }
    
    // Process using injected strategy
    return s.processor.ProcessPayment(amount)
}

// Add new payment methods without modifying existing code
type CreditCardProcessor struct{}
func (c CreditCardProcessor) ProcessPayment(amount float64) error {
    fmt.Printf("Processing credit card payment: $%.2f\n", amount)
    return nil
}

type PayPalProcessor struct{}
func (p PayPalProcessor) ProcessPayment(amount float64) error {
    fmt.Printf("Processing PayPal payment: $%.2f\n", amount)
    return nil
}

type CryptoProcessor struct{} // New payment method - no changes needed!
func (c CryptoProcessor) ProcessPayment(amount float64) error {
    fmt.Printf("Processing crypto payment: $%.2f\n", amount)
    return nil
}
```

---

## Slide 9: Liskov Substitution Principle (LSP)

### L - Liskov Substitution Principle

> **"Subtypes must be substitutable for their base types without altering program correctness"**

Named after **Barbara Liskov** (1987)

**What it means:**
- Derived classes must be usable through base class interface
- Child classes shouldn't break parent class behavior
- Preconditions cannot be strengthened
- Postconditions cannot be weakened

**❌ BAD Example - Violating LSP:**
```go
// Bad: Square violates LSP when substituted for Rectangle

type Rectangle struct {
    width  float64
    height float64
}

func (r *Rectangle) SetWidth(width float64) {
    r.width = width
}

func (r *Rectangle) SetHeight(height float64) {
    r.height = height
}

func (r *Rectangle) Area() float64 {
    return r.width * r.height
}

// Square IS-A Rectangle? Mathematically yes, but...
type Square struct {
    Rectangle
}

// This breaks LSP!
func (s *Square) SetWidth(width float64) {
    s.width = width
    s.height = width // Forces height to equal width
}

func (s *Square) SetHeight(height float64) {
    s.width = height // Forces width to equal height
    s.height = height
}

// Client code that demonstrates the violation
func TestRectangle(r *Rectangle) {
    r.SetWidth(5)
    r.SetHeight(4)
    
    // Expected: 5 * 4 = 20
    // But if r is actually a Square: 4 * 4 = 16
    // Violates LSP!
    if r.Area() != 20 {
        panic("Area calculation failed!")
    }
}
```

---

## Slide 10: LSP - Correct Implementation

### Liskov Substitution - The Right Way

**✅ GOOD Example - Following LSP:**
```go
// Good: Proper abstraction that respects LSP

// Define behavior, not structure
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    width  float64
    height float64
}

func NewRectangle(width, height float64) *Rectangle {
    return &Rectangle{width: width, height: height}
}

func (r *Rectangle) Area() float64 {
    return r.width * r.height
}

func (r *Rectangle) Perimeter() float64 {
    return 2 * (r.width + r.height)
}

type Square struct {
    side float64
}

func NewSquare(side float64) *Square {
    return &Square{side: side}
}

func (s *Square) Area() float64 {
    return s.side * s.side
}

func (s *Square) Perimeter() float64 {
    return 4 * s.side
}

// Both can be used interchangeably
func CalculateAreas(shapes []Shape) float64 {
    total := 0.0
    for _, shape := range shapes {
        total += shape.Area()
    }
    return total
}

// Another Example: Bird hierarchy
// BAD - Classic LSP violation
type Bird interface {
    Fly() error
    Walk() error
}

type Sparrow struct{}
func (s Sparrow) Fly() error { return nil }
func (s Sparrow) Walk() error { return nil }

type Penguin struct{}
func (p Penguin) Fly() error { 
    return errors.New("penguins can't fly") // LSP violation!
}
func (p Penguin) Walk() error { return nil }

// GOOD - Respecting LSP with better abstraction
type Walker interface {
    Walk() error
}

type Flyer interface {
    Fly() error
}

type FlyingBird interface {
    Walker
    Flyer
}

type WalkingBird interface {
    Walker
}

type Sparrow2 struct{}
func (s Sparrow2) Fly() error { return nil }
func (s Sparrow2) Walk() error { return nil }

type Penguin2 struct{}
func (p Penguin2) Walk() error { return nil }
// Penguin doesn't implement Fly() - no violation!

// Preconditions and Postconditions Example
type Account interface {
    Withdraw(amount float64) error
    Deposit(amount float64) error
    Balance() float64
}

type BasicAccount struct {
    balance float64
}

func (a *BasicAccount) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("amount must be positive")
    }
    if amount > a.balance {
        return errors.New("insufficient funds")
    }
    a.balance -= amount
    return nil
}

type PremiumAccount struct {
    BasicAccount
    overdraftLimit float64
}

// Respects LSP - doesn't strengthen preconditions
func (p *PremiumAccount) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("amount must be positive")
    }
    // Weaker precondition - allows overdraft
    if amount > p.balance+p.overdraftLimit {
        return errors.New("exceeds overdraft limit")
    }
    p.balance -= amount
    return nil
}
```

---

## Slide 11: Interface Segregation Principle (ISP)

### I - Interface Segregation Principle

> **"Clients should not be forced to depend on methods they don't use"**

**What it means:**
- Many small interfaces > one large interface
- Interface pollution creates unnecessary dependencies
- Role-based interfaces

**❌ BAD Example - Fat Interface:**
```go
// Bad: Fat interface forcing unnecessary implementations

type Worker interface {
    Work() error
    Eat() error
    Sleep() error
    GetPaid() float64
    TakeVacation(days int) error
    AttendMeeting() error
    SubmitTimesheet() error
    GetPromoted() error
}

// Human worker - uses all methods
type HumanWorker struct {
    name   string
    salary float64
}

func (h *HumanWorker) Work() error { return nil }
func (h *HumanWorker) Eat() error { return nil }
func (h *HumanWorker) Sleep() error { return nil }
func (h *HumanWorker) GetPaid() float64 { return h.salary }
func (h *HumanWorker) TakeVacation(days int) error { return nil }
func (h *HumanWorker) AttendMeeting() error { return nil }
func (h *HumanWorker) SubmitTimesheet() error { return nil }
func (h *HumanWorker) GetPromoted() error { return nil }

// Robot worker - doesn't need all methods!
type RobotWorker struct {
    model string
}

func (r *RobotWorker) Work() error { return nil }
func (r *RobotWorker) Eat() error { 
    return errors.New("robots don't eat") // Forced to implement!
}
func (r *RobotWorker) Sleep() error { 
    return errors.New("robots don't sleep") // Forced to implement!
}
func (r *RobotWorker) GetPaid() float64 { 
    return 0 // Robots don't get paid?
}
func (r *RobotWorker) TakeVacation(days int) error { 
    return errors.New("robots don't take vacations")
}
func (r *RobotWorker) AttendMeeting() error { return nil }
func (r *RobotWorker) SubmitTimesheet() error { return nil }
func (r *RobotWorker) GetPromoted() error {
    return errors.New("robots don't get promoted")
}
```

---

## Slide 12: ISP - Correct Implementation

### Interface Segregation - The Right Way

**✅ GOOD Example - Segregated Interfaces:**
```go
// Good: Small, focused interfaces

// Basic work capability
type Worker interface {
    Work() error
}

// Biological needs
type BiologicalBeing interface {
    Eat() error
    Sleep() error
}

// Financial aspects
type Payable interface {
    GetPaid() float64
    SubmitTimesheet() error
}

// Career progression
type CareerDevelopment interface {
    AttendMeeting() error
    GetPromoted() error
    TakeVacation(days int) error
}

// Human implements multiple interfaces
type HumanEmployee struct {
    name   string
    salary float64
    energy int
}

func (h *HumanEmployee) Work() error {
    if h.energy < 10 {
        return errors.New("too tired to work")
    }
    h.energy -= 10
    return nil
}

func (h *HumanEmployee) Eat() error {
    h.energy += 20
    return nil
}

func (h *HumanEmployee) Sleep() error {
    h.energy = 100
    return nil
}

func (h *HumanEmployee) GetPaid() float64 {
    return h.salary
}

func (h *HumanEmployee) SubmitTimesheet() error {
    return nil
}

func (h *HumanEmployee) AttendMeeting() error {
    return nil
}

func (h *HumanEmployee) GetPromoted() error {
    h.salary *= 1.1
    return nil
}

func (h *HumanEmployee) TakeVacation(days int) error {
    fmt.Printf("%s is on vacation for %d days\n", h.name, days)
    return nil
}

// Robot only implements what it needs
type RobotWorker struct {
    model      string
    workHours  int
}

func (r *RobotWorker) Work() error {
    r.workHours++
    return nil
}

func (r *RobotWorker) SubmitTimesheet() error {
    fmt.Printf("Robot %s worked %d hours\n", r.model, r.workHours)
    return nil
}

// Freelancer implements different combination
type Freelancer struct {
    name       string
    hourlyRate float64
    hoursWorked float64
}

func (f *Freelancer) Work() error {
    f.hoursWorked++
    return nil
}

func (f *Freelancer) GetPaid() float64 {
    return f.hourlyRate * f.hoursWorked
}

func (f *Freelancer) SubmitTimesheet() error {
    return nil
}

// Client code uses only what it needs
func ProcessWorkday(workers []Worker) {
    for _, w := range workers {
        w.Work()
    }
}

func ProcessPayroll(employees []Payable) {
    for _, e := range employees {
        e.SubmitTimesheet()
        fmt.Printf("Payment: $%.2f\n", e.GetPaid())
    }
}

func HandleBiologicalNeeds(beings []BiologicalBeing) {
    for _, b := range beings {
        b.Eat()
        b.Sleep()
    }
}

// Real-world example: Database operations
type Reader interface {
    Read(id string) (interface{}, error)
}

type Writer interface {
    Write(data interface{}) error
}

type Deleter interface {
    Delete(id string) error
}

// Some clients only need to read
type ReportGenerator struct {
    reader Reader
}

// Some need full CRUD
type Repository interface {
    Reader
    Writer
    Deleter
}

// Some might be read-only
type ReadOnlyRepository struct {
    db *sql.DB
}

func (r *ReadOnlyRepository) Read(id string) (interface{}, error) {
    // Implementation
    return nil, nil
}

// No need to implement Write or Delete!
```

---

## Slide 13: Dependency Inversion Principle (DIP)

### D - Dependency Inversion Principle

> **"High-level modules should not depend on low-level modules. Both should depend on abstractions."**
> **"Abstractions should not depend on details. Details should depend on abstractions."**

**What it means:**
- Depend on interfaces, not concrete implementations
- Business logic shouldn't know about infrastructure details
- Inversion of traditional dependency flow

**❌ BAD Example - Direct Dependencies:**
```go
// Bad: High-level business logic depends on low-level details

// Low-level module (infrastructure)
type MySQLDatabase struct {
    connection *sql.DB
}

func (db *MySQLDatabase) Query(query string) ([]map[string]interface{}, error) {
    // MySQL specific implementation
    rows, err := db.connection.Query(query)
    if err != nil {
        return nil, err
    }
    // ... process rows
    return nil, nil
}

// High-level module (business logic) - BAD!
type UserService struct {
    db *MySQLDatabase // Direct dependency on concrete implementation!
}

func NewUserService() *UserService {
    // Tightly coupled to MySQL
    mysqlDB := &MySQLDatabase{
        connection: createMySQLConnection(),
    }
    return &UserService{db: mysqlDB}
}

func (s *UserService) GetUser(id int) (*User, error) {
    // Business logic coupled to MySQL syntax
    query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", id)
    results, err := s.db.Query(query)
    
    if err != nil {
        return nil, err
    }
    
    // Process MySQL specific results
    return parseUserFromMySQL(results), nil
}

// Problems:
// 1. Can't change database without changing UserService
// 2. Can't unit test without MySQL database
// 3. Business logic knows about SQL
```

---

## Slide 14: DIP - Correct Implementation

### Dependency Inversion - The Right Way

**✅ GOOD Example - Depending on Abstractions:**
```go
// Good: Both high-level and low-level depend on abstractions

// Domain layer defines the interface (abstraction)
type UserRepository interface {
    FindByID(id int) (*User, error)
    FindByEmail(email string) (*User, error)
    Save(user *User) error
    Delete(id int) error
}

// Domain entity (high-level)
type User struct {
    ID       int
    Username string
    Email    string
    Password string
}

// Business logic (high-level module)
type UserService struct {
    repo UserRepository // Depends on abstraction!
    hasher PasswordHasher // Another abstraction
    notifier Notifier    // Another abstraction
}

// Constructor with dependency injection
func NewUserService(repo UserRepository, hasher PasswordHasher, notifier Notifier) *UserService {
    return &UserService{
        repo:     repo,
        hasher:   hasher,
        notifier: notifier,
    }
}

func (s *UserService) CreateUser(username, email, password string) (*User, error) {
    // Business logic - no infrastructure details!
    
    // Check if user exists
    existing, _ := s.repo.FindByEmail(email)
    if existing != nil {
        return nil, errors.New("user already exists")
    }
    
    // Hash password using abstraction
    hashedPassword, err := s.hasher.Hash(password)
    if err != nil {
        return nil, err
    }
    
    // Create user
    user := &User{
        Username: username,
        Email:    email,
        Password: hashedPassword,
    }
    
    // Save using abstraction
    if err := s.repo.Save(user); err != nil {
        return nil, err
    }
    
    // Notify using abstraction
    s.notifier.SendWelcomeEmail(user.Email)
    
    return user, nil
}

// Infrastructure layer implements the abstraction
type MySQLUserRepository struct {
    db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
    return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) FindByID(id int) (*User, error) {
    var user User
    err := r.db.QueryRow(
        "SELECT id, username, email, password FROM users WHERE id = ?",
        id,
    ).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
    
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    
    return &user, nil
}

func (r *MySQLUserRepository) Save(user *User) error {
    result, err := r.db.Exec(
        "INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
        user.Username, user.Email, user.Password,
    )
    if err != nil {
        return err
    }
    
    id, _ := result.LastInsertId()
    user.ID = int(id)
    return nil
}

// Can easily swap implementations
type MongoUserRepository struct {
    collection *mongo.Collection
}

func (r *MongoUserRepository) FindByID(id int) (*User, error) {
    var user User
    err := r.collection.FindOne(
        context.Background(),
        bson.M{"_id": id},
    ).Decode(&user)
    
    if err == mongo.ErrNoDocuments {
        return nil, nil
    }
    return &user, err
}

func (r *MongoUserRepository) Save(user *User) error {
    _, err := r.collection.InsertOne(context.Background(), user)
    return err
}

// In-memory implementation for testing
type InMemoryUserRepository struct {
    users map[int]*User
    mu    sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
    return &InMemoryUserRepository{
        users: make(map[int]*User),
    }
}

func (r *InMemoryUserRepository) FindByID(id int) (*User, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    
    user, exists := r.users[id]
    if !exists {
        return nil, nil
    }
    return user, nil
}

func (r *InMemoryUserRepository) Save(user *User) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    if user.ID == 0 {
        user.ID = len(r.users) + 1
    }
    r.users[user.ID] = user
    return nil
}

// Wiring it all together (usually in main or DI container)
func InitializeApp(config *Config) *UserService {
    var repo UserRepository
    
    switch config.DatabaseType {
    case "mysql":
        db := initMySQL(config.MySQLConfig)
        repo = NewMySQLUserRepository(db)
    case "mongo":
        db := initMongo(config.MongoConfig)
        repo = NewMongoUserRepository(db)
    case "memory":
        repo = NewInMemoryUserRepository()
    }
    
    hasher := NewBcryptHasher()
    notifier := NewEmailNotifier(config.SMTPConfig)
    
    return NewUserService(repo, hasher, notifier)
}
```

---

## Slide 15: SOLID Working Together

### How SOLID Principles Complement Each Other

```go
// Example: E-Commerce Order Processing System

// 1. Single Responsibility - Each class has one job
type Order struct {
    ID         string
    CustomerID string
    Items      []OrderItem
    Status     OrderStatus
    Total      float64
}

type OrderItem struct {
    ProductID string
    Quantity  int
    Price     float64
}

// 2. Interface Segregation - Small, focused interfaces
type OrderRepository interface {
    Save(order *Order) error
    FindByID(id string) (*Order, error)
}

type PaymentProcessor interface {
    ProcessPayment(amount float64, method PaymentMethod) error
}

type InventoryManager interface {
    ReserveItems(items []OrderItem) error
    ReleaseItems(items []OrderItem) error
}

type NotificationService interface {
    SendOrderConfirmation(order *Order) error
}

// 3. Dependency Inversion - Depend on abstractions
type OrderService struct {
    orderRepo    OrderRepository      // Interface
    payment      PaymentProcessor     // Interface
    inventory    InventoryManager     // Interface
    notification NotificationService  // Interface
}

// 4. Open/Closed - Extensible without modification
type DiscountStrategy interface {
    Calculate(order *Order) float64
}

type PercentageDiscount struct {
    percentage float64
}

func (d PercentageDiscount) Calculate(order *Order) float64 {
    return order.Total * (d.percentage / 100)
}

type FixedAmountDiscount struct {
    amount float64
}

func (d FixedAmountDiscount) Calculate(order *Order) float64 {
    return d.amount
}

type LoyaltyPointsDiscount struct {
    pointsService LoyaltyService
}

func (d LoyaltyPointsDiscount) Calculate(order *Order) float64 {
    points := d.pointsService.GetPoints(order.CustomerID)
    return float64(points) * 0.01 // 1 cent per point
}

// 5. Liskov Substitution - All implementations are interchangeable
func NewOrderService(
    repo OrderRepository,
    payment PaymentProcessor,
    inventory InventoryManager,
    notification NotificationService,
) *OrderService {
    return &OrderService{
        orderRepo:    repo,
        payment:      payment,
        inventory:    inventory,
        notification: notification,
    }
}

func (s *OrderService) ProcessOrder(order *Order, discount DiscountStrategy) error {
    // Apply discount (Open/Closed - any discount strategy works)
    if discount != nil {
        discountAmount := discount.Calculate(order)
        order.Total -= discountAmount
    }
    
    // Reserve inventory (Single Responsibility)
    if err := s.inventory.ReserveItems(order.Items); err != nil {
        return fmt.Errorf("inventory reservation failed: %w", err)
    }
    
    // Process payment (Dependency Inversion)
    if err := s.payment.ProcessPayment(order.Total, CreditCard); err != nil {
        // Rollback inventory
        s.inventory.ReleaseItems(order.Items)
        return fmt.Errorf("payment failed: %w", err)
    }
    
    // Save order (Interface Segregation)
    order.Status = OrderConfirmed
    if err := s.orderRepo.Save(order); err != nil {
        return fmt.Errorf("failed to save order: %w", err)
    }
    
    // Send notification (Liskov Substitution - any notifier works)
    go s.notification.SendOrderConfirmation(order)
    
    return nil
}

// Different implementations can be swapped (LSP + DIP)
type EmailNotificationService struct {
    smtpClient SMTPClient
}

func (e *EmailNotificationService) SendOrderConfirmation(order *Order) error {
    // Email implementation
    return e.smtpClient.Send(/* email details */)
}

type SMSNotificationService struct {
    smsClient SMSClient
}

func (s *SMSNotificationService) SendOrderConfirmation(order *Order) error {
    // SMS implementation
    return s.smsClient.Send(/* sms details */)
}

// Testing becomes easy
func TestOrderProcessing(t *testing.T) {
    // Use test doubles (DIP makes this possible)
    mockRepo := &MockOrderRepository{}
    mockPayment := &MockPaymentProcessor{}
    mockInventory := &MockInventoryManager{}
    mockNotifier := &MockNotificationService{}
    
    service := NewOrderService(mockRepo, mockPayment, mockInventory, mockNotifier)
    
    order := &Order{
        ID:    "123",
        Total: 100.00,
        Items: []OrderItem{{ProductID: "P1", Quantity: 2, Price: 50}},
    }
    
    // Test with different discount strategies (OCP)
    percentDiscount := PercentageDiscount{percentage: 10}
    err := service.ProcessOrder(order, percentDiscount)
    
    assert.NoError(t, err)
    assert.Equal(t, 90.00, order.Total)
}
```

---

## Slide 16: Common SOLID Violations

### Recognizing SOLID Violations in Code

```go
// Code Smells that Indicate SOLID Violations

// 1. SRP Violation - God Class
type UserManager struct {
    db *sql.DB
}

// Too many responsibilities in one class
func (u *UserManager) CreateUser() {}
func (u *UserManager) AuthenticateUser() {}
func (u *UserManager) SendEmail() {}
func (u *UserManager) GenerateReport() {}
func (u *UserManager) BackupDatabase() {}
func (u *UserManager) ValidatePassword() {}
func (u *UserManager) UploadProfilePicture() {}
func (u *UserManager) CalculateAge() {}

// 2. OCP Violation - Switch Statements
func ProcessPayment(method string, amount float64) error {
    switch method {
    case "credit_card":
        // process credit card
    case "paypal":
        // process paypal
    case "bitcoin":  // Must modify code to add new payment method
        // process bitcoin
    }
    // Adding new payment method requires modifying this function
    return nil
}

// 3. LSP Violation - Unexpected Behavior
type Bird interface {
    Fly() error
}

type Ostrich struct{} // Ostrich is a bird but...
func (o Ostrich) Fly() error {
    panic("Ostriches can't fly!") // Violates LSP!
}

// 4. ISP Violation - Unused Interface Methods
type DatabaseOperations interface {
    Create() error
    Read() error
    Update() error
    Delete() error
    Backup() error
    Restore() error
    Replicate() error
    Shard() error
}

type SimpleCache struct{}
// Forced to implement methods it doesn't need
func (s SimpleCache) Backup() error { return errors.New("not implemented") }
func (s SimpleCache) Restore() error { return errors.New("not implemented") }
func (s SimpleCache) Replicate() error { return errors.New("not implemented") }
func (s SimpleCache) Shard() error { return errors.New("not implemented") }

// 5. DIP Violation - New Keyword in Business Logic
type OrderService struct{}

func (o *OrderService) CreateOrder() {
    db := &MySQLDatabase{} // Creating concrete instance
    logger := &FileLogger{} // Direct instantiation
    emailer := &SMTPEmailer{} // Tight coupling
    
    // Business logic tightly coupled to implementations
}
```

---

## Slide 17: Refactoring to SOLID

### Step-by-Step Refactoring Guide

```go
// BEFORE: Violating multiple SOLID principles
type BlogPost struct {
    Title   string
    Content string
    Author  string
}

func (b *BlogPost) Save() error {
    db, _ := sql.Open("mysql", "...")
    _, err := db.Exec("INSERT INTO posts...", b.Title, b.Content)
    return err
}

func (b *BlogPost) SendEmailToSubscribers() {
    // Get subscribers from database
    // Send email to each
}

func (b *BlogPost) GenerateHTML() string {
    return "<h1>" + b.Title + "</h1><p>" + b.Content + "</p>"
}

func (b *BlogPost) GeneratePDF() []byte {
    // Generate PDF
    return nil
}

// AFTER: Following SOLID principles

// Step 1: Extract Entity (SRP)
type BlogPost struct {
    ID        string
    Title     string
    Content   string
    Author    Author
    CreatedAt time.Time
    Tags      []string
}

// Step 2: Create Repository Interface (DIP + SRP)
type BlogPostRepository interface {
    Save(post *BlogPost) error
    FindByID(id string) (*BlogPost, error)
    FindByAuthor(authorID string) ([]*BlogPost, error)
}

// Step 3: Create Renderer Interface (OCP + ISP)
type Renderer interface {
    Render(post *BlogPost) ([]byte, error)
}

type HTMLRenderer struct {
    template *template.Template
}

func (r *HTMLRenderer) Render(post *BlogPost) ([]byte, error) {
    var buf bytes.Buffer
    err := r.template.Execute(&buf, post)
    return buf.Bytes(), err
}

type PDFRenderer struct {
    pdfGenerator PDFGenerator
}

func (r *PDFRenderer) Render(post *BlogPost) ([]byte, error) {
    return r.pdfGenerator.Generate(post)
}

type JSONRenderer struct{}

func (r *JSONRenderer) Render(post *BlogPost) ([]byte, error) {
    return json.Marshal(post)
}

// Step 4: Create Notification Service (SRP + DIP)
type NotificationService interface {
    NotifyNewPost(post *BlogPost) error
}

type EmailNotificationService struct {
    subscriberRepo SubscriberRepository
    emailSender    EmailSender
}

func (s *EmailNotificationService) NotifyNewPost(post *BlogPost) error {
    subscribers, err := s.subscriberRepo.GetAll()
    if err != nil {
        return err
    }
    
    for _, subscriber := range subscribers {
        s.emailSender.Send(EmailMessage{
            To:      subscriber.Email,
            Subject: "New Blog Post: " + post.Title,
            Body:    post.Content,
        })
    }
    
    return nil
}

// Step 5: Create Blog Service (Orchestration)
type BlogService struct {
    postRepo     BlogPostRepository
    notifier     NotificationService
    cache        CacheService
}

func NewBlogService(
    repo BlogPostRepository,
    notifier NotificationService,
    cache CacheService,
) *BlogService {
    return &BlogService{
        postRepo: repo,
        notifier: notifier,
        cache:    cache,
    }
}

func (s *BlogService) PublishPost(post *BlogPost) error {
    // Validate
    if err := s.validatePost(post); err != nil {
        return err
    }
    
    // Save
    if err := s.postRepo.Save(post); err != nil {
        return err
    }
    
    // Cache
    s.cache.Set("post:"+post.ID, post, 1*time.Hour)
    
    // Notify asynchronously
    go s.notifier.NotifyNewPost(post)
    
    return nil
}

func (s *BlogService) validatePost(post *BlogPost) error {
    if post.Title == "" {
        return errors.New("title is required")
    }
    if post.Content == "" {
        return errors.New("content is required")
    }
    if len(post.Content) < 100 {
        return errors.New("content must be at least 100 characters")
    }
    return nil
}

// Step 6: Wire everything together
func InitializeBlogSystem(config *Config) *BlogService {
    // Choose implementations based on config
    var repo BlogPostRepository
    if config.Database == "mysql" {
        repo = NewMySQLBlogRepository(config.MySQLConn)
    } else {
        repo = NewMongoBlogRepository(config.MongoConn)
    }
    
    emailSender := NewSMTPEmailSender(config.SMTPConfig)
    subscriberRepo := NewSubscriberRepository(config.Database)
    notifier := &EmailNotificationService{
        subscriberRepo: subscriberRepo,
        emailSender:    emailSender,
    }
    
    cache := NewRedisCache(config.RedisConn)
    
    return NewBlogService(repo, notifier, cache)
}
```

---

## Slide 18: Testing with SOLID

### How SOLID Makes Testing Easier

```go
// Testing becomes simple with SOLID principles

// 1. Testing with Dependency Injection (DIP)
type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) FindByID(id int) (*User, error) {
    args := m.Called(id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) Save(user *User) error {
    args := m.Called(user)
    return args.Error(0)
}

func TestUserService_CreateUser(t *testing.T) {
    // Arrange
    mockRepo := new(MockUserRepository)
    mockHasher := new(MockPasswordHasher)
    mockNotifier := new(MockNotifier)
    
    service := NewUserService(mockRepo, mockHasher, mockNotifier)
    
    mockRepo.On("FindByEmail", "test@example.com").Return(nil, nil)
    mockRepo.On("Save", mock.Anything).Return(nil)
    mockHasher.On("Hash", "password123").Return("hashed_password", nil)
    mockNotifier.On("SendWelcomeEmail", "test@example.com").Return(nil)
    
    // Act
    user, err := service.CreateUser("testuser", "test@example.com", "password123")
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, "testuser", user.Username)
    assert.Equal(t, "test@example.com", user.Email)
    
    mockRepo.AssertExpectations(t)
    mockHasher.AssertExpectations(t)
    mockNotifier.AssertExpectations(t)
}

// 2. Testing with Interface Segregation (ISP)
// Only mock what you need
type ReaderMock struct {
    mock.Mock
}

func (r *ReaderMock) Read(id string) (interface{}, error) {
    args := r.Called(id)
    return args.Get(0), args.Error(1)
}

func TestReportGenerator(t *testing.T) {
    // Only need to mock Reader interface, not entire repository
    mockReader := new(ReaderMock)
    generator := &ReportGenerator{reader: mockReader}
    
    mockReader.On("Read", "123").Return(&ReportData{
        Title: "Test Report",
        Value: 100,
    }, nil)
    
    report := generator.Generate("123")
    assert.Equal(t, "Test Report: 100", report)
}

// 3. Testing Open/Closed Principle (OCP)
func TestDiscountStrategies(t *testing.T) {
    order := &Order{
        Total: 100,
        Items: []OrderItem{
            {ProductID: "P1", Quantity: 2, Price: 50},
        },
    }
    
    testCases := []struct {
        name     string
        strategy DiscountStrategy
        expected float64
    }{
        {
            name:     "Percentage Discount",
            strategy: PercentageDiscount{percentage: 10},
            expected: 10,
        },
        {
            name:     "Fixed Amount Discount",
            strategy: FixedAmountDiscount{amount: 15},
            expected: 15,
        },
        {
            name:     "No Discount",
            strategy: nil,
            expected: 0,
        },
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            var discount float64
            if tc.strategy != nil {
                discount = tc.strategy.Calculate(order)
            }
            assert.Equal(t, tc.expected, discount)
        })
    }
}

// 4. Testing Liskov Substitution (LSP)
func TestShapeCalculations(t *testing.T) {
    shapes := []Shape{
        Rectangle{Width: 5, Height: 10},
        Circle{Radius: 5},
        Triangle{Base: 10, Height: 5},
    }
    
    // All shapes should work the same way
    for _, shape := range shapes {
        area := shape.Area()
        perimeter := shape.Perimeter()
        
        assert.Greater(t, area, 0.0)
        assert.Greater(t, perimeter, 0.0)
    }
}

// 5. Integration Test with Test Doubles
type TestDoubles struct {
    repo     *InMemoryUserRepository
    hasher   *SimpleHasher
    notifier *MockNotifier
}

func setupTestDoubles() *TestDoubles {
    return &TestDoubles{
        repo:     NewInMemoryUserRepository(),
        hasher:   &SimpleHasher{},
        notifier: &MockNotifier{},
    }
}

func TestUserServiceIntegration(t *testing.T) {
    // Setup
    doubles := setupTestDoubles()
    service := NewUserService(doubles.repo, doubles.hasher, doubles.notifier)
    
    // Test user creation
    user1, err := service.CreateUser("user1", "user1@test.com", "pass1")
    assert.NoError(t, err)
    assert.NotNil(t, user1)
    
    // Test duplicate email
    user2, err := service.CreateUser("user2", "user1@test.com", "pass2")
    assert.Error(t, err)
    assert.Nil(t, user2)
    
    // Verify user was saved
    savedUser, err := doubles.repo.FindByID(user1.ID)
    assert.NoError(t, err)
    assert.Equal(t, user1.Email, savedUser.Email)
}

// Benchmark tests are easier with SOLID
func BenchmarkUserCreation(b *testing.B) {
    repo := NewInMemoryUserRepository()
    hasher := &SimpleHasher{}
    notifier := &NoOpNotifier{}
    
    service := NewUserService(repo, hasher, notifier)
    
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        email := fmt.Sprintf("user%d@test.com", i)
        service.CreateUser("user", email, "password")
    }
}
```

---

## Slide 19: SOLID in Real Projects

### Case Study: Payment Processing System

```go
// Real-world example: Payment processing system following SOLID

// Domain Entities (SRP)
type Payment struct {
    ID            string
    Amount        Money
    Currency      string
    Method        PaymentMethod
    Status        PaymentStatus
    CustomerID    string
    OrderID       string
    ProcessedAt   *time.Time
    FailureReason string
}

type Money struct {
    Amount   float64
    Currency string
}

type PaymentMethod string

const (
    CreditCard   PaymentMethod = "credit_card"
    DebitCard    PaymentMethod = "debit_card"
    PayPal       PaymentMethod = "paypal"
    BankTransfer PaymentMethod = "bank_transfer"
    Cryptocurrency PaymentMethod = "crypto"
)

// Interfaces (ISP + DIP)
type PaymentGateway interface {
    ProcessPayment(payment *Payment) (*PaymentResult, error)
    SupportsMethod(method PaymentMethod) bool
}

type PaymentRepository interface {
    Save(payment *Payment) error
    FindByID(id string) (*Payment, error)
    FindByOrderID(orderID string) ([]*Payment, error)
}

type FraudDetector interface {
    AssessRisk(payment *Payment) (RiskScore, error)
}

type PaymentNotifier interface {
    NotifySuccess(payment *Payment) error
    NotifyFailure(payment *Payment, reason string) error
}

type CurrencyConverter interface {
    Convert(amount Money, targetCurrency string) (Money, error)
}

// Implementations (OCP - Easy to add new payment methods)
type StripeGateway struct {
    apiKey string
    client *stripe.Client
}

func (s *StripeGateway) ProcessPayment(payment *Payment) (*PaymentResult, error) {
    charge, err := s.client.Charges.New(&stripe.ChargeParams{
        Amount:      stripe.Int64(int64(payment.Amount.Amount * 100)),
        Currency:    stripe.String(payment.Currency),
        Description: stripe.String("Order: " + payment.OrderID),
    })
    
    if err != nil {
        return &PaymentResult{Success: false, Error: err.Error()}, nil
    }
    
    return &PaymentResult{
        Success:       true,
        TransactionID: charge.ID,
    }, nil
}

func (s *StripeGateway) SupportsMethod(method PaymentMethod) bool {
    return method == CreditCard || method == DebitCard
}

type PayPalGateway struct {
    clientID string
    secret   string
}

func (p *PayPalGateway) ProcessPayment(payment *Payment) (*PaymentResult, error) {
    // PayPal specific implementation
    return &PaymentResult{Success: true}, nil
}

func (p *PayPalGateway) SupportsMethod(method PaymentMethod) bool {
    return method == PayPal
}

// New payment method - no existing code changes needed (OCP)
type CryptoGateway struct {
    walletAddress string
    network       string
}

func (c *CryptoGateway) ProcessPayment(payment *Payment) (*PaymentResult, error) {
    // Cryptocurrency payment implementation
    return &PaymentResult{Success: true}, nil
}

func (c *CryptoGateway) SupportsMethod(method PaymentMethod) bool {
    return method == Cryptocurrency
}

// Payment Service (SRP - Orchestration only)
type PaymentService struct {
    gateways   []PaymentGateway
    repository PaymentRepository
    fraud      FraudDetector
    notifier   PaymentNotifier
    converter  CurrencyConverter
}

func NewPaymentService(
    gateways []PaymentGateway,
    repo PaymentRepository,
    fraud FraudDetector,
    notifier PaymentNotifier,
    converter CurrencyConverter,
) *PaymentService {
    return &PaymentService{
        gateways:   gateways,
        repository: repo,
        fraud:      fraud,
        notifier:   notifier,
        converter:  converter,
    }
}

func (s *PaymentService) ProcessPayment(payment *Payment) error {
    // 1. Fraud detection
    riskScore, err := s.fraud.AssessRisk(payment)
    if err != nil {
        return fmt.Errorf("fraud assessment failed: %w", err)
    }
    
    if riskScore > HighRisk {
        payment.Status = PaymentFailed
        payment.FailureReason = "High risk transaction"
        s.repository.Save(payment)
        s.notifier.NotifyFailure(payment, payment.FailureReason)
        return errors.New("payment rejected due to high risk")
    }
    
    // 2. Currency conversion if needed
    if payment.Currency != "USD" {
        converted, err := s.converter.Convert(
            payment.Amount,
            "USD",
        )
        if err != nil {
            return fmt.Errorf("currency conversion failed: %w", err)
        }
        payment.Amount = converted
    }
    
    // 3. Find appropriate gateway (LSP - all gateways are interchangeable)
    gateway := s.findGateway(payment.Method)
    if gateway == nil {
        return fmt.Errorf("no gateway available for %s", payment.Method)
    }
    
    // 4. Process payment
    result, err := gateway.ProcessPayment(payment)
    if err != nil {
        payment.Status = PaymentFailed
        payment.FailureReason = err.Error()
        s.repository.Save(payment)
        s.notifier.NotifyFailure(payment, err.Error())
        return err
    }
    
    // 5. Update payment status
    if result.Success {
        payment.Status = PaymentCompleted
        payment.ProcessedAt = timePtr(time.Now())
    } else {
        payment.Status = PaymentFailed
        payment.FailureReason = result.Error
    }
    
    // 6. Save to repository
    if err := s.repository.Save(payment); err != nil {
        return fmt.Errorf("failed to save payment: %w", err)
    }
    
    // 7. Send notification
    if result.Success {
        go s.notifier.NotifySuccess(payment)
    } else {
        go s.notifier.NotifyFailure(payment, result.Error)
    }
    
    return nil
}

func (s *PaymentService) findGateway(method PaymentMethod) PaymentGateway {
    for _, gateway := range s.gateways {
        if gateway.SupportsMethod(method) {
            return gateway
        }
    }
    return nil
}

// Retry mechanism with exponential backoff (Additional feature)
type RetryablePaymentService struct {
    PaymentService
    maxRetries int
}

func (s *RetryablePaymentService) ProcessPaymentWithRetry(payment *Payment) error {
    var lastError error
    
    for attempt := 0; attempt < s.maxRetries; attempt++ {
        if err := s.ProcessPayment(payment); err != nil {
            lastError = err
            time.Sleep(time.Duration(math.Pow(2, float64(attempt))) * time.Second)
            continue
        }
        return nil
    }
    
    return fmt.Errorf("payment failed after %d attempts: %w", s.maxRetries, lastError)
}

// Usage
func main() {
    // Setup dependencies
    stripeGateway := &StripeGateway{apiKey: "sk_test_..."}
    paypalGateway := &PayPalGateway{clientID: "...", secret: "..."}
    cryptoGateway := &CryptoGateway{walletAddress: "..."}
    
    repository := NewPostgresPaymentRepository(db)
    fraudDetector := NewMachineLearningFraudDetector()
    notifier := NewEmailNotifier()
    converter := NewExchangeRateConverter()
    
    // Create service with all dependencies
    service := NewPaymentService(
        []PaymentGateway{stripeGateway, paypalGateway, cryptoGateway},
        repository,
        fraudDetector,
        notifier,
        converter,
    )
    
    // Process payment
    payment := &Payment{
        ID:         uuid.New().String(),
        Amount:     Money{Amount: 99.99, Currency: "USD"},
        Method:     CreditCard,
        CustomerID: "customer-123",
        OrderID:    "order-456",
    }
    
    if err := service.ProcessPayment(payment); err != nil {
        log.Printf("Payment failed: %v", err)
    }
}
```

---

## Slide 20: SOLID Metrics & Benefits

### Measuring SOLID Impact

```go
// Metrics showing benefits of SOLID principles

type CodeMetrics struct {
    CyclomaticComplexity int
    LinesOfCode         int
    TestCoverage        float64
    Dependencies        int
    CouplingFactor      float64
}

// Before SOLID
type MonolithicService struct {
    // 3000+ lines of code
    // 15+ responsibilities
    // 50+ methods
}

beforeMetrics := CodeMetrics{
    CyclomaticComplexity: 45,    // Very complex
    LinesOfCode:         3000,   // Large class
    TestCoverage:        35.0,   // Hard to test
    Dependencies:        25,     // Many dependencies
    CouplingFactor:      0.85,   // Highly coupled
}

// After SOLID
type UserService struct {/* ... */}
type UserRepository struct {/* ... */}
type EmailService struct {/* ... */}
type ValidationService struct {/* ... */}

afterMetrics := CodeMetrics{
    CyclomaticComplexity: 8,      // Simple
    LinesOfCode:         150,     // Small, focused classes
    TestCoverage:        95.0,    // Easy to test
    Dependencies:        3,       // Few dependencies
    CouplingFactor:      0.15,    // Loosely coupled
}

// Improvement Percentages
improvements := map[string]float64{
    "Complexity Reduction":  82.2,  // 82% less complex
    "Code Size Reduction":   95.0,  // 95% smaller classes
    "Test Coverage Increase": 171.4, // 171% more coverage
    "Dependency Reduction":  88.0,  // 88% fewer dependencies
    "Coupling Reduction":    82.4,  // 82% less coupled
}

// Development Speed Metrics
type DevelopmentMetrics struct {
    FeatureDeliveryTime   time.Duration
    BugFixTime           time.Duration
    OnboardingTime       time.Duration
    CodeReviewTime       time.Duration
}

withoutSOLID := DevelopmentMetrics{
    FeatureDeliveryTime: 2 * 7 * 24 * time.Hour,  // 2 weeks
    BugFixTime:         3 * 24 * time.Hour,       // 3 days
    OnboardingTime:     4 * 7 * 24 * time.Hour,   // 4 weeks
    CodeReviewTime:     4 * time.Hour,            // 4 hours
}

withSOLID := DevelopmentMetrics{
    FeatureDeliveryTime: 3 * 24 * time.Hour,      // 3 days
    BugFixTime:         4 * time.Hour,            // 4 hours
    OnboardingTime:     1 * 7 * 24 * time.Hour,   // 1 week
    CodeReviewTime:     30 * time.Minute,         // 30 minutes
}

// ROI Calculation
type ROICalculation struct {
    InitialRefactoringCost float64
    MonthlySavings        float64
    PaybackPeriod         int // months
}

solidROI := ROICalculation{
    InitialRefactoringCost: 50000,  // $50k investment
    MonthlySavings:        15000,   // $15k/month saved
    PaybackPeriod:         4,       // ROI in 4 months
}
```

**Real-World Benefits:**

| Metric | Before SOLID | After SOLID | Improvement |
|--------|--------------|-------------|-------------|
| Bug Density | 15 bugs/KLOC | 3 bugs/KLOC | 80% reduction |
| Test Coverage | 35% | 95% | 171% increase |
| Build Time | 15 min | 2 min | 87% faster |
| Deploy Frequency | Monthly | Daily | 30x increase |
| MTTR | 4 hours | 30 min | 87% reduction |
| Developer Satisfaction | 4/10 | 8/10 | 100% increase |

---

## Slide 21: SOLID Design Patterns

### How Design Patterns Implement SOLID

```go
// Common Design Patterns that follow SOLID principles

// 1. Strategy Pattern (OCP + DIP)
type CompressionStrategy interface {
    Compress(data []byte) []