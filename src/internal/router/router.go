package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/guillermoBallester/devOpsMastery/src/internal/handler"
	"github.com/guillermoBallester/devOpsMastery/src/internal/service"
	"time"
)

type Router struct {
	router *chi.Mux
}

func NewRouter() *Router {
	r := chi.NewRouter()

	setupMiddleware(r)

	healthRoutes := NewHealthRoutes(handler.NewHealthHandler())
	apiRoutes := NewAPIRoutes(handler.NewHelloHandler(service.NewHelloService()))

	healthRoutes.Register(r)
	apiRoutes.Register(r)

	return &Router{
		router: r,
	}
}

// Handler returns the router as an HTTP handler
func (r *Router) Handler() chi.Router {
	return r.router
}

func setupMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
}
