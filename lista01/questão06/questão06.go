package main

import "fmt"

func main() {
    var (
        ncasos, F float64
        C float64
        )
        
  fmt.Println("temperatura em F")
  fmt.Scanf(&ncasos)
  
  for n := 0; < ncasos; n++ {
      
      fmt.Scanf("%v %f")
      C = 5 * (F - 32) / 9
      fmt.Printf("%v FAHRENHEIT EQUIVALE A %f CELSIUS")
  }
}