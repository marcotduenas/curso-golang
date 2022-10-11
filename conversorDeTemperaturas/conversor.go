package main 

import "fmt"

func converterFarenheitCelsius(temperaturaFarenheit int) float64 {
    return float64(temperaturaFarenheit - 32)/1.8 
}

func main(){
    minhaConversao := converterFarenheitCelsius(100)
    fmt.Printf("%.2f\n",minhaConversao)
}
