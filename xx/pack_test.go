package xx_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/zut/x/xx"
)

type Item struct {
	Name        string
	privateName string
	NameOther   string
}
type Item2 struct {
	Name  string
	Name2 string
	Name3 string
}

func TestPack(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		oo := Item{
			Name:        xx.RandomLetter(10),
			privateName: "privateName",
			NameOther:   "NameOther",
		}
		b, err := xx.Pack(oo)
		t.Assert(err, nil)
		var aa Item2
		err = xx.UnpackTo(b, &aa)
		t.Assert(err, nil)
		t.Assert(aa.Name, oo.Name)
	})
}

func Test_UnPack(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		b, err := xx.Pack(123)
		t.Assert(err, nil)
		value, e2 := xx.Unpack(b)
		t.Assert(e2, nil)
		t.Assert(value, 123)
	})
}

func Test_UnpackTo(t *testing.T) {
}
