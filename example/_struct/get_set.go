package main

import (
	"github.com/gogf/gf/os/glog"
	"reflect"
)

type ChPosMapFq struct {
	LCH  float64 `example:"2400"`
	MCH  float64 `example:"2412"`
	HCH  float64 `example:"2424"`
	HOPP float64
	CH12 float64 `example:"2413"`
	CH13 float64 `example:"2414"`
	CH26 float64 `example:"2480"`
}

func (v *ChPosMapFq) GetValue(field string) float64 {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Float()
}
func (v *ChPosMapFq) SetValue(field string) {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	f.SetFloat(1111)
}
func main() {
	aa := ChPosMapFq{
		LCH:  1,
		MCH:  2,
		HCH:  3,
		HOPP: 4,
		CH12: 5,
		CH13: 6,
		CH26: 7,
	}
	glog.Info(aa.LCH)
	aa.SetValue("LCH")
	glog.Info(aa.GetValue("LCH"))
}
