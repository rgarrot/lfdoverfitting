package num

import (
	"testing"
)

func TestZeros(t *testing.T) {
	A := Zeros(5, 2)
	if A.shape != [2]int{5, 2} {
		t.Fail()
	}
	for row := 0; row < A.shape[0]; row++ {
		if len(A.data[row]) != 2 {
			t.Fail()
		}
		for col := 0; col < A.shape[1]; col++ {
			if A.data[row][col] != 0.0 {
				t.Fail()
			}
		}
	}
}

func TestOnes(t *testing.T) {
	A := Ones(5, 2)
	if A.shape != [2]int{5, 2} {
		t.Fail()
	}
	for row := 0; row < A.shape[0]; row++ {
		if len(A.data[row]) != 2 {
			t.Fail()
		}
		for col := 0; col < A.shape[1]; col++ {
			if A.data[row][col] != 1.0 {
				t.Fail()
			}
		}
	}
}

func TestTimesScalar(t *testing.T) {
	c := 3.14
	B := Ones(5, 2).TimesScalar(c)
	for row := 0; row < B.shape[0]; row++ {
		for col := 0; col < B.shape[1]; col++ {
			if B.data[row][col] != c {
				t.Fail()
			}
		}
	}
}

func TestPlus(t *testing.T) {
	if !Ones(3, 3).Plus(Ones(3, 3)).Equals(Ones(3, 3).TimesScalar(2.0)) {
		t.Fail()
	}
}

func TestEye(t *testing.T) {
	A := Eye(5)
	if !A.Times(A).Equals(A) {
		t.Fail()
	}
}

func TestRand(t *testing.T) {
	A := Rand(5, 5)
	for r := 0; r < A.shape[0]; r++ {
		for c := 0; c < A.shape[1]; c++ {
			val := A.data[r][c]
			if val < 0 || val > 1 {
				t.Fail()
			}
		}
	}
	if !A.Minus(A).Equals(Zeros(5, 5)) {
		t.Fail()
	}
}

func TestSwap(t *testing.T) {
	A := Eye(3)
	B := Matrix{
		[2]int{3, 3},
		[][]float64{{0.0, 0.0, 1.0}, {0.0, 1.0, 0.0}, {1.0, 0.0, 0.0}}}
	A.swap(0, 2)

	if !A.Equals(B) {
		t.Fail()
	}
}

func TestrowScalar(t *testing.T) {
	A := Eye(3)
	B := Matrix{
		[2]int{3, 3},
		[][]float64{{3.0, 0.0, 0.0}, {0.0, 2.0, 0.0}, {0.0, 0.0, -3.14}}}
	A.rowScalar(3.0, 0)
	A.rowScalar(2.0, 1)
	A.rowScalar(-3.14, 2)
	if !A.Equals(B) {
		t.Fail()
	}
}

func TestInverse(t *testing.T) {
	A := Randn(3, 3)
	Ainv := A.Inverse()
	if !A.Times(Ainv).Equals(Eye(3)) {
		t.Fail()
	}
	if !Ainv.Times(A).Equals(Eye(3)) {
		t.Fail()
	}
}

func TestLinearCombination(t *testing.T) {
	A := Eye(3)
	B := Matrix{
		[2]int{3, 3},
		[][]float64{{1.0, 0.0, 0.0}, {1.0, 1.0, 0.0}, {1.0, 2.0, 1.0}}}
	A.linearCombination(1.0, 0, 2)
	A.linearCombination(2.0, 1, 2)
	A.linearCombination(1.0, 0, 1)
	if !A.Equals(B) {
		t.Fail()
	}
}
