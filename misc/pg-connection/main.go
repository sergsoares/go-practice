// https://golangdocs.com/golang-postgresql-example
// pgcli postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
// sudo docker run -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "postgres"
	dbname = "postgres"
)

func main() {
	fmt.Println("Starting")

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", conn)
	CheckError(err)

	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected")
}

func CheckError(err error){
	if err != nil {
		panic(err)
	}
}