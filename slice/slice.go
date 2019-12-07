// Package slice provides some helper methods for working with slices.
package slice

// HasString checks if 'val' is in 'slice'.
func HasString(val string, slice []string) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}

// HasRune checks if 'val' is in 'slice'.
func HasRune(val rune, slice []rune) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}

// HasInt checks if 'val' is in 'slice'.
func HasInt(val int, slice []int) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}

// HasInt8 checks if 'val' is in 'slice'.
func HasInt8(val int8, slice []int8) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}

// HasInt16 checks if 'val' is in 'slice'.
func HasInt16(val int16, slice []int16) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}

// HasInt32 checks if 'val' is in 'slice'.
func HasInt32(val int32, slice []int32) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}

// HasInt64 checks if 'val' is in 'slice'.
func HasInt64(val int64, slice []int64) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}

// HasUint checks if 'val' is in 'slice'.
func HasUint(val uint, slice []uint) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}

// HasUint8 checks if 'val' is in 'slice'.
func HasUint8(val uint8, slice []uint8) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}

// HasUint16 checks if 'val' is in 'slice'.
func HasUint16(val uint16, slice []uint16) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}

// HasUint32 checks if 'val' is in 'slice'.
func HasUint32(val uint32, slice []uint32) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}

// HasUint64 checks if 'val' is in 'slice'.
func HasUint64(val uint64, slice []uint64) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}

// HasFloat32 checks if 'val' is in 'slice'.
func HasFloat32(val float32, slice []float32) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}

// HasFloat64 checks if 'val' is in 'slice'.
func HasFloat64(val float64, slice []float64) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}
