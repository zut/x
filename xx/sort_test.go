package xx_test

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/zut/x/xx"
	"testing"
)

func TestSortMapByKey(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := map[string]interface{}{"b": 23, "a": 2, "c": 25}
		v1 := xx.SortMapByKey(v0)
		t.Assert(v1, g.SliceStr{"a", "b", "c"})
	})
}

func TestSortMapByValue(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := map[string]float64{"b": 23, "c": 2, "a": 25}
		v1 := xx.SortMapByValue(v0)
		t.Assert(v1, g.SliceStr{"a", "b", "c"})
	})
}
