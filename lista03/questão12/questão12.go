package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	lista := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&lista[i])
	}

	for i := 1; i < n; i++ {
		key := lista[i]
		j := i - 1
		for j >= 0 && lista[j] > key {
			lista[j+1] = lista[j]
			j = j - 1
		}
		lista[j+1] = key
	}
	for _, num := range lista {
		fmt.Println(num)
	}
}
