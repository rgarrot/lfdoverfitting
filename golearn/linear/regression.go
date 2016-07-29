package linear

import (
	"math"

	"github.com/rgarrot/lfdoverfitting/golearn/num"
)

func GenerateLinearData(sd float64, dim, obs int) (num.Matrix, num.Matrix, num.Matrix) {
	x, alpha := num.Rand(obs, dim), num.Randn(dim, 1)
	y := x.Times(alpha).Plus(num.Randn(obs, 1).TimesScalar(sd))
	return x, y, alpha
}

type LinearModel struct {
	Coefs num.Matrix
	Error float64
}

func (m *LinearModel) Fit(x, y num.Matrix) {
	_ = m.FitTransform(x, y)
}

func (m LinearModel) Transform(x num.Matrix) num.Matrix {
	return x.Times(m.Coefs)
}

func (m *LinearModel) FitTransform(x, y num.Matrix) num.Matrix {
	xT := x.Transpose()
	m.Coefs = (xT.Times(x)).Inverse().Times(xT).Times(y)
	predicted := x.Times(m.Coefs)
	m.Error = 0.0
	for j := 0; j < predicted.Shape[0]; j++ {
		m.Error += math.Pow(predicted.Data[j][0]-y.Data[j][0], 2)
	}
	m.Error /= float64(predicted.Shape[0])

	return predicted
}
