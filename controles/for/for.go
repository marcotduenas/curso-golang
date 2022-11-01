package main

import (
    "fmt"
    "time"
)

func main(){
    var i byte = 1
    for i <= 10{
        fmt.Println(i)
        i++
    }

    for i := 0; i <= 20; i += 2 {
        fmt.Printf("%d\n", i)
    }

    for i := 1; i <= 10; i++{
        if i%2 == 0 {
            fmt.Printf("%d, Par\n", i)
        }else{
            fmt.Printf("%d, Impar\n", i)
        }

    }

    for{
        fmt.Println("Para sempre...")
        time.Sleep(time.Second)
    }
}
