package main
import "fmt"

func main() {
    var nota float64
  fmt.Println("INSIRA SUA NOTA")
  fmt.Scan(&nota)
  if nota >= 0 && nota <= 5.9 {
      fmt.Print("N0TA =", nota, " CONCEITO = D")
  } else if nota >= 6 && nota <= 7.5 {
      fmt.Print("N0TA =", nota, " CONCEITO = C")
      } else if nota >= 7.6 && nota <= 8.9 {
          fmt.Print("N0TA =", nota, " CONCEITO = B")
             } else if nota >= 9 && nota <=10 {
                 fmt.Print("N0TA =", nota, " CONCEITO = A")
             }
          }