package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/zut/x/xlog"
)

func main() {
	xlog.Info(1)
	xlog.Error(2)
	xlog.Error(g.Map{"1": 123123, "222": "abc"})
}
