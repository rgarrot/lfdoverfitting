package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
	criaMatrizLegendre(100)
}

func main() {
	var inicio = time.Now()
	fmt.Printf("\nv48\n")

	var b = geraBase(2, 20, 0.0)
	g2 := polyfit(b, 2)
	g10 := polyfit(b, 10)

	writeBase(b)

	fmt.Printf("f: %v \n\n", b.F)
	fmt.Printf("g2: %v \n\n", g2)
	fmt.Printf("g10: %v", g10)
	// plotBase(b, yP2, yP10)

	fmt.Printf("tempo total:  %s", time.Since(inicio))

}

func eout(f []float64, g []float64) float64 {
	return esp(g, g) - 2*esp(g, f) + esp(f, f)
}
