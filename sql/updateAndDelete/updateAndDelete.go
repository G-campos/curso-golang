package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const dbAdd = "root:user_pwd@/cursogo"

func main() {
	db, err := sql.Open("mysql", dbAdd)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Uóxinton Istive", 1)
	stmt.Exec("Sheristone Uasleska", 2)

	//delete
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)

}
