package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	initLogger()
	e := echo.New()
	e.GET("/", echoHandler, addLoggerToContext)

	log.Fatal().Err(e.Start(":8080")).Send()
}

func initLogger() {
	log.Logger = zerolog.New(os.Stdout).With().Str("app", "test_echo").Logger()
}

func addLoggerToContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		l := log.Logger.With().Str("middleware", "add logger to context").Logger()
		c.SetRequest(c.Request().WithContext(l.WithContext(c.Request().Context())))
		return next(c)
	}
}

func echoHandler(c echo.Context) error {
	logger := zerolog.Ctx(c.Request().Context())
	logger.Info().Str("endpoint", c.Path()).Msg("c in handler")
	return nil
}
