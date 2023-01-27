package xlog_test

import (
	"github.com/gogf/gf/test/gtest"
	"github.com/zut/x/xlog"
	"testing"
)

func TestLogger_Info(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		xlog.Info(1)
	})
}
