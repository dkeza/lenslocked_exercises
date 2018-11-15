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
		//id       int
		//name     string
		//email    string
		psqlInfo string
		err      error
		db       *sql.DB
		//		row      *sql.Row
	)
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	type User struct {
		ID    int
		name  string
		email string
	}
	var users []User
	rows, err := db.Query(`SELECT id, name, email FROM users WHERE id = $1 OR 1=1`, 1)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.name, &user.email)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No rows")
			} else {
				panic(err)
			}
		}
		users = append(users, user)

	}

	if rows.Err() != nil {
		// Handle error here!
	}
	fmt.Println(users)
	// Query one row
	//
	// row = db.QueryRow(`
	// 		SELECT id, name, email FROM users WHERE id = $1
	// 	`, 10)
	// err = row.Scan(&id, &name, &email)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		fmt.Println("No rows")
	// 	} else {
	// 		panic(err)
	// 	}
	// }

	//fmt.Println(id, name, email)
}
