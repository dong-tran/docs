package handler

import (
"net/http"

"github.com/dong-tran/docs/integration-example/usecase"
"github.com/labstack/echo/v4"
)

// OrderHandler - Presentation layer (Clean Architecture)
type OrderHandler struct {
	orderUseCase *usecase.OrderUseCase
}

func NewOrderHandler(orderUseCase *usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{orderUseCase: orderUseCase}
}

type CreateOrderRequest struct {
	CustomerID string               `json:"customer_id"`
	Items      []OrderItemRequest   `json:"items"`
}

type OrderItemRequest struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
}

type ProcessPaymentRequest struct {
	PaymentMethod string `json:"payment_method"`
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	var req CreateOrderRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Convert to DTO
	dto := usecase.CreateOrderDTO{
		CustomerID: req.CustomerID,
		Items:      make([]usecase.OrderItemDTO, len(req.Items)),
	}

	for i, item := range req.Items {
		dto.Items[i] = usecase.OrderItemDTO{
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			Price:       item.Price,
			Currency:    item.Currency,
		}
	}

	order, err := h.orderUseCase.CreateOrder(dto)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
"id":          order.ID().String(),
		"customer_id": order.CustomerID().String(),
		"total":       order.TotalAmount().Amount(),
		"currency":    order.TotalAmount().Currency(),
		"status":      order.Status(),
	})
}

func (h *OrderHandler) ProcessPayment(c echo.Context) error {
	orderID := c.Param("id")
	
	var req ProcessPaymentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := h.orderUseCase.ProcessPayment(orderID, req.PaymentMethod); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "payment processed"})
}

func (h *OrderHandler) GetOrder(c echo.Context) error {
	orderID := c.Param("id")
	
	order, err := h.orderUseCase.GetOrder(orderID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "order not found"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
"id":          order.ID().String(),
		"customer_id": order.CustomerID().String(),
		"total":       order.TotalAmount().Amount(),
		"currency":    order.TotalAmount().Currency(),
		"status":      order.Status(),
	})
}
