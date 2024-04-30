//questão 23
package main

import "fmt"

func somaDivisores(numero int) int {
	sum := 1
	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			sum += i
			if q := numero / i; q != i {
				sum += q
			}
		}
	}
	return sum
}

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	for i := 1; count < n; i++ {
		sum1 := somaDivisores(i)
		sum2 := somaDivisores(sum1)

		if i == sum2 && i < sum1 {
			fmt.Printf("(%d,%d)\n", i, sum1)
			count++
		}
	}
}
