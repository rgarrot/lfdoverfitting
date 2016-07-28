package legendre

//Calculate legendre polynomial of degree k
func Legendre(k int, x float64) float64 {
  switch k {
    case 0:
      return float64(1)
    case 1:
      return x
    default:
      return legendre_recursive(k, x, 1, x, 0, 1)
  }
}

func legendre_recursive(k int, x float64, kMenos1 int, lkMenos1 float64, kMenos2 int, lkMenos2 float64) float64 {
    var kAtual float64 = float64(kMenos1+1);
    var lkAtual float64 = (((2 * kAtual - 1) / kAtual ) * x * lkMenos1) - (((kAtual - 1)/kAtual) * lkMenos2)
    if kMenos1 == (k-1) {
      return lkAtual
    } else {
      return legendre_recursive(k, x, (kMenos1+1), lkAtual, kMenos1, lkMenos1);
    }
    return 0
}
