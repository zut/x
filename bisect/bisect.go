package bisect

import "sort"
// https://docs.python.org/ 3 /library/bisect.html

// all(val >= x for val in a)
func Left(s []float64, v float64) int {
	return sort.Search(len(s), func(i int) bool {
		return s[i] >= v
	})
}
// all(val > x for val in a)
func Right(s []float64, v float64) int {
	return sort.Search(len(s), func(i int) bool {
		return s[i] > v
	})
}
