package main
import "fmt"

func main() {
    var n1, n2, n3, n4 int
        
  fmt.Println("Digite as horas, os minutos e os segundos")
  fmt.Scan(&n1, &n2, &n3)
  n4 = ((n1 * 60 *60) + (n2 * 60) + n3)
  
  fmt.Print("O TEMPO EM SEGUNDOS É ", n4)
}