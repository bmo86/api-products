package database

import (
	"context"
	"crud-t/models"
	"database/sql"
	"log"

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

func(repo *PostgresRepository) GetByIdUser(ctx context.Context, id string) (*models.User, error){
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil{
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil{
		return nil, err
	}

	return &user, nil

}


func (repo *PostgresRepository) GetByEmailUser(ctx context.Context, email string) (*models.User, error)  {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email, pass FROM users WHERE email= $1", email)

	defer func() {
		err = rows.Close()
		if err != nil{
			log.Fatal(err)
		}
	}()
	var user = models.User{}

	for rows.Next(){
		if err = rows.Scan(&user.Id, &user.Email, &user.Pass); err == nil{
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil{
		return &user, nil
	}
	return &user, nil
}

