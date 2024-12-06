package middleware

import (
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/config"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/logger"
)

type MiddlewareManager struct {
	cfg    *config.Config
	logger logger.Logger
}

// NewMiddlewareManager creates and returns a new instance of MiddlewareManager.
func NewMiddlewareManager(cfg *config.Config, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{cfg: cfg, logger: logger}
}
