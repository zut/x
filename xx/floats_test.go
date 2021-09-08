package xx_test

import (
	"fmt"
	"github.com/gogf/gf/test/gtest"
	"github.com/zut/x/model"
	"github.com/zut/x/xx"
	"testing"
)

func Test_FloatsDiff(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := []float64{1, 2, 3}
		s2 := []float64{3, 4, 6}
		v0 := 3.0
		v1 := xx.DiffTowSliceMax(s1, s2)
		t.AssertEQ(v1, v0)
	})
	gtest.C(t, func(t *gtest.T) {
		xx.T1()
		s1 := xx.RandomF64s(1, 2, 1e5)
		s2 := xx.RandomF64s(1, 2, 1e5)
		v1 := xx.DiffTowSliceMax(s1, s2)
		fmt.Println(v1)
		xx.T2()
	})
}

func Test_Float64sReverse(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3}
		v1 := xx.F64sReverse(v0)
		t.AssertEQ(v1, []float64{3, 2, 1})
	})
}

func Test_Float64sInsert(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3}
		v1 := []float64{1, 2, 111, 3}
		v0 = xx.InsertSF64(v0, 2, 111)
		t.AssertEQ(v0, v1)
	})
}
func Test_Float64sGTIdx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3, 4, 5}
		v1 := xx.F64sGTIdx(v0, 2)
		t.AssertEQ(v1, 2)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{6, 5, 4, 3, 2, 1}
		v1 := xx.F64sGTIdx(xx.F64sReverse(v0), 2)
		t.AssertEQ(6-v1-1, 3)
	})
}
func Test_Float64sSort(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{2, 1, 3}
		vTarget := []float64{1, 2, 3}
		v1 := xx.Float64sSort(v0)
		t.AssertEQ(v1, vTarget)
	})
}

func Test_LogToLine(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := 11.0
		v1 := xx.LogToLine(v0)
		t.AssertEQ(xx.R6(v1), 12.589254)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := 10.0
		v1 := xx.LogToLine(v0)
		t.AssertEQ(v1, 10.0)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := 1.0
		v1 := xx.LogToLine(v0)
		t.AssertEQ(xx.R6(v1), 1.258925)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := 0.0
		v1 := xx.LogToLine(v0)
		t.AssertEQ(v1, 1.0)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := -1.0
		v1 := xx.LogToLine(v0)
		t.AssertEQ(xx.R6(v1), 0.794328)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := -100.0
		v1 := xx.LogToLine(v0)
		t.AssertEQ(v1, 0.0000000001)
	})
}

func Test_LineToLog(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := 11.0
		v1 := xx.LineToLog(v0)
		t.AssertEQ(xx.R6(v1), 10.413927)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := 10.0
		v1 := xx.LineToLog(v0)
		t.AssertEQ(v1, 10.0)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := 1.0
		v1 := xx.LineToLog(v0)
		t.AssertEQ(v1, 0.0)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := 0.0
		v1 := xx.LineToLog(v0)
		t.AssertEQ(v1, 0.0)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := -1.0
		v1 := xx.LineToLog(v0)
		t.AssertEQ(v1, 0.0)
	})
}

func Test_GetOutliers(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v1 := model.Outlier{
			Data: []float64{1, 201, 115, 83, 2, 201, 201, 1, 201},
		}
		v1.Calc()
		t.Assert(v1.ValList, []float64{1, 2, 1})
		t.Assert(v1.IdxList, []float64{0, 4, 7})
	})
	gtest.C(t, func(t *gtest.T) {
		v1 := model.Outlier{
			Data: []float64{19, 9, 16, 382, 7, 12, 6, 13, 4, 5},
		}
		v1.Calc()
		t.Assert(v1.ValList, []float64{382})
		t.Assert(v1.IdxList, []float64{3})
	})
	gtest.C(t, func(t *gtest.T) {
		v1 := model.Outlier{
			Data: []float64{1, 2, 3, 4, 1, 2, 3, 4, 5, 1, 200},
		}
		v1.Calc()
		t.Assert(v1.ValList, []float64{200})
		t.Assert(v1.IdxList, []float64{10})
	})
}
