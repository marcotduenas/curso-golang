package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool){
    comprarTV50 := trab1 && trab2
    comprarTV32 := trab1 != trab2 // OU exclusivo
    comprarSorvete := trab1 || trab2

    return comprarTV50, comprarTV32, comprarSorvete
}


func main(){
    tv50, tv32, sorvete := compras(true, true)
    fmt.Printf("TV50: %t, TV32: %t, Sorverte: %t, Saudavel: %t\n", tv50, tv32, sorvete, !sorvete)
}
