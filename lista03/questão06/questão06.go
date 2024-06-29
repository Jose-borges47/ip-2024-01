//6
package main 

import "fmt"

func main() {
  var num, soma int
  fmt.Scan(&num)
  n := make([]int, num)
  
  for i := range n {
    fmt.Scan(&n[i])
  }
  for _, x := range n {
    soma += x 
  }
  fmt.Println(soma)
}