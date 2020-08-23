package main

import (
	"github.com/fogleman/gg"
	"github.com/zut/x/xlog"
	"os"
)

func main() {
	im, err := gg.LoadImage("example/img/1.png")
	if err != nil {
		xlog.Fatal(err)
	}
	const W = 1200
	const H = 900
	dc := gg.NewContext(W, H)
	//dc.SetHexColor("ffffff")
	//dc.Clear()
	dc.DrawImage(im, 0, 0)
	dc.SetHexColor("FF0000")
	dc.SetLineWidth(1)
	dc.DrawLine(0, 0, W, H)
	dc.DrawLine(W, 0, 0, H)
	//dc.DrawLine(400, 300, 800, 600)
	//dc.DrawLine(800, 300, 400, 600)
	dc.Stroke()
	dc.SavePNG("out.png")

	f, err := os.Create("image.png")
	if err != nil {
		xlog.Fatal(err)
	}

	dc.EncodePNG(f)

	if err := f.Close(); err != nil {
		xlog.Fatal(err)
	}

}
