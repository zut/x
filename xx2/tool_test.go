package xx2

import (
	"github.com/zut/x/xlog"
	"testing"

	"github.com/gogf/gf/test/gtest"
)

func TestGetDiskInfo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := GetPsInfo()
		xlog.Info(a)
	})
}

func TestIp2Location(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s, err := Ip2Location([]string{
			"113.104.251a.44",
			"27.38.155.8",
			"163.125.234.241",
		})
		t.Assert(err, nil)
		t.Assert(len(s), 3)

		t.Assert(s[0].Country, "Invalid IP address.")
		t.Assert(s[0].CountryShort, "Invalid IP address.")

		t.Assert(s[1].Country, "China")
		t.Assert(s[1].CountryShort, "CN")
		t.Assert(s[1].Province, "Guangdong")
		t.Assert(s[1].City, "Shenzhen")

		t.Assert(s[2].City, "Shenzhen")
	})
}
