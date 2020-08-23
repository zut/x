package main

import (
	"github.com/zut/x/xlog"
	"gonum.org/v1/gonum/floats"
)

func main() {
	aa := []float64{1, 3, 6, 2, 4}
	idx := make([]int, len(aa))
	floats.Argsort(aa, idx)
	xlog.Info(idx)
}
