package domain

import "time"

type TelemetryEvent struct {
	SensorID  string         `json:"sensor_id"`
	Timestamp time.Time      `json:"timestamp"`
	Value     float64        `json:"value"`
	Metadata  map[string]any `json:"metadata"`
}

type SensorRange struct {
	Min float64
	Max float64
}

func (r SensorRange) IsOutOfRange(value float64) bool {
	return value < r.Min || value > r.Max
}