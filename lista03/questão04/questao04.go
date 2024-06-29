//4 package main

import "fmt"

func main() {
  var x int
  fmt.Scan(&x)
  for {
  if x < 5000 {
    break
  } else {
    fmt.Println("deve ser menor que 5000")
  }
  }
  if x < 0{return}
  
  num := make([]int, x)
  for i := range num {
    fmt.Scan(&num[i]) 
  }
  for i := range num {
    fmt.Printf("%v ", num[i])
  }
}