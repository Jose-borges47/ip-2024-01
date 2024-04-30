//questão 26
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, s float64
	var q int
	fmt.Println("Digite o ângulo em radianos e a quantidade de termos da série")
	fmt.Scan(&x, &q)

	for i := 0; i <= q; i++ {
		fatorial := 1.0
		for n := 1; n <= 2*i+1; n++ {
			fatorial *= float64(n)
		}
		s += math.Pow(-1, float64(i)) * math.Pow(x, float64(2*i+1)) / fatorial
	}
	fmt.Printf("\nO seno de %.2f radianos é %.6f.\n", x, s)
}
