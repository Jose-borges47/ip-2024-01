// 10
package main

import "fmt"

func main() {
  var x, y, m, so int 
  fmt.Println("DIGITE A QUANTIDADE DE NOTAS E EM SEGUINTE AS NOTAS")
  fmt.Scan(&x)
  s := make([]int, 10)
  for i := 0; i < x; i++ {
    fmt.Scan(&y)
    s[y - 1]++
  if y > m {
    m = y
    so = i
  }
}
fmt.Printf("NOTA %d, %d vezes\n NOTA %d, indice %d\n", y, s[y - 1], m, so)
}