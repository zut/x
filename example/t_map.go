package main

import (
	"github.com/zut/x/xlog"

	"github.com/gogf/gf/frame/g"
)

func main() {
	s4 := g.Slice{1, 2, 3, 4, 5, 6, 7}
	xlog.Info(s4[0:2])
	xlog.Info(s4[2:4])
	aMap := map[string]map[string]int{"d": {"": 1, "a": 2}}
	if val, ok := aMap["11"]; !ok {
		xlog.Info(val, ok)
	} else {
		xlog.Info(val, ok)
	}
	if val, ok := aMap[""]; !ok {
		xlog.Info(val, ok)
	} else {
		xlog.Info(val, ok)
	}
	xlog.Info(aMap["cddddd"]["c"]) // 0
}
