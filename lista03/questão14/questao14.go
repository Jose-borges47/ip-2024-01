package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	amostra := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&amostra[i])
	}

	for i := 1; i < n; i++ {
		key := amostra[i]
		j := i - 1
		for j >= 0 && amostra[j] > key {
			amostra[j+1] = amostra[j]
			j = j - 1
		}
		amostra[j+1] = key
	}

	var mediana float64
	if n%2 == 1 {
		mediana = float64(amostra[n/2])
	} else {
		mediana = (float64(amostra[n/2-1]) + float64(amostra[n/2])) / 2.0
	}
	fmt.Printf("%.2f\n", mediana)
}