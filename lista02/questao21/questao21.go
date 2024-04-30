//questão 21
package main

import (
	"fmt"
)

func main() {
	for {
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}

		var sequencia []float64
		for i := 0; i < n; i++ {
			var num float64
			fmt.Scan(&num)
			sequence = append(sequencia, num)
		}

		ordenado := true
		for i := 1; i < len(sequencia); i++ {
			if sequencia[i-1] >= sequencia[i] {
				ordenado = false
				break
			}
		}

		if ordenaso {
			fmt.Println("ORDENADA")
		} else {
			fmt.Println("DESORDENADA")
		}
	}
}
