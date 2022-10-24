package main

import "fmt"

func main(){
    salarios := map[string]float64{"Marco": 1000.15, "Leticia": 2000.75, "Fred": 3500.50}

    salarios["Zeca"] = 1350.0
    delete(salarios, "Ines")

    for nome, salario := range salarios{
        fmt.Println(nome, salario)
    }
}
