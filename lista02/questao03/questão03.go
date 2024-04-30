//questão 03
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int64
	fmt.Println("Digite o número a ser calculado o fatorial")
	fmt.Scan(&n)

	r := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		r.Mul(r, big.NewInt(i))
	}

	fmt.Printf("%d! = %s\n", n, r.String())
}
