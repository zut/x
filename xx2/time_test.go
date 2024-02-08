package xx2

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestT1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		T1()
		Sleep(0.1)
		v := T2()
		t.AssertGE(v, 0.1)
		t.AssertLT(v, 0.2)
	})
}

func TestT2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		T1()
		Sleep(0.1)
		v := T2()
		t.AssertGE(v, 0.1)
		t.AssertLT(v, 0.2)
	})
}

func TestSleep(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		T1()
		Sleep(0.1)
		v := T2()
		t.AssertGE(v, 0.1)
		t.AssertLT(v, 0.2)
	})
}

func TestZzz(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		Zzz()
	})
}
