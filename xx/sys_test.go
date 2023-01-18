package xx

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestGetDiskInfo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := GetPsInfo()
		glog.Info(gctx.New(), a)
	})
}

func TestIp2Location(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s, err := Ip2Location(ST{
			"113.104.251a.44",
			"27.38.155.8",
			"163.125.234.241",
		})
		glog.Info(gctx.New(), s)
		glog.Info(gctx.New(), err)
	})
}
