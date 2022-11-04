package main

import "fmt"

func obterValorAprovado(numero int) int{
    defer fmt.Println("FIM!")
    if numero >= 5000{
        fmt.Println("Alto")
        return 5000
    }else{
        fmt.Println("Baixo")
        return numero
    }
}

func main(){
    fmt.Println(obterValorAprovado(1000))
}
