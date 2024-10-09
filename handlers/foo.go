package handlers

import (
	"github.com/ADahjer/togos/views/foo"
	"github.com/labstack/echo/v4"
)

type FooHandler struct{}

func NewFooHandler() *FooHandler {
	return &FooHandler{}
}

func (h *FooHandler) registerRoutes(e *echo.Group) {
	e.GET("", Bar)
}

func (h *FooHandler) basePath() string {
	return "/foo"
}

func Bar(c echo.Context) error {
	return Render(c, foo.Index())
}
