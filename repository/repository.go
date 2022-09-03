package repository

import (
	"context"
	"crud-t/models"
)


type Repo interface{
	
	NewUser(ctx context.Context,  user *models.User) error
	Close() error

}

var implementation Repo

//set repo
func SetRepo(repo Repo)  {
	implementation = repo
}

//close connection db
func Close() error {
	return implementation.Close()
}

func NewUser(ctx context.Context, user models.User) error{
	return implementation.NewUser(ctx, &user)
}
