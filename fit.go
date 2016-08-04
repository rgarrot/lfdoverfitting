package main

import (
	"math"

	"github.com/gonum/matrix/mat64"
)

func polyfit(b Base, n int) []float64 {
	x := xPolyMatrix(b, n)
	y := mat64.NewDense(len(b.Y), 1, b.Y)
	var result mat64.Dense
	result.Solve(x, y)
	return result.RawMatrix().Data
}

func xPolyMatrix(b Base, n int) mat64.Matrix {
	m := len(b.X)
	x := make([]float64, (n+1)*m)
	for r := 0; r < m; r++ {
		for c := 0; c < (n + 1); c++ {
			x[r*(n+1)+c] = math.Pow(b.X[r], float64(c))
		}
	}
	return mat64.NewDense(m, (n + 1), x)
}
