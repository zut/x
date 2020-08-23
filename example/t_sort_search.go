package main

import (
	"github.com/zut/x/xlog"
	"sort"
)

func main() {
	aa := []float64{1, 2, 3, 4, 5,1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	xlog.Info(sort.Search(len(aa), func(i int) bool { return aa[i] >= 2.9 }))      //2
	xlog.Info(sort.Search(len(aa), func(i int) bool { return aa[i] >= 3 }))        //2
	xlog.Info(sort.Search(len(aa), func(i int) bool { return aa[i] >= 3.1 }))      //3
	xlog.Info(sort.Search(len(aa), func(i int) bool { return aa[i] >= 10000000 })) //10
}
