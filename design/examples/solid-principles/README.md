# SOLID Principles Examples in Go

Demonstrating all five SOLID principles with practical examples.

## Examples

### 1. Single Responsibility Principle (SRP)
**srp/user_service.go** - Each class has one reason to change

### 2. Open/Closed Principle (OCP)
**ocp/notification.go** - Open for extension, closed for modification

### 3. Liskov Substitution Principle (LSP)
**lsp/shapes.go** - Subtypes must be substitutable for their base types

### 4. Interface Segregation Principle (ISP)
**isp/worker.go** - Clients shouldn't depend on interfaces they don't use

### 5. Dependency Inversion Principle (DIP)
**dip/payment.go** - Depend on abstractions, not concretions

## Running

```bash
go test ./...
```

## Key Takeaways

- **SRP**: One class, one responsibility
- **OCP**: Extend behavior without modification
- **LSP**: Substitutability without breaking functionality  
- **ISP**: Small, focused interfaces
- **DIP**: Depend on abstractions for flexibility
