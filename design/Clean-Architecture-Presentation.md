# Clean Architecture
## Building Maintainable and Testable Software Systems

**Presenter:** Dong Tran  
**Date:** November 20, 2025  
**Time:** 02:45 UTC

---

## Slide 1: Title Slide

# Clean Architecture
### The Dependency Rule and Building Maintainable Systems

**Presenter:** Dong Tran  
**Date:** November 20, 2025

*"The goal of software architecture is to minimize the human resources required to build and maintain the required system."*  
— Robert C. Martin

---

## Slide 2: Agenda

### Today's Journey

1. **The Problems We're Solving**
2. **What is Clean Architecture?**
3. **The Dependency Rule**
4. **The Four Layers**
5. **Implementation in Go**
6. **Real-World Example: Task Management System**
7. **Testing Strategies**
8. **Clean Architecture vs Other Patterns**
9. **Best Practices & Anti-Patterns**
10. **Migration Strategy**
11. **Q&A Session**

---

## Slide 3: The Problems We're Solving

### Common Architecture Pain Points

**Without Proper Architecture:**
- 🔥 **Framework Lock-in** - Married to specific frameworks
- 🔥 **Database Coupling** - Business logic mixed with SQL
- 🔥 **Testing Nightmare** - Can't test without the database
- 🔥 **Rigid Systems** - Hard to add new features
- 🔥 **Technical Debt** - Changes ripple through entire system

**The Cost:**
- 70% of software cost is maintenance
- 60% of maintenance is understanding existing code
- Simple changes take weeks instead of hours

---

## Slide 4: What is Clean Architecture?

### Clean Architecture Overview

**Definition:**
An architecture pattern that separates concerns through well-defined layers with a strict dependency rule.

**Created by Robert C. Martin (Uncle Bob) in 2012**

**Core Principles:**
1. **Independent of Frameworks** - Architecture doesn't depend on libraries
2. **Testable** - Business rules can be tested without UI, Database, or external elements
3. **Independent of UI** - UI can change without changing the system
4. **Independent of Database** - Switch from Oracle to MongoDB without changing business rules
5. **Independent of External Services** - Business rules don't know about the outside world

---

## Slide 5: The Dependency Rule

### The Dependency Rule

**The Golden Rule:**
> Source code dependencies must point only inward, toward higher-level policies

```
┌─────────────────────────────────────┐
│                                     │
│         Entities (Core)             │ ← Most Stable
│                                     │
├─────────────────────────────────────┤
│                                     │
│         Use Cases                   │ ← Business Rules
│                                     │
├─────────────────────────────────────┤
│                                     │
│     Interface Adapters              │ ← Translation Layer
│                                     │
├─────────────────────────────────────┤
│                                     │
│    Frameworks & Drivers             │ ← Most Volatile
│                                     │
└─────────────────────────────────────┘
        ↑                    ↑
    Dependencies        Stability
    Point Inward       Increases
```

**Key Points:**
- Inner circles don't know about outer circles
- Data flows across boundaries through simple data structures
- We use interfaces (dependency inversion) when needed

---

## Slide 6: The Four Layers Explained

### Layer 1: Entities (Enterprise Business Rules)

**What They Are:**
- Core business objects
- Domain models with business logic
- The most stable layer

**Characteristics:**
- No dependencies on any other layer
- Pure business logic
- Could be used by multiple applications

**Example:**
```go
// Entity - Pure business logic
type Task struct {
    ID          string
    Title       string
    Description string
    Status      TaskStatus
    Priority    Priority
    DueDate     time.Time
    CreatedAt   time.Time
}

func (t *Task) IsOverdue() bool {
    return t.Status != Completed && 
           time.Now().After(t.DueDate)
}

func (t *Task) CanBeAssignedTo(user User) bool {
    return user.HasPermission(PermissionAssignTask) &&
           t.Status == Open
}
```

---

## Slide 7: The Four Layers - Use Cases

### Layer 2: Use Cases (Application Business Rules)

**What They Are:**
- Application-specific business rules
- Orchestrate data flow to/from entities
- Implement use case scenarios

**Characteristics:**
- Contain application business logic
- Call entity methods
- Don't know about delivery mechanism

**Example:**
```go
// Use Case - Application business logic
type CreateTaskUseCase struct {
    taskRepo TaskRepository
    userRepo UserRepository
    notifier Notifier
}

type CreateTaskInput struct {
    Title       string
    Description string
    AssigneeID  string
    DueDate     time.Time
}

func (uc *CreateTaskUseCase) Execute(
    ctx context.Context,
    input CreateTaskInput,
) (*Task, error) {
    // Validate input
    if err := uc.validate(input); err != nil {
        return nil, err
    }
    
    // Create entity
    task := &Task{
        ID:          GenerateID(),
        Title:       input.Title,
        Description: input.Description,
        Status:      Open,
        DueDate:     input.DueDate,
        CreatedAt:   time.Now(),
    }
    
    // Business rule check
    if input.AssigneeID != "" {
        assignee, err := uc.userRepo.FindByID(ctx, input.AssigneeID)
        if err != nil {
            return nil, err
        }
        
        if !task.CanBeAssignedTo(assignee) {
            return nil, ErrCannotAssign
        }
    }
    
    // Save to repository
    if err := uc.taskRepo.Save(ctx, task); err != nil {
        return nil, err
    }
    
    // Send notification
    uc.notifier.NotifyTaskCreated(task)
    
    return task, nil
}
```

---

## Slide 8: The Four Layers - Interface Adapters

### Layer 3: Interface Adapters

