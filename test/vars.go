package test

import (
	"errors"
	"time"
)

var (
	// Now represents the current time.
	Now = time.Now()
	// FixedTime represents a fixed time "2019-03-03T13:51:06.7369193+00:00".
	FixedTime time.Time
	// Error is a default error for mocks.
	Error = errors.New("unittest")
)

func init() {
	FixedTime, _ = time.Parse(time.RFC3339Nano, "2019-03-03T13:51:06.7369193+00:00")
}
