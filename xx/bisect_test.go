package xx_test

import (
	"github.com/gogf/gf/test/gtest"
	"github.com/zut/x/xx"
	"testing"
)

func Test_BisectLeft(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3}
		vTarget := 1
		v1 := xx.BisectLeft(v0, 1.5)
		t.AssertEQ(v1, vTarget)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3}
		vTarget := 1
		v1 := xx.BisectLeft(v0, 2)
		t.AssertEQ(v1, vTarget)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3}
		vTarget := 2
		v1 := xx.BisectLeft(v0, 2.1)
		t.AssertEQ(v1, vTarget)
	})
}

func Test_Right(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3}
		vTarget := 1
		v1 := xx.BisectRight(v0, 1.5)
		t.AssertEQ(v1, vTarget)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3}
		vTarget := 2 // right
		v1 := xx.BisectRight(v0, 2)
		t.AssertEQ(v1, vTarget)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3}
		vTarget := 2
		v1 := xx.BisectRight(v0, 2.1)
		t.AssertEQ(v1, vTarget)
	})
}
