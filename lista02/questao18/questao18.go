//questão 18

package main

import "fmt"

func main() {
	var n1, n2, n3 int
	fmt.Println("Digite três números para descobrir o MMC deles")
	fmt.Scan(&n1, &n2, &n3)
	cont := 2
	mmc := 1

	for {
		if n1 == 1 && n2 == 1 && n3 == 1 {
			break
		}
		if n1%cont != 0 && n2%cont != 0 && n3%cont != 0 {
			cont++
		} else {
			mmc *= cont
			fmt.Printf("%v %v %v | %v\n", n1, n2, n3, cont)
			if n1%cont == 0 {
				n1 = (n1 / cont)
			}
			if n2%cont == 0 {
				n2 = (n2 / cont)
			}
			if n3%cont == 0 {
				n3 = (n3 / cont)
			}
		}
	}
	fmt.Printf("MMC : %d\n", mmc)
}
