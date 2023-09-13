package xx

import (
	"fmt"
	"github.com/gogf/gf/text/gregex"
	"math"
	"sort"

	"github.com/gogf/gf/util/gconv"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/floats/scalar"
)

func Diff(s []float64) float64 {
	return Max(s) - Min(s)
}
func DiffOneValue(s []float64, v float64) []float64 {
	diff := make([]float64, len(s))
	for n, i := range s {
		diff[n] = math.Abs(i - v)
	}
	return diff
}

func DiffTowSliceMax(s1, s2 []float64, validPercent ...float64) float64 {
	if len(s1) != len(s2) {
		return 1e6
	}
	diff := make([]float64, len(s1))
	for n, i := range s1 {
		diff[n] = math.Abs(i - s2[n])
	}
	if FirstF64(validPercent) == 0 {
		return Max(diff)
	}
	sort.Float64s(diff)
	validPos := gconv.Int(float64(len(s1)) * validPercent[0] / 100)
	//xlog.Info(diff[validPos:len(diff)-1])
	return diff[validPos]
}

// math.Pow(10, i/10.0)
func LogToLine(i float64) float64 { // 如果保留小数位, 可能太小 变0
	a := math.Pow(10, i/10.0)
	return IfF64(math.IsInf(a, 0), 999999, a)
	//return Round(math.Pow(10, i/10.0), decimals...)
}
func LogToLines(i []float64) []float64 {
	i2 := make([]float64, len(i))
	for index, value := range i {
		i2[index] = LogToLine(value)
	}
	return i2
}

// 10 * math.Log10(i)
func LineToLog(i float64) float64 {
	if i <= 0 {
		//xlog.Warningf("LineToLog(%v) Error %v", i, vv.Skip1)
		return 0
	}
	a := 10 * math.Log10(i)
	return IfF64(math.IsInf(a, 0), 999999, a)
}
func LineToLogs(i []float64) []float64 {
	i2 := make([]float64, len(i))
	for index, value := range i {
		i2[index] = LineToLog(value)
	}
	return i2
}
func LogToLineSumLineToLog(i []float64) float64 {
	return LineToLog(floats.Sum(LogToLines(i)))
}
func YListLogToLineAndP2PLineToLog(yList [][]float64) []float64 {
	yListLine := make([][]float64, len(yList))
	for n, y := range yList {
		yListLine[n] = LogToLines(y)
	}
	ySum := make([]float64, len(yListLine[0]))
	// 点对点相加
	vLineListOfPos := make([]float64, len(yListLine))
	for n := range ySum {
		for nOfy := range vLineListOfPos {
			vLineListOfPos[nOfy] = yListLine[nOfy][n]
		}
		ySum[n] = LineToLog(floats.Sum(vLineListOfPos))
	}
	return ySum
}
func YListP2PSum(yList [][]float64) []float64 {
	ySum := make([]float64, len(yList[0]))
	// 点对点相加
	vLineListOfPos := make([]float64, len(yList))
	for n := range ySum {
		for nOfy := range vLineListOfPos {
			vLineListOfPos[nOfy] = yList[nOfy][n]
		}
		ySum[n] = floats.Sum(vLineListOfPos)
	}
	return ySum
}
func ClearAndF64(i interface{}) float64 {
	s, err := gregex.MatchString(`[+\-]\d[\d\.]*`, gconv.String(i))
	if err == nil && len(s) > 0 {
		return gconv.Float64(s[0])
	}
	return gconv.Float64(i)
}
func F64(i interface{}) float64 {
	// NAV returned by FETCh and Calculate: Measurement state OFF
	// NCAP returned by FETCh and Calculate: Disabled view
	// INV returned: Underflow, overflow, sync error, trigger timeout
	switch i {
	case "INV":
		return 999999
	case "NAV", "NCAP", "OFF":
		return 0
	case "ON":
		return 1
	}
	return gconv.Float64(i)
}
func F64s(s []string) []float64 {
	s2 := make([]float64, len(s))
	for n := range s {
		s2[n] = F64(s[n])
	}
	return s2
}

// Range 0,3 > [0,1,2]
func Range(start, stop int) []float64 {
	a := make([]float64, stop-start)
	for i := 0; i < stop-start; i++ {
		a[i] = F64(start + i)
	}
	return a
}

// RangeInt 0,3 > [0,1,2]
func RangeInt(start, stop int) []int {
	a := make([]int, stop-start)
	for i := 0; i < stop-start; i++ {
		a[i] = start + i
	}
	return a
}

