package main

import (
	"time"

	"github.com/leesper/go_rng"
)

func genDist(name string, n int, s, a float64) []float64 {
	points := make([]float64, n)
	switch name {
	case "weibull":
		g := rng.NewWeibullGenerator(time.Now().UnixNano())
		for i := 0; i < len(points); i++ {
			x := g.Weibull(1, 0.9)
			points[i] = x
		}
	case "bernoulli":
		g := rng.NewBernoulliGenerator(time.Now().UnixNano())
		p := rng.NewUniformGenerator(time.Now().UnixNano())
		for i := 0; i < len(points); i++ {
			x := g.Bernoulli_P(p.Float64())
			if x {
				points[i] = 1
			}
		}
	}

	return points
}
