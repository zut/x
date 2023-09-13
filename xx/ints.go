package xx

import (
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/util/gconv"
	"github.com/pkg/errors"
	"strconv"
)

func MinInt(s []int) int {
	x, _, _, _ := MinMaxInt(s)
	return x
}
func MinIntIdx(s []int) int {
	_, minIdx, _, _ := MinMaxInt(s)
	return minIdx
}
func MaxInt(s []int) int {
	_, _, x, _ := MinMaxInt(s)
	return x
}

func MaxInt64(s ...int64) int64 {
	if len(s) == 0 {
		return 0
	}
	x := s[0]
	for _, num := range s {
		if num > x {
			x = num
		}
	}
	return x
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
	minVal, minIdx, maxVal, maxIdx := s[0], 0, s[0], 0
	for n, v := range s {
		if minVal > v {
			minVal = v
			minIdx = n
		}
		if maxVal < v {
			maxVal = v
			maxIdx = n
		}
	}
	return minVal, minIdx, maxVal, maxIdx
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

func InRangeInt(i, minVal, maxVal int) int {
	if i > maxVal {
		return maxVal
	} else if i < minVal {
		return minVal
	}
	return i
}

func InRangeInt64(i, minVal, maxVal int64) int64 {
	if i > maxVal {
		return maxVal
	} else if i < minVal {
		return minVal
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

func HexToInt64(i string) (int64, error) {
	v, err := strconv.ParseInt(i, 16, 64)
	if err != nil {
		return 0, errors.Errorf("HexToInt64 error: %v (%v)", err, i)
	}
	return v, nil
}
