package main

import (
	"fmt"
	"log"

	"github.com/Wallace-Pereira1/ecostream-telemetry-engine/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("EcoStream: High-Throughput Telemetry Engine\n")
	fmt.Printf("environment: %s\n", cfg.AppEnv)
	fmt.Printf("http port: %s\n", cfg.HTTPPort)
}