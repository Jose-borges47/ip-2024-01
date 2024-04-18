package main
import "fmt"

func main() {
    var n, nq int
  fmt.Println("Digite um número par e a quantidade a serem gerados")
  fmt.Scan(&n)
  
  if n%2 != 0 {
      fmt.Println("O primeiro numekr não é par")
  } else {
      fmt.Scan(&nq)
      
      for i := 0; i < nq; i++ {
          fmt.Printf("%v ", n)
          n = n + 2
      }
}
}