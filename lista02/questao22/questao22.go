//Questã0 22
package main

import "fmt"

func main() {
	var n float64
	fmt.Println("Digite um número real")
	fmt.Scan(&n)
	for i := 1.0; 1 > 0; i++ {
		if n*i == float64(int(n*i)) {

			fmt.Printf("%.f/%.f", n*i, i)
			return

		}
	}
}
