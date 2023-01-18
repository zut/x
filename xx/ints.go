package xx

import (
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gconv"
)

func MinInt(s []int) int {
	min, _, _, _ := MinMaxInt(s)
	return min
}
func MinIntIdx(s []int) int {
	_, minIdx, _, _ := MinMaxInt(s)
	return minIdx
}
func MaxInt(s []int) int {
	_, _, max, _ := MinMaxInt(s)
	return max
}

func MaxIntIdx(s []int) int {
	_, _, _, maxIdx := MinMaxInt(s)
	return maxIdx
}

// MinMaxInt return min,minIdx,max,maxIdx
func MinMaxInt(s []int) (int, int, int, int) {
	if len(s) == 0 {
		return 999999, 0, 999999, 0
	}
	min, minIdx, max, maxIdx := s[0], 0, s[0], 0
	for n, v := range s {
		if min > v {
			min = v
			minIdx = n
		}
		if max < v {
			max = v
			maxIdx = n
		}
	}
	return min, minIdx, max, maxIdx
}

func IntDiffOneValue(s []int, v int) []int {
	diff := make([]int, len(s))
	for n, i := range s {
		diff[n] = AbsInt(i - v)
	}
	return diff
}

func AbsInt(x int) int {
	return AbsDiffInt(x, 0)
}

func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

// ---

func Int(i interface{}) int {
	return gconv.Int(i)
}
func ReplaceAndInt(i interface{}) int {
	if i2, err := gregex.ReplaceString(`\D`, "", Str(i)); err == nil && len(i2) > 0 {
		return gconv.Int(i2)
	}
	return gconv.Int(i)
}
func Int64(i interface{}) int64 {
	return gconv.Int64(i)
}

func InRangeInt(i, min, max int) int {
	if i > max {
		return max
	} else if i < min {
		return min
	}
	return i
}
func InRangeInt64(i, min, max int64) int64 {
	if i > max {
		return max
	} else if i < min {
		return min
	}
	return i
}

func IntListReverse(i []int) []int {
	length := len(i)
	i2 := make([]int, length)
	for index, value := range i {
		i2[length-index-1] = value
	}
	return i2
}
