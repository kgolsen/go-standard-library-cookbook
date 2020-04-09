package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	call = "SELECT * FROM format_name($1,$2,$3);"
)

type Result struct {
	Name     string
	Category int
}

func createConnection() *sql.DB {
	connStr := "postgres://pgsql:pgsql@192.168.88.191:5432/testinggo"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	db := createConnection()
	defer db.Close()

	r := Result{}
	err := db.QueryRow(call, "John", "Doe", 42).Scan(&r.Name)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result: %+v\n", r)
}
