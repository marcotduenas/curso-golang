package main

import(
    "net/http"
    "log"
)

func main(){
    fs := http.FileServer(http.Dir("publico"))
    http.Handle("/", fs)
    log.Println("Executando servidor")
    log.Fatal(http.ListenAndServe(":3000", nil))
}
