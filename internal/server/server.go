package server

import (
	"context"

	_ "github.com/jumayevgadam/book-app-with-refreshtoken/docs"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/config"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/database"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/errlst"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/logger"
	"github.com/labstack/echo/v4"
)

// Server struct.
type Server struct {
	Echo      *echo.Echo
	Cfg       *config.Config
	DataStore database.DataStore
	Logger    logger.Logger
}

// NewServer func.
func NewServer(cfg *config.Config, dataStore database.DataStore, logger logger.Logger) *Server {
	return &Server{
		Echo:      echo.New(),
		Cfg:       cfg,
		DataStore: dataStore,
		Logger:    logger,
	}
}

func (s *Server) Run() error {
	err := s.MapHandlers()
	if err != nil {
		return errlst.ParseErrors(err)
	}

	// Configure Echo's HTTP server
	s.Echo.Server.ReadTimeout = s.Cfg.Server.ReadTimeOut
	s.Echo.Server.WriteTimeout = s.Cfg.Server.WriteTimeOut

	err = s.Echo.Start(":" + s.Cfg.Server.HTTPPort)
	if err != nil {
		return errlst.ParseErrors(err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.Echo.Shutdown(ctx)
	if err != nil {
		return errlst.ParseErrors(err)
	}

	return nil
}
