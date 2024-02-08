package xx2

import (
	"fmt"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestIfError(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(IfError(false, "1"), nil)
		t.Assert(IfError(true, "1"), fmt.Errorf("1"))
	})
}
