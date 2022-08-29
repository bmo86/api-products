package database

import (
	_ "github.com/lib/pq"
	"database/sql"
)

type PostgresRepository struct{
	db *sql.DB
}

func NewPostgresRepo(url string) (*PostgresRepository, error)  {
	db, err := sql.Open("postgres", url)

	if err != nil{
		return nil, err
	}

	return &PostgresRepository{db}, nil

}

func (repo *PostgresRepository) Close() error{
	return repo.db.Close()
}