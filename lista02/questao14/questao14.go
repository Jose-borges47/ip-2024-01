// questão 14
package main

import "fmt"

func main() {
	var m, n int
	fmt.Println("Digite o número de linhas: ")
	fmt.Scanln(&m)
	fmt.Println("Digite o número de colunas: ")
	fmt.Scanln(&n)

	fmt.Println("Pares de índices inferiores à diagonal principal:")

	for i := 2; i <= m; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf("(%d,%d)", i, j)
			if j != i-1 {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}
