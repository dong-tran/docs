package main

import (
"github.com/labstack/echo/v4"
"github.com/labstack/echo/v4/middleware"
"io"
"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route to User Service
	e.Any("/api/users/*", func(c echo.Context) error {
return proxy(c, "http://localhost:8081")
})

	// Route to Product Service
	e.Any("/api/products/*", func(c echo.Context) error {
return proxy(c, "http://localhost:8082")
})

	// Route to Order Service
	e.Any("/api/orders/*", func(c echo.Context) error {
return proxy(c, "http://localhost:8083")
})

	e.Start(":8080")
}

func proxy(c echo.Context, target string) error {
	req := c.Request()
	resp, err := http.DefaultClient.Do(&http.Request{
		Method: req.Method,
		URL:    req.URL,
		Header: req.Header,
		Body:   req.Body,
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return c.Blob(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
