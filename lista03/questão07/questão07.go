package main

import (
	"fmt"
)

func main() {
	for {
		var n int
		// Ler o tamanho do vetor
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		// Ler os elementos do vetor
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}

		// Encontrar o maior elemento
		maxValue := arr[0]
		for _, v := range arr {
			if v > maxValue {
				maxValue = v
			}
		}

		// Calcular a frequência dos valores menores ou iguais a i
		freq := make([]int, maxValue+1)
		for _, v := range arr {
			freq[v]++
		}

		// Calcular a frequência acumulada
		for i := 1; i <= maxValue; i++ {
			freq[i] += freq[i-1]
		}

		// Imprimir a saída
		for i := 0; i <= maxValue; i++ {
			fmt.Printf("(%d) %d\n", i, freq[i])
		}
	}
}
