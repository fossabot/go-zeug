// Package env provides a type for the current runtime environment (not env vars!).
//
// The idea is to initialize the Env very early in your code to make decisions based on
// the current runtime env. Because Env is implementing the Unmarshaler interface it's
// compatible with https://github.com/vrischmann/envconfig and that's the idea how to use it.
//
// Example
//
// With envconfig:
//  var conf struct {
//      Env *env.Env `envconfig:"default=dev"`
//      // ...
//      Host string `envconfig:"default=localhost"`
//      // ...
//  }
package env

import (
	"fmt"

	"github.com/demaggus83/go-zeug/pkg/slice"
)

const (
	// Dev is used for development.
	Dev = "dev"
	// Test is used for testing.
	Test = "test"
	// Stage is used for staging.
	Stage = "stage"
	// Prod is used for production.
	Prod = "production"
)

var validEnvs = []string{Dev, Test, Stage, Prod}

// IsValidEnv returns true if env is valid.
func IsValidEnv(env string) bool {
	return slice.HasString(env, validEnvs)
}

// Env is a helper type for checking the current runtime env.
type Env string

// Unmarshal implements the Unmarshaler interface.
func (e *Env) Unmarshal(s string) error {
	if !IsValidEnv(s) {
		return fmt.Errorf("given env '%s' is invalid. Allowed %s", s, validEnvs)
	}
	*e = Env(s)
	return nil
}

// IsDev returns true if the env is dev.
func (e *Env) IsDev() bool {
	return *e == Dev
}

// IsTest returns true if the env is test.
func (e *Env) IsTest() bool {
	return *e == Test
}

// IsStage returns true if the env is stage.
func (e *Env) IsStage() bool {
	return *e == Stage
}

// IsProd returns true if the env is production.
func (e *Env) IsProd() bool {
	return *e == Prod
}

// Is returns true if the env is one of envs.
func (e *Env) Is(envs []Env) bool {
	for _, t := range envs {
		if *e == t {
			return true
		}
	}
	return false
}
