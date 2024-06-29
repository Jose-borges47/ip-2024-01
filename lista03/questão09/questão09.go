// 9
package main

import "fmt"
import "math"

func main() {
  var x int
  var y float64
  var z []float64
  var sla float64
  fmt.Scan(&x)
  for i := 0; i < x; i++ {
    for a := 0; a < 3; a++ {
      fmt.Scan(&y)
      z = append(z, y)
    }
  }
  for i := 0; i < x * 3 - 3; i += 3 { 
    sla = (math.Sqrt(float64((z[i]) - z[i + 3])*(z[i] - z[i + 3]) + (z[i + 1] - z[ i + 4]) * (z[i + 1]-
((z[i] - z[i + 4]) + (z[i + 1] - z[i + 4]) * (z[i + 2] - z[i + 5])))))
fmt.Printf("%.2f", sla)
}
}