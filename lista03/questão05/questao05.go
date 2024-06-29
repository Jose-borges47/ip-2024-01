//5
package main 

import "fmt" 

func main() {
  var n, num, so int 
  for {
    rapaz := 0 
    fmt.Scan(&n) 
    if n == 0 {
      break
    } 
    for i := 0; i < n; i++ {
      fmt.Scan(&num)
      if num > rapaz {
        rapaz = num
        so = i 
      }
    }
    fmt.Printf("%v %v \n", so, rapaz)
  }
}