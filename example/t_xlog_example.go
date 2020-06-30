package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/zut/x/xlog"
)

func f2() {
	defer end()
	xlog.Debug(1, "Skip1")
	xlog.Info(2, "Skip2")
	xlog.Warning(3)
	xlog.Error(g.Map{"1": 123123, "222": "abc"})
	//xlog.Panic(5)
	//xlog.Fatal(4) // 没有 defer end
}
func f1() {
	f2()
}
func end() {
	fmt.Println("end ...")
}
func main() {
	err := xlog.SetConfigWithMap(g.Map{
		"StLevel": 1,
	})
	if err != nil {
		panic(err)
	}
	glog.Info(1)
	f1()
}
