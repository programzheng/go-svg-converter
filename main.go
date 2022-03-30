package main

import (
	"image"
	"image/png"
	"os"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

func main() {
	in, err := os.Open("in.svg")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	icon, err := oksvg.ReadIconStream(in)
	if err != nil {
		panic(err)
	}
	w := int(icon.ViewBox.W)
	h := int(icon.ViewBox.H)

	icon.SetTarget(0, 0, float64(w), float64(h))
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())), 1)

	out, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = png.Encode(out, rgba)
	if err != nil {
		panic(err)
	}
}
