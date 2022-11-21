package main

import(
    "time"
    "fmt"
)



func rotina(c chan int){
    time.Sleep(time.Second)
    c <- 1
    fmt.Println("So depois que for lido")    
}


func main(){
    c := make(chan int)

    go rotina(c)

    fmt.Println(<-c)
    fmt.Println("Foi lido")
    fmt.Println(<-c)
    fmt.Println("Fim")
}
