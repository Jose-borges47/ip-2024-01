// questão 4

package main

import "fmt"

func main() {
	var n, i, K, s float64

	fmt.Print("Digite um número : ")
	fmt.Scanln(&n)

	fmt.Print("Digite o valor inicial : ")
	fmt.Scanln(&i)

	fmt.Print("Digite a quantidade de valores : ")
	fmt.Scanln(&K)

	fmt.Print("Digite o incremento : ")
	fmt.Scanln(&s)

	fmt.Println("Tabuada de soma:")
	for j := i; j <= i+((K-1)*s); j += s {
		r := n + j
		fmt.Printf("%.2f + %.2f = %.2f\n", n, j, r)
	}

	fmt.Println("Tabuada de subtração:")
	for j := i; j <= i+((K-1)*s); j += s {
		r := n - j
		fmt.Printf("%.2f - %.2f = %.2f\n", n, j, r)
	}

	fmt.Println("Tabuada de multiplicação:")
	for j := i; j <= i+((K-1)*s); j += s {
		r := n * j
		fmt.Printf("%.2f * %.2f = %.2f\n", n, j, r)
	}

}
