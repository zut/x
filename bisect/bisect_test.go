package bisect_test

import (
	"github.com/gogf/gf/test/gtest"
	"github.com/zut/x/bisect"
	"testing"
)

func Test_Left(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3}
		vTarget := 1
		v1 := bisect.Left(v0, 1.5)
		t.AssertEQ(v1, vTarget)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3}
		vTarget := 1
		v1 := bisect.Left(v0, 2)
		t.AssertEQ(v1, vTarget)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3}
		vTarget := 2
		v1 := bisect.Left(v0, 2.1)
		t.AssertEQ(v1, vTarget)
	})
}

func Test_Right(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3}
		vTarget := 1
		v1 := bisect.Right(v0, 1.5)
		t.AssertEQ(v1, vTarget)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3}
		vTarget := 2 // right
		v1 := bisect.Right(v0, 2)
		t.AssertEQ(v1, vTarget)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3}
		vTarget := 2
		v1 := bisect.Right(v0, 2.1)
		t.AssertEQ(v1, vTarget)
	})
}
