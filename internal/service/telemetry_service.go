package service

import (
	"context"
	"fmt"

	"github.com/Wallace-Pereira1/ecostream-telemetry-engine/internal/domain"
)

type TelemetryService interface {
	Process(ctx context.Context, event domain.TelemetryEvent) error
}

type telemetryService struct {
	ranges map[string]domain.SensorRange
}

func NewTelemetryService() TelemetryService {
	return &telemetryService{
		ranges: map[string]domain.SensorRange{
			"temperature": {Min: -20, Max: 120},
			"pressure":    {Min: 10, Max: 300},
			"voltage":     {Min: 100, Max: 240},
		},
	}
}

func (s *telemetryService) Process(ctx context.Context, event domain.TelemetryEvent) error {
	_ = ctx

	if event.SensorID == "" {
		return fmt.Errorf("sensor_id is required")
	}

	if event.Timestamp.IsZero() {
		return fmt.Errorf("timestamp is required")
	}

	if sensorRange, ok := s.ranges[event.SensorID]; ok {
		if sensorRange.IsOutOfRange(event.Value) {
			return fmt.Errorf("sensor value out of allowed range for sensor %s", event.SensorID)
		}
	}

	return nil
}