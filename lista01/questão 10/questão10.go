package main
import "fmt"

func main() {
    var a, b, c, d, ma int
  fmt.Println("digite o valor de a, b, c e d")
  fmt.Scan(&a, &b, &c, &d)
  ma = (a * d - b * c)
  fmt.Print("O VALOR DO DETERMINANTE É ", ma)
}