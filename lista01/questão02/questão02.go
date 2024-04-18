package main

import "fmt"

func main () {
    var (
        nCasos int
        nPessoas float64
        nP, nG, nA, nC float64
        renda float64
    )
    fmt.Scan(&nCasos)

    for n := 0; n < nCasos; n++ {

    fmt.Scanf("%v %f %f %f %f", &nPessoas, &nP, &nG, &nA, &nC)

    renda = (nP / 100 * nPessoas * 1 + nG / 100 * nPessoas * 5 + nA / 100 * nPessoas * 10 + nC / 100 * nPessoas * 20)
    fmt.Printf("a renda do jogo %v é %v\n", n+1, renda)
    }
}