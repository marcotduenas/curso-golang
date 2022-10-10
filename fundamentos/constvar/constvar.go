package main

import (
    "fmt"
    "math"
)
func main(){
     const PI float64 = 3.1415
     var raio = 3.2 //float64 automatico pelo compilador
     area := PI * math.Pow(raio, 2)
     fmt.Println("A área a circunferência é", area)

     const(
         a = 1
         b = 2
     )
     var e, f bool = true, false
     g, h := 1, 2
     fmt.Println(a, b, e, f, g, h)
}
