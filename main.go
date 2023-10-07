package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("initializing echo server")
	e := echo.New()

	fmt.Println("initializing backend handler")
	backendHandler := NewHandler()

	fmt.Println("register GET /api/v1/users endpoint")
	e.GET("/api/v1/users", backendHandler.getUsers)

	fmt.Println("starting server")
	e.Logger.Fatal(e.Start(":1323"))
}
