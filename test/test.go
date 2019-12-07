// Package test provides some helpers and vars usefully for testing.
package test

import (
	"bytes"
	"errors"
	"testing"
)

// IsError checks if 'err' is the 'expected' one and calls t.Fatalf if not.
func IsError(t *testing.T, expected, err error) {
	if !errors.Is(err, expected) {
		t.Fatalf(`Expected error "%s" but got "%s"`, expected, err)
	}
}

// NoError checks if 'err' is 'nil' and calls t.Fatalf if not.
func NoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf(`Did not expected to receive error "%s"`, err)
	}
}

// NotNil checks if 'val' is 'nil' and calls t.Fatalf(msg, val) if so.
func NotNil(t *testing.T, val interface{}, msg string) {
	if val == nil {
		t.Fatalf(msg, val)
	}
}

// True checks if val is true and calls t.Fatalf if not.
func True(t *testing.T, val bool) {
	if !val {
		t.Fatalf(`Expected 'true'`)
	}
}

// False checks if val is false and calls t.Fatalf if not.
func False(t *testing.T, val bool) {
	if val {
		t.Fatalf(`Expected 'false'`)
	}
}

// Equals checks if expected != val and calls t.Fatalf(msg, expected, val) if so.
func Equals(t *testing.T, expected, val interface{}, msg string, args ...interface{}) {
	if expected != val {
		tmp := make([]interface{}, 0, 2+len(args))
		tmp = append(tmp, expected, val)
		tmp = append(tmp, args...)
		t.Fatalf(msg, tmp...)
	}
}

// EqualBytes checks if the bytes are unequal and calls t.Fatalf(msg, expected, val) if so.
func EqualBytes(t *testing.T, expected, val []byte, msg string) {
	if !bytes.Equal(expected, val) {
		t.Fatalf(msg, expected, val)
	}
}
