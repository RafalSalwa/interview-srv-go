package util

import "net/http"

type Error struct {
	Location     string `json:"location"`
	LocationType string `json:"locationType"`
	Reason       string `json:"reason"`
	Message      string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Code    int32  `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type NotFoundResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message,omitempty"`
}

type UnauthorizedResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewErrorResponse(statusCode int32, reason, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    statusCode,
		Reason:  reason,
		Message: message,
	}
}

func NewNotFoundResponse() *NotFoundResponse {
	return &NotFoundResponse{
		http.StatusNotFound,
		"Not found",
	}
}
func NewUnauthorizedResponse() *UnauthorizedResponse {
	return &UnauthorizedResponse{
		http.StatusUnauthorized,
		"Not authorized",
		"",
	}
}

func NewBadRequestErrorResponse(msg string) *ErrorResponse {
	if msg == "" {
		msg = "Your request is invalid"
	}

	return NewErrorResponse(
		http.StatusBadRequest,
		"Bad request",
		msg,
	)
}

func NewUnauthorizedErrorResponse(msg string) *ErrorResponse {
	return NewErrorResponse(
		http.StatusUnauthorized,
		"Unauthorized",
		msg,
	)
}

func NewAccessDeniedErrorResponse() *ErrorResponse {
	return NewErrorResponse(
		http.StatusForbidden,
		"Forbidden",
		"",
	)
}

func NewInternalServerErrorErrorResponse() *ErrorResponse {
	return NewErrorResponse(
		http.StatusInternalServerError,
		"Internal server error",
		"",
	)
}

func NewServiceUnavailableErrorResponse() *ErrorResponse {
	return NewErrorResponse(
		http.StatusServiceUnavailable,
		"Service unavailable",
		"",
	)
}
