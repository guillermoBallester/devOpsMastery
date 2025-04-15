package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/guillermoBallester/devOpsMastery/src/internal/handler"
)

// HealthRoutes defines the health check routes
type HealthRoutes struct {
	handler *handler.HealthHandler
}

// NewHealthRoutes creates a new health routes registrar
func NewHealthRoutes(handler *handler.HealthHandler) *HealthRoutes {
	return &HealthRoutes{
		handler: handler,
	}
}

// Register registers all health routes on the provided router
func (r *HealthRoutes) Register(router chi.Router) {
	router.Get("/health", r.handler.Check)
}
