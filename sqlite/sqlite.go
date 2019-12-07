// Package sqlite have some helpers and types for working with sqlite.
package sqlite

import (
	"errors"
)

var (
	ErrSQLTimeUnmarshallType = errors.New("incompatible type for SQLTime")
)
