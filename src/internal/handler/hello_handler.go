package handler

import (
	"github.com/guillermoBallester/devOpsMastery/src/internal/response"
	"github.com/guillermoBallester/devOpsMastery/src/internal/service"
	"net/http"
	"time"
)

type HelloHandler struct {
	helloService *service.HelloService
}

func NewHelloHandler(helloService *service.HelloService) *HelloHandler {
	return &HelloHandler{
		helloService: helloService,
	}
}

func (h *HelloHandler) SayHello(w http.ResponseWriter, r *http.Request) {
	message, err := h.helloService.GetHello(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
	}

	response.Success(w, map[string]string{
		"message":   message,
		"status":    "success",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}
