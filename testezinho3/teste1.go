package main

import (
    "bufio"
    "fmt"
    "os"
)

func inverter(input string) string {
    x := []rune(input)
    for i, a := 0, len(x)-1; i < a; i, a = i+1, a-1 {
        x[i], x[a] = x[a], x[i]
    }
    return string(x)
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Digite a frase ao contrário:")

    input, _ := reader.ReadString('\n')

    input = input[:len(input)-1]

    invertido := inverter(input)
    fmt.Println("Frase ao contrário:", invertido)
}