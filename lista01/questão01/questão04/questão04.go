package main

import "fmt"

func main() {

    var salario, kW, cConsumo, preço, cDesconto float64

    fmt.Println("Digite o valor do salário mínimo e a quantidade de kW gasta")
    fmt.Scan(&salario, &kW)
    preço = ((kW * 70) / 10000)
    cConsumo = (preço * salario)
    cDesconto = (cConsumo * 0.9)
    fmt.Print("Custo por kW : R$ ", preço, " Custo do consumo : R$ ", cConsumo, " Custo com desconto : R$ ", cDesconto)

}