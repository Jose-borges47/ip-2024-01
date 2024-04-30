//questão 27
package main

import "fmt"

func main() {
	var n int
	fmt.Println("Digite um número inteiro")
	for {
		fmt.Scan(&n)
		if n < 1 {
			fmt.Printf("fatoração não é possível para o número %v", n)
		} else {
			break
		}
	}
	if n == 1 {
		fmt.Print(" 1 = 1")
	} else {
		fmt.Printf("%v = ", n)
		fatorial := 2
		for {
			if n%fatorial == 0 {
				n = n / fatorial
				fmt.Printf("%v", fatorial)
				if n == 1 {
					break
				}
				fmt.Print(" x ")
			} else {
				fatorial++
			}
		}
	}
}
