package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

func main(){
    db, err := sql.Open("mysql", "root:12345678@tcp(localhost:1433)/cursogo")
    if err != nil{
        log.Fatal(err)
    }    
    defer db.Close()

    tx, _ := db.Begin()
    stmt, _ := tx.Prepare("INSERT INTO usuarios(id, nome) values(?,?)")
    
    stmt.Exec(69, "Carlos")
    stmt.Exec(420, "Jonas")
    _, err = stmt.Exec(1, "Tiago")

    if err != nil{
        tx.Rollback()
        log.Fatal(err)
    }
    tx.Commit()

}
        
