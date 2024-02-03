package errs

import "net/http"

type APIError struct {
	Msg  string `json:"message"`
	Code int    `json:"code"`
}

func (e *APIError) Error() string {
	return e.Msg
}

func NewAPIError(msg string, code int) *APIError {
	return &APIError{
		Msg:  msg,
		Code: code,
	}
}

func BadRequest(msg string) *APIError {
	return NewAPIError(msg, http.StatusBadRequest)
}

func TooManyRequests(msg string) *APIError {
	return NewAPIError(msg, http.StatusTooManyRequests)
}

func Unauthorized(msg string) *APIError {
	return NewAPIError(msg, http.StatusUnauthorized)
}

func NotFound(msg string) *APIError {
	return NewAPIError(msg, http.StatusNotFound)
}

func Forbidden(msg string) *APIError {
	return NewAPIError(msg, http.StatusForbidden)
}
