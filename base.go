package main

import (
	"math"
	"math/rand"

	"github.com/rgarrot/lfdoverfitting/legendre"
)

//Base gerada pela soma de funções de legendre. X inputs, Y outputs, A coefs.
type Base struct {
	A []float64 //constantes a's normalizadas
	X []float64 //vetor de entrada
	Y []float64 //saida
}

//Gera uma base com n instancias baseado na função alvo gerada pelo somatorio de polinômios de legendre + ruido
//y_n = f(x_n) + sigma * e_n
//f(x) = sum_{q=0}^{qf} ( a_q * Legendre_q(x) )
func geraBase(qf int, n int, sigma float64) Base {
	var b = Base{}
	b.A = make([]float64, qf+1)
	b.X = make([]float64, n)
	b.Y = make([]float64, n)

	//calcula fator de normalização
	c := 0.0
	for i := 0; i <= qf; i++ {
		c += 1.0 / (2.0*float64(i) + 1.0)
	}
	c = math.Sqrt(c)

	//gera coeficientes
	for j := 0; j <= qf; j++ {
		b.A[j] = r(true) / c
	}

	//gera coeficientes e vetores de entrada e saída
	for i := 0; i < n; i++ {
		b.X[i] = r(false)
		f := 0.0
		for j := 0; j <= qf; j++ {
			f += b.A[j] * legendre.Legendre(j, b.X[i])
		}

		b.Y[i] = f //+ sigma*r(true)
	}

	return b
}

//Gera um número randômico
//norm = false ==> distribuição uniforme [-1;1]
//norm = true ==> distribuição normal padrão
func r(norm bool) float64 {
	if norm {
		return rand.NormFloat64()
	}
	return -1.0 + 2.0*rand.Float64()
}
