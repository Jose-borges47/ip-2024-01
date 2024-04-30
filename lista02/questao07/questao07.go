//questão 7
package main

import (
	"fmt"
)

func main() {
	var (
		n      int
		sPar   float64
		sImpar float64
		qPar   float64
		qImpar float64
	)
	fmt.Println("Digite os números")
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		if n%2 == 0 {
			sPar += float64(n)
			qPar++
		} else {
			sImpar += float64(n)
			qImpar++
		}

	}

	var mPar, mImpar float64

	if qPar > 0.0 {
		mPar = sPar / float64(qPar)
	}
	if qImpar > 0.0 {
		mImpar = sImpar / float64(qImpar)

	}

	fmt.Printf("MEDIA PAR = %.2f\n", mPar)
	fmt.Printf("MEDIA IMPAR = %.2f\n", mImpar)
}
