package main

import (
	"fmt"
)

func main() {
	v := []int{1, 2, 3, 4, 5}
	fmt.Println(somarAte(v))
}

func somarAte(n []int) int {
	if len(n)-1 == 0 {
		return 0
	}
	return n[0] + somarAte(n[1:])
}