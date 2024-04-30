//questão 15
package main

import "fmt"

func main() {
	fmt.Println("Digite o número de testes e em seguida os números a serem comparados em cada teste")
	var a, b, t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&a, &a)
		reva := reverter(a)
		revb := reverter(b)

		if reva > revb {
			fmt.Println(reverter(reva))
		} else {
			fmt.Println(reverter(revb))
		}
	}
}

func reverter(n int) int {
	rev := 0
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev
}
