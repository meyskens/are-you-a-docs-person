package main

import (
	"database/sql"
	"fmt"

	"github.com/davecgh/go-spew/spew"

	_ "github.com/go-sql-driver/mysql"
)

var repoList = []string{}
var db *sql.DB
var docFiles = map[string]bool{
	"Markdown": true,
	"HTML":     true,
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root@/")
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
		repoList = append(repoList, repoID)
	}

	for _, repoID := range repoList {
		files := getFileCount(repoID)
		mdCount := getMarkdownFileCount(repoID)
		//getCodeFileComments(repoID)
		fmt.Printf("%s is %f %% markdown\n", repoID, (float64(mdCount)/float64(files))*100.0)
	}
}

func getMarkdownFileCount(repoID string) int {
	q, err := db.Query(fmt.Sprintf(`
	SELECT COUNT(*)
	FROM
		commit_files
	NATURAL JOIN
		ref_commits
	WHERE
		ref_commits.ref_name = 'HEAD'
		AND ref_commits.history_index = 0
		AND commit_files.repository_id = '%s'
		AND commit_files.file_path NOT REGEXP '^vendor.*'
		AND language(commit_files.file_path) = 'Markdown'
	`, repoID))
	if err != nil {
		panic(err)
	}
	q.Next()
	count := 0
	q.Scan(&count)
	return count
}

func getCodeFileComments(repoID string) []fileComments {
	comments := []fileComments{}
	q, err := db.Query(fmt.Sprintf(`
	SELECT 
		file_path,
		language(commit_files.file_path),
		uast_xpath(uast_mode("annotated", blobs.blob_content, language(file_path)), "//*[@roleComment]")
	FROM
		commit_files
	NATURAL JOIN
		ref_commits
	NATURAL JOIN
		blobs
	WHERE
		ref_commits.ref_name = 'HEAD'
		AND ref_commits.history_index = 0
		AND commit_files.repository_id = '%s'
		AND commit_files.file_path NOT REGEXP '^vendor.*'
		AND language(commit_files.file_path) != 'Markdown'
		AND is_binary(blobs.blob_content) = false
	`, repoID))
	if err != nil {
		panic(err)
	}

	for q.Next() {
		f := fileComments{}
		q.Scan(&f.FilePath, &f.Language, &f.UASTs)
		comments = append(comments, f)
		spew.Dump(f)
	}
	return comments
}

func getFileCount(repoID string) int {
	q, err := db.Query(fmt.Sprintf(`
	SELECT 
		COUNT(*)
	FROM
		commit_files
	NATURAL JOIN
		ref_commits
	WHERE
		ref_commits.ref_name = 'HEAD'
		AND ref_commits.history_index = 0
		AND commit_files.repository_id = '%s'
		AND commit_files.file_path NOT REGEXP '^vendor.*'
	`, repoID))
	if err != nil {
		panic(err)
	}
	q.Next()
	count := 0
	q.Scan(&count)
	return count
}
