package main

import(
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

    stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
    stmt.Exec("Bastião Lopes", 69)
    stmt.Exec("Tião", 420)

    stmt2, _ := db.Prepare("delete from usuarios where id = ?")
    stmt2.Exec(8)



}