**What They Are:**
- Convert data between use cases and external systems
- Controllers, Presenters, Gateways

**Characteristics:**
- Know about use cases
- Convert data formats
- No business logic

**Example:**
```go
// HTTP Controller - Interface Adapter
type TaskController struct {
    createTask *CreateTaskUseCase
    listTasks  *ListTasksUseCase
}

type CreateTaskRequest struct {
    Title       string    `json:"title" binding:"required"`
    Description string    `json:"description"`
    AssigneeID  string    `json:"assignee_id"`
    DueDate     string    `json:"due_date"`
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
    var req CreateTaskRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    // Convert HTTP request to use case input
    dueDate, _ := time.Parse(time.RFC3339, req.DueDate)
    input := CreateTaskInput{
        Title:       req.Title,
        Description: req.Description,
        AssigneeID:  req.AssigneeID,
        DueDate:     dueDate,
    }
    
    // Execute use case
    task, err := c.createTask.Execute(ctx, input)
    if err != nil {
        ctx.JSON(500, gin.H{"error": err.Error()})
        return
    }
    
    // Convert entity to HTTP response
    ctx.JSON(201, c.toResponse(task))
}

func (c *TaskController) toResponse(task *Task) map[string]interface{} {
    return map[string]interface{}{
        "id":          task.ID,
        "title":       task.Title,
        "description": task.Description,
        "status":      task.Status.String(),
        "due_date":    task.DueDate.Format(time.RFC3339),
        "is_overdue":  task.IsOverdue(),
    }
}
```

---

## Slide 9: The Four Layers - Frameworks & Drivers

### Layer 4: Frameworks & Drivers

**What They Are:**
- Frameworks and tools
- Database, Web framework, UI
- External services

**Characteristics:**
- Most detailed layer
- Most likely to change
- Glue code

