package main

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"

	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"
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
func histChart(title string, points []float64, stacked, counts, shifted bool, img *Img) {

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
	img.Plot(&hc)

}

//
// Full fletched scatter plots
//
func scatterChart(xval, yval []float64, img *Img) {

	pl := chart.ScatterChart{Title: "Scatter + Lines"}
	pl.XRange.Label, pl.YRange.Label = "X - Value", "Y - Value"
	pl.Key.Pos = "itl"
	// pl.XRange.TicSetting.Delta = 5
	pl.XRange.TicSetting.Grid = 1

	pl.AddDataPair("Distribution", xval, yval, chart.PlotStyleLines,
		chart.Style{Symbol: '%', LineWidth: 2, LineColor: color.NRGBA{0xa0, 0x00, 0x00, 0xff}, LineStyle: chart.SolidLine})

	pl.XRange.ShowZero = true
	pl.XRange.TicSetting.Mirror = 1
	pl.YRange.TicSetting.Mirror = 1
	pl.XRange.TicSetting.Grid = 1
	pl.XRange.Label = "X - Interations"
	pl.YRange.Label = "Y - Values"
	pl.Key.Cols = 2
	pl.Key.Pos = "orb"

	img.Plot(&pl)
}

// genFloats return range of float64
func genFloats(n int) []float64 {
	r := make([]float64, n, n)
	for i := 0; i < n; i++ {
		r[i] = float64(i)
	}
	return r
}
