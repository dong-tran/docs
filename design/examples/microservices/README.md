# Microservices Architecture Example

Demonstrates microservices patterns with Go and Echo.

## Services

- **User Service** (Port 8081): Manages users
- **Product Service** (Port 8082): Manages products
- **Order Service** (Port 8083): Manages orders
- **API Gateway** (Port 8080): Routes requests

## Running

```bash
# Terminal 1 - User Service
cd user-service && go run main.go

# Terminal 2 - Product Service
cd product-service && go run main.go

# Terminal 3 - Order Service
cd order-service && go run main.go

# Terminal 4 - API Gateway
cd api-gateway && go run main.go
```

## Testing

```bash
# Via API Gateway
curl http://localhost:8080/api/users/1
curl http://localhost:8080/api/products
curl -X POST http://localhost:8080/api/orders \
  -H "Content-Type: application/json" \
  -d '{"user_id":"1","product_id":"1","total":999.99}'
```

## Key Concepts

- Service Independence
- API Gateway Pattern
- Service Discovery
- Inter-service Communication
- Distributed Data Management
