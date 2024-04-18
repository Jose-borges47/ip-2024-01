package main
import "fmt"

func main() {
    var (
        f, p float64
        x, mm float64
        )
  fmt.Println("insira F e P")
  fmt.Scan(&f, &p)
  x = (5 * (f - 32) / 9)
  mm = (p * 25.4)
  
  fmt.Print("O valor em celsius = ", x, "\nA quantidade de chuva = ", mm)
}