package config

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config struct.
type Config struct {
	// Postgres options.
	Postgres struct {
		Host     string `envconfig:"DB_HOST" validate:"required"`
		Port     string `envconfig:"DB_PORT" validate:"required"`
		Name     string `envconfig:"DB_NAME" validate:"required"`
		User     string `envconfig:"DB_USER" validate:"required"`
		Password string `envconfig:"DB_PASSWORD" validate:"required"`
		SslMode  string `envconfig:"DB_SSLMODE" validate:"required"`
	}
	// Redis options.
	Redis struct {
		Address  string `envconfig:"REDIS_ADDRESS" validate:"required"`
		Password string `envconfig:"REDIS_PASSWORD"`
	}
	// Server options.
	Server struct {
		HTTPPort          string        `envconfig:"HTTP_PORT" validate:"required"`
		MetricsPort       string        `envconfig:"METRICS_PORT" validate:"required"`
		Mode              string        `envconfig:"SERVER_MODE" validate:"required"`
		ReadTimeOut       time.Duration `envconfig:"READ_TIMEOUT" validate:"required"`
		WriteTimeOut      time.Duration `envconfig:"WRITE_TIMEOUT" validate:"required"`
		CtxDefaultTimeOut time.Duration `envconfig:"CTX_DEFAULT_TIMEOUT" validate:"required"`
	}
	// Jwt options.
	JWT struct {
		TokenSecret string `envconfig:"JWT_SECRET_KEY" validate:"required"`
	}
	// Logger options.
	Logger struct {
		Development       bool   `envconfig:"LOG_DEVELOPMENT" validate:"required"`
		DisableCaller     bool   `envconfig:"LOG_DISABLE_CALLER"`
		DisableStackTrace bool   `envconfig:"LOG_DISABLE_STACK_TRACE"`
		Encoding          string `envconfig:"LOG_ENCODING" validate:"required"`
		Level             string `envconfig:"LOG_LEVEL" validate:"required"`
	}
}

// LoadConfig loads config options.
func LoadConfig() (*Config, error) {
	// Read .env file with this method.
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("config.LoadConfig.Load: %w", err)
	}

	// Get instance of config file.
	var c Config
	// Populate the specified struct based on environment variables.
	if err := envconfig.Process("", &c); err != nil {
		return nil, fmt.Errorf("envconfig.Process: %w", err)
	}

	// Validate the Config.
	if err := validator.New().Struct(c); err != nil {
		return nil, fmt.Errorf("error in validating Config struct: %w", err)
	}

	return &c, nil
}
