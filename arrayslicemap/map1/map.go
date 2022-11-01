package main

import "fmt"

func main(){
//  var aprovados map[int]string
//  mapas devem ser inicializados
    
    aprovados := make(map[int]string)
    aprovados[12345678] = "Maria"
    aprovados[87654321] = "Carlos"
    aprovados[65656565] = "Ana"

    fmt.Println(aprovados)


    for cpf, nome := range aprovados{
        fmt.Printf("%s (CPF: %d)\n", nome, cpf)
    }

    fmt.Println(aprovados[12345678])
    delete(aprovados, 65656565)
    fmt.Println(aprovados[65656565])
}
