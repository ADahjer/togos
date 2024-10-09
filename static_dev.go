//+build dev
//go:build dev
// +build dev

package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func public(e *echo.Echo) {
	log.Println("Building static files for development...")
	e.GET("/public/*", func(c echo.Context) error {
		fs := http.FileServer(http.Dir("public"))
		http.StripPrefix("/public", fs).ServeHTTP(c.Response(), c.Request())
		return nil
	})
}
