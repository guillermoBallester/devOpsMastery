package connection

import (
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"sync"
	"time"
)

type Info struct {
	RequestID     string
	RemoteAddress string
	Method        string
	Path          string
	StartTime     time.Time
	Done          chan struct{}
}

type Manager struct {
	activeConnections map[string]*Info
	mu                sync.RWMutex
	maxConnections    int
	//Metrics Collector
}

// NewManager creates a new connection manager
func NewManager(maxConnections int) *Manager {
	if maxConnections <= 0 {
		maxConnections = 1000
	}

	return &Manager{
		activeConnections: make(map[string]*Info),
		maxConnections:    maxConnections,
	}
}

// Middleware returns an HTTP middleware that tracks connections
func (m *Manager) Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := middleware.GetReqID(r.Context())

			connInfo := &Info{
				RequestID:     requestID,
				RemoteAddress: r.RemoteAddr,
				Method:        r.Method,
				Path:          r.URL.Path,
				StartTime:     time.Now(),
				Done:          make(chan struct{}),
			}

			m.registerConnection(requestID, connInfo)
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			defer func() {
				m.unregisterConnection(requestID)
				close(connInfo.Done)
			}()

			//TODO Some connection metrics?

			next.ServeHTTP(ww, r)
		})
	}
}

// registerConnection adds a connection to the active connections map
func (m *Manager) registerConnection(requestID string, info *Info) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.activeConnections[requestID] = info
}

// unregisterConnection removes a connection from the active connections map
func (m *Manager) unregisterConnection(requestID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.activeConnections, requestID)
}

// GetActiveConnectionCount returns the number of active connections
func (m *Manager) GetActiveConnectionCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.activeConnections)
}
