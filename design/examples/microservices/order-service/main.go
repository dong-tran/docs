package main

import (
"github.com/labstack/echo/v4"
"net/http"
)

type Order struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	ProductID string  `json:"product_id"`
	Total     float64 `json:"total"`
}

func main() {
	e := echo.New()

	e.POST("/orders", func(c echo.Context) error {
var order Order
if err := c.Bind(&order); err != nil {
			return err
		}
		order.ID = "order-123"
		return c.JSON(http.StatusCreated, order)
	})

	e.GET("/orders/:id", func(c echo.Context) error {
order := Order{
ID:        c.Param("id"),
UserID:    "user-1",
ProductID: "product-1",
Total:     999.99,
}
return c.JSON(http.StatusOK, order)
})

	e.Start(":8083")
}
