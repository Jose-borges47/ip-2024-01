// questão 1

package main

import "fmt"

func main() {
	const (
		tP    = 8
		tL    = 5
		tH    = 128
		pM    = 0.75 * tH
		nMinA = 6.0
		pP    = 0.7
		pL    = 0.15
		pT    = 0.15
	)

	var (
		matricula, pre int
		pro, lis, tb   float64
	)
	fmt.Println("Digite a matrícula do aluno, as notas das 8 provas, as notas das 5 listas, a nota do trabalho final e o valor correspondente a presença do aluno")
	for {
		fmt.Scan(&matricula)
		if matricula == -1 {
			break
		}

		sPro := 0.0
		for i := 0; i < tP; i++ {
			fmt.Scan(&pro)
			sPro += pro
		}
		mePro := sPro / float64(tP)

		sLis := 0.0
		for i := 0; i < tL; i++ {
			fmt.Scan(&lis)
			sLis += lis
		}
		meLis := sLis / float64(tL)

		fmt.Scan(&tb, &pre)

		nFinal := pP*mePro + pL*meLis + pT*tb

		situacao := ""
		if nFinal >= nMinA && float64(pre) >= pM {
			situacao = "APROVADO"
		} else if float64(pre) < pM && nFinal < nMinA {
			situacao = "REPROVADO POR NOTA E POR FREQUENCIA"
		} else if float64(pre) < pM {
			situacao = "REPROVADO POR FREQUENCIA"
		} else if nFinal < nMinA {
			situacao = "REPROVADO POR NOTA"
		}

		fmt.Printf("Matricula: %d, Nota Final: %.2f, Situação Final: %s\n", matricula, nFinal, situacao)
	}
}
