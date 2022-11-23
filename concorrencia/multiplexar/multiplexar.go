package main
 
import (
    "fmt"
)

func encaminhar(origem <- chan string, destino chan string){
    for{
        destino <- <- origem
    }
}


//multiplexar - misturar mensagens num unico canal
func  juntar(entrada1, entrada2 <- chan string) <- chan string{
    c := make(chan string)
    go encaminhar(entrada1, c)
    go encaminhar(entrada2, c)

    return c
}

func main(){
    
    c := juntar(
        generator.Titulo("https://www.gnu.org", "https://metager.org"),
        generator.Titulo("https://www.archlinux.org", "https://gentoo.org"),
    )
    fmt.Println(c)
}
