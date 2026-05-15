package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Wallace-Pereira1/ecostream-telemetry-engine/internal/config"
	"github.com/Wallace-Pereira1/ecostream-telemetry-engine/internal/handler"
	"github.com/Wallace-Pereira1/ecostream-telemetry-engine/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	telemetryService := service.NewTelemetryService()
	telemetryHandler := handler.NewTelemetryHandler(telemetryService)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/telemetry", telemetryHandler.Ingest)

	addr := ":" + cfg.HTTPPort

	fmt.Println("EcoStream: High-Throughput Telemetry Engine")
	fmt.Printf("environment: %s\n", cfg.AppEnv)
	fmt.Printf("http port: %s\n", cfg.HTTPPort)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}