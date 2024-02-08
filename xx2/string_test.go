package xx2

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestSubStringHan(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(SubStringHan("1", 1), "1")
	})
}

func TestSafeFilename(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(SafeFilename("1"), "1")
	})
}
