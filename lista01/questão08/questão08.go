package main
import "fmt"

func main() {
    var (
        r, h, pi, ab, al, at, custo float64
        )
  fmt.Println("insira o raio e a altura da lata")
  fmt.Scan(&r, &h)
  pi = (3.14159)
  ab = (pi * (r * r))
  al = (2 * pi * r * h)
  at = (2 * ab + al)
  custo = (at * 100)
  fmt.Print("O custo da lata = ", custo)
  
}