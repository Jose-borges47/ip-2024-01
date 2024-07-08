package main 

import "fmt"

func main() {
  k := 3
  n := 2
  fmt.Println(elevar(k, n))
}

func elevar(k, n int) int {
  if n == 0 {
    return 1
  }
  return k * elevar(k, n-1)
}