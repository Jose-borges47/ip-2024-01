package main

import (
    "fmt"
    "strconv"
)
func main() {
    
    var n1, n2, n3 int
    
    fmt.Println("DIGITE TRÊS NÚMEROS")
    fmt.Scan(&n1, &n2, &n3)
    if n1 > 0 && n1 < 10 && n2 < 10 && n3 < 10 {
        s := fmt.Sprintf("%v%v%v", n1, n2, n3) {
          n, _ := strconv.Atoi(s)
          fmt.Println(n, n * n)
          } else {
              fmt.Println("DÍGITO INVALIDO")
              return
          }      
}