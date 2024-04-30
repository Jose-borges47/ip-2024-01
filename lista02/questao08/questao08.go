//questão 8
package main

import "fmt"

func main() {
	var n int
	fmt.Println("Digite a quantidade de times")
	fmt.Scanln(&n)

	if n <= 1 {
		fmt.Println("Campeonato invalido!")
		return
	}

	fmt.Println("Saída:")
	count := 1
	for i := 1; i <= n; i++ {
		for a := i + 1; a <= n; a++ {
			fmt.Printf("Final %d: Time%d X Time%d\n", count, i, a)
			count++
		}
	}
}
