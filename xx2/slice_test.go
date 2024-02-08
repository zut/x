package xx2

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestCopySlice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// int
		t.AssertEQ(CopySlice([]int{1, 2, 3}), []int{1, 2, 3})
		// string
		t.AssertEQ(CopySlice([]string{"1", "2", "3"}), []string{"1", "2", "3"})
		// float
		t.AssertEQ(CopySlice([]float64{1.1, 2.2, 3.3}), []float64{1.1, 2.2, 3.3})
	})
}

func TestDefaultSlice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// int
		t.AssertEQ(DefaultSlice([]int{1, 2, 3}, []int{1, 2, 3}), []int{1, 2, 3})
		t.AssertEQ(DefaultSlice([]int{1, 2, 3}, []int{1, 2, 3}, 1), []int{1, 2, 3})
		// string
		t.AssertEQ(DefaultSlice([]string{"1", "2", "3"}, []string{"1", "2", "3"}), []string{"1", "2", "3"})
		t.AssertEQ(DefaultSlice([]string{"1", "2", "3"}, []string{"1", "2", "3"}, 1), []string{"1", "2", "3"})
		// float
		t.AssertEQ(DefaultSlice([]float64{1.1, 2.2, 3.3}, []float64{1.1, 2.2, 3.3}), []float64{1.1, 2.2, 3.3})
		t.AssertEQ(DefaultSlice([]float64{1.1, 2.2, 3.3}, []float64{1.1, 2.2, 3.3}, 1), []float64{1.1, 2.2, 3.3})
	})
}

func TestEqualSlice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// same
		t.Assert(EqualSlice([]int{1, 2, 3}, []int{1, 2, 3}), true)
		t.Assert(EqualSlice([]string{"1", "2", "3"}, []string{"1", "2", "3"}), true)
		t.Assert(EqualSlice([]float64{1.1, 2.2, 3.3}, []float64{1.1, 2.2, 3.3}), true)
		// not same
		t.Assert(EqualSlice([]int{1, 2, 3}, []int{1, 2, 3, 4}), false)
		t.Assert(EqualSlice([]int{1, 2, 3}, []int{1, 2, 4}), false)
		t.Assert(EqualSlice([]string{"1", "2", "3"}, []string{"1", "2", "3", "4"}), false)
		t.Assert(EqualSlice([]float64{1.1, 2.2, 3.3}, []float64{1.1, 2.2, 3.3, 4.4}), false)
	})
}

func TestFirst(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
	})
}
