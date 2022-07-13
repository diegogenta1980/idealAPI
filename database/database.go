package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1s450v3rFl0w!"
	dbname   = "assetsdb"
)

var DbConnection *sql.DB

func SetupDatabase() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DbConnection, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		println("erro db")
	}
	println("db inicializado")
}
