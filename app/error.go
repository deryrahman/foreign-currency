package app

import "errors"

var (
	// ErrNotFound is used if any resource is not found
	ErrNotFound = errors.New("not found")
	// ErrExist is used when added existing resource
	ErrExist = errors.New("resource exist")
)
