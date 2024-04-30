//questão 10
package main

import (
	"fmt"
)

func main() {
	var (
		a, b, c   float64
		matricula int
	)
	fmt.Println("Digite a matrícula, o número de horas trabalhadas e o valor recebido por hora trabalhada")
	for {
		_, err := fmt.Scanf("%d %f %f", &matricula, &a, &b)
		if err != nil {
			fmt.Print("entrada inválida", err)
			return
		}

		if matricula == 0 {
			break
		}
		c = (a * b)
		fmt.Printf("matricula %d salario %.2f \n", matricula, c)
		fmt.Scanln()
	}
}
