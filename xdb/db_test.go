package xdb_test

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/zut/x/xdb"
	"testing"
)

func init() {
	if err := xdb.Open(); err != nil {
		glog.Fatal(context.TODO(), err)
	}
	//defer xdb.Close()
	//xdb.FlushDB()
	glog.Info(context.TODO(), xdb.Info())
}

func TestOpen(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
	})
}
