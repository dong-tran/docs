# Clean Architecture Example - Task Management System

This example demonstrates Clean Architecture principles with Go, Echo, and SQLx.

## Architecture Layers

```
┌─────────────────────────────────────────┐
│         Frameworks & Drivers            │
│    (HTTP Handlers, Database, Echo)      │
└─────────────────────────────────────────┘
              ↓ depends on
┌─────────────────────────────────────────┐
│        Interface Adapters               │
│    (Controllers, Presenters, Gateways)  │
└─────────────────────────────────────────┘
              ↓ depends on
┌─────────────────────────────────────────┐
│        Application Business Rules       │
│           (Use Cases)                   │
└─────────────────────────────────────────┘
              ↓ depends on
┌─────────────────────────────────────────┐
│       Enterprise Business Rules         │
│            (Entities)                   │
└─────────────────────────────────────────┘
```

## Project Structure

```
.
├── domain/              # Enterprise Business Rules (innermost layer)
│   └── task.go         # Task entity with business rules
├── usecase/            # Application Business Rules
│   └── task_usecase.go # Use cases for task operations
├── repository/         # Interface Adapters - Data Access
│   └── task_repository.go
├── handler/            # Frameworks & Drivers - HTTP handlers
│   └── task_handler.go
├── infrastructure/     # Frameworks & Drivers - External concerns
│   └── database.go
└── main.go            # Application entry point
```

## Running the Example

```bash
# Install dependencies
go mod download

# Run the application
go run main.go

# The server will start on http://localhost:8080
```

## API Endpoints

- `POST /tasks` - Create a new task
- `GET /tasks/:id` - Get a task by ID
- `GET /tasks` - List all tasks
- `PUT /tasks/:id` - Update a task
- `DELETE /tasks/:id` - Delete a task

## Testing with curl

```bash
# Create a task
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Clean Architecture","description":"Study the principles"}'

# Get all tasks
curl http://localhost:8080/tasks

# Get a specific task
curl http://localhost:8080/tasks/1

# Update a task
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Master Clean Architecture","description":"Apply in projects","completed":true}'

# Delete a task
curl -X DELETE http://localhost:8080/tasks/1
```

## Key Clean Architecture Principles Demonstrated

1. **Dependency Rule**: Dependencies point inward. Domain layer has no dependencies.
2. **Independence**: Business logic is independent of frameworks, UI, and database.
3. **Testability**: Each layer can be tested independently.
4. **Flexibility**: Easy to swap implementations (e.g., change database or framework).
