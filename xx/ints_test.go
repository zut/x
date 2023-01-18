package xx_test

import (
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/zut/x/xx"
	"testing"
)

func TestMinInt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(1, xx.MinInt(xx.SI{1, 2, 3}))
		t.AssertEQ(1, xx.MinInt(xx.SI{3, 2, 1}))
	})
}
func TestMinIntIdx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(0, xx.MinIntIdx(xx.SI{1, 2, 3}))
		t.AssertEQ(2, xx.MinIntIdx(xx.SI{3, 2, 1}))
	})
}

func TestMaxInt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(3, xx.MaxInt(xx.SI{1, 2, 3}))
		t.AssertEQ(3, xx.MaxInt(xx.SI{3, 2, 1}))
	})
}
func TestMaxIntIdx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(2, xx.MaxIntIdx(xx.SI{1, 2, 3}))
		t.AssertEQ(0, xx.MaxIntIdx(xx.SI{3, 2, 1}))
	})
}

func TestMinMaxInt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		min, minIdx, max, maxIdx := xx.MinMaxInt(xx.SI{1, 2, 3})
		t.AssertEQ(1, min)
		t.AssertEQ(0, minIdx)
		t.AssertEQ(3, max)
		t.AssertEQ(2, maxIdx)
	})
}

func TestIntDiffOneValue(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := xx.IntDiffOneValue(xx.SI{1, 2, 3}, 1)
		t.AssertEQ(0, s[0])
		t.AssertEQ(1, s[1])
		t.AssertEQ(2, s[2])
	})
}

func TestAbsInt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(0, xx.AbsInt(0))
		t.AssertEQ(1, xx.AbsInt(1))
		t.AssertEQ(1, xx.AbsInt(-1))
	})
}
func TestAbsDiffInt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(1, xx.AbsDiffInt(0, 1))
		t.AssertEQ(1, xx.AbsDiffInt(1, 0))
	})
}
