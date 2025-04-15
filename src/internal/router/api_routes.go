package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/guillermoBallester/devOpsMastery/src/internal/handler"
)

// APIRoutes defines the API routes
type APIRoutes struct {
	helloHandler *handler.HelloHandler
}

// NewAPIRoutes creates a new API routes registrar
func NewAPIRoutes(helloHandler *handler.HelloHandler) *APIRoutes {
	return &APIRoutes{
		helloHandler: helloHandler,
	}
}

// Register registers all API routes on the provided router
func (r *APIRoutes) Register(router chi.Router) {
	router.Route("/api/v1", func(router chi.Router) {
		// Hello world endpoint
		router.Get("/hello", r.helloHandler.SayHello)

		//TODO  Add more
	})
}
