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
