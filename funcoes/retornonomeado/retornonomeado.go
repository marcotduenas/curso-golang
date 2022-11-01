package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int){
    segundo = p2
    primeiro = p1
    return
}

func main(){
    x, y := trocar(5, 3)
    fmt.Println(x,y)

    segundo, primeiro := trocar(7, 1)
    fmt.Println(segundo, primeiro)
}