func F64sReverse(i []float64, decimals ...int) []float64 {
	length := len(i)
	i2 := make([]float64, length)
	for index, value := range i {
		i2[length-index-1] = Round(value, decimals...)
	}
	return i2
}
func InsertSF64(s []float64, index int, value float64) []float64 {
	s2 := CopySF64(s)
	tmp := append([]float64{}, s2[index:]...)
	s2 = append(s2[0:index], value)
	s2 = append(s2, tmp...)
	return s2
}
func InsertSStr(s []string, index int, value string) []string {
	s2 := CopyST(s)
	tmp := append([]string{}, s2[index:]...)
	s2 = append(s2[0:index], value)
	s2 = append(s2, tmp...)
	return s2
}
func DelSStr(s []string, delValue string) []string {
	s2 := make([]string, 0)
	for _, i := range s {
		if i == delValue {
			continue
		}
		s2 = append(s2, i)
	}
	return s2
}
func IntSAdd(i []int, a int) []int {
	i2 := make([]int, len(i))
	for n, v := range i {
		i2[n] = v + a
	}
	return i2
}
func F64sAdd(i []float64, a float64, decimals ...int) []float64 {
	i2 := make([]float64, len(i))
	for index, value := range i {
		i2[index] = Round(value+a, decimals...)
		i2[index] = IfF64(math.IsInf(i2[index], 0), 999999, i2[index])
	}
	return i2
}

func Max(s []float64) float64 {
	if len(s) == 0 {
		return 999999
	}
	a := floats.Max(s)
	return IfF64(math.IsInf(a, 0), 999999, a)
}

func MaxF64(s ...float64) float64 {
	if len(s) == 0 {
		return 999999
	}
	return floats.Max(s)
}

func Min(s []float64) float64 {
	if len(s) == 0 {
		return 999999
	}
	a := floats.Min(s)
	return IfF64(math.IsNaN(a), 999999, a)
}
func MinF64(s ...float64) float64 {
	if len(s) == 0 {
		return 999999
	}
	return floats.Min(s)
}
func ToMinF64(y []float64, minVal float64) []float64 {
	for n, i := range y {
		y[n] = MaxF64(i, minVal)
	}
	return y
}
func ToMaxF64(y []float64, maxVal float64) []float64 {
	for n, i := range y {
		y[n] = MinF64(i, maxVal)
	}
	return y
}

func InRangeF64(i, minVal, maxVal float64) float64 {
	if i > maxVal {
		return maxVal
	} else if i < minVal {
		return minVal
	}
	return i
}

func IsInRangeF64(i, minVal, maxVal float64) bool {
	return i >= minVal && i <= maxVal
}

func MaxIdx(s []float64) int {
	return floats.MaxIdx(s)
}
func MinIdx(s []float64) int {
	return floats.MinIdx(s)
}

// Greater Than
func F64sGTIdx(s []float64, limit float64) int {
	for n, i := range s {
		if i > limit {
			return n
		}
	}
	return len(s)
}

// https://golang.org/src/sort/example_test.go
// min > max
func Float64sSort(s []float64) []float64 {
	d := CopySF64(s)
	sort.Float64s(d)
	return d
}

// Float64sSortMaxToMin max > min
func Float64sSortMaxToMin(s []float64) []float64 {
	return ReverseF64(Float64sSort(s))
}

func Avg(s []float64) float64 {
	return floats.Sum(s) / F64(len(s))
}

func Abs(a float64) float64 {
	return math.Abs(a)
}

// Sum Sum
func Sum(s []float64) float64 {
	return floats.Sum(s)
}

// ArgSort min > max
func ArgSort(s []float64) []int {
	idx := make([]int, len(s))
	floats.Argsort(CopySF64(s), idx)
	return idx
}

// ArgSortMaxToMin max > min
func ArgSortMaxToMin(s []float64) []int {
	return IntListReverse(ArgSort(s))
}
func Float64sAddRandom(i []float64, x1 float64, x2 float64, decimals ...int) []float64 {
	i2 := make([]float64, len(i))
	for index, value := range i {
		i2[index] = Round(value+RandomF64(x1, x2), decimals...)
	}
	return i2
}

func ZFill(i float64, decimals int) string {
	return fmt.Sprintf(fmt.Sprintf("%%.%vf", decimals), i)
}

// Round decimals: 6
func Round(i float64, decimals ...int) float64 {
	return scalar.Round(i, OrInt(6, decimals...))
}

