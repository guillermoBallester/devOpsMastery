package connection

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestConnectionManager(t *testing.T) {
	manager := NewManager(10)

	if count := manager.GetActiveConnectionCount(); count != 0 {
		t.Errorf("Expected 0 active connections, got %d", count)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	})

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(manager.Middleware())
	r.Get("/test", testHandler)

	server := httptest.NewServer(r)
	defer server.Close()

	go func() {
		_, err := http.Get(server.URL + "/test")
		if err != nil {
			t.Errorf("Error making test request: %v", err)
		}
	}()

	time.Sleep(50 * time.Millisecond)

	if count := manager.GetActiveConnectionCount(); count != 1 {
		t.Errorf("Expected 1 active connection, got %d", count)
	}

	time.Sleep(200 * time.Millisecond)

	if count := manager.GetActiveConnectionCount(); count != 0 {
		t.Errorf("Expected 0 active connections after completion, got %d", count)
	}
}
