package slice_test

import (
	"fmt"
	"testing"

	"github.com/demaggus83/go-zeug/pkg/slice"
	"github.com/demaggus83/go-zeug/pkg/test"
)

func ExampleHasString() {
	s := []string{"hello", "world"}
	fmt.Println(slice.HasString("hello", s))
	// Output: true
}

func TestHasString(t *testing.T) {
	t.Parallel()

	s := []string{"hello", "world"}
	test.True(t, slice.HasString("hello", s))
	test.False(t, slice.HasString("no", s))
}

func TestHasInt(t *testing.T) {
	t.Parallel()

	s := []int{1, 2}
	test.True(t, slice.HasInt(1, s))
	test.False(t, slice.HasInt(3, s))
}

func TestHasInt8(t *testing.T) {
	t.Parallel()

	s := []int8{1, 2}
	test.True(t, slice.HasInt8(1, s))
	test.False(t, slice.HasInt8(3, s))
}

func TestHasInt16(t *testing.T) {
	t.Parallel()

	s := []int16{1, 2}
	test.True(t, slice.HasInt16(1, s))
	test.False(t, slice.HasInt16(3, s))
}

func TestHasInt32(t *testing.T) {
	t.Parallel()

	s := []int32{1, 2}
	test.True(t, slice.HasInt32(1, s))
	test.False(t, slice.HasInt32(3, s))
}

func TestHasInt64(t *testing.T) {
	t.Parallel()

	s := []int64{1, 2}
	test.True(t, slice.HasInt64(1, s))
	test.False(t, slice.HasInt64(3, s))
}

func TestHasUint(t *testing.T) {
	t.Parallel()

	s := []uint{1, 2}
	test.True(t, slice.HasUint(1, s))
	test.False(t, slice.HasUint(3, s))
}

func TestHasUint8(t *testing.T) {
	t.Parallel()

	s := []uint8{1, 2}
	test.True(t, slice.HasUint8(1, s))
	test.False(t, slice.HasUint8(3, s))
}

func TestHasUint16(t *testing.T) {
	t.Parallel()

	s := []uint16{1, 2}
	test.True(t, slice.HasUint16(1, s))
	test.False(t, slice.HasUint16(3, s))
}

func TestHasUint32(t *testing.T) {
	t.Parallel()

	s := []uint32{1, 2}
	test.True(t, slice.HasUint32(1, s))
	test.False(t, slice.HasUint32(3, s))
}

func TestHasUint64(t *testing.T) {
	t.Parallel()

	s := []uint64{1, 2}
	test.True(t, slice.HasUint64(1, s))
	test.False(t, slice.HasUint64(3, s))
}

func TestHasFloat32(t *testing.T) {
	t.Parallel()

	s := []float32{1, 2}
	test.True(t, slice.HasFloat32(1, s))
	test.False(t, slice.HasFloat32(3, s))
}

func TestHasFloat64(t *testing.T) {
	t.Parallel()

	s := []float64{1, 2}
	test.True(t, slice.HasFloat64(1, s))
	test.False(t, slice.HasFloat64(3, s))
}

func TestHasRune(t *testing.T) {
	t.Parallel()

	s := []rune{'a', 'b'}
	test.True(t, slice.HasRune('a', s))
	test.False(t, slice.HasRune('c', s))
}
