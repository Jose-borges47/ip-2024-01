package main

import (
	"fmt"
)

func troca(vetor []int, i, j int) {
	vetor[i], vetor[j] = vetor[j], vetor[i]
}

func trocaOpostosSeMenor(vetor []int, tamanho int) {
	for i := 0; i < tamanho/2; i++ {
		oposto := tamanho - 1 - i
		if vetor[i] < vetor[oposto] {
			troca(vetor, i, oposto)
		}
	}
}

func main() {
	var numCasos int
	fmt.Scan(&numCasos)

	for i := 0; i < numCasos; i++ {
		var n int
		fmt.Scan(&n)

		vetor := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&vetor[j])
		}

		trocaOpostosSeMenor(vetor, n)

		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(vetor[j])
		}
		fmt.Println()
	}
}