package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {

	return 0, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {

	return nil, nil
}

/*
Return the lastest 10 snippets
*/
func (m *SnippetModel) Latest(title string, content string, expires int) (*[]Snippet, error) {

	return nil, nil
}
