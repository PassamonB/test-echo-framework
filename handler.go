package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	User struct {
		Name  string `json:"name" form:"name"`
		Email string `json:"email" form:"email"`
	}
	handler struct{}
)

func NewHandler() handler {
	return handler{}
}

func (h *handler) getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, []User{
		{
			Name:  "test user name",
			Email: "email@test.com",
		},
	})
}
