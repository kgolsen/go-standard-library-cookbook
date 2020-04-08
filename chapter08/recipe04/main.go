package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	trunc = "TRUNCATE TABLE post;"
	ins   = "INSERT INTO post VALUES ($1, $2, $3)"
)

var testTable = []struct {
	ID      int
	Title   string
	Content string
}{
	{1, "Title One", "Content one"},
	{2, "Title Two", "Content two"},
	{3, "Title Three", "Content three"},
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

	_, err := db.Exec(trunc)
	if err != nil {
		panic(err)
	}

	stm, err := db.Prepare(ins)
	if err != nil {
		panic(err)
	}

	inserted := int64(0)
	for _, val := range testTable {
		fmt.Printf("Inserting record %d\n", val.ID)
		r, err := stm.Exec(val.ID, val.Title, val.Content)
		if err != nil {
			fmt.Printf("Failed to insert record %d: %v\n", val.ID, err)
		}
		if affected, err := r.RowsAffected(); err == nil {
			inserted = inserted + affected
		}
	}

	fmt.Printf("Inserted %d rows\n", inserted)

}
