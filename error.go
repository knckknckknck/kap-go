package kap

import (
	"errors"
	"fmt"
)

// Sentinel errors for KAP API error codes (ER001–ER008).
var (
	ErrNoPermission     = errors.New("kap: no service access permission")
	ErrUnauthorized     = errors.New("kap: unauthorized request")
	ErrIPRestricted     = errors.New("kap: unregistered IP address")
	ErrInvalidToken     = errors.New("kap: invalid token")
	ErrIPVerification   = errors.New("kap: IP verification failed")
	ErrTokenExpired     = errors.New("kap: token has expired")
	ErrTokenValidation  = errors.New("kap: token could not be validated")
	ErrNotFound         = errors.New("kap: authorization token is not valid")
	ErrUnexpectedStatus = errors.New("kap: unexpected HTTP status")
)

// errorCodeSentinel maps KAP error codes to sentinel errors.
var errorCodeSentinel = map[string]error{
	"ER001": ErrNoPermission,
	"ER002": ErrUnauthorized,
	"ER003": ErrIPRestricted,
	"ER004": ErrInvalidToken,
	"ER005": ErrIPVerification,
	"ER006": ErrTokenExpired,
	"ER007": ErrTokenValidation,
	"ER008": ErrNotFound,
}

// APIError represents an error returned by the KAP API.
type APIError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("kap: api error %s (HTTP %d): %s", e.Code, e.HTTPStatus, e.Message)
}

// Unwrap returns the sentinel error matching the API error code, or
// ErrUnexpectedStatus if the code is not recognized.
func (e *APIError) Unwrap() error {
	if sentinel, ok := errorCodeSentinel[e.Code]; ok {
		return sentinel
	}
	return ErrUnexpectedStatus
}

// RequestError represents a transport-level failure (e.g. network timeout,
// DNS resolution, or response decoding error).
type RequestError struct {
	Method string
	Path   string
	Err    error
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("kap: %s %s: %v", e.Method, e.Path, e.Err)
}

func (e *RequestError) Unwrap() error {
	return e.Err
}
