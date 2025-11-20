package main

import (
"github.com/labstack/echo/v4"
"net/http"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	e := echo.New()

	e.GET("/products/:id", func(c echo.Context) error {
product := Product{
ID:    c.Param("id"),
Name:  "Laptop",
Price: 999.99,
}
return c.JSON(http.StatusOK, product)
})

	e.GET("/products", func(c echo.Context) error {
products := []Product{
{ID: "1", Name: "Laptop", Price: 999.99},
{ID: "2", Name: "Mouse", Price: 29.99},
}
return c.JSON(http.StatusOK, products)
})

	e.Start(":8082")
}
