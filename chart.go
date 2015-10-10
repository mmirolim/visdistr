package main

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"

	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"
)

var (
	data1  = []float64{15e-7, 30e-7, 35e-7, 50e-7, 70e-7, 75e-7, 80e-7, 32e-7, 35e-7, 70e-7, 65e-7}
	data10 = []float64{34567, 35432, 37888, 39991, 40566, 42123, 44678}

	data2 = []float64{10e-7, 11e-7, 12e-7, 22e-7, 25e-7, 33e-7}
	data3 = []float64{50e-7, 55e-7, 55e-7, 60e-7, 50e-7, 65e-7, 60e-7, 65e-7, 55e-7, 50e-7}
)

var Background = color.RGBA{0xff, 0xff, 0xff, 0xff}

// -------------------------------------------------------------------------
// Img

// Img helps saving plots of size WxH in a NxM grid layout
// in several formats
type Img struct {
	N, M, W, H, Cnt int
	I               *image.RGBA
}

func NewImg(n, m, w, h int) *Img {
	dumper := Img{N: n, M: m, W: w, H: h}

	dumper.I = image.NewRGBA(image.Rect(0, 0, n*w, m*h))
	bg := image.NewUniform(color.RGBA{0xff, 0xff, 0xff, 0xff})
	draw.Draw(dumper.I, dumper.I.Bounds(), bg, image.ZP, draw.Src)

	return &dumper
}

// Plot a chart
func (d *Img) Plot(c chart.Chart) {
	row, col := d.Cnt/d.N, d.Cnt%d.N

	igr := imgg.AddTo(d.I, col*d.W, row*d.H, d.W, d.H, color.RGBA{0xff, 0xff, 0xff, 0xff}, nil, nil)
	c.Plot(igr)

	d.Cnt++

}

// gaussian distribution with n samples, stddev of s, offset of a, forced to [l,u]
func gauss(n int, s, a, l, u float64) []float64 {
	points := make([]float64, n)
	for i := 0; i < len(points); i++ {
		x := rand.NormFloat64()*s + a
		if x < l {
			x = l
		} else if x > u {
			x = u
		}
		points[i] = x
	}
	return points
}

//
// Histograms Charts
//
func histChart(name, title string, points []float64, stacked, counts, shifted bool, dumper *Img) {

	hc := chart.HistChart{Title: title, Stacked: stacked, Counts: counts, Shifted: shifted}
	hc.XRange.Label = "Sample Value"
	if counts {
		hc.YRange.Label = "Total Count"
	} else {
		hc.YRange.Label = "Rel. Frequency [%]"
	}
	hc.Key.Hide = true
	hc.AddData("Sample 1", points,
		chart.Style{ /*LineColor: color.NRGBA{0xff,0x00,0x00,0xff}, LineWidth: 1, LineStyle: 1, FillColor: color.NRGBA{0xff,0x80,0x80,0xff}*/ })
	hc.Kernel = chart.BisquareKernel //  chart.GaussKernel // chart.EpanechnikovKernel // chart.RectangularKernel // chart.BisquareKernel
	dumper.Plot(&hc)

}
