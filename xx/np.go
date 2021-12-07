package xx

import (
	"gonum.org/v1/gonum/floats"
)

func WhereGtSI(s []int, limit, a, b int) []int {
	d := make([]int, len(s))
	for n, i := range s {
		d[n] = IfInt(i > limit, a, b)
	}
	return d
}
func WhereGtNum(s []float64, limit float64) (num int) {
	for n := range s {
		if s[n] > limit {
			num++
		}
	}
	return
}
func WhereGT(s []float64, limit, a, b float64) []float64 {
	d := make([]float64, len(s))
	for n, i := range s {
		d[n] = IfF64(i > limit, a, b)
	}
	return d
}
func WhereGT2(s []float64, limit float64) []float64 {
	d := make([]float64, len(s))
	for n, i := range s {
		d[n] = IfF64(i > limit, i, 0)
	}
	return d
}
func WhereGE(s []float64, limit, a, b float64) []float64 {
	d := make([]float64, len(s))
	for n, i := range s {
		d[n] = IfF64(i >= limit, a, b)
	}
	return d
}

func WhereLT(s []float64, limit, a, b float64) []float64 {
	d := make([]float64, len(s))
	for n, i := range s {
		d[n] = IfF64(i < limit, a, b)
	}
	return d
}

func WhereLtNum(s []float64, limit float64) (num int) {
	for n := range s {
		if s[n] < limit {
			num++
		}
	}
	return
}

// 2维变1维，然后相加
func SumWhereGT(src [][]float64, limit, a, b float64) []float64 {
	if len(src) == 0 {
		return nil
	}
	dst := make([]float64, len(src[0]))
	valueListOfPerIdx := make([]float64, len(src))
	for n := range src[0] {
		for idx := range valueListOfPerIdx {
			valueListOfPerIdx[idx] = src[idx][n]
		}
		dst[n] = IfF64(floats.Sum(valueListOfPerIdx) > limit, a, b)
	}
	return dst
}

// 一维变二维
func Split01(src []float64) [][]float64 {
	if len(src) == 0 {
		return nil
	}
	dst := make([][]float64, 0)
	last := -1.0
	lastIdx := 0
	endN := len(src) - 1
	for n, i := range src {
		//if n == len(src)-2 {
		//	xlog.Info(i)
		//}
		if last != i { // 不同,改变了
			if n != 0 { // 第一个不要进来
				dst = append(dst, src[lastIdx:n])
			}
			last = i
			lastIdx = n
			if n == endN { // 最后一个只有1个的时候, 也需要加入, 否者漏点1个
				dst = append(dst, src[lastIdx:])
			}
		} else {
			if n == endN {
				dst = append(dst, src[lastIdx:])
			}
		}
	}
	return dst
}

// 拉点
func PullUp01(src [][]float64, point int) [][]float64 {
	if len(src) == 0 {
		return nil
	}
	dst := make([][]float64, 0)
	for n, i := range src {
		if i[0] == 0 && len(i) < point {
			for m := range i {
				i[m] = 1
			}
		}
		if n == 0 || dst[len(dst)-1][0] != i[0] {
			dst = append(dst, i)
		} else {
			dst[len(dst)-1] = append(dst[len(dst)-1], i...)
		}
	}
	return dst
}
