package response

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail contains error information
type ErrorDetail struct {
	Code      string            `json:"code"`
	Message   string            `json:"message"`
	Details   []ValidationError `json:"details,omitempty"`
	RequestID string            `json:"request_id,omitempty"`
}

// ValidationError represents a validation error for a specific field
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// SuccessResponse represents a success response
type SuccessResponse struct {
	Data interface{} `json:"data,omitempty"`
}

// WriteJSON writes a JSON response
func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// WriteError writes an error response
func WriteError(w http.ResponseWriter, status int, code, message string, requestID string) error {
	return WriteJSON(w, status, ErrorResponse{
		Error: ErrorDetail{
			Code:      code,
			Message:   message,
			RequestID: requestID,
		},
	})
}

// WriteValidationError writes a validation error response
func WriteValidationError(w http.ResponseWriter, errors []ValidationError, requestID string) error {
	return WriteJSON(w, http.StatusBadRequest, ErrorResponse{
		Error: ErrorDetail{
			Code:      "VALIDATION_ERROR",
			Message:   "Invalid request parameters",
			Details:   errors,
			RequestID: requestID,
		},
	})
}

// WriteSuccess writes a success response
func WriteSuccess(w http.ResponseWriter, status int, data interface{}) error {
	return WriteJSON(w, status, SuccessResponse{Data: data})
}
