package main

import "fmt"

func main() {
  var n, m, maior int 
  for { 
    fmt.Scan(&n)
    if n >= 1 && n <= 1000 {
      break
    } else { 
      fmt.Println("O número precisa ser entre 1 e 1000")
    }
  }
  num := make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&num[i])
  }
  fmt.Scan(&m)
  for a := 0; a < n; a++ { 
    if num[a] >= m {
      maior++ 
    }
  }
  fmt.Println(maior)
}