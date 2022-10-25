package main

import "fmt"

func main(){
    funcsPorLetra := map[string]map[string]float64{
        "G": {
            "Gabriel": 1125.50,
            "Guga": 1100.2,
        },

        "J": {
            "Joao": 2500.24,
        },

        "P": {
            "Pedro Junior": 1200.0,
        },
    }

    delete(funcsPorLetra, "P")
    for letra, funcs := range funcsPorLetra{
        fmt.Println(letra, funcs)
    }
}
