//+build prod
//go:build prod
// +build prod

package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed public
var publicFS embed.FS

func public(e *echo.Echo) {
	log.Println("serving static files")
	e.GET("/public/*", echo.WrapHandler(http.FileServer(http.FS(publicFS))))
}
