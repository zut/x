package xdb_test

import (
	"github.com/gogf/gf/test/gtest"
	"github.com/zut/x/xdb"
	"github.com/zut/x/xlog"
	"testing"
)

func init() {
	if err := xdb.Open(); err != nil {
		xlog.Fatal(err)
	}
	//defer xdb.Close()
	//xdb.FlushDB()
	xlog.Info(xdb.Info())
}

func TestOpen(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
	})
}
