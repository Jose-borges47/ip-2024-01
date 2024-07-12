package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Scan(&N)

		vetor := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&vetor[i])
		}
		minDistancia := math.MaxInt32
		comparacoes := 0

		for i := 1; i < N; i++ {
			distancia := abs(vetor[i] - vetor[i-1])
			comparacoes++
			if distancia < minDistancia {
				minDistancia = distancia
			}
		}
		fmt.Printf("%d %d\n", minDistancia, comparacoes)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
