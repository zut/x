package xx2

import (
	"github.com/samber/lo"
)

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// InRange 限定范围 min <= i <= max
func InRange[T Comparable](i, minVal, maxVal T) T {
	if i > maxVal {
		return maxVal
	} else if i < minVal {
		return minVal
	}
	return i
}

// Default provides a default value for various types when the main value is its type's zero value.
func Default[T comparable](value T, defaultValue T) T {
	var zeroValue T // Automatically initialized to zero value of T
	return lo.Ternary(value != zeroValue, value, defaultValue)
}
