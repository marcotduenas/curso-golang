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
    db, err := sql.Open("mysql", "root:12345678@tcp(172.17.0.2:3306)/cursogo")
    if err != nil{
        log.Fatal(err)
    }
    defer db.Close()

    rows, _ := db.Query("select id, nome from usuarios where id = ?", 1)
    defer rows.Close()

    for rows.Next(){
        var u usuario
        rows.Scan(&u.id, &u.nome)
        fmt.Println(u)
    }


}
