package service

import (
	"boilerplate/pkg/repository"
	"github.com/gofiber/fiber/v2"
)

func GetHomeMessage() string {
	// Add business logic and interaction with repositories here
	if 2 == 2 {
		panic(fiber.ErrForbidden)
	}
	return repository.GetMessageFromDB()
}
