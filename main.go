package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@/")
	if err != nil {
		panic(err)
	}

	repos, err := db.Query("select repository_id from repositories")
	if err != nil {
		panic(err)
	}

	for repos.Next() {
		var repoID string
		repos.Scan(&repoID)
		fmt.Println(repoID)
	}
}
