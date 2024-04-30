//questão 02
package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("DIGITE A POPULAÇÃO DO PAÍS A")
	fmt.Scan(&a)
	fmt.Println("DIGITE A POPULAÇÃO DO PAÍS B")
	fmt.Scan(&b)

	n := 0
	for a < b {
		a = (a + (a * 3 / 100))
		b = (b + (b * 15 / 1000))
		n++
	}
	fmt.Printf("ANOS = %d", n)
}
