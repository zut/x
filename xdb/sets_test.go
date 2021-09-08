package xdb_test

import (
	"github.com/gogf/gf/test/gtest"
	"github.com/zut/x/xdb"
	"testing"
)


func TestSAdd(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.SAdd(key, value), nil)
	})
}

func TestSCard(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
	})
}

func TestSDiff(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.SAdd(key, value), nil)
	})
}
