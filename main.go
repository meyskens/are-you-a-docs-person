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

	/*repos, err := db.Query("select repository_id from repositories")
	if err != nil {
		panic(err)
	}

	for repos.Next() {
		var repoID string
		repos.Scan(&repoID)
		repoList = append(repoList, repoID)
	}

	for _, repoID := range repoList {

	}*/

	codeCommits := map[string]int{}
	docsCommits := map[string]int{}

	for _, commit := range getCommits("Cast") {
		//files := getCommitFiles("Cast", commit.CommitHash)
		fmt.Printf("%s\n", commit.CommitHash)
		tree := getTreeEntry("Cast", commit.TreeHash)
		spew.Dump(tree)
		/*docfiles := 0
		for _, file := range files {
			if _, ok := docFiles[file.Language]; ok {
				docfiles++
			}
		}
		if docfiles > 0 {
			docsCommits[commit.CommitAuthorName]++
		}
		if len(files) > docfiles {
			codeCommits[commit.CommitAuthorName]++
		}*/
	}

	spew.Dump(codeCommits)
	spew.Dump(docsCommits)
}

func getCommits(repoID string) []commit {
	commitList := []commit{}
	commits, err := db.Query(fmt.Sprintf("SELECT repository_id, commit_hash, commit_author_name, commit_author_email, commit_author_when, tree_hash FROM commits WHERE repository_id='%s'", repoID))
	if err != nil {
		panic(err)
	}
	for commits.Next() {
		c := commit{}
		err := commits.Scan(&c.RepositoryID, &c.CommitHash, &c.CommitAuthorName, &c.CommitAuthorEmail, &c.CommitAuthorWhen, &c.TreeHash)
		if err != nil {
			panic(err)
		}
		commitList = append(commitList, c)
	}

	return commitList
}

func getTreeEntry(repoID, treeHash string) []treeEntry {
	treeList := []treeEntry{}
	trees, err := db.Query(fmt.Sprintf("SELECT tree_entry_name, blob_hash FROM tree_entries WHERE repository_id='%s' AND tree_hash='%s'", repoID, treeHash))
	if err != nil {
		panic(err)
	}
	for trees.Next() {
		t := treeEntry{}
		err := trees.Scan(&t.TreeEntryName, &t.BlobHash)
		if err != nil {
			panic(err)
		}
		treeList = append(treeList, t)
	}

	return treeList
}

func getCommitFiles(repoID, commitHash string) []commitFile {
	filesList := []commitFile{}
	files, err := db.Query(fmt.Sprintf(`
	select
		file_path, language(file_path) as lang, blob_hash, tree_hash
	FROM
		commit_files
	WHERE
	lang IS NOT NULL
	AND commit_files.repository_id = '%s'
	AND commit_files.commit_hash = '%s'
	`, repoID, commitHash))
	if err != nil {
		panic(err)
	}

	for files.Next() {
		f := commitFile{}
		err := files.Scan(&f.Path, &f.Language, &f.BlobHash, &f.TreeHash)
		if err != nil {
			panic(err)
		}
		filesList = append(filesList, f)
	}

	return filesList
}
