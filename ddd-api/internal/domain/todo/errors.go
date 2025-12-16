package todo

import "errors"

var (
	ErrInvalidID    = errors.New("invalid id")
	ErrInvalidTitle = errors.New("invalid title")
	ErrNotFound     = errors.New("not found")
)
