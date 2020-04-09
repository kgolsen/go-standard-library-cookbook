package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	sel = "SELECT id, title, content FROM post WHERE ID = $1;"
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

func parseWithRawBytes(rows *sql.Rows, cols []string) map[string]interface{} {
	values := make([]sql.RawBytes, len(cols))
	pointers := make([]interface{}, len(values))
	for i := range values {
		pointers[i] = &values[i]
	}
	if err := rows.Scan(pointers...); err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	for i, col := range values {
		if col == nil {
			m[cols[i]] = nil
		} else {
			m[cols[i]] = string(col)
		}
	}
	return m
}

func parseToMap(rows *sql.Rows, cols []string) map[string]interface{} {
	values := make([]interface{}, len(cols))
	pointers := make([]interface{}, len(cols))
	for i := range values {
		pointers[i] = &values[i]
	}

	if err := rows.Scan(pointers...); err != nil {
		panic(err)
	}

	m := make(map[string]interface{})
	for i, colName := range cols {
		if values[i] == nil {
			m[colName] = nil
		} else {
			m[colName] = values[i]
		}
	}
	return m
}

func main() {

	db := createConnection()
	defer db.Close()

	rows, err := db.Query(sel, 1)
	if err != nil {
		panic(err)
	}
	cols, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		m := parseWithRawBytes(rows, cols)
		fmt.Println(m)
		m = parseToMap(rows, cols)
		fmt.Println(m)
	}

}
