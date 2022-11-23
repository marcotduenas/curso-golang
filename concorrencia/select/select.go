package main

import(
    "fmt"
    //"time"
)

func main(){
    c1 := make(chan string)
    c2 := make(chan string)

    go func(){
        c1 <- "One"
    }()
    
    go func(){
        c2 <- "Two"
    }()

    select{
    case msg1 := <-c1:
        fmt.Println("Recebido", msg1)
    case msg2 := <-c2:
        fmt.Println("Recebido", msg2)
    }
}
