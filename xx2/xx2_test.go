package xx2

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestInRange(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(InRange(-1, 0, 2), 0) // < min
		t.AssertEQ(InRange(0, 0, 2), 0)  // = min
		t.AssertEQ(InRange(1, 0, 2), 1)  // in range
		t.AssertEQ(InRange(2, 0, 2), 2)  // == max
		t.AssertEQ(InRange(3, 0, 2), 2)  // > max
	})
}

func TestDefault(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// string
		t.AssertEQ(Default("", "default"), "default")
		t.AssertEQ(Default("test", "default"), "test")
		// int
		t.AssertEQ(Default(0, 1), 0)
		t.AssertEQ(Default(1, 1), 1)
		// int64
		t.AssertEQ(Default(int64(0), int64(1)), int64(0))
		t.AssertEQ(Default(int64(1), int64(1)), int64(1))
		// float
		t.AssertEQ(Default(float32(0), float32(1)), float32(0))
		t.AssertEQ(Default(float32(1), float32(1)), float32(1))
		// float64
		t.AssertEQ(Default(float64(0), float64(1)), float64(0))
		t.AssertEQ(Default(float64(1), float64(1)), float64(1))

	})
}
