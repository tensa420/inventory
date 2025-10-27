package entity

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrInternal       = errors.New("internal error")
	ErrDataBaseFailed = errors.New("database error")
)
