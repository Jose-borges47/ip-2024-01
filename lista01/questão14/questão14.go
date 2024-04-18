package main
import (
    "fmt"
    "math"
    )

func main() {
    var v, ab, h, a float64
  fmt.Println("Insira a altura da pirâmide e a aresta do hexágono")
  fmt.Scan(&h, &a)
  ab = ((3 * (a*a) * math.Sqrt(3)) / 2)
  v = (ab * h) / 3
  fmt.Printf("O VOLUME DA PIRÂMIDE É %.2f METROS CÚBICOS", v)
          }