package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	selOne = "SELECT id, title, content FROM post WHERE id = $1;"
	insert = "INSERT INTO post(id, title, content) VALUES ($1, $2, $3);"
)

type Post struct {
	ID      int
	Title   string
	Content string
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

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	newPost := Post{4, "Title four", "Content four"}
	_, err = tx.Exec(insert, newPost.ID, newPost.Title, newPost.Content)
	if err != nil {
		panic(err)
	}
	p := Post{}
	// note use of db to query - this occurs outside the open transaction
	if err = db.QueryRow(selOne, 4).Scan(&p.ID, &p.Title, &p.Content); err != nil {
		fmt.Println("Error in db.QueryRow: " + err.Error())
	}
	fmt.Println(p)
	// switch to using tx - occurs within the open transaction
	if err = tx.QueryRow(selOne, 4).Scan(&p.ID, &p.Title, &p.Content); err != nil {
		fmt.Println("Error in tx.QueryRow: " + err.Error())
	}
	fmt.Println(p)
	// rollback the insert
	if err = tx.Rollback(); err != nil {
		panic(err)
	}

}
