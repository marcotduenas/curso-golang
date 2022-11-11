package main

import(
    "fmt"
    "encoding/json"
)

type produto struct{
    ID int `json:"id"`
    Nome string `json:"nome"`
    Preco float64 `json:"preco"`
    Tags []string `json:"tags"`
}


func main(){
    p1 := produto{1, "Notebok", 1999.50, []string{"Promoção", "Tecnologia"}}
    p1Json, _ := json.Marshal(p1)
    fmt.Println(string(p1Json))

    var p2 produto
    jsonString := `{"id":2,"nome":"Caneta","preco":2.50,"tags":["Papelaria"]}`
    json.Unmarshal([]byte(jsonString), &p2)
    fmt.Println(p2.Tags[0])

}
