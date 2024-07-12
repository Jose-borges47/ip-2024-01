package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	V := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&V[i])
	}
	W := make([]int, n)
	for i := 0; i < n; i++ {
		W[i] = V[n-1-i]
	}
	maiorV := V[0]
	for _, v := range V {
		if v > maiorV {
			maiorV = v
		}
	}
	menorW := W[0]
	for _, w := range W {
		if w < menorW {
			menorW = w
		}
	}
	for i, v := range V {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
	for i, w := range W {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(w)
	}
	fmt.Println()
	fmt.Println(maiorV)
	fmt.Println(menorW)
}