package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/zut/x/xlog"
)

func f2() {
	defer end()
	xlog.Debug("111 Skip1")
	xlog.Debug(1, "1 Skip1")
	xlog.Info(2, "Skip2")
	xlog.Error(3, "Skip1")
	xlog.Warning(3)
	xlog.Error(g.Map{"1": 123123, "222": "abc"})
	//xlog.Panic(5)
	//xlog.Fatal(4) // 没有 defer end 123
}
func f1() {
	f2()
}
func end() {
	fmt.Println("end ...")
}
func main() {
	err := xlog.SetConfigWithMap(g.Map{
		"StLevel": 2,
	})
	if err != nil {
		panic(err)
	}
	xlog.Info(1)
	f1()
}
