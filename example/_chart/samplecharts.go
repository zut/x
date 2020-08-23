package main

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"
	"github.com/vdobler/chart/svgg"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

// Dumper helps saving plots of size WxH in a NxM grid layout
// in several formats
type Dumper struct {
	N, M, W, H, Cnt  int
	S                *svg.SVG
	I                *image.RGBA
	svgFile, imgFile *os.File
}

func NewDumper(name string, n, m, w, h int) *Dumper {
	var err error
	dumper := Dumper{N: n, M: m, W: w, H: h}

	dumper.svgFile, err = os.Create(name + ".svg")
	if err != nil {
		panic(err)
	}
	dumper.S = svg.New(dumper.svgFile)
	dumper.S.Start(n*w, m*h)
	dumper.S.Title(name)
	dumper.S.Rect(0, 0, n*w, m*h, "fill: #ffffff")

	dumper.imgFile, err = os.Create(name + ".png")
	if err != nil {
		panic(err)
	}
	dumper.I = image.NewRGBA(image.Rect(0, 0, n*w, m*h))
	bg := image.NewUniform(color.RGBA{0xff, 0xff, 0xff, 0xff})
	draw.Draw(dumper.I, dumper.I.Bounds(), bg, image.Point{}, draw.Src)

	return &dumper
}
func (d *Dumper) Close() {
	png.Encode(d.imgFile, d.I)
	d.imgFile.Close()

	d.S.End()
	d.svgFile.Close()

}

func (d *Dumper) Plot(c chart.Chart) {
	row, col := d.Cnt/d.N, d.Cnt%d.N
	igr := imgg.AddTo(d.I, col*d.W, row*d.H, d.W, d.H, color.RGBA{0xff, 0xff, 0xff, 0xff}, nil, nil)
	c.Plot(igr)
	sgr := svgg.AddTo(d.S, col*d.W, row*d.H, d.W, d.H, "", 12, color.RGBA{0xff, 0xff, 0xff, 0xff})
	c.Plot(sgr)
	d.Cnt++
}

//
// Scatter plots with different tic/grid settings
//
func scatterTics() {
	dumper := NewDumper("z", 1, 1, 1200, 900)
	defer dumper.Close()

	c := chart.ScatterChart{Title: "123123"}
	c.XRange.Fixed(0, 10, math.Pi)
	c.YRange.Fixed(-1.25, 1.25, 0.5)
	c.XRange.TicSetting.Format = func(f float64) string {
		w := int(180*f/math.Pi + 0.5)
		return fmt.Sprintf("%d°", w)
	}
	c.AddFunc("Sin(x)", func(x float64) float64 { return math.Sin(x) }, chart.PlotStyleLines,
		chart.Style{Symbol: '@', LineWidth: 1, LineColor: color.NRGBA{0x00, 0x00, 0xcc, 0xff}, LineStyle: 0})
	c.AddFunc("Cos(x)", func(x float64) float64 { return math.Cos(x) }, chart.PlotStyleLines,
		chart.Style{Symbol: '%', LineWidth: 1, LineColor: color.NRGBA{0x00, 0xcc, 0x00, 0xff}, LineStyle: 0})
	dumper.Plot(&c)

}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func PlotYList(title string, yList [][]float64, pointLimit ...int) {
	dumper := NewDumper("z", 1, 1, 1200, 900)
	defer dumper.Close()
	c := chart.ScatterChart{Title: title}
	if len(pointLimit) == 0 {
		pointLimit = append(pointLimit, 10001)
	}
	for n, y := range yList {
		x := make([]float64, Min(len(y), pointLimit[0]))
		for xn := range x {
			x[xn] = float64(xn + 1)
		}
		c.AddDataPair(
			fmt.Sprintf("Trace {%v}", n+1),
			x, y,
			chart.PlotStyleLines,
			chart.Style{Symbol: '@', LineWidth: 1, LineColor: color.NRGBA{0x00, 0x00, 0xcc, 0xff},Font: chart.Font{
				Size:  120,
			}},
		)
	}
	//c.AddFunc("Sin(x)", yList[0], chart.PlotStyleLines,
	//	chart.Style{Symbol: '@', LineWidth: 1, LineColor: color.NRGBA{0x00, 0x00, 0xcc, 0xff}})

	dumper.Plot(&c)

}

func main() {
	//scatterTics()
	PlotYList("test", [][]float64{{1, 2, 3, 5, 6, 7, 3}} )
}
