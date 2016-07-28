package main

import (
	"fmt"
	"golearn/linear"
	"golearn/num"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func main() {
	var inicio = time.Now()
	fmt.Printf("\nv14\n")

	var b = geraBase(5, 120, 0.0)

	xg2 := g2(b)
	xg10 := g10(b)
	y := y(b)
	model := linear.LinearModel{}
	model.Fit(xg2, y)
	//fmt.Printf("G2 Mean Squared Error: %.7f \n", model.Error)
	yP2 := xg2.Times(model.Coefs)
	model.Fit(xg10, y)
	//fmt.Printf("G10 Mean Squared Error: %.7f \n", model.Error)
	yP10 := xg10.Times(model.Coefs)

	plotBase(b, yP2, yP10)

	//fmt.Printf("%f", b.Y)
	//fmt.Printf("X %d, Y %d, A %d \n", len(b.X), len(b.Y), len(b.A))
	/*
		xg2 := g2(b)
		xg10 := g10(b)
		y := y(b)
		model := linear.LinearModel{}
		model.Fit(xg2, y)
		//fmt.Printf("Computed a=\n%v\n", model.Coefs)
		fmt.Printf("G2 Mean Squared Error: %.7f \n", model.Error)
		model.Fit(xg10, y)
		//fmt.Printf("Computed a=\n%v\n", model.Coefs)
		fmt.Printf("G10 Mean Squared Error: %.7f \n", model.Error)
	*/
	fmt.Printf("tempo total:  %s", time.Since(inicio))

}

func g2(b Base) num.Matrix {
	m := len(b.X)
	result := num.Zeros(m, 3)

	for row := 0; row < m; row++ {
		result.Data[row][0] = 1
		result.Data[row][1] = b.X[row]
		result.Data[row][2] = math.Pow(b.X[row], 2)
	}
	return result
}

func g10(b Base) num.Matrix {
	m := len(b.X)
	result := num.Zeros(m, 11)

	for row := 0; row < m; row++ {
		result.Data[row][0] = 1
		result.Data[row][1] = b.X[row]
		result.Data[row][2] = math.Pow(b.X[row], 2)
		result.Data[row][3] = math.Pow(b.X[row], 3)
		result.Data[row][4] = math.Pow(b.X[row], 4)
		result.Data[row][5] = math.Pow(b.X[row], 5)
		result.Data[row][6] = math.Pow(b.X[row], 6)
		result.Data[row][7] = math.Pow(b.X[row], 7)
		result.Data[row][8] = math.Pow(b.X[row], 8)
		result.Data[row][9] = math.Pow(b.X[row], 9)
		result.Data[row][10] = math.Pow(b.X[row], 10)
	}
	return result
}

func y(b Base) num.Matrix {
	m := len(b.X)
	result := num.Zeros(m, 1)
	for row := 0; row < m; row++ {
		result.Data[row][0] = b.Y[row]
	}
	return result
}
