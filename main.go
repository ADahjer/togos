package main

import (
	"context"
	"log"

	"github.com/ADahjer/togos/config"
	"github.com/ADahjer/togos/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

func NewEchoServer(lc fx.Lifecycle, cfg *config.Config) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	// middlewares
	e.Use(middleware.LoggerWithConfig(*cfg.LOGGER_CONFIG))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.HTTPErrorHandler = handlers.ErrorHandlerFunc

	public(e)
	// health check
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Healthy")
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go e.Start(":" + cfg.PORT)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down server...")
			return e.Shutdown(ctx)
		},
	})

	return e

}

func main() {
	fx.New(
		fx.Provide(
			config.SetupConfig,
			handlers.RegisterHandlers,
			NewEchoServer,
		),
		fx.Invoke(
			handlers.SetUpRoutes,
		),
	).Run()
}
