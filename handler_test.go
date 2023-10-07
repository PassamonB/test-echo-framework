package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	// Setup
	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	context := e.NewContext(request, response)

	// Action
	h := &handler{}
	h.getUsers(context)

	// Assertions
	expectedResult := `[{"name":"test user name","email":"email@test.com"}]`
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, expectedResult, strings.TrimSuffix(response.Body.String(), "\n"))
}
