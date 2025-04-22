package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/guillermoBallester/devOpsMastery/src/internal/connection"
	"github.com/guillermoBallester/devOpsMastery/src/internal/handler"
	"github.com/guillermoBallester/devOpsMastery/src/internal/service"
	"github.com/guillermoBallester/devOpsMastery/src/internal/telemetry"
	"github.com/riandyrn/otelchi"
	"time"
)

type Router struct {
	router  *chi.Mux
	connMgr *connection.Manager
}

func NewRouter(tl *telemetry.Telemetry) *Router {
	r := chi.NewRouter()

	connMgr := connection.NewManager(1000)
	setupMiddleware(r, connMgr)

	healthRoutes := NewHealthRoutes(handler.NewHealthHandler())
	apiRoutes := NewAPIRoutes(handler.NewHelloHandler(service.NewHelloService(tl)))

	healthRoutes.Register(r)
	apiRoutes.Register(r)

	return &Router{
		router:  r,
		connMgr: connMgr,
	}
}

// Handler returns the router as an HTTP handler
func (r *Router) Handler() chi.Router {
	return r.router
}

func setupMiddleware(r *chi.Mux, mgr *connection.Manager) {
	r.Use(otelchi.Middleware("devops-mastery", otelchi.WithChiRoutes(r)))
	r.Use(mgr.Middleware())
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

func (r *Router) GetConnectionManager() *connection.Manager {
	return r.connMgr
}
