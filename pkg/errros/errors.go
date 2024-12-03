package errros

import "fmt"

// API Error codes
const (
	UnknownErr            = "unknown"
	InternalServerErr     = "internal_server_error"
	BadRequestErr         = "bad_request"
	UnauthorizedErr       = "unauthorized"
	ForbiddenErr          = "forbidden"
	NotFoundErr           = "not_found"
	InvalidJSONErr        = "invalid_json"
	ValidationErr         = "validation_error"
	ConflictErr           = "conflict"
	TooManyRequestsErr    = "too_many_requests"
	ServiceUnavailableErr = "service_unavailable"
	ConnectionErr         = "connection_error"
	TimeoutErr            = "timeout"
)

type APIError struct {
	Error     string                 `json:"error"`
	ErrorCode string                 `json:"error_code"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

// InternalError constructs a generic internal error.
func InternalError() *APIError {
	return Errorf("internal error").WithCode(InternalServerErr)
}

// Errorf creates an APIError w/ the formateed message.
func Errorf(msg string, vars ...interface{}) *APIError {
	return &APIError{Error: fmt.Sprintf(msg, vars...)}
}

// Error wraps the error into an API error.
func Error(err error) *APIError {
	if err == nil {
		return nil
	}

	return &APIError{Error: err.Error()}
}

// WithCode adds an error code to an APIError
func (e *APIError) WithCode(code string) *APIError {
	e.ErrorCode = code
	return e
}

// ValidationAPIError defines validation error api layer
type ValidationAPIError struct {
	APIError
	Fields map[string]string `json:"fields,omitempty"`
}
