package xx2

// CopySlice creates and returns a copy of the provided slice.
func CopySlice[T any](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}

// DefaultSlice provides a copy of the given slice if it's not empty,
// or a copy of the default value slice otherwise, based on the copy flag.
// The copy flag is optional and defaults to false if not provided.
func DefaultSlice[T any](s []T, defaultS []T, notCopy ...int) []T {
	shouldCopy := len(notCopy) == 0 || notCopy[0] != 1
	if len(s) > 0 {
		if shouldCopy {
			return CopySlice(s)
		}
		return s
	}
	if shouldCopy {
		return CopySlice(defaultS)
	}
	return defaultS
}

// EqualSlice Compare two slices to see if they are the same, in the same order and with the same value.
func EqualSlice[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// First returns the first element of a collection or error if empty.
func First[T any](s []T) T {
	if len(s) == 0 {
		var zeroValue T
		return zeroValue
	}
	return s[0]
}
