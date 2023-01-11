package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dbAdd = "root:user_pwd@/cursogo"

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", dbAdd)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//update
	rows, _ := db.Query("select id, nome from usuarios where id > ?", 1)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}

}
