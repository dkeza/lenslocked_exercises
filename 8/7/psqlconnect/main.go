package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "lenslocked_dev"
)

func main() {
	var (
		id       int
		psqlInfo string
		err      error
		db       *sql.DB
		row      *sql.Row
	)
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	row = db.QueryRow(`
			INSERT INTO users (name, email) 
				VALUES ($1, $2) RETURNING id
		`, "My Name", "my@name.com")
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}