**Example:**
```go
// PostgreSQL Repository - Framework Layer
type PostgresTaskRepository struct {
    db *sql.DB
}

func (r *PostgresTaskRepository) Save(ctx context.Context, task *Task) error {
    query := `
        INSERT INTO tasks (id, title, description, status, priority, due_date, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        ON CONFLICT (id) DO UPDATE SET
            title = $2,
            description = $3,
            status = $4,
            priority = $5,
            due_date = $6
    `
    
    _, err := r.db.ExecContext(ctx, query,
        task.ID,
        task.Title,
        task.Description,
        task.Status,
        task.Priority,
        task.DueDate,
        task.CreatedAt,
    )
    
    return err
}

func (r *PostgresTaskRepository) FindByID(ctx context.Context, id string) (*Task, error) {
    query := `
        SELECT id, title, description, status, priority, due_date, created_at
        FROM tasks
        WHERE id = $1
    `
    
    var task Task
    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &task.ID,
        &task.Title,
        &task.Description,
        &task.Status,
        &task.Priority,
        &task.DueDate,
        &task.CreatedAt,
    )
    
    if err == sql.ErrNoRows {
        return nil, ErrTaskNotFound
    }
    
    return &task, err
}
```

---

## Slide 10: Project Structure in Go

### Clean Architecture Project Structure

```
project/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
│
├── internal/
│   ├── domain/                     # Layer 1: Entities
│   │   ├── entity/
│   │   │   ├── task.go
│   │   │   ├── user.go
│   │   │   └── project.go
│   │   └── valueobject/
│   │       ├── status.go
│   │       └── priority.go
│   │
│   ├── usecase/                    # Layer 2: Use Cases
│   │   ├── task/
│   │   │   ├── create.go
│   │   │   ├── update.go
│   │   │   ├── delete.go
│   │   │   └── list.go
│   │   ├── user/
│   │   └── interfaces.go           # Repository interfaces
│   │
│   ├── adapter/                    # Layer 3: Interface Adapters
│   │   ├── controller/
│   │   │   ├── http/
│   │   │   │   ├── task.go
│   │   │   │   └── user.go
│   │   │   └── grpc/
│   │   ├── presenter/
│   │   │   └── task_presenter.go
│   │   └── gateway/
│   │       └── notification.go
│   │
│   └── infrastructure/             # Layer 4: Frameworks & Drivers
│       ├── persistence/
│       │   ├── postgres/
│       │   │   ├── task_repository.go
│       │   │   └── user_repository.go
│       │   └── mongodb/
│       ├── web/
│       │   └── router.go
│       └── config/
│           └── config.go
│
├── pkg/                            # Shared packages
│   ├── logger/
│   └── validator/
│
├── go.mod
└── go.sum
```

---

## Slide 11: Dependency Injection & Wiring

### Wiring It All Together

```go
// cmd/api/main.go - Application bootstrap
package main

import (
    "database/sql"
    "log"
    
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
    
    "project/internal/adapter/controller/http"
    "project/internal/infrastructure/persistence/postgres"
    "project/internal/usecase/task"
)

func main() {
    // Setup infrastructure
    db := setupDatabase()
    
    // Initialize repositories (Layer 4)
    taskRepo := postgres.NewTaskRepository(db)
    userRepo := postgres.NewUserRepository(db)
    
    // Initialize use cases (Layer 2)
    createTaskUC := task.NewCreateTaskUseCase(
        taskRepo,
        userRepo,
        setupNotifier(),
    )
    listTasksUC := task.NewListTasksUseCase(taskRepo)
    
    // Initialize controllers (Layer 3)
    taskController := http.NewTaskController(
        createTaskUC,
        listTasksUC,
    )
    
    // Setup routes (Layer 4)
    router := gin.Default()
    
    api := router.Group("/api/v1")
    {
        api.POST("/tasks", taskController.CreateTask)
        api.GET("/tasks", taskController.ListTasks)
        api.GET("/tasks/:id", taskController.GetTask)
        api.PUT("/tasks/:id", taskController.UpdateTask)
        api.DELETE("/tasks/:id", taskController.DeleteTask)
    }
    
    log.Fatal(router.Run(":8080"))
}

func setupDatabase() *sql.DB {
    db, err := sql.Open("postgres", 
        "postgres://user:pass@localhost/taskdb?sslmode=disable")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    
    if err := db.Ping(); err != nil {
        log.Fatal("Failed to ping database:", err)
    }
    
    return db
}
```

---

## Slide 12: Repository Pattern & Interfaces

### Repository Pattern in Clean Architecture

```go
// internal/usecase/interfaces.go - Use Case Layer
package usecase

import (
    "context"
    "project/internal/domain/entity"
)

// Repository interfaces belong to use case layer
// Implementation belongs to infrastructure layer

type TaskRepository interface {
    Save(ctx context.Context, task *entity.Task) error
    FindByID(ctx context.Context, id string) (*entity.Task, error)
    FindAll(ctx context.Context, filter TaskFilter) ([]*entity.Task, error)
    Update(ctx context.Context, task *entity.Task) error
    Delete(ctx context.Context, id string) error
}

type UserRepository interface {
    FindByID(ctx context.Context, id string) (*entity.User, error)
    FindByEmail(ctx context.Context, email string) (*entity.User, error)
    Save(ctx context.Context, user *entity.User) error
}

type Notifier interface {
    NotifyTaskCreated(task *entity.Task) error
    NotifyTaskAssigned(task *entity.Task, user *entity.User) error
}

// Filter object for queries
type TaskFilter struct {
    Status     *entity.TaskStatus
    AssigneeID *string
    ProjectID  *string
    Limit      int
    Offset     int
}
```

---

## Slide 13: Testing Strategy

### Testing in Clean Architecture

**Testing Pyramid:**
```
        ┌─────┐
        │ E2E │       ← Few tests
       ┌┴─────┴┐
       │ Integ │      ← Some tests
      ┌┴───────┴┐
      │  Unit   │     ← Many tests
     └───────────┘
```

**Unit Testing Use Cases:**
```go
// internal/usecase/task/create_test.go
package task

import (
    "context"
    "testing"
    "time"
    
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    
    "project/internal/domain/entity"
)

// Mock repository
type MockTaskRepository struct {
    mock.Mock
}

func (m *MockTaskRepository) Save(ctx context.Context, task *entity.Task) error {
    args := m.Called(ctx, task)
    return args.Error(0)
}

// Test use case
func TestCreateTaskUseCase_Execute(t *testing.T) {
    // Arrange
    mockRepo := new(MockTaskRepository)
    mockUserRepo := new(MockUserRepository)
    mockNotifier := new(MockNotifier)
    
    useCase := NewCreateTaskUseCase(mockRepo, mockUserRepo, mockNotifier)
    
    input := CreateTaskInput{
        Title:       "Test Task",
        Description: "Test Description",
        DueDate:     time.Now().Add(24 * time.Hour),
    }
    
    // Set expectations
    mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
    mockNotifier.On("NotifyTaskCreated", mock.Anything).Return(nil)
    
    // Act
    task, err := useCase.Execute(context.Background(), input)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, task)
    assert.Equal(t, "Test Task", task.Title)
    assert.Equal(t, entity.Open, task.Status)
    
    // Verify expectations
    mockRepo.AssertExpectations(t)
    mockNotifier.AssertExpectations(t)
}

func TestCreateTaskUseCase_ValidationError(t *testing.T) {
    // Test validation failures
    useCase := NewCreateTaskUseCase(nil, nil, nil)
    
    testCases := []struct {
        name  string
        input CreateTaskInput
        error string
    }{
        {
            name: "Empty title",
            input: CreateTaskInput{
                Title: "",
            },
            error: "title is required",
        },
        {
            name: "Past due date",
            input: CreateTaskInput{
                Title:   "Task",
                DueDate: time.Now().Add(-24 * time.Hour),
            },
            error: "due date must be in the future",
        },
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            _, err := useCase.Execute(context.Background(), tc.input)
            assert.Error(t, err)
            assert.Contains(t, err.Error(), tc.error)
        })
    }
}
```

---

## Slide 14: Integration Testing

### Integration Testing with Clean Architecture

```go
// internal/adapter/controller/http/task_integration_test.go
package http_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    
    "project/internal/adapter/controller/http"
    "project/internal/infrastructure/persistence/memory"
    "project/internal/usecase/task"
)

func TestTaskController_Integration(t *testing.T) {
    // Setup with in-memory repository
    taskRepo := memory.NewTaskRepository()
    userRepo := memory.NewUserRepository()
    notifier := &NoOpNotifier{}
    
    // Wire up real dependencies
    createUC := task.NewCreateTaskUseCase(taskRepo, userRepo, notifier)
    listUC := task.NewListTasksUseCase(taskRepo)
    
    controller := http.NewTaskController(createUC, listUC)
    
    // Setup router
    gin.SetMode(gin.TestMode)
    router := gin.New()
    router.POST("/tasks", controller.CreateTask)
    router.GET("/tasks", controller.ListTasks)
    
    t.Run("Create and List Tasks", func(t *testing.T) {
        // Create task
        createReq := map[string]interface{}{
            "title":       "Integration Test Task",
            "description": "Testing the full flow",
            "due_date":    "2025-12-01T00:00:00Z",
        }
        
        body, _ := json.Marshal(createReq)
        req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
        req.Header.Set("Content-Type", "application/json")
        w := httptest.NewRecorder()
        
        router.ServeHTTP(w, req)
        
        assert.Equal(t, http.StatusCreated, w.Code)
        
        var createResp map[string]interface{}
        json.Unmarshal(w.Body.Bytes(), &createResp)
        assert.Equal(t, "Integration Test Task", createResp["title"])
        
        // List tasks
        req = httptest.NewRequest("GET", "/tasks", nil)
        w = httptest.NewRecorder()
        
        router.ServeHTTP(w, req)
        
        assert.Equal(t, http.StatusOK, w.Code)
        
        var listResp []map[string]interface{}
        json.Unmarshal(w.Body.Bytes(), &listResp)
        assert.Len(t, listResp, 1)
        assert.Equal(t, "Integration Test Task", listResp[0]["title"])
    })
}
```

---

## Slide 15: Clean Architecture vs Other Architectures

### Comparing Architecture Patterns

| Aspect | Clean Architecture | Hexagonal | Onion | Traditional Layered |
|--------|-------------------|-----------|-------|-------------------|
| **Core Concept** | Dependency Rule | Ports & Adapters | Domain at Center | Top-down layers |
| **Layers** | 4 concentric | 3 (Domain, App, Infra) | 4 concentric | 3-4 horizontal |
| **Dependencies** | Inward only | Inward only | Inward only | Top to bottom |
| **Business Logic** | Entities & Use Cases | Domain | Domain Model | Mixed |
| **Testing** | Excellent | Excellent | Excellent | Moderate |
| **Complexity** | High | Medium | Medium | Low |
| **Learning Curve** | Steep | Moderate | Moderate | Easy |
| **Best For** | Large, complex systems | Medium systems | DDD projects | Simple CRUD |

**Visual Comparison:**
```
Clean Architecture        Hexagonal              Traditional Layered
    ┌─────┐                ┌─────┐                  ┌─────┐
    │Core │                │Domain│                  │ UI  │
   ┌┴─────┴┐              ┌┴─────┴┐                ┌┴─────┴┐
   │UseCase│              │  App  │                │Business│
  ┌┴───────┴┐            ┌┴───────┴┐              ┌┴────────┴┐
  │Interface│            │  Infra  │              │   Data    │
 ┌┴─────────┴┐          └─────────┘              └───────────┘
 │ Framework │              
 └───────────┘
```

---

## Slide 16: Real-World Example - Complete Use Case

### Complete Task Management Feature

```go
// internal/domain/entity/task.go
package entity

import (
    "errors"
    "time"
)

type TaskStatus int

const (
    Open TaskStatus = iota
    InProgress
    Completed
    Cancelled
)

type Priority int

const (
    Low Priority = iota
    Medium
    High
    Critical
)

type Task struct {
    ID          string
    Title       string
    Description string
    Status      TaskStatus
    Priority    Priority
    AssigneeID  string
    ProjectID   string
    Tags        []string
    DueDate     time.Time
    CompletedAt *time.Time
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// Business rules
func (t *Task) Complete() error {
    if t.Status == Cancelled {
        return errors.New("cannot complete cancelled task")
    }
    
    if t.Status == Completed {
        return errors.New("task already completed")
    }
    
    now := time.Now()
    t.Status = Completed
    t.CompletedAt = &now
    t.UpdatedAt = now
    
    return nil
}

func (t *Task) AssignTo(userID string) error {
    if t.Status == Completed || t.Status == Cancelled {
        return errors.New("cannot assign completed or cancelled task")
    }
    
    t.AssigneeID = userID
    t.UpdatedAt = time.Now()
    
    if t.Status == Open {
        t.Status = InProgress
    }
    
    return nil
}

func (t *Task) UpdatePriority(priority Priority) {
    t.Priority = priority
    t.UpdatedAt = time.Now()
}

func (t *Task) IsOverdue() bool {
    return t.Status != Completed && 
           t.Status != Cancelled &&
           time.Now().After(t.DueDate)
}

func (t *Task) DaysUntilDue() int {
    if t.IsOverdue() {
        return 0
    }
    
    duration := time.Until(t.DueDate)
    return int(duration.Hours() / 24)
}
```

---

## Slide 17: Complex Use Case Implementation

### Advanced Use Case with Multiple Dependencies

```go
// internal/usecase/task/assign_task.go
package task

import (
    "context"
    "errors"
    "time"
    
    "project/internal/domain/entity"
    "project/internal/usecase"
)

type AssignTaskUseCase struct {
    taskRepo     usecase.TaskRepository
    userRepo     usecase.UserRepository
    projectRepo  usecase.ProjectRepository
    notifier     usecase.Notifier
    auditLogger  usecase.AuditLogger
}

type AssignTaskInput struct {
    TaskID     string
    AssigneeID string
    AssignerID string
    Notes      string
}

func (uc *AssignTaskUseCase) Execute(
    ctx context.Context,
    input AssignTaskInput,
) error {
    // 1. Load task
    task, err := uc.taskRepo.FindByID(ctx, input.TaskID)
    if err != nil {
        return err
    }
    
    // 2. Load assignee
    assignee, err := uc.userRepo.FindByID(ctx, input.AssigneeID)
    if err != nil {
        return err
    }
    
    // 3. Load assigner
    assigner, err := uc.userRepo.FindByID(ctx, input.AssignerID)
    if err != nil {
        return err
    }
    
    // 4. Check permissions
    if !assigner.CanAssignTasks() {
        return errors.New("user lacks permission to assign tasks")
    }
    
    // 5. Check if assignee is part of the project
    if task.ProjectID != "" {
        project, err := uc.projectRepo.FindByID(ctx, task.ProjectID)
        if err != nil {
            return err
        }
        
        if !project.HasMember(assignee.ID) {
            return errors.New("assignee is not a member of the project")
        }
    }
    
    // 6. Check assignee workload
    assignedTasks, err := uc.taskRepo.FindAll(ctx, usecase.TaskFilter{
        AssigneeID: &assignee.ID,
        Status:     &entity.InProgress,
    })
    
    if err != nil {
        return err
    }
    
    if len(assignedTasks) >= assignee.MaxConcurrentTasks {
        return errors.New("assignee has reached maximum concurrent tasks")
    }
    
    // 7. Perform assignment (domain logic)
    previousAssignee := task.AssigneeID
    if err := task.AssignTo(assignee.ID); err != nil {
        return err
    }
    
    // 8. Save changes
    if err := uc.taskRepo.Update(ctx, task); err != nil {
        return err
    }
    
    // 9. Send notifications
    if err := uc.notifier.NotifyTaskAssigned(task, assignee); err != nil {
        // Log error but don't fail the operation
        // Notification is not critical
    }
    
    // 10. Audit log
    uc.auditLogger.Log(ctx, usecase.AuditEntry{
        Action:     "TASK_ASSIGNED",
        EntityType: "task",
        EntityID:   task.ID,
        UserID:     assigner.ID,
        Changes: map[string]interface{}{
            "previous_assignee": previousAssignee,
            "new_assignee":      assignee.ID,
            "notes":             input.Notes,
        },
        Timestamp: time.Now(),
    })
    
    return nil
}
```

---

## Slide 18: Presenter Pattern

### Presenter Pattern for Output Formatting

```go
// internal/adapter/presenter/task_presenter.go
package presenter

import (
    "time"
    "project/internal/domain/entity"
)

type TaskPresenter struct{}

// Different view models for different contexts
type TaskSummaryView struct {
    ID       string `json:"id"`
    Title    string `json:"title"`
    Status   string `json:"status"`
    Priority string `json:"priority"`
    DueDate  string `json:"due_date"`
}

type TaskDetailView struct {
    ID          string              `json:"id"`
    Title       string              `json:"title"`
    Description string              `json:"description"`
    Status      string              `json:"status"`
    Priority    string              `json:"priority"`
    Assignee    *UserSummaryView    `json:"assignee,omitempty"`
    Project     *ProjectSummaryView `json:"project,omitempty"`
    Tags        []string            `json:"tags"`
    DueDate     string              `json:"due_date"`
    IsOverdue   bool                `json:"is_overdue"`
    DaysUntilDue int                `json:"days_until_due"`
    CompletedAt *string             `json:"completed_at,omitempty"`
    CreatedAt   string              `json:"created_at"`
    UpdatedAt   string              `json:"updated_at"`
}

type TaskListView struct {
    Tasks      []TaskSummaryView `json:"tasks"`
    TotalCount int               `json:"total_count"`
    Page       int               `json:"page"`
    PageSize   int               `json:"page_size"`
}

func (p *TaskPresenter) PresentSummary(task *entity.Task) TaskSummaryView {
    return TaskSummaryView{
        ID:       task.ID,
        Title:    task.Title,
        Status:   p.statusToString(task.Status),
        Priority: p.priorityToString(task.Priority),
        DueDate:  task.DueDate.Format(time.RFC3339),
    }
}

func (p *TaskPresenter) PresentDetail(
    task *entity.Task,
    assignee *entity.User,
    project *entity.Project,
) TaskDetailView {
    view := TaskDetailView{
        ID:           task.ID,
        Title:        task.Title,
        Description:  task.Description,
        Status:       p.statusToString(task.Status),
        Priority:     p.priorityToString(task.Priority),
        Tags:         task.Tags,
        DueDate:      task.DueDate.Format(time.RFC3339),
        IsOverdue:    task.IsOverdue(),
        DaysUntilDue: task.DaysUntilDue(),
        CreatedAt:    task.CreatedAt.Format(time.RFC3339),
        UpdatedAt:    task.UpdatedAt.Format(time.RFC3339),
    }
    
    if task.CompletedAt != nil {
        completed := task.CompletedAt.Format(time.RFC3339)
        view.CompletedAt = &completed
    }
    
    if assignee != nil {
        view.Assignee = &UserSummaryView{
            ID:   assignee.ID,
            Name: assignee.Name,
        }
    }
    
    if project != nil {
        view.Project = &ProjectSummaryView{
            ID:   project.ID,
            Name: project.Name,
        }
    }
    
    return view
}

func (p *TaskPresenter) statusToString(status entity.TaskStatus) string {
    switch status {
    case entity.Open:
        return "open"
    case entity.InProgress:
        return "in_progress"
    case entity.Completed:
        return "completed"
    case entity.Cancelled:
        return "cancelled"
    default:
        return "unknown"
    }
}

func (p *TaskPresenter) priorityToString(priority entity.Priority) string {
    switch priority {
    case entity.Low:
        return "low"
    case entity.Medium:
        return "medium"
    case entity.High:
        return "high"
    case entity.Critical:
        return "critical"
    default:
        return "medium"
    }
}
```

---

## Slide 19: Error Handling

### Clean Error Handling Across Layers

```go
// internal/domain/errors/errors.go
package errors

import "fmt"

// Domain errors
type DomainError struct {
    Code    string
    Message string
}

func (e DomainError) Error() string {
    return e.Message
}

var (
    ErrTaskNotFound = DomainError{
        Code:    "TASK_NOT_FOUND",
        Message: "task not found",
    }
    
    ErrInvalidTaskStatus = DomainError{
        Code:    "INVALID_TASK_STATUS",
        Message: "invalid task status transition",
    }
    
    ErrUserNotAuthorized = DomainError{
        Code:    "USER_NOT_AUTHORIZED",
        Message: "user is not authorized for this action",
    }
)

// Use case errors
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Message)
}

type BusinessRuleError struct {
    Rule    string
    Message string
}

func (e BusinessRuleError) Error() string {
    return fmt.Sprintf("business rule violation: %s", e.Message)
}

// Error handler in controller
func HandleError(err error) (int, map[string]interface{}) {
    switch e := err.(type) {
    case DomainError:
        switch e.Code {
        case "TASK_NOT_FOUND":
            return 404, map[string]interface{}{
                "error": e.Message,
                "code":  e.Code,
            }
        case "USER_NOT_AUTHORIZED":
            return 403, map[string]interface{}{
                "error": e.Message,
                "code":  e.Code,
            }
        default:
            return 400, map[string]interface{}{
                "error": e.Message,
                "code":  e.Code,
            }
        }
        
    case ValidationError:
        return 400, map[string]interface{}{
            "error": "validation_error",
            "field": e.Field,
            "message": e.Message,
        }
        
    case BusinessRuleError:
        return 422, map[string]interface{}{
            "error": "business_rule_violation",
            "rule":  e.Rule,
            "message": e.Message,
        }
        
    default:
        // Log unexpected errors
        return 500, map[string]interface{}{
            "error": "internal_server_error",
            "message": "An unexpected error occurred",
        }
    }
}
```

---

## Slide 20: Migration Strategy

### Migrating to Clean Architecture

**Phase 1: Identify Boundaries**
```
Current Monolith          →    Identify Layers
┌──────────────┐               ┌──────────────┐
│              │               │   UI/API     │
│   Mixed      │               ├──────────────┤
│   Concerns   │      →        │   Business   │
│              │               ├──────────────┤
│              │               │   Database   │
└──────────────┘               └──────────────┘
```

**Phase 2: Extract Domain Logic**
```go
// Before: Mixed concerns
type TaskService struct {
    db *sql.DB
}

func (s *TaskService) CreateTask(title string) error {
    // Validation mixed with SQL
    if title == "" {
        return errors.New("title required")
    }
    
    _, err := s.db.Exec(
        "INSERT INTO tasks (title) VALUES (?)",
        title,
    )
    return err
}

// After: Separated concerns
// Entity with business rules
type Task struct {
    Title string
}

func NewTask(title string) (*Task, error) {
    if title == "" {
        return nil, errors.New("title required")
    }
    return &Task{Title: title}, nil
}

// Use case
type CreateTaskUseCase struct {
    repo TaskRepository
}

func (uc *CreateTaskUseCase) Execute(title string) error {
    task, err := NewTask(title)
    if err != nil {
        return err
    }
    return uc.repo.Save(task)
}
```

---

## Slide 21: Migration Strategy - Practical Steps

### Step-by-Step Migration Guide

**1. Start with New Features**
- Implement new features using Clean Architecture
- Don't refactor everything at once

**2. Create Anti-Corruption Layer**
```go
// Adapter for legacy code
type LegacyTaskAdapter struct {
    legacyDB *sql.DB
}

func (a *LegacyTaskAdapter) Save(ctx context.Context, task *entity.Task) error {
    // Translate to legacy format
    return a.legacyDB.Exec(
        "INSERT INTO old_tasks ...",
        task.ID, task.Title,
    )
}

func (a *LegacyTaskAdapter) FindByID(ctx context.Context, id string) (*entity.Task, error) {
    // Query legacy DB and translate to entity
    var legacyTask struct {
        ID    int
        Title string
    }
    
    err := a.legacyDB.QueryRow(
        "SELECT * FROM old_tasks WHERE id = ?",
        id,
    ).Scan(&legacyTask.ID, &legacyTask.Title)
    
    if err != nil {
        return nil, err
    }
    
    // Convert to clean entity
    return &entity.Task{
        ID:    strconv.Itoa(legacyTask.ID),
        Title: legacyTask.Title,
    }, nil
}
```

**3. Gradual Refactoring Timeline**
- Month 1-2: Identify and document boundaries
- Month 3-4: Extract entities and use cases
- Month 5-6: Create repository interfaces
- Month 7-8: Refactor controllers
- Month 9-12: Complete migration and remove legacy code

---

## Slide 22: Best Practices

### Clean Architecture Best Practices

**✅ DO's:**

1. **Keep Entities Pure**
   - No framework dependencies
   - No database annotations
   - Only business logic

2. **Use Dependency Injection**
   - Constructor injection preferred
   - Interfaces in inner layers
   - Implementations in outer layers

3. **Test from Inside Out**
   - Start with entity tests
   - Then use case tests
   - Finally integration tests

4. **Keep Use Cases Focused**
   - One use case = one business action
   - Clear input/output DTOs
   - No UI concerns

5. **Version Your APIs**
   - Maintain backward compatibility
   - Use versioned endpoints
   - Document changes

**Code Example:**
```go
// ✅ Good: Clear separation
type CreateUserUseCase struct {
    userRepo UserRepository
    hasher   PasswordHasher
    emailer  EmailService
}

func (uc *CreateUserUseCase) Execute(input CreateUserInput) (*User, error) {
    // Clear, focused logic
    hashedPassword, err := uc.hasher.Hash(input.Password)
    if err != nil {
        return nil, err
    }
    
    user := &User{
        Email:    input.Email,
        Password: hashedPassword,
    }
    
    if err := uc.userRepo.Save(user); err != nil {
        return nil, err
    }
    
    uc.emailer.SendWelcome(user.Email)
    
    return user, nil
}
```

---

## Slide 23: Anti-Patterns to Avoid

### Common Clean Architecture Anti-Patterns

**❌ DON'Ts:**

1. **Don't Skip Layers**
```go
// ❌ Bad: Controller directly accessing repository
func (c *Controller) GetTask(id string) {
    task := c.db.Query("SELECT * FROM tasks WHERE id = ?", id)
    // Skipping use case layer
}

// ✅ Good: Going through use case
func (c *Controller) GetTask(id string) {
    task, err := c.getTaskUseCase.Execute(id)
    // Proper layering
}
```

2. **Don't Put Business Logic in Controllers**
```go
// ❌ Bad: Business logic in controller
func (c *Controller) CreateTask(req Request) {
    if req.DueDate.Before(time.Now()) {
        // Business rule in wrong layer!
        return BadRequest("Due date must be in future")
    }
}

// ✅ Good: Business logic in entity
func (t *Task) SetDueDate(date time.Time) error {
    if date.Before(time.Now()) {
        return ErrInvalidDueDate
    }
    t.DueDate = date
    return nil
}
```

3. **Don't Create Anemic Entities**
```go
// ❌ Bad: Anemic entity (just data)
type User struct {
    ID    string
    Email string
    Age   int
}

// ✅ Good: Rich entity with behavior
type User struct {
    id    string
    email string
    age   int
}

func (u *User) CanVote() bool {
    return u.age >= 18
}

func (u *User) ChangeEmail(email string) error {
    if !isValidEmail(email) {
        return ErrInvalidEmail
    }
    u.email = email
    return nil
}
```

---

## Slide 24: Performance Considerations

### Performance in Clean Architecture

**Challenge:** Multiple layers can impact performance

**Solutions:**

1. **Caching Strategy**
```go
type CachedTaskRepository struct {
    cache Cache
    repo  TaskRepository
}

func (r *CachedTaskRepository) FindByID(ctx context.Context, id string) (*Task, error) {
    // Check cache first
    if task, found := r.cache.Get(id); found {
        return task.(*Task), nil
    }
    
    // Load from repository
    task, err := r.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    // Store in cache
    r.cache.Set(id, task, 5*time.Minute)
    
    return task, nil
}
```

2. **Query Optimization**
```go
// Use case specific queries
type TaskQueryService struct {
    db *sql.DB
}

func (s *TaskQueryService) GetTaskDashboard(userID string) (*DashboardView, error) {
    // Optimized query for specific view
    query := `
        SELECT 
            COUNT(*) FILTER (WHERE status = 'open') as open_count,
            COUNT(*) FILTER (WHERE status = 'completed') as completed_count,
            COUNT(*) FILTER (WHERE due_date < NOW() AND status != 'completed') as overdue_count
        FROM tasks
        WHERE assignee_id = $1
    `
    // Direct query for read model
}
```

3. **Lazy Loading**
```go
type Task struct {
    id       string
    title    string
    comments func() ([]Comment, error) // Lazy load
}

func (t *Task) GetComments() ([]Comment, error) {
    if t.comments == nil {
        return nil, nil
    }
    return t.comments()
}
```

---

## Slide 25: Monitoring & Observability

### Observability in Clean Architecture

```go
// internal/usecase/middleware/instrumentation.go
package middleware

import (
    "context"
    "time"
    
    "github.com/prometheus/client_golang/prometheus"
    "go.opentelemetry.io/otel/trace"
)

type InstrumentedUseCase struct {
    useCase     UseCase
    tracer      trace.Tracer
    metrics     *Metrics
}

type Metrics struct {
    requestDuration *prometheus.HistogramVec
    requestCount    *prometheus.CounterVec
    errorCount      *prometheus.CounterVec
}

func (i *InstrumentedUseCase) Execute(ctx context.Context, input interface{}) (interface{}, error) {
    // Start span
    ctx, span := i.tracer.Start(ctx, "usecase.execute")
    defer span.End()
    
    // Record metrics
    start := time.Now()
    
    result, err := i.useCase.Execute(ctx, input)
    
    duration := time.Since(start).Seconds()
    
    labels := prometheus.Labels{
        "use_case": i.useCase.Name(),
        "status":   getStatus(err),
    }
    
    i.metrics.requestDuration.With(labels).Observe(duration)
    i.metrics.requestCount.With(labels).Inc()
    
    if err != nil {
        i.metrics.errorCount.With(labels).Inc()
        span.RecordError(err)
    }
    
    return result, err
}

// Logging middleware
type LoggingMiddleware struct {
    logger Logger
    next   UseCase
}

func (m *LoggingMiddleware) Execute(ctx context.Context, input interface{}) (interface{}, error) {
    m.logger.Info("Executing use case",
        "use_case", m.next.Name(),
        "input", input,
    )
    
    result, err := m.next.Execute(ctx, input)
    
    if err != nil {
        m.logger.Error("Use case failed",
            "use_case", m.next.Name(),
            "error", err,
        )
    } else {
        m.logger.Info("Use case completed",
            "use_case", m.next.Name(),
            "result", result,
        )
    }
    
    return result, err
}
```

---

## Slide 26: Real Production Example

### Production-Ready Clean Architecture

```go
// Complete production setup with all concerns

// cmd/api/wire.go - Dependency injection with Wire
//go:build wireinject

package main

import (
    "github.com/google/wire"
    
    "project/internal/adapter/controller/http"
    "project/internal/infrastructure/persistence/postgres"
    "project/internal/infrastructure/cache/redis"
    "project/internal/usecase/task"
)

func InitializeApp() (*Application, error) {
    wire.Build(
        // Infrastructure providers
        postgres.NewDB,
        redis.NewCache,
        
        // Repository providers
        postgres.NewTaskRepository,
        postgres.NewUserRepository,
        
        // Use case providers
        task.NewCreateTaskUseCase,
        task.NewListTasksUseCase,
        task.NewAssignTaskUseCase,
        
        // Controller providers
        http.NewTaskController,
        
        // Application
        NewApplication,
    )
    
    return nil, nil
}

// Application configuration
type Application struct {
    Server     *http.Server
    DB         *sql.DB
    Cache      cache.Cache
    Config     *config.Config
    Shutdowner Shutdowner
}

func (app *Application) Start() error {
    // Start health checks
    go app.startHealthChecks()
    
    // Start metrics server
    go app.startMetricsServer()
    
    // Start main server
    return app.Server.ListenAndServe()
}

func (app *Application) Shutdown(ctx context.Context) error {
    // Graceful shutdown
    if err := app.Server.Shutdown(ctx); err != nil {
        return err
    }
    
    if err := app.DB.Close(); err != nil {
        return err
    }
    
    return app.Cache.Close()
}

// Health checks
func (app *Application) startHealthChecks() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        checks := map[string]string{
            "database": app.checkDatabase(),
            "cache":    app.checkCache(),
        }
        
        status := http.StatusOK
        for _, check := range checks {
            if check != "healthy" {
                status = http.StatusServiceUnavailable
                break
            }
        }
        
        w.WriteHeader(status)
        json.NewEncoder(w).Encode(checks)
    })
}
```

---

## Slide 27: Key Takeaways

### Clean Architecture Key Takeaways

**🎯 Core Benefits:**
1. **Testability** - Test business logic without external dependencies
2. **Flexibility** - Change frameworks without changing business logic
3. **Maintainability** - Clear separation of concerns
4. **Independence** - Business logic doesn't depend on delivery mechanism

**📊 Metrics from Real Projects:**
- **70% reduction** in bug density in business logic
- **50% faster** feature development after initial setup
- **90% code coverage** achievable for business logic
- **40% reduction** in time to onboard new developers

**🚀 When to Use:**
- Long-lived projects (> 1 year)
- Complex business logic
- Multiple delivery mechanisms
- Need for high testability
- Team > 3 developers

**⚠️ When to Avoid:**
- Simple CRUD applications
- Prototypes or MVPs
- Small scripts or utilities
- Team unfamiliar with the pattern

**Remember:**
> "Architecture is about intent" - Uncle Bob

The architecture should scream what the system does, not what frameworks it uses.

---

## Slide 28: Questions & Resources

### Questions & Discussion

**Let's Discuss:**
- What challenges do you see implementing Clean Architecture?
- Which parts of our current system would benefit most?
- How can we start small and iterate?
- What concerns do you have about the complexity?

### Resources

**Books:**
- "Clean Architecture" - Robert C. Martin
- "Get Your Hands Dirty on Clean Architecture" - Tom Hombergs
- "Implementing Clean Architecture" - Manfred Steyer

**Online Resources:**
- [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) - Original blog post
- [Clean Architecture in Go](https://github.com/bxcodec/go-clean-arch) - Example implementation
- [Wild Workouts](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example) - Production example

**Videos:**
- Robert C. Martin - Clean Architecture and Design
- GoLab 2019 - Clean Architecture in Go

**Contact:**
Dong Tran  
dong.tran@company.com  
GitHub: @dong-tran

---

## Slide 29: Thank You!

### Thank You!

**Next Steps:**
1. **Pilot Project** - Identify a bounded context for Clean Architecture
2. **Study Group** - Weekly Clean Architecture discussions
3. **Code Review** - Focus on separation of concerns
4. **Documentation** - Create team guidelines

**Action Items:**
- [ ] Review existing codebase for refactoring opportunities
- [ ] Set up example repository with Clean Architecture
- [ ] Schedule follow-up workshop on testing strategies
- [ ] Create ADR (Architecture Decision Record) for adoption

**Final Thought:**
> "The only way to go fast, is to go well." - Robert C. Martin

Clean Architecture is an investment in the future maintainability of your software.

---

## Appendix A: Complete Example Repository

The complete working example is available at:
```
github.com/dong-tran/clean-architecture-go-example
```

Features:
- Full CRUD operations
- Authentication & Authorization
- Database migrations
- Docker setup
- CI/CD pipeline
- Integration tests
- Performance tests
- Documentation

---

## Appendix B: Architecture Decision Record (ADR)

```markdown
# ADR-001: Adopt Clean Architecture

## Status
Proposed

## Context
Our current monolithic application has become difficult to maintain and test.
Business logic is scattered across controllers and database layers.

## Decision
We will adopt Clean Architecture for new features and gradually refactor existing code.

## Consequences
### Positive
- Better testability
- Clear separation of concerns
- Framework independence
- Better onboarding for new developers

### Negative
- Initial complexity
- Learning curve for team
- More boilerplate code
- Longer initial development time

## Alternatives Considered
1. Hexagonal Architecture
2. Traditional Layered Architecture
3. Continue with current approach

## References
- Clean Architecture by Robert C. Martin
- Team discussion notes from 2025-11-15
```