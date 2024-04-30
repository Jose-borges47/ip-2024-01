// questão 5

package main

import "fmt"

func main() {
	var n, comp, compSeg int
	fmt.Scan(&n)

	sequencia := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&sequencia[i])
	}

	comp = 0
	compSeg = 0
	for i := 1; i < n; i++ {
		if sequencia[i] > sequencia[i-1] {
			compSeg++
			if compSeg > comp {
				comp = compSeg
			}
		} else {
			compSeg = 0
		}
	}

	fmt.Printf("O comprimento do segmento crescente maximo e: %d\n", comp)
}
