package main

import(
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "fmt"
)

type usuario struct{
    id int 
    nome string
}

func main(){
    db, err := sql.Open("mysql", "root:12345678@tcp(localhost:1433)/cursogo")
    if err != nil{
        log.Fatal(err)
    }
    defer db.Close()

    rows, _ := db.Query("select id, nome from usuarios where id > ?", 5)
    defer rows.Close()

    for rows.Next(){
        var u usuario
        rows.Scan(&u.id, &u.nome)
        fmt.Println(u)
    }


}
