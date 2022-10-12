package main

import(
    "fmt"
    "reflect"
    "math"
)


func main(){
    fmt.Println(1 ,2, 1000)
    fmt.Println("Tipo é", reflect.TypeOf(10))

    var b byte = 255
    fmt.Println("Tipo é", reflect.TypeOf(b))

    i1 := math.MaxInt64
    fmt.Println(i1)
    fmt.Println(reflect.TypeOf(i1))

    var i2 rune = 'a'
    fmt.Println(reflect.TypeOf(i2))
    fmt.Println(i2)

    var x = float32(42.50)
    fmt.Println(reflect.TypeOf(x))
    fmt.Println(42.50) //float63

    bo := true
    fmt.Println(reflect.TypeOf(bo))
    fmt.Println(bo)

    s1 := "Olá, meu nome é Marco"
    fmt.Println(s1 + "!")
    fmt.Println(len(s1))
    fmt.Println(reflect.TypeOf(len(s1)))

    s2 := `Meu 
    nome
        e
            Marcola`
    fmt.Println(len(s2))

    char := 'a' //Nao existe char reservado, vira int32
    fmt.Println(char)
    fmt.Println(reflect.TypeOf(char))
}
