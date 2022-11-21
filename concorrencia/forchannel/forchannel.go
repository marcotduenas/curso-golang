package main

import(
    "fmt"
    "time"
)

func printNum(n int, c chan int){
    for i := 0; i <= n; i++{
        c <- i
        time.Sleep(time.Millisecond * 100)
    }
    close(c)
}


func main(){
    c := make(chan int, 30)
    go printNum(1000, c)
    for num := range c{
        fmt.Printf("%d ", num) 
    }
    fmt.Printf("\nFim.\n")
}

