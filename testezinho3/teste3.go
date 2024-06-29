package main

import "fmt"

func inverternums(x []int) []int {
	for i := 0; i < len(x)-1; i++ {
		for a := 0; a < len(x)-i-1; a++ {
			if x[a] > x[a+1] {
				x[a], x[a+1] = x[a+1], x[a]
			}
		}
	}
	return x
}

func main() {
	array := []int{32, 92, 43, 2, 54, 14, 12}
	fmt.Println(inverternums(array))
}
