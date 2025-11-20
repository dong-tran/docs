package main

import (
"log"

"github.com/dong-tran/docs/clean-architecture-example/handler"
"github.com/dong-tran/docs/clean-architecture-example/infrastructure"
"github.com/dong-tran/docs/clean-architecture-example/repository"
"github.com/dong-tran/docs/clean-architecture-example/usecase"
"github.com/labstack/echo/v4"
"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize database (outermost layer)
	db, err := infrastructure.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Dependency injection from outer to inner layers
	taskRepo := repository.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskUseCase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUseCase)

	// Setup Echo framework
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.POST("/tasks", taskHandler.CreateTask)
	e.GET("/tasks/:id", taskHandler.GetTask)
	e.GET("/tasks", taskHandler.GetAllTasks)
	e.PUT("/tasks/:id", taskHandler.UpdateTask)
	e.DELETE("/tasks/:id", taskHandler.DeleteTask)

	// Start server
	log.Println("Server starting on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
