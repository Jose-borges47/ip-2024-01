// questão 19
package main

import "fmt"

func main() {
	var v, f int
	fmt.Println("Digite um número inteiro")
	fmt.Scan(&v)
	in := 1
	for i := 1; i <= v; i++ {
		fmt.Printf("%v * %v * %v = ", i, i, i)
		f = in + (2 * (i - 1))
		for n := in; n <= f; n += 2 {
			fmt.Printf("%v", n)
			if n != f {
				fmt.Print(" + ")
			} else {
				fmt.Print("\n")
			}
		}
		in = f + 2
	}
}
