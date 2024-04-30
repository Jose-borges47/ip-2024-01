//questão 9
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Println("Digite um número para saber se ele é primo")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Número inválido")
		return
	}

	if n <= 1 {
		fmt.Println("Não primo")
		return
	}

	if n == 2 || n == 3 {
		fmt.Println("primo")
		return
	}

	if n%2 == 0 || n%3 == 0 {
		fmt.Println("não primo")
		return
	}

	for i := 5; i <= int(math.Sqrt(float64(n))); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			fmt.Println("NÃO PRIMO")
			return
		}
	}

	fmt.Println("PRIMO")
}
