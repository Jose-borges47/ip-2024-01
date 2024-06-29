// 8
package main

import "fmt"

func main() { 
  var numero int
  resultado := []int{}
  for {
  _, erro := fmt.Scanln(&numero)
  if erro != nil { 
    break
}
resultado = append(resultado, numero)
}
for i := range resultado {
  fmt.Printf("%b\n", resultado[i])
}
}