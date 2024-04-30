//questão 13
package main

import "fmt"

func main() {
	var n, r int
	fmt.Println("Digite um número inteiro")
	fmt.Scan(&n)
	r = 33*n + 31*(n*2)
	fmt.Print(r)
}
