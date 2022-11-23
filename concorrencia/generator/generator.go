package main

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
)

// Funcao que retorna titulo de pagina html
func Titulo(urls ...string) <- chan string{
    c := make(chan string)
    for _, url := range urls{
        go func(url string){
            resp, _ := http.Get(url)
            html, _ := ioutil.ReadAll(resp.Body)

            r, _ := regexp.Compile("<title>(.*?)<\\/title>")
            c <- r.FindStringSubmatch(string(html))[1]
        }(url)
    }
        return c
}

func main(){
    t1 := titulo("https://gnu.org", "https://www.gentoo.org")
    t2 := titulo("https://metager.org", "https://archlinux.org")
    fmt.Println("Primeiros:", <-t1, "|", <-t2)
    fmt.Println("Segundos:", <-t1, "|", <-t2)
}
