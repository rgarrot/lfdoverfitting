package main

func esp(f []float64, g []float64) float64 {
	return intMinus1To1Poly(mulPoly(f, g))
}

//Multiplica f(x) * g(x)
//f e g são os indices dos polinômios. Ex.: f[0]x^0 + f[1]x^1 + ... + f[n]x^n.
func mulPoly(f []float64, g []float64) []float64 {
	fg := make([]float64, (len(f)-1)*(len(g)-1)+1)
	for i := 0; i < len(f); i++ {
		for j := 0; j < len(g); j++ {
			fg[i+j] += f[i] * g[j]
		}
	}
	return fg
}

// Integral{-1^1}( f(x) )
// f é o vetor de indices de um polinômio. Ex.: f[0]x^0 + f[1]x^1 + ... + f[n]x^n.
func intMinus1To1Poly(f []float64) float64 {
	result := 0.0
	for i := 0; i < len(f); i += 2 {
		result += 2 * f[i] / float64(i+1)
	}
	return result
}
