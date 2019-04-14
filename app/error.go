package app

import "errors"

var (
	// ErrNotFound is used if any resource is not found
	ErrNotFound = errors.New("not found")
	// ErrExist is used when added existing resource
	ErrExist = errors.New("resource exist")
)

// ErrorResponse is a struct for any error response
type ErrorResponse struct {
	ErrMsg string `json:"err_msg"`
}
