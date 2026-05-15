package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppEnv           string
	HTTPPort         string
	DatabaseURL      string
	DBMaxConns       int32
	DBMinConns       int32
	DBMinIdleConns   int32
	DBMaxConnLife    time.Duration
	DBMaxConnIdle    time.Duration
	DBHealthCheck    time.Duration
	WorkerCount      int
	WorkerQueueSize  int
	BatchSize        int
	BatchFlushPeriod time.Duration
}

func Load() (*Config, error) {
	cfg := &Config{
		AppEnv:           getEnv("APP_ENV", "development"),
		HTTPPort:         getEnv("HTTP_PORT", "8080"),
		DatabaseURL:      getEnv("DATABASE_URL", ""),
		DBMaxConns:       int32(getEnvAsInt("DB_MAX_CONNS", 30)),
		DBMinConns:       int32(getEnvAsInt("DB_MIN_CONNS", 10)),
		DBMinIdleConns:   int32(getEnvAsInt("DB_MIN_IDLE_CONNS", 10)),
		DBMaxConnLife:    getEnvAsDuration("DB_MAX_CONN_LIFETIME", time.Hour),
		DBMaxConnIdle:    getEnvAsDuration("DB_MAX_CONN_IDLE_TIME", 15*time.Minute),
		DBHealthCheck:    getEnvAsDuration("DB_HEALTH_CHECK_PERIOD", time.Minute),
		WorkerCount:      getEnvAsInt("WORKER_COUNT", 8),
		WorkerQueueSize:  getEnvAsInt("WORKER_QUEUE_SIZE", 50000),
		BatchSize:        getEnvAsInt("BATCH_SIZE", 500),
		BatchFlushPeriod: getEnvAsDuration("BATCH_FLUSH_PERIOD", 200*time.Millisecond),
	}

	if cfg.HTTPPort == "" {
		return nil, fmt.Errorf("HTTP_PORT is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getEnvAsInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}

func getEnvAsDuration(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}

	return parsed
}