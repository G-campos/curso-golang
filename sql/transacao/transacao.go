package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const dbAdd = "root:user_pwd@/cursogo"

func main() {
	db, err := sql.Open("mysql", dbAdd)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?, ?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, err = stmt.Exec(6, "Tiago") //chave duplicada

	if err != nil {
		tx.Rollback() //se der B.O. desfaz tudo
		log.Fatal(err)
	}

	tx.Commit()
}
