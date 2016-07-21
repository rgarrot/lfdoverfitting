package main

import ("fmt"
        "time"
        "math/rand"
        //"github.com/montanaflynn/stats"
)

func init(){
    rand.Seed(int64(time.Now().Nanosecond()))
}

func main(){
  var inicio = time.Now()

  var segundoMomento = 0.0
  var primeiroMomento = 0.0

  for j:=0; j<100; j++ {
    var b = geraBase(1, 10000, 0)
    for i:=0; i<len(b.Y); i++ {
      primeiroMomento += b.Y[i]
      segundoMomento += b.Y[i] * b.Y[i]
    }
  }
  primeiroMomento /= float64(10000 * 100)
  segundoMomento /= float64(10000 * 100)

  fmt.Printf("Primeiro Momento = %f, Segundo Momento = %f \n", primeiroMomento, segundoMomento)
  fmt.Printf("tempo total:  %s", time.Since(inicio))
}
