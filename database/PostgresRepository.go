package database

import "database/sql"

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}
