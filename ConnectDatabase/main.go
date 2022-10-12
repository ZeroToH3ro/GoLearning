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
	password = "minh21052002"
	dbname   = "connect_database_to_go"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	sqlStatement := `
INSERT INTO users (age, email, first_name, last_name)
VALUES ($1, $2, $3, $4)
RETURNING id`
	id := 2
	err = db.QueryRow(sqlStatement, 31, "jon12@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)

	updateSqlStatement := `
UPDATE users
SET first_name = $2, last_name = $3
WHERE id = $1;`
	_, err = db.Exec(updateSqlStatement, 1, "NewFirst", "NewLast")
	if err != nil {
		panic(err)
	}
	/*
		deleteSqlStatement := `
		DELETE FROM users
		WHERE id = $1;`
		_, err = db.Exec(deleteSqlStatement, 1)
		if err != nil {
			panic(err)
		}
	*/

}
