package main

import ("fmt"
        "time"
        "math/rand"
        "github.com/rgarrot/lfdoverfitting/legendre"
)

type Base struct {
  A []float64; //constantes a's normalizadas
  X []float64; //vetor de entrada
  Y []float64; //saida
}

func init(){
    rand.Seed(int64(time.Now().Nanosecond()))
}

func main(){

  var b = geraBaseLegendre(50, 100, 0.1)

  var inicio = time.Now()

  fmt.Println()
  fmt.Println(b.X)
  fmt.Println(b.Y)
  fmt.Println()
  fmt.Println(b.A)

  fmt.Printf("tempo total:  %s", time.Since(inicio))
}

/* Gera uma base com n instancias baseado na função alvo gerada pelo somatorio de polinômios de legendre + ruido
   y_n = f(x_n) + sigma * e_n
   f(x) = sum_{q=0}^{qf} ( a_q * Legendre_q(x) )
*/
func geraBaseLegendre (qf int, n int, sigma float64) Base {
    var b = Base{}
    var f float64;
    b.A = make([]float64, qf);
    b.X = make([]float64, n);
    b.Y = make([]float64, n);

    // gera vetor de entrada X
    for i:=0; i<n; i++ {
      b.X[i] = r(false);
    }

    //gera vetor de pesos
    for i:=0; i<qf; i++ {
      b.A[i] = r(true);
    }

    //gera vetor de saida
    for i:=0; i<n; i++ {
      f = 0;
      for j:=0; j<qf; j++{
        f += b.A[j] * legendre.Legendre(j, b.X[i]);
      }
      b.Y[i] = f + sigma * r(true);
    }

    return b;
}

/* Gera um número randômico
   norm = false ==> distribuição uniforme [-1;1]
   norm = true ==> distribuição normal padrão
*/

func r(norm bool) float64 {
  if norm {
    return rand.NormFloat64();
  } else {
    return -1.0 + 2.0 * rand.Float64();
  }
}
