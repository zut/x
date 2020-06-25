package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/zut/x/xlog"
)

func f2() {
	xlog.Debug(1, "Skip1")
	xlog.Info(2, "Skip2")
	xlog.Warning(3)
	xlog.Error(g.Map{"1": 123123, "222": "abc"})
	xlog.Fatal(4)
	xlog.Panic(5)
}
func f1() {
	f2()
}
func main() {
	f1()
}
