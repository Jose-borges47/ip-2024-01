// questão 6
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	seq := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&seq[i])
	}

	tamSubseq := 1
	SubseqA := 1
	for i := 1; i < n; i++ {
		if seq[i] == seq[i-1] {
			SubseqA++
			if SubseqA > tamSubseq {
				tamSubseq = SubseqA
			}
		} else {
			SubseqA = 1
		}
	}

	fmt.Printf("A maior subsequência de elementos iguais possui %d elementos.\n", tamSubseq)
}
