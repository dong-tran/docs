package main

import (
"github.com/labstack/echo/v4"
"net/http"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()

	e.GET("/users/:id", func(c echo.Context) error {
user := User{
ID:    c.Param("id"),
Name:  "John Doe",
Email: "john@example.com",
}
return c.JSON(http.StatusOK, user)
})

	e.POST("/users", func(c echo.Context) error {
var user User
if err := c.Bind(&user); err != nil {
			return err
		}
		user.ID = "123"
		return c.JSON(http.StatusCreated, user)
	})

	e.Start(":8081")
}
