package repository

import "errors"

var (
	ErrNotFound       = errors.New("error: not found")
	ErrAliasWasExists = errors.New("this alias is empty")
)
