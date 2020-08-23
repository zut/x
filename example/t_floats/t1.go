package main

import (
	"github.com/zut/x/xlog"
	"gonum.org/v1/gonum/floats"
	"sort"
)

func main() {
	aa := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dst := make([]float64, len(aa))
	floats.CumSum(dst, aa)
	xlog.Info(dst)

	xlog.Info(sort.Search(len(aa), func(i int) bool { return aa[i] >= 2.9 }))
	xlog.Info(sort.Search(len(aa), func(i int) bool { return aa[i] >= 3 }))
	xlog.Info(sort.Search(len(aa), func(i int) bool { return aa[i] >= 3.1 }))
	//print(np.cumsum(aa))
	//print(np.searchsorted(np.cumsum(aa), 2))
}
