package num

import (
	"fmt"
	"math/rand"
)

type Matrix struct {
	Shape [2]int
	Data  [][]float64
}

func Abs(x float64) float64 {
	if x > 0 {
		return x
	}
	return -x
}

func AlmostEqual(a, b float64) bool {
	return Abs(a-b) < 0.00001
}

func (m Matrix) rowScalar(multiple float64, row int) {
	for col := 0; col < m.Shape[1]; col++ {
		m.Data[row][col] *= multiple
	}
}

func (m Matrix) linearCombination(multiple float64, j, k int) {
	for col := 0; col < m.Shape[1]; col++ {
		m.Data[k][col] += multiple * m.Data[j][col]
	}
}

func (m Matrix) swap(j, k int) {
	row_j := m.Data[j]
	m.Data[j] = m.Data[k]
	m.Data[k] = row_j
}

func (m Matrix) ChoosePivot(col int) int {
	pivotVal := 0.0
	pivotRow := 0
	for row := 0; row < m.Shape[0]; row++ {
		if Abs(m.Data[row][col]) > pivotVal {
			pivotRow = row
			pivotVal = m.Data[row][col]
		}
	}
	if pivotVal == 0.0 {
		panic("Matrix is singular")
	}
	return pivotRow
}

func (m Matrix) Copy() Matrix {
	Data := make([][]float64, m.Shape[0])
	for row := 0; row < m.Shape[0]; row++ {
		Data[row] = make([]float64, m.Shape[1])
		for col := 0; col < m.Shape[1]; col++ {
			Data[row][col] = m.Data[row][col]
		}
	}

	return Matrix{[2]int{m.Shape[0], m.Shape[1]}, Data}
}

func (m Matrix) Inverse() Matrix {
	c := m.Copy()
	inv := Eye(m.Shape[0])
	for col := 0; col < m.Shape[1]; col++ {
		pivot := c.ChoosePivot(col)
		c.swap(pivot, col)
		inv.swap(pivot, col)

		inv.rowScalar(1.0/c.Data[col][col], col)
		c.rowScalar(1.0/c.Data[col][col], col)
		for row := 0; row < m.Shape[0]; row++ {
			if row != col {
				inv.linearCombination(-c.Data[row][col], col, row)
				c.linearCombination(-c.Data[row][col], col, row)
			}
		}
	}
	return inv
}

func (m Matrix) String() string {
	str := ""
	for r := 0; r < m.Shape[0]; r++ {
		if r < 5 || m.Shape[0] < 10 {
			str += "| "
			for c := 0; c < m.Shape[1]; c++ {
				str += fmt.Sprintf("%0.2f ", m.Data[r][c])
			}
			str += fmt.Sprintf("|\n")
		} else if r == 5 {
			str += fmt.Sprintf("  ...  \n")
		} else if m.Shape[0]-r < 3 {
			str += fmt.Sprintf("| ")
			for c := 0; c < m.Shape[1]; c++ {
				str += fmt.Sprintf("%0.2f ", m.Data[r][c])
			}
			str += fmt.Sprintf("|\n")
		}
	}
	return str
}

func (m Matrix) Transpose() Matrix {
	result := Zeros(m.Shape[1], m.Shape[0])
	for r := 0; r < m.Shape[0]; r++ {
		for c := 0; c < m.Shape[1]; c++ {
			result.Data[c][r] = m.Data[r][c]
		}
	}
	return result
}

func (m Matrix) Equals(om Matrix) bool {
	if m.Shape != om.Shape {
		return false
	}

	for r := 0; r < m.Shape[0]; r++ {
		for c := 0; c < m.Shape[1]; c++ {
			if !AlmostEqual(m.Data[r][c], om.Data[r][c]) {
				return false
			}
		}
	}
	return true
}

func (m Matrix) Times(om Matrix) Matrix {
	if m.Shape[1] != om.Shape[0] {
		panic("Incompatible dimensions")
	}
	result := Zeros(m.Shape[0], om.Shape[1])
	for rowIdx := 0; rowIdx < m.Shape[0]; rowIdx++ {
		for colIdx := 0; colIdx < om.Shape[1]; colIdx++ {
			for r := 0; r < om.Shape[0]; r++ {
				result.Data[rowIdx][colIdx] += m.Data[rowIdx][r] * om.Data[r][colIdx]
			}
		}
	}
	return result
}

func (m Matrix) PlusScalar(c float64) Matrix {
	result := Zeros(m.Shape[0], m.Shape[1])
	for r := 0; r < m.Shape[0]; r++ {
		for col := 0; col < m.Shape[1]; col++ {
			result.Data[r][col] = c + m.Data[r][col]
		}
	}
	return result
}

func (m Matrix) TimesScalar(c float64) Matrix {
	result := Zeros(m.Shape[0], m.Shape[1])
	for r := 0; r < m.Shape[0]; r++ {
		for col := 0; col < m.Shape[1]; col++ {
			result.Data[r][col] = c * m.Data[r][col]
		}
	}
	return result
}

func (m Matrix) Plus(om Matrix) Matrix {
	if m.Shape != om.Shape {
		panic("Incompatible dimensions")
	}
	result := Zeros(m.Shape[0], m.Shape[1])

	for r := 0; r < m.Shape[0]; r++ {
		for c := 0; c < m.Shape[1]; c++ {
			result.Data[r][c] = m.Data[r][c] + om.Data[r][c]
		}
	}
	return result
}

func (m Matrix) Minus(om Matrix) Matrix {
	if m.Shape != om.Shape {
		panic("Incompatible dimensions")
	}
	return m.Plus(om.TimesScalar(-1.0))
}

func Zeros(m, n int) Matrix {
	result := Matrix{
		[2]int{m, n},
		make([][]float64, m),
	}
	for row := 0; row < m; row++ {
		result.Data[row] = make([]float64, n)
	}
	return result
}

func Ones(m, n int) Matrix {
	result := Zeros(m, n)

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			result.Data[row][col] = 1.0
		}
	}
	return result
}

func Eye(m int) Matrix {
	result := Zeros(m, m)
	for row := 0; row < m; row++ {
		result.Data[row][row] = 1.0
	}
	return result
}

func Randn(m, n int) Matrix {
	result := Zeros(m, n)

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			result.Data[r][c] = rand.NormFloat64()
		}
	}
	return result
}

func Rand(m, n int) Matrix {
	result := Zeros(m, n)

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			result.Data[r][c] = rand.Float64()
		}
	}
	return result
}
