package main

import "fmt"

func main () {
    var (
        n1, n2, n3 float64
        x float64
    )
    fmt.Println("Digite tres valores")
    fmt.Scanf("%f", &n1)
    fmt.Scanf("%f", &n2)
    fmt.Scanf("%f", &n3)

    x = (n1 + n2 + n3) / 3
    fmt.Println("MEDIA = ", x)

    if x >= 6 {
    fmt.Println("APROVADO")

    }   else {
    fmt.Println("REPROVADO")}
}