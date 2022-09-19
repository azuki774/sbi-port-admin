package usecase

import "errors"

var (
	ErrNotFound    error = errors.New("not found")
	ErrInvalidDate error = errors.New("invalid date")
)
