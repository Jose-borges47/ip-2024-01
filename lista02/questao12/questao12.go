// questão 12
package main

import "fmt"

func main() {
	var valorIngresso, valorInicial, valorFinal float64
	fmt.Println("Digite o valor do ingresso, o valor inicial e o valor final")
	_, err := fmt.Scan(&valorIngresso, &valorInicial, &valorFinal)
	if err != nil {
		fmt.Println("número inválido", err)
		return
	}
	if valorInicial >= valorFinal {
		fmt.Println("Intervalo inválido")
		return
	}
	mLucro := 0.0
	mValor := 0.0
	mqIng := -1.0

	for p := valorInicial; p <= valorFinal; p += 1.0 {
		var mIng float64
		if p < 5.0 {
			mIng = 120.0 + 25.0
		} else {
			mIng = 120.0 - 30.0
		}
		dFixa := 200 + mIng*0.05
		r := p * mIng
		lucro := r - dFixa - valorIngresso*mIng

		if lucro > mLucro {
			mLucro = lucro
			mValor = p
			mqIng = mIng
		}
		fmt.Printf("V: %.2f, N: %.0f, L: %.2f\n", p, mIng, lucro)
	}
	if mqIng != -1.0 {
		fmt.Printf("Melhor valor final: %.2f\n", mValor)
		fmt.Printf("Lucro: %.2f\n", mLucro)
		fmt.Printf("Numero de ingressos: %.0f\n", mqIng)
	} else {
		fmt.Println("Não foi possível encontrar um lucro máximo")
	}
}
