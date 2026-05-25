package domain

import "errors"

var (
	ErrInvalidApplication = errors.New("invalid application")
	ErrNotFound           = errors.New("application not found")
)
