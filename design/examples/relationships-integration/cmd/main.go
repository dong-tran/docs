package main

import (
"log"

"github.com/dong-tran/docs/integration-example/handler"
"github.com/dong-tran/docs/integration-example/infrastructure"
"github.com/dong-tran/docs/integration-example/repository"
"github.com/dong-tran/docs/integration-example/shared/patterns"
"github.com/dong-tran/docs/integration-example/usecase"
"github.com/labstack/echo/v4"
"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize infrastructure
	db, err := infrastructure.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Setup event system (Observer pattern)
	eventPublisher := patterns.NewEventPublisher()
	eventPublisher.Subscribe(&infrastructure.EmailNotificationHandler{})
	eventPublisher.Subscribe(&infrastructure.LoggingHandler{})
	eventPublisher.Subscribe(&infrastructure.AnalyticsHandler{})

	// Setup factories (Factory pattern)
	paymentFactory := patterns.NewPaymentFactory()

	// Dependency injection (DIP)
	orderRepo := repository.NewOrderRepository(db)
	orderUseCase := usecase.NewOrderUseCase(orderRepo, paymentFactory, eventPublisher)
	orderHandler := handler.NewOrderHandler(orderUseCase)

	// Setup Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.POST("/orders", orderHandler.CreateOrder)
	e.GET("/orders/:id", orderHandler.GetOrder)
	e.POST("/orders/:id/payment", orderHandler.ProcessPayment)

	log.Println("🚀 Integration Example Server starting on :8080")
	log.Println("📚 Demonstrates: Clean Architecture + DDD + SOLID + Design Patterns + Microservices concepts")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
