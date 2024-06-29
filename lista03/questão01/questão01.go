package main 

import (
  "fmt"
)
func main() {
  var q, t, n, f int 
  fmt.Scan(&q)
  var r []int
  for i := 0; i < q; i++ {
    fmt.Scan(&t) 
    r = append(r, t)
  }
  fmt.Scan(&n)
  for a := 0; a < n; a++ {
    f = 0
    fmt.Scan(&t)
    for x := 0; x < q; x++ {
      if t == r[x] {
        f = 1 
        fmt.Println("ACHEI")
        break
      }
    }
    if f == 0 {
      fmt.Println("NÃO ACHEI")
    }
  }
}
