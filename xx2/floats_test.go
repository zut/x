package xx2

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestRound(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 2
		t.Assert(Round(0.11111111, 2), 0.11)
		t.Assert(Round(0.55555555, 2), 0.56)
		t.Assert(Round(0.99999999, 2), 1)
		// 3
		t.Assert(Round(0.11111111, 3), 0.111)
		t.Assert(Round(0.55555555, 3), 0.556)
		t.Assert(Round(0.99999999, 3), 1)
	})
}

func TestRound0(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(Round0(0.11111111), 0)
		t.Assert(Round0(0.55555555), 1)
		t.Assert(Round0(0.99999999), 1)
	})
}

func TestRound1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(Round1(0.11111111), 0.1)
		t.Assert(Round1(0.55555555), 0.6)
		t.Assert(Round1(0.99999999), 1)
	})
}

func TestRound2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(Round2(0.11111111), 0.11)
		t.Assert(Round2(0.55555555), 0.56)
		t.Assert(Round2(0.99999999), 1)
	})
}

func TestRound3(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(Round3(0.11111111), 0.111)
		t.Assert(Round3(0.55555555), 0.556)
		t.Assert(Round3(0.99999999), 1)
	})
}

func TestRound4(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(Round4(0.11111111), 0.1111)
		t.Assert(Round4(0.55555555), 0.5556)
		t.Assert(Round4(0.99999999), 1)
	})
}

func TestRound5(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(Round5(0.11111111), 0.11111)
		t.Assert(Round5(0.55555555), 0.55556)
		t.Assert(Round5(0.99999999), 1)
	})
}

func TestRound6(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(Round6(0.11111111), 0.111111)
		t.Assert(Round6(0.55555555), 0.555556)
		t.Assert(Round6(0.99999999), 1)
	})
}
