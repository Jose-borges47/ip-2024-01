package main
import "fmt"

func main() {
    var r, n, a, pa, soma int
  fmt.Println("Digite o valor inicial da PA, a razão e o número de elementos")
  fmt.Scan(&a, &r, &n)
  
  pa = a
  soma = 0
  
  for i := 0; i < n; i++ {
      pa = a + r*i
      soma = pa + soma
  }
  fmt.Print(soma)
}