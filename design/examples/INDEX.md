# Examples Index

Quick reference guide for all design pattern and architecture examples.

## Directory Structure

```
examples/
├── clean-architecture/          # Clean Architecture Pattern
│   ├── domain/                  # Enterprise Business Rules
│   ├── usecase/                 # Application Business Rules
│   ├── repository/              # Interface Adapters
│   ├── handler/                 # Frameworks & Drivers
│   ├── infrastructure/          # External Concerns
│   └── main.go                  # Entry Point
│
├── ddd/                         # Domain-Driven Design
│   ├── domain/
│   │   ├── model/               # Entities & Value Objects
│   │   ├── repository/          # Repository Interfaces
│   │   └── service/             # Domain Services
│   ├── application/             # Application Services
│   ├── infrastructure/          # Persistence & HTTP
│   └── cmd/                     # Entry Points
│
├── solid-principles/            # SOLID Principles
│   ├── srp/                     # Single Responsibility
│   ├── ocp/                     # Open/Closed
│   ├── lsp/                     # Liskov Substitution
│   ├── isp/                     # Interface Segregation
│   └── dip/                     # Dependency Inversion
│
├── design-patterns/             # GoF Design Patterns
│   ├── creational/              # Singleton, Factory, Builder
│   ├── structural/              # Adapter, Decorator
│   └── behavioral/              # Strategy, Observer
│
├── microservices/               # Microservices Architecture
│   ├── user-service/            # Port 8081
│   ├── product-service/         # Port 8082
│   ├── order-service/           # Port 8083
│   └── api-gateway/             # Port 8080
│
└── relationships-integration/   # Integration Example
    └── (Shows how all patterns work together)
```

## Examples by Complexity

### Beginner
1. **SOLID Principles** - Start here to understand foundational concepts
2. **Design Patterns** - Learn reusable solutions

### Intermediate
3. **Clean Architecture** - Understand layered architecture
4. **DDD** - Model business domains effectively

### Advanced
5. **Microservices** - Build distributed systems
6. **Integration** - Combine all concepts

## Examples by Use Case

### Building a Single Service
- ✅ Clean Architecture
- ✅ DDD
- ✅ SOLID Principles
- ✅ Design Patterns

### Building Multiple Services
- ✅ Microservices
- ✅ Integration Example

### Learning Patterns
- ✅ Design Patterns (23 GoF patterns)
- ✅ SOLID Principles (5 principles)

## Technology Stack per Example

| Example | Go | Echo | SQLx | SQLite | UUID |
|---------|-----|------|------|--------|------|
| Clean Architecture | ✅ | ✅ | ✅ | ✅ | ❌ |
| DDD | ✅ | ✅ | ✅ | ❌ | ✅ |
| SOLID | ✅ | ❌ | ❌ | ❌ | ❌ |
| Design Patterns | ✅ | ❌ | ❌ | ❌ | ❌ |
| Microservices | ✅ | ✅ | ❌ | ❌ | ❌ |
| Integration | ✅ | ✅ | ✅ | ❌ | ✅ |

## Running Examples Summary

### Clean Architecture
```bash
cd clean-architecture && go run main.go
# http://localhost:8080
```

### DDD
```bash
cd ddd && go run cmd/main.go
# Implementation pending - see README
```

### SOLID Principles
```bash
cd solid-principles && go test ./...
```

### Design Patterns
```bash
cd design-patterns && go test ./...
```

### Microservices (4 terminals needed)
```bash
# Terminal 1: cd microservices/user-service && go run main.go
# Terminal 2: cd microservices/product-service && go run main.go
# Terminal 3: cd microservices/order-service && go run main.go
# Terminal 4: cd microservices/api-gateway && go run main.go
# Access via: http://localhost:8080
```

### Integration
```bash
cd relationships-integration && go run cmd/main.go
# See README for details
```

## File Count

- **Go Files**: 20+
- **README Files**: 7
- **Directories**: 28
- **Lines of Code**: 1500+

## Concepts Covered

### Architecture Patterns
- [x] Clean Architecture
- [x] Domain-Driven Design
- [x] Microservices Architecture
- [x] Layered Architecture
- [x] Hexagonal Architecture (implied in Clean)

### SOLID Principles
- [x] Single Responsibility Principle
- [x] Open/Closed Principle
- [x] Liskov Substitution Principle
- [x] Interface Segregation Principle
- [x] Dependency Inversion Principle

### Design Patterns (GoF)
**Creational**:
- [x] Singleton
- [x] Factory Method
- [x] Builder

**Structural**:
- [x] Adapter
- [x] Decorator

**Behavioral**:
- [x] Strategy
- [x] Observer

### DDD Building Blocks
- [x] Entities
- [x] Value Objects
- [x] Aggregates
- [x] Repositories
- [x] Domain Services
- [x] Application Services

## API Endpoints Reference

### Clean Architecture (Port 8080)
```
POST   /tasks          Create task
GET    /tasks/:id      Get task
GET    /tasks          List tasks
PUT    /tasks/:id      Update task
DELETE /tasks/:id      Delete task
```

### Microservices (via API Gateway - Port 8080)
```
GET    /api/users/:id        User service
POST   /api/users            User service
GET    /api/products/:id     Product service
GET    /api/products         Product service
POST   /api/orders           Order service
GET    /api/orders/:id       Order service
```

## Testing Examples

All examples include either:
- Runnable applications with REST APIs
- Test files demonstrating usage
- README with curl examples

## Next Steps

1. Read the main `README.md` for overview
2. Choose an example based on your learning goal
3. Read that example's specific README
4. Run the example
5. Modify and experiment
6. Refer back to the presentations in `/design/`

## Related Documentation

- `/design/Clean-Architecture-Presentation.md`
- `/design/DDD-Complete-Presentation.md`
- `/design/GoF-Design-Patterns-Presentation.md`
- `/design/Microservices-Architecture-Presentation.md`
- `/design/SOLID-Principles-Presentation.md`
- `/design/Relationships-Integration-Presentation.md`

---

**Created**: November 20, 2025  
**Author**: Dong Tran  
**Language**: Go 1.21  
**Framework**: Echo v4  
**Database**: SQLx + SQLite
