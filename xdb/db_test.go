package xdb_test

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/zut/x/xdb"
	"testing"
)

func init() {
	if err := xdb.Open(); err != nil {
		glog.Fatal(gctx.New(), err)
	}
	//defer xdb.Close()
	//xdb.FlushDB()
	glog.Info(gctx.New(), xdb.Info())
}

func TestOpen(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
	})
}
