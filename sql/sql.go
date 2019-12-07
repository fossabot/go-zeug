// Package sql have some helpers and types for sql.
package sql

// Scanner Interface will help with some more generic parse functions.
type Scanner interface {
	Scan(args ...interface{}) error
}
