package main

import (
	"fmt"
)

func main() {
  
	var n int
	fmt.Scan(&n)

	freq := make(map[int]int)

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		freq[num]++
	}

	var maxFreq int
	var minValue int = 101 

	for value, count := range freq {
		if count > maxFreq || (count == maxFreq && value < minValue) {
			maxFreq = count
			minValue = value
		}
	}

	fmt.Println(minValue)
	fmt.Println(maxFreq)
}
