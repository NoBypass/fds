package repository

import "github.com/surrealdb/surrealdb.go"

type Repository struct {
	*surrealdb.DB
}

// New creates a new repository
func New(db *surrealdb.DB) *Repository {
	return &Repository{
		db,
	}
}
