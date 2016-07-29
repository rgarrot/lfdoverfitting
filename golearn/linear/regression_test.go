package linear

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	x, y, alpha := GenerateLinearData(0, 1, 100)
	model := new(LinearModel)
	model.Fit(x, y)
	if !alpha.Equals(model.Coefs) {
		fmt.Printf("Coefficients: %v\n", model.Coefs)
		t.Fail()
	}
}
