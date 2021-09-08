package xx_test

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"github.com/zut/x/xx"
	"testing"
)

func TestRandomFloat(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := xx.RandomF64(1, 2)
		t.AssertGE(v0, 1)
		t.AssertLE(v0, 2)
	})
}
func TestRandomFloats(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := xx.RandomF64s(1, 2, 5)
		fmt.Println(v0)
		t.AssertGE(v0[0], 1)
		t.AssertLE(v0[0], 2)
	})
}
func TestMinInt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := xx.MinInt(xx.SI{1, 2, 3})
		t.AssertEQ(1, v0)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := xx.MinInt(xx.SI{1, 2, 3})
		t.AssertEQ(1, v0)
	})
}

func TestJoin(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := xx.Join(111)
		t.AssertEQ("111", v0)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := xx.Join(1, 2, 3)
		t.AssertEQ("1_2_3", v0)
	})
}

func TestCompareStruct(t *testing.T) {
	type Foo struct {
		Name string
		Date int
	}
	gtest.C(t, func(t *gtest.T) {
		f1 := &Foo{"Drew", 30}
		f2 := &Foo{"Drew", 50}
		v1 := xx.CompareStruct(f1, f2)
		t.Assert(v1, []string{"Date"})
	})
	gtest.C(t, func(t *gtest.T) {
		f1 := &Foo{"Drew", 1}
		f2 := &Foo{"Drew", 1}
		v1 := xx.CompareStruct(f1, f2)
		t.Assert(v1, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		f1 := &Foo{"Drew", 1}
		v1 := xx.CompareStruct(f1, nil)
		t.Assert(v1, []string{"Name", "Date"})
	})
	gtest.C(t, func(t *gtest.T) {
		f1 := &Foo{"Drew", 1}
		v1 := xx.CompareStruct(f1, []string{})
		t.Assert(v1, []string{"Name", "Date"})
	})
}

func TestReverse(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := g.Slice{1, 2, 3}
		vTarget := g.Slice{3, 2, 1}
		v1 := xx.Reverse(v0)
		t.Assert(vTarget, v1)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := g.Slice{1, 2, 3, 4}
		vTarget := g.Slice{4, 3, 2, 1}
		v1 := xx.Reverse(v0)
		t.Assert(vTarget, v1)
	})
}

func TestReverseString(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := []string{"1", "2", "3"}
		vTarget := []string{"3", "2", "1"}
		v1 := xx.ReverseStr(v0)
		t.Assert(vTarget, v1)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := []string{"1", "2", "3", "4"}
		vTarget := []string{"4", "3", "2", "1"}
		v1 := xx.ReverseStr(v0)
		t.Assert(vTarget, v1)
	})
}

func TestReverseFloat64(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3}
		vTarget := []float64{3, 2, 1}
		v1 := xx.ReverseF64(v0)
		t.Assert(vTarget, v1)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3, 4}
		vTarget := []float64{4, 3, 2, 1}
		v1 := xx.ReverseF64(v0)
		t.Assert(vTarget, v1)
	})
}

func TestCrossPoint(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 3, 5, 1}
		vTarget := 2
		v1, _ := xx.CrossPoint(v0, 2)
		t.Assert(vTarget, v1)
	})
	gtest.C(t, func(t *gtest.T) {
		v0 := []float64{1, 2, 3, 4, 5, 6, 1, 5, 2, 2, 2, 2, 1}
		vTarget := 4
		v1, _ := xx.CrossPoint(v0, 3)
		t.Assert(vTarget, v1)
	})

}

func TestCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type Item struct {
			Name string
			Data []int
		}
		v0 := Item{
			Name: "v0",
			Data: []int{1, 2, 3},
		}
		//v1 := Item{}
		v1 := new(Item)
		err := xx.Copy(v0, v1)
		t.Assert(err, nil)
		t.Assert(v0.Name, v1.Name)
		t.Assert(v0.Data, v1.Data)
		v1.Name = "v1"
		v1.Data[0] = 111
		t.Assert(v0.Name, "v0")
		t.Assert(v0.Data[0], 1)

	})
}
