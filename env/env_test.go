package env_test

import (
	"testing"

	"github.com/demaggus83/go-zeug/pkg/env"
	"github.com/demaggus83/go-zeug/pkg/test"
)

func ExampleEnv() {
	e := env.Env(env.Dev)
	if e.IsDev() {
		// change the log level to debug
		// add sql tracing
	}

	if e.IsProd() {
		// add cache to a service
	}
}

func ExampleEnv_Is() {
	e := env.Env(env.Dev)
	if e.Is([]env.Env{env.Dev, env.Stage}) {
		// we are on DEV oder STAGE
	}
}

var envTestTable = []struct {
	Env      string
	Expected bool
}{
	{env.Dev, true},
	{"DEV", false},
	{"Dev", false},
	{env.Test, true},
	{env.Stage, true},
	{env.Prod, true},
	{"invalid", false},
}

func TestIsValidEnv(t *testing.T) {
	t.Parallel()

	for _, tt := range envTestTable {
		tt := tt
		t.Run(tt.Env, func(t *testing.T) {
			t.Parallel()
			test.Equals(t, tt.Expected, env.IsValidEnv(tt.Env), `Expected IsValidEnv() to be "%t"`)
		})
	}
}

func TestUnmarshalling(t *testing.T) {
	for _, tt := range envTestTable {
		tt := tt

		t.Run(tt.Env, func(t *testing.T) {
			t.Parallel()
			var e env.Env

			err := e.Unmarshal(tt.Env)
			if tt.Expected {
				test.NoError(t, err)
			}

			if !tt.Expected {
				test.NotNil(t, err, `Expected to receive an error but got nil`)
			}
		})
	}
}

func TestIsDev(t *testing.T) {
	e := env.Env(env.Dev)
	test.Equals(t, true, e.IsDev(), `Expected IsDev() to be "%t" but got "%t"`)
	test.Equals(t, false, e.IsTest(), `Expected IsTest() to be "%t" but got "%t"`)
	test.Equals(t, false, e.IsStage(), `Expected IsStage() to be "%t" but got "%t"`)
	test.Equals(t, false, e.IsProd(), `Expected IsProd() to be "%t" but got "%t"`)
}

func TestIsTest(t *testing.T) {
	e := env.Env(env.Test)
	test.Equals(t, false, e.IsDev(), `Expected IsDev() to be "%t" but got "%t"`)
	test.Equals(t, true, e.IsTest(), `Expected IsTest() to be "%t" but got "%t"`)
	test.Equals(t, false, e.IsStage(), `Expected IsStage() to be "%t" but got "%t"`)
	test.Equals(t, false, e.IsProd(), `Expected IsProd() to be "%t" but got "%t"`)
}

func TestIsStage(t *testing.T) {
	e := env.Env(env.Stage)
	test.Equals(t, false, e.IsDev(), `Expected IsDev() to be "%t" but got "%t"`)
	test.Equals(t, false, e.IsTest(), `Expected IsTest() to be "%t" but got "%t"`)
	test.Equals(t, true, e.IsStage(), `Expected IsStage() to be "%t" but got "%t"`)
	test.Equals(t, false, e.IsProd(), `Expected IsProd() to be "%t" but got "%t"`)
}

func TestIsProd(t *testing.T) {
	e := env.Env(env.Prod)
	test.Equals(t, false, e.IsDev(), `Expected IsDev() to be "%t" but got "%t"`)
	test.Equals(t, false, e.IsTest(), `Expected IsTest() to be "%t" but got "%t"`)
	test.Equals(t, false, e.IsStage(), `Expected IsStage() to be "%t" but got "%t"`)
	test.Equals(t, true, e.IsProd(), `Expected IsProd() to be "%t" but got "%t"`)
}

func TestIs(t *testing.T) {
	e := env.Env(env.Prod)
	test.Equals(t, false, e.Is([]env.Env{env.Dev, env.Stage}), `Expected Is() to be "%t" but got "%t"`)
	test.Equals(t, true, e.Is([]env.Env{env.Dev, env.Prod}), `Expected Is() to be "%t" but got "%t"`)
}
