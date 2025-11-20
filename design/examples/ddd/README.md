# Domain-Driven Design (DDD) Example - E-Commerce Product Catalog

This example demonstrates DDD tactical patterns with Go, Echo, and SQLx.

## DDD Concepts Demonstrated

### Building Blocks

1. **Entities**: Product (has identity)
2. **Value Objects**: Money, Category (immutable, no identity)
3. **Aggregates**: Product is an aggregate root
4. **Repositories**: ProductRepository interface
5. **Domain Services**: PricingService
6. **Application Services**: ProductService

## Project Structure

```
.
├── domain/
│   ├── model/              # Entities and Value Objects
│   │   └── product.go
│   ├── repository/         # Repository interfaces
│   │   └── product_repository.go
│   └── service/            # Domain services
│       └── pricing_service.go
├── application/            # Application services
│   └── product_service.go
├── infrastructure/
│   ├── persistence/        # Repository implementations
│   └── http/              # HTTP handlers
└── cmd/                   # Application entry point
```

## Running

```bash
go mod download
go run cmd/main.go
```

## API Examples

```bash
# Create product
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop",
    "description": "High-performance laptop",
    "price": 999.99,
    "currency": "USD",
    "category": "Electronics"
  }'

# Get all products
curl http://localhost:8080/products

# Apply discount
curl -X POST http://localhost:8080/products/{id}/discount \
  -H "Content-Type: application/json" \
  -d '{"discount": 10}'
```

## Key DDD Principles

1. **Ubiquitous Language**: Code reflects business domain
2. **Rich Domain Model**: Business logic in domain objects
3. **Bounded Contexts**: Clear boundaries
4. **Aggregate Roots**: Consistency boundaries
