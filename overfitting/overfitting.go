package main

import ("fmt"
        "time"
        "github.com/rgarrot/yaser/legendre"
)

func main(){

  var inicio = time.Now()
  var inicioCada = time.Now()
  var x float64 = 0.5
  for i:=0;i<100;i++ {
    inicioCada = time.Now()
    fmt.Printf("%f  ,", legendre.Legendre(i,x))
    fmt.Printf("i = %d, tempo = %s \n", i, time.Since(inicioCada))
  }

  fmt.Printf("%s %s","tempo total: ", time.Since(inicio))
}
