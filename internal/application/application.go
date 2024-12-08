package application

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/config"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/database/postgres"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/initializers"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/server"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/errlist"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/logger"
)

func BootStrap(ctx context.Context) error {
	cfgs, err := config.LoadConfig()
	if err != nil {
		return errlist.ParseErrors(err)
	}

	appLogger := logger.NewApiLogger(cfgs)
	appLogger.InitLogger()
	appLogger.Infof("Mode: %s", cfgs.Server.Mode)

	psqlDB, err := initializers.GetDBConnection(ctx, *cfgs)
	if err != nil {
		return errlist.ParseErrors(err)
	}
	appLogger.Infof("successfully connected\n")

	defer func() {
		if err := psqlDB.Close(); err != nil {
			appLogger.Warnf("error in closing DB: %v", err)
		}
	}()

	dataStore := postgres.NewDataStore(psqlDB)

	srv := server.NewServer(cfgs, dataStore, appLogger)

	// Start server in a goroutine
	serverErrors := make(chan error, 1)
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()
	appLogger.Info("Server Started\n")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-shutdown:
		appLogger.Infof("Caught signal: %v, initiating graceful shutdown...\n", sig)
	case err := <-serverErrors:
		appLogger.Errorf("Server error: %v\n", err.Error())
		return errlist.ParseErrors(err)
	}

	// Gracefully stop server.
	ctx, cancel := context.WithTimeout(ctx, cfgs.Server.CtxDefaultTimeOut)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		errlist.ParseErrors(err)
	}

	appLogger.Info("Server stopped gracefully\n")
	return nil
}
