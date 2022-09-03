package database

import (
	"context"
	"crud-t/models"
	"database/sql"

	_ "github.com/lib/pq"
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

func(repo *PostgresRepository) NewUser(ctx context.Context, user *models.User) error{
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id, email, pass, name, lastname, date_brithday) VALUES ($1, $2, $3, $4, $5, $6)", user.Id, user.Email, user.Pass, user.Name, user.LastName, user.DateBrithday) 
	return err
}
