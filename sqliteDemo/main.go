package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // 在函数结束时关闭数据库连接
	sqlStmt := `CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s", err, sqlStmt)
	}
	stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("Alice")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Name: %s", id, name)
	}
}
