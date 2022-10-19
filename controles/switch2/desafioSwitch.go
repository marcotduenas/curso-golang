package main

import "fmt"


func main(){
    nota := 11
    switch{
        case nota >=9 && nota <= 10:
            fmt.Println("A")
        case nota < 9 && nota >= 8:
            fmt.Println("B")
        case nota < 8 && nota >= 7:
            fmt.Println("C")
        case nota < 7 && nota >= 5:
            fmt.Println("D")
        case nota < 5 && nota <= 4:
            fmt.Println("E")
        default:
            fmt.Println("Invalido, tente novamente")

    }
}
