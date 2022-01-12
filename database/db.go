package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func StartDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:hbkpatrick@tcp(localhost:3306)/crud_usuario")
	if err != nil {
		log.Fatal(err.Error())
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr.Error())
	}

	return db
}
