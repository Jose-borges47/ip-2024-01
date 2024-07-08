package main 

import "fmt"

func main() {
  x := []int{1, 2, 3, 4, 5}
  inverter(x, 0, len(x)-1)
  fmt.Println(x)
}
func inverter(v []int, x, y int) {
  if x >= y { 
    return 
}
v[x], v[y] = v[y], v[x] 
inverter(v, x+1, y-1)
}