// questão 11
package main

import (
	"fmt"
)

func main() {
	var n, e, r float64
	fmt.Println("Digite um número e a precisão da raiz quadrada")
	fmt.Scan(&n, &e)
	r = 1
	a := n - (r * r)
	for a > e {
		r = ((n / r) + r) / 2
		a = n - (r * r)
		if a > 0 {
			a = (n - (r*r)*-1)
		}
		fmt.Printf("r: %.9f erro: %.9f", r, a)
	}
}
