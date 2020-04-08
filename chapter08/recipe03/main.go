package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	sel   = "SELECT * FROM post;"
	trunc = "TRUNCATE TABLE post;"
	ins   = "INSERT INTO post(id, title, content) VALUES (1, 'Title 1', 'Content 1'), (2, 'Title 2', 'Content 2')"
)

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
	fmt.Println("Table truncated")
	r, err := db.Exec(ins)
	if err != nil {
		panic(err)
	}
	affected, err := r.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted %d rows\n", affected)

	rs, err := db.Query(sel)
	if err != nil {
		panic(err)
	}
	defer rs.Close()
	cols, err := rs.Columns()
	if err != nil {
		panic(err)
	}
	colt, err := rs.ColumnTypes()
	if err != nil {
		panic(err)
	}
	x := ""
	for idx, col := range cols {
		x = x + col + "(" + colt[idx].DatabaseTypeName() + ") "
	}
	fmt.Printf("Columns: %v\n", x)
	count := 0
	for rs.Next() {
		count++
	}
	fmt.Printf("%d rows returned\n", count)

}
