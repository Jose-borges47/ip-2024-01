// 3
package main

import "fmt"

func main() {
    var n int
    for {
    fmt.Scan(&n)
    if n < 5000 {
      break 
    } else { 
      fmt.Println("Apenas menor que 5000")
    } 
    }
    if n < 0{return}
    num := make([]int, n)
    for i := 0; i < n; i++ {
      fmt.Scan(&num[i])
    }
    for a := n - 1; a >= 0; a-- {
    fmt.Printf("%v\n", num[a])
    }
}