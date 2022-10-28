package main
import "fmt"


func media(numeros ...float64) float64{
    total := 0.0
    for _, num := range numeros{
        total += num
    }

    return total / float64(len(numeros))
}

func main(){
    fmt.Printf("Media: %.2f", media(8.5, 9.8, 1.3, 2.7))
}