// R0 Round 1
func R0(i float64) float64 {
	return scalar.Round(i, 0)
}

// Round 1
func R1(i float64) float64 {
	return scalar.Round(i, 1)
}

// Round 2
func R2(i float64) float64 {
	return scalar.Round(i, 2)
}

// Round 3
func R3(i float64) float64 {
	return scalar.Round(i, 3)
}

func ToMs(i float64) float64 {
	return R3(i * 1000)
}
func ToPercent(i float64) float64 {
	return R2(i * 100)
}

// Deprecated: Use Other
func DeprecatedTest(i float64) float64 {
	return scalar.Round(i, 3)
}

// IntZ2 ZFill
func IntZ2(i int) string {
	return fmt.Sprintf("%.2d", i)
}

// IntZ3 ZFill
func IntZ3(i int) string {
	return fmt.Sprintf("%.3d", i)
}

// IntZ8 ZFill
func IntZ8(i int) string {
	return fmt.Sprintf("%.8d", i)
}

// Z1 ZFill
func Z1(i float64) string {
	return fmt.Sprintf("%.1f", i)
}

// Z2 ZFill
func Z2(i float64) string {
	return fmt.Sprintf("%.2f", i)
}

// Z3 ZFill
func Z3(i float64) string {
	return fmt.Sprintf("%.3f", i)
}

// Z4 ZFill
func Z4(i float64) string {
	return fmt.Sprintf("%.4f", i)
}

// R4  Round 4
func R4(i float64) float64 {
	return scalar.Round(i, 4)
}

// R5 Round 5
func R5(i float64) float64 {
	return scalar.Round(i, 5)
}

// R6 Round 6
func R6(i float64) float64 {
	return scalar.Round(i, 6)
}

// R8 Round 8
func R8(i float64) float64 {
	return scalar.Round(i, 8)
}

// R9 Round 9
func R9(i float64) float64 {
	return scalar.Round(i, 9)
}

// Rounds Copy
func Rounds(s []float64, decimals ...int) []float64 {
	s2 := make([]float64, len(s))
	for index, value := range s {
		s2[index] = Round(value, decimals...)
	}
	return s2
}

// Copy
func R0s(s []float64) []float64 {
	return Rounds(s, 0)
}
func R1s(s []float64) []float64 {
	return Rounds(s, 1)
}

// Copy
func R2s(s []float64) []float64 {
	return Rounds(s, 2)
}
func R3s(s []float64) []float64 {
	return Rounds(s, 3)
}
func R6s(s []float64) []float64 {
	return Rounds(s, 6)
}
func CopySF64(src []float64) []float64 {
	dst := make([]float64, len(src))
	copy(dst, src)
	return dst
}
func CopySInt(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}
func CopyST(s []string) []string {
	d := make([]string, len(s))
	copy(d, s)
	return d
}

func Linspace(start, end float64, num int) []float64 {
	if num < 0 {
		panic("Linspace.num<0")
	}
	result := make([]float64, num)
	step := (end - start) / float64(num-1)
	for i := range result {
		result[i] = start + float64(i)*step
	}
	return result
}

func CalcPercent(i interface{}, total interface{}) int {
	if F64(total) == 0 {
		return 0
	}
	return int(F64(i) / F64(total) * 100)
}

func ProductSF64(ssf [][]float64) [][]float64 {
	var toRet [][]float64
	if len(ssf) == 1 {
		for _, a := range ssf[0] {
			toRet = append(toRet, []float64{a})
		}
		return toRet
	}
	t := ProductSF64(ssf[1:])
	for _, a := range ssf[0] {
		for _, perm := range t {
			toRetAdd := append([]float64{a}, perm...)
			toRet = append(toRet, toRetAdd)
		}
	}
	return toRet
}

func CompareF64(a, b float64, Operation string) bool {
	// LT：less than 小于
	// LE：less than or equal to 小于等于
	// EQ：equal to 等于
	// NE：not equal to 不等于
	// GE：greater than or equal to 大于等于
	// GT：greater than 大于
	switch Operation {
	case "LT":
		return a < b
	case "LE":
		return a <= b
	case "EQ":
		return a == b
	case "NE":
		return a != b
	case "GE":
		return a >= b
	case "GT":
		return a > b
	}
	return false
}

func PermutationsInt(s []int) [][]int {
	var helper func([]int, int)
	var res [][]int
	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(s, len(s))
	return res
}

func MakeFloat64(length int, v float64) []float64 {
	s := make([]float64, length)
	for i := range s {
		s[i] = v
	}
	return s
}
