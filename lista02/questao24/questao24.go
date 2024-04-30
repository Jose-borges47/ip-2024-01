//questão 24
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, r, c, n float64
	fmt.Println("Digite o ângulo em radiano e a quantidade de termos da série")
	fmt.Scan(&x, &n)

	r = 1
	for c = 1; c < n; c++ {
		den := 1.0
		for j := 2.0; j <= 2*c; j++ {
			den *= j
		}
		temp := math.Pow(x, 2*c) / den

		if int(c)%2 == 0 {
			r += temp
		} else {
			r -= temp
		}
	}
	fmt.Printf("cos (%.2f) = %.6f", x, r)
}
