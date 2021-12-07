package xx

import (
	"testing"

	"github.com/gogf/gf/test/gtest"
	"github.com/zut/x/xlog"
)

func TestGetDiskInfo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := GetPsInfo()
		xlog.Info(a)
	})
}

func TestIp2Location(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s, err := Ip2Location(ST{
			"113.104.251a.44",
			"27.38.155.8",
			"163.125.234.241",
		})
		xlog.Info(s)
		xlog.Info(err)
	})
}
