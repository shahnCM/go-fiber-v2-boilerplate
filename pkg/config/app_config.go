package config

import (
	"boilerplate/pkg/errorhandler"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	preformEnabled, _ := strconv.ParseBool(os.Getenv("SERVER_READ_TIMEOUT"))
	appName := os.Getenv("APP_NAME")

	// Return Fiber configuration.
	return fiber.Config{
		AppName:       appName,
		Prefork:       preformEnabled,
		CaseSensitive: true,
		StrictRouting: true,
		ReadTimeout:   time.Second * time.Duration(readTimeoutSecondsCount),
		ErrorHandler:  errorhandler.CustomFiberErrorHandler,
	}
}

func RecoveryConfig() recover.Config {
	return recover.Config{
		EnableStackTrace: true,
	}
}
