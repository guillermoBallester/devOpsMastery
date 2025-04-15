package handler

import (
	"github.com/guillermoBallester/devOpsMastery/src/internal/response"
	"net/http"
)

// HealthHandler handles health check requests
type HealthHandler struct{}

// NewHealthHandler creates a new health handler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Check handles the health check endpoint
func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	response.Success(w, map[string]string{
		"status": "ok",
	})
}
