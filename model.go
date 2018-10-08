package main

type commit struct {
	RepositoryID      string
	CommitHash        string
	CommitAuthorName  string
	CommitAuthorEmail string
	CommitAuthorWhen  []byte
	TreeHash          string
}

type commitFile struct {
	Path     string
	Language string
	BlobHash string
	TreeHash string
}

type treeEntry struct {
	TreeEntryName string
	BlobHash      string
}
