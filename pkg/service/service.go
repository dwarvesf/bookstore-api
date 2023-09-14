package service

import (
	"github.com/dwarvesf/df-bookstore-srv/pkg/config"
)

// Service for app
type Service struct {
}

// New will return the services in app
func New(cfg *config.Config) Service {

	return Service{}
}
