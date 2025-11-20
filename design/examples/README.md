# Design Patterns and Architecture Examples

Comprehensive Go examples demonstrating software design patterns and architectural principles using **Echo web framework** and **SQLx** for database operations.

## 📁 Repository Structure

```
examples/
├── clean-architecture/          # Clean Architecture with Task Management
├── ddd/                        # Domain-Driven Design with E-Commerce
├── solid-principles/           # All 5 SOLID principles
├── design-patterns/            # Gang of Four patterns
├── microservices/              # Microservices architecture
└── relationships-integration/  # How they all work together
```

## 🎯 Examples Overview

### 1. Clean Architecture (`clean-architecture/`)
**Topic**: Building Maintainable and Testable Software Systems

**Demonstrates**:
- The Dependency Rule
- Four layers (Entities, Use Cases, Interface Adapters, Frameworks & Drivers)
- Task Management REST API
- Independence from frameworks and databases

**Tech Stack**: Go, Echo, SQLx, SQLite

**Run**:
```bash
cd clean-architecture
go mod download
go run main.go
```

**API**: `http://localhost:8080`

---

### 2. Domain-Driven Design (`ddd/`)
**Topic**: Building Software That Reflects Business Reality

**Demonstrates**:
- Entities (Product)
- Value Objects (Money, Category)
- Aggregates & Aggregate Roots
- Domain Services (PricingService)
- Application Services
- Repository Pattern

**Tech Stack**: Go, Echo, SQLx

**Run**:
```bash
cd ddd
go mod download
go run cmd/main.go
```

---

### 3. SOLID Principles (`solid-principles/`)
**Topic**: The Foundation of Object-Oriented Design

**Demonstrates**:
- **S**ingle Responsibility Principle (SRP)
- **O**pen/Closed Principle (OCP)
- **L**iskov Substitution Principle (LSP)
- **I**nterface Segregation Principle (ISP)
- **D**ependency Inversion Principle (DIP)

Each principle has both BAD and GOOD examples showing violations and correct implementations.

**Run**:
```bash
cd solid-principles
go test ./...
```

---

### 4. GoF Design Patterns (`design-patterns/`)
**Topic**: 23 Classic Solutions to Recurring Design Problems

**Demonstrates**:

**Creational Patterns**:
- Singleton
- Factory Method
- Builder

**Structural Patterns**:
- Adapter
- Decorator

**Behavioral Patterns**:
- Strategy
- Observer

**Run**:
```bash
cd design-patterns
go test ./...
```

---

### 5. Microservices Architecture (`microservices/`)
**Topic**: Building Scalable Distributed Systems

**Demonstrates**:
- Service Independence
- API Gateway Pattern
- Inter-service Communication
- Distributed Architecture

**Services**:
- User Service (Port 8081)
- Product Service (Port 8082)
- Order Service (Port 8083)
- API Gateway (Port 8080)

**Tech Stack**: Go, Echo

**Run** (4 terminals):
```bash
# Terminal 1
cd microservices/user-service && go run main.go

# Terminal 2
cd microservices/product-service && go run main.go

# Terminal 3
cd microservices/order-service && go run main.go

# Terminal 4
cd microservices/api-gateway && go run main.go
```

**Access**: `http://localhost:8080/api/*`

---

### 6. Relationships & Integration (`relationships-integration/`)
**Topic**: Connecting Architecture Patterns, Design Principles, and Development Practices

**Demonstrates**:
- How Clean Architecture, DDD, SOLID, Design Patterns, and Microservices work together
- Integration points between different architectural concepts
- Building cohesive systems through pattern synergy

---

## 🚀 Quick Start

### Prerequisites
- Go 1.21 or higher
- Basic understanding of REST APIs
- Familiarity with Go modules

### Running Examples

Each example is self-contained with its own `go.mod` file:

```bash
# Choose an example
cd <example-directory>

# Install dependencies
go mod download

# Run the application
go run main.go  # or specific command from example README
```

### Testing Examples

```bash
# For examples with tests
go test ./...

# With verbose output
go test -v ./...
```

## 📚 Learning Path

**Recommended Order**:

1. **SOLID Principles** - Foundation concepts
2. **Design Patterns** - Reusable solutions
3. **Clean Architecture** - System structure
4. **Domain-Driven Design** - Business logic modeling
5. **Microservices** - Distributed systems
6. **Integration** - Putting it all together

## 🔗 Related Presentations

Each example corresponds to a presentation in `/design/`:

- `Clean-Architecture-Presentation.md`
- `DDD-Complete-Presentation.md`
- `GoF-Design-Patterns-Presentation.md`
- `Microservices-Architecture-Presentation.md`
- `SOLID-Principles-Presentation.md`
- `Relationships-Integration-Presentation.md`

## 🛠️ Technologies Used

- **Language**: Go 1.21
- **Web Framework**: Echo v4
- **Database**: SQLx + SQLite (for examples with persistence)
- **ID Generation**: Google UUID
- **Architecture**: RESTful APIs

## �� Key Concepts Demonstrated

### Architectural Patterns
- Layered Architecture
- Hexagonal Architecture (Ports & Adapters)
- CQRS concepts
- Event-Driven patterns
- Service-Oriented Architecture

### Design Principles
- SOLID principles
- DRY (Don't Repeat Yourself)
- KISS (Keep It Simple, Stupid)
- YAGNI (You Aren't Gonna Need It)
- Separation of Concerns
- Dependency Injection

### Domain Modeling
- Ubiquitous Language
- Bounded Contexts
- Entities vs Value Objects
- Aggregates
- Domain Events
- Domain Services

## 💡 Best Practices

All examples demonstrate:
- ✅ Clean code principles
- ✅ Testable architecture
- ✅ Proper error handling
- ✅ Interface-based design
- ✅ Dependency injection
- ✅ Clear separation of concerns
- ✅ Idiomatic Go

## 🤝 Contributing

These examples are meant to be educational. Feel free to:
- Extend them with additional features
- Add more test cases
- Implement additional patterns
- Create variations for different use cases

## 📖 Additional Resources

- [Clean Architecture by Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Domain-Driven Design by Eric Evans](https://www.domainlanguage.com/ddd/)
- [Design Patterns: Elements of Reusable Object-Oriented Software](https://en.wikipedia.org/wiki/Design_Patterns)
- [Building Microservices by Sam Newman](https://samnewman.io/books/building_microservices/)

## 📧 Contact

**Presenter**: Dong Tran  
**Date**: November 20, 2025

---

**Note**: Each example directory contains its own detailed README with specific instructions, API documentation, and architecture diagrams.
