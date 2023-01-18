package xx

import "sort"

// https://docs.python.org/ 3 /library/bisect.html

// BisectLeft all(val >= x for val in a) // https://docs.python.org/ 3 /library/bisect.html
func BisectLeft(s []float64, v float64) int {
	return sort.Search(len(s), func(i int) bool {
		return s[i] >= v
	})
}

// BisectRight  all(val > x for val in a) // https://docs.python.org/ 3 /library/bisect.html
func BisectRight(s []float64, v float64) int {
	return sort.Search(len(s), func(i int) bool {
		return s[i] > v
	})
}
