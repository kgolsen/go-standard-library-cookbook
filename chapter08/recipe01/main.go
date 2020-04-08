package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://pgsql:pgsql@192.168.88.191:5432/testinggo"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Ping")
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Pong")
}
