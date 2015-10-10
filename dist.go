package main

import (
	"time"

	"github.com/leesper/go_rng"
)

func genDist(name string, n int, s, a float64) []float64 {
	points := make([]float64, n, n)
	switch name {
	case "weibull":
		g := rng.NewWeibullGenerator(time.Now().UnixNano())
		for i := 0; i < len(points); i++ {
			x := g.Weibull(1, 0.8)*s + a
			points[i] = x
		}
	}

	return points
}
