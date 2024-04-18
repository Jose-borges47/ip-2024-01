package main
import "fmt"

func main() {
    var salario, salarioc float64
    fmt.Println("DIGITE SEU SALÁRIO")
    fmt.Scan(&salario)
    if salario <= 300 {
        salarioc = (salario * 1.5)
        fmt.Print("SALÁRIO COM REAJUSTE = ", salarioc)
    } else if salario >= 300 {
        salarioc = (salario * 1.3)
        fmt.Print("SALÁRIO COM REAJUSTE = ", salarioc)
    }
}