package main

import(
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result{
    result, err := db.Exec(sql)
    if err != nil{
        panic(err)
    }

    return result
}

func main(){
    db, err := sql.Open("mysql", "root:12345678@tcp(172.17.0.2:3306)/")
    if err != nil{
        panic(err)
    }

    defer db.Close()

    exec(db, "create database if not exists cursogo")
    exec(db, "use cursogo")
    exec(db, `create table usuarios(
        id integer auto_increment,
        nome varchar(80),
        PRIMARY KEY (id)
    )`)
    exec(db, "insert into usuarios(nome) values ('João Silva')")
}
