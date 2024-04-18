package main
import (
    "fmt"
    "math"
)

func main() {
    var n, r, numq float64
  fmt.Println("Digite o número")
  fmt.Scan(&n)
  
  numq = 2
  if n <= 5 || n >= 2000 {
      fmt.Println("Número inválido")
      return
  } else {
      for i := 2.0; i <= n; i = i + 2 {
          r = math.Pow(i, numq)
          fmt.Printf("%v ^ %v = %v\n", i, numq, r)
      }
  
}
}