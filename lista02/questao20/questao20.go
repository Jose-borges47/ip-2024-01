// questão 20
package main

import "fmt"

func main() {
	var soma, n int
	fmt.Println("Digite para saber se o número é perfeito")
	fmt.Scan(&n)
	fmt.Print(n, " = ")
	for i := 1; i < n; i++ {
		if n%i == 0 {
			soma += i
			fmt.Printf("%v", i)
			if i != n-1 {
				fmt.Print(" + ")
			}
		}
	}
	if soma == n {
		fmt.Printf("= %v (número perfeito)", soma)
	} else {
		fmt.Printf("= %v (não é um número perfeito)", soma)
	}
}
