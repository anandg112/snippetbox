package models

import (
	"database/sql"
	"time"
)

// Define a Snippet type to hold the data for an individual snippet.
// The fields in the Snippet type correspond to the fields in our Snippets table.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
type snippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the database.
func (m *snippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Return a specific snippet based on its id.
func (m *snippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

// Return the 10 most recently created snippets.
func (m *snippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
