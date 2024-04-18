package main
import "fmt"

func main() {
    var soma float64
	var num float64
  fmt.Println("Digite o número a ser calculado o somatório")
  fmt.Scan(&num)
  
  if num < 1.0 {
      fmt.Println("NÚMERO INVÁLIDO")
  } else {
      for n := 1.0; n <= num; n++ {
          soma = soma + (1 / n)
      }
    }
      
      fmt.Printf("%.6f\n", soma)
}