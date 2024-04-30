//questão 17
package main

import (
	"fmt"
	"math"
)

func main() {
	var codigo int
	var pCom, pVen float64
	var numVendas, lme10, lent10e20, lmaq20 int
	var maLcodigo int
	var maisVencodigo int
	var maLucro, ttCompra, ttVenda float64

	fmt.Println("Digite ")

	for {
		_, err := fmt.Scan(&codigo, &pCom, &pVen, &numVendas)
		if err != nil {
			break
		}

		lucro := (pVen - pCom) / pCom * 100
		if lucro < 10 {
			lme10++
		} else if lucro >= 10 && lucro <= 20 {
			lent10e20++
		} else {
			lmaq20++
		}

		if lucro > maLucro {
			maLucro = lucro
			maLcodigo = codigo
		}

		if float64(numVendas) > ttVenda {
			maisVencodigo = codigo
			ttVenda = float64(numVendas)
		}

		ttCompra += pCom * float64(numVendas)
		ttVenda += pVen * float64(numVendas)

		lucroTotal := (ttVenda - ttCompra) / ttCompra * 100

		fmt.Printf("Quantidade de mercadorias que geraram lucro menor que 10%%: %d\n", lme10)
		fmt.Printf("Quantidade de mercadorias que geraram lucro maior ou igual a 10%% e menor ou igual a 20%%: %d\n", lent10e20)
		fmt.Printf("Quantidade de mercadorias que geraram lucro maior do que 20%%: %d\n", lmaq20)
		fmt.Printf("Codigo da mercadoria que gerou maior lucro: %d\n", maLcodigo)
		fmt.Printf("Codigo da mercadoria mais vendida: %d\n", maisVencodigo)
		fmt.Printf("Valor total de compras: %.2f, valor total de vendas: %.2f e percentual de lucro total: %.2f%%\n", ttCompra, ttVenda, math.Abs(lucroTotal))
	}
}
