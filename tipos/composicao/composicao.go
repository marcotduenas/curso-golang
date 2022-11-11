package main

import "fmt"

type esportivo interface{
    ligarTurbo() bool
}

type luxuoso interface{
    fazerBaliza()
}

type esportivoLuxuoso interface{
    esportivo
    luxuoso
}

type bmw7 struct{
    turboLigado bool
}

func (b bmw7) ligarTurbo()bool{
    return true
}

func (b bmw7) fazerBaliza(){
    fmt.Println("baliza")
}

func main(){
    var carro1 esportivoLuxuoso = bmw7{false}
    fmt.Println(carro1.ligarTurbo())
    carro1.fazerBaliza()
}
