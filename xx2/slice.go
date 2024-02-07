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

// First returns the first element of a collection or error if empty.
func First[T any](s []T) T {
	if len(s) == 0 {
		var zeroValue T
		return zeroValue
	}
	return s[0]
}
