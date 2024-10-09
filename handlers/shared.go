package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	registerRoutes(e *echo.Group)
	basePath() string
}

func RegisterHandlers() []Handler {
	return []Handler{
		NewFooHandler(),
	}
}

func SetUpRoutes(e *echo.Echo, h []Handler) {
	for _, handler := range h {
		group := e.Group(handler.basePath())
		handler.registerRoutes(group)
	}
}

func ErrorHandlerFunc(err error, c echo.Context) {
	switch err.(type) {
	default:
		c.JSON(http.StatusInternalServerError, echo.Map{"message": "error not handled yet", "error": err.Error()})
	}
}

// Handle templates
func Render(e echo.Context, c templ.Component) error {
	return c.Render(e.Request().Context(), e.Response().Writer)
}
