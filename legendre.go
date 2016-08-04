package main

import (
	"math"
	"math/big"
)

//MatrizLegendre com os coeficientes das funções de legendre
var MatrizLegendre [][]*big.Float

const prec = 200

//CriaMatrizLegendre calcula coeficientes da MatrizLegendre
func criaMatrizLegendre(n int) {
	n++
	MatrizLegendre = make([][]*big.Float, n)
	for i := 0; i < n; i++ {
		MatrizLegendre[i] = make([]*big.Float, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			MatrizLegendre[i][j] = new(big.Float).SetPrec(prec).SetFloat64(0.0)
		}
	}

	MatrizLegendre[0][0] = new(big.Float).SetPrec(prec).SetFloat64(1.0)
	MatrizLegendre[1][1] = new(big.Float).SetPrec(prec).SetFloat64(1.0)

	for k := 2; k < n; k++ {
		for i := 0; i < k; i++ {
			a := new(big.Float).SetPrec(prec).Set(MatrizLegendre[k-1][i])
			b := new(big.Float).SetPrec(prec).SetFloat64((2.0*float64(k) - 1.0) / float64(k))
			a.Mul(a, b)
			MatrizLegendre[k][i+1].Add(MatrizLegendre[k][i+1], a)

			c := new(big.Float).SetPrec(prec).Set(MatrizLegendre[k-2][i])
			d := ((float64(k) - 1.0) / float64(k))
			c.Mul(c, new(big.Float).SetFloat64(d))
			MatrizLegendre[k][i].Sub(MatrizLegendre[k][i], c)

			//MatrizLegendre[k][i+1] += MatrizLegendre[k-1][i] * ((new(big.Float).SetFloat64(float64(2*k)) - new(big.Float).SetFloat64(1.0)) / new(big.Float).SetFloat64(float64(k)))
			//MatrizLegendre[k][i] -= MatrizLegendre[k-2][i] * ((float64(k) - 1.0) / float64(k))
		}
	}
}

// Legendre polinomio de legendre grau k no ponto x
// Apenas se a MatrizLegendre estiver inicializada até este valor.
func legendre(k int, x float64) float64 {
	//if len(MatrizLegendre) < k {
	//	return 0.0
	//	}
	result := new(big.Float).SetPrec(prec).SetFloat64(0.0)
	for i := 0; i <= k; i++ {
		a := new(big.Float).SetPrec(prec).SetFloat64(math.Pow(x, float64(i)))
		a.Mul(a, MatrizLegendre[k][i])
		result.Add(result, a)
	}
	resultFloat, _ := result.Float64()
	return resultFloat
}
