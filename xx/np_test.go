package xx

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestSumWhere01(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v1 := [][]float64{
			{1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1},
			{1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1},
			{1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1},
		}
		v2 := []float64{
			1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1}
		v3 := SumWhereGT(v1, 0, 1, 0)
		t.AssertEQ(v2, v3)
	})
}

func TestSplit01(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v1 := []float64{
			1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1}
		v2 := [][]float64{{1}, {0}, {1}, {0, 0, 0}, {1, 1, 1, 1, 1, 1, 1, 1, 1}}
		v3 := Split01(v1)
		t.Assert(v2, v3)
	})
}

func TestPullUp01(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v1 := [][]float64{{1}, {0}, {1}, {0, 0, 0}, {1, 1, 1, 1, 1, 1, 1, 1, 1}}
		v2 := [][]float64{{1, 1, 1}, {0, 0, 0}, {1, 1, 1, 1, 1, 1, 1, 1, 1}}
		v3 := PullUp01(v1, 3)
		t.AssertEQ(v2, v3)
	})
}
