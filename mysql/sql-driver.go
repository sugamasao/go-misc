package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Blog is table
type Blog struct {
	ID          int
	Title       string
	PublishedAt mysql.NullTime // Nullableな時刻の場合
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func main() {
	fmt.Println("生のsql-driverを使ったサンプル")
	// 時刻をtimeにバインドするためには `parseTime=true` が必要
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:13306)/example?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// prepared statementを使わない例もできるが、まあこれで統一しておくのがよかろう
	stmt, err := db.Prepare(`
SELECT
  *
FROM
  blog
`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var blog Blog
		err := rows.Scan(&(blog.ID), &(blog.Title), &(blog.PublishedAt), &(blog.CreatedAt), &(blog.UpdatedAt))
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(blog)
	}
}
