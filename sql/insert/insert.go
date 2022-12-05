package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

func main(){
    db, err := sql.Open("mysql", "root:12345678@tcp(localhost:1433)/cursogo")
    if err != nil{
        panic(err)
    }    
    defer db.Close()

    stmt, _ := db.Prepare("INSERT INTO usuarios(nome) values(?)")
    stmt.Exec("Maria")
    stmt.Exec("Paulo")

    res, _ := stmt.Exec("Pedro")

    stmt.Exec("Jorge")

    id, _ := res.LastInsertId()
    fmt.Println(id)

    linhas, err := res.RowsAffected()
    fmt.Println(linhas)
}
        
