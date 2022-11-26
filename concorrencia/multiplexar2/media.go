package main

import(
    "fmt"
    "time"
)

func media(n1, n2, n3 float64) <- chan float64{
    c := make(chan float64)
    go func(){
        for{
            soma := (n1 + n2 + n3)/3
            time.Sleep(time.Second)
            c <- soma
        }
    }()
    return c
}


func mediaPorAluno(entrada1, entrada2 <- chan float64) <- chan float64{
    n := make(chan float64)
    go func(){
        for{
            select{
            case nota := <- entrada1:
                n <- nota
            case nota := <- entrada2:
                n <- nota
            }
        }
    }()
    return n
}


func main(){
    i := mediaPorAluno(media(7.5, 8, 5.25), media(6.25, 6.75, 7))
    fmt.Println(<-i, "|", <-i)
}
