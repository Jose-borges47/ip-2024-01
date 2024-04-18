package main
import "fmt"

func main() {
    var n int
  fmt.Println("DIGITE UM NÚMERO")
  fmt.Scan(&n)
  
  if (n%3 == 0 && n%5 == 0) {
  fmt.Printf("O NÚMERO É DIVISÍVEL")
  
  } else {
    fmt.Printf("O NÚMERO NÃO É DIVISÍVEL")
  }
}