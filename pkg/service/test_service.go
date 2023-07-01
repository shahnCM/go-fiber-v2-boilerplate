package service

import (
	"boilerplate/pkg/repository"
)

func GetHomeMessage() string {
	// Add business logic and interaction with repositories here
	return repository.GetMessageFromDB()
}
