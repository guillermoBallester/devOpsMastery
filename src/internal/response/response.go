package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON renders a JSON response with given status code and payload
func JSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	// Set content type header
	w.Header().Set("Content-Type", "application/json")

	// Set status code
	w.WriteHeader(statusCode)

	// Marshal and write response
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		// Log the error but don't expose it to the client
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, `{"error":"Internal server error"}`, http.StatusInternalServerError)
	}
}

// Success is a convenience function for successful responses
func Success(w http.ResponseWriter, payload interface{}) {
	JSON(w, http.StatusOK, payload)
}

// Error renders an error response with appropriate status code
func Error(w http.ResponseWriter, statusCode int, message string) {
	errorResponse := map[string]string{
		"error": message,
	}
	JSON(w, statusCode, errorResponse)
}
