package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Wallace-Pereira1/ecostream-telemetry-engine/internal/domain"
	"github.com/Wallace-Pereira1/ecostream-telemetry-engine/internal/service"
)

type TelemetryHandler struct {
	service service.TelemetryService
}

func NewTelemetryHandler(service service.TelemetryService) *TelemetryHandler {
	return &TelemetryHandler{
		service: service,
	}
}

func (h *TelemetryHandler) Ingest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return
	}

	defer r.Body.Close()

	var event domain.TelemetryEvent

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&event); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	if err := h.service.Process(r.Context(), event); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusAccepted, map[string]any{
		"message": "telemetry received",
		"event":   event,
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}