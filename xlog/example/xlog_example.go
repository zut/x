package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/zut/x/xlog"
)

func main() {
	xlog.Debug(1)
	xlog.Info(2)
	xlog.Warning(3)
	xlog.Error(g.Map{"1": 123123, "222": "abc"})
	xlog.Fatal(4)
	xlog.Panic(5)
}
