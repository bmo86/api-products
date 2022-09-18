package repository

import (
	"context"
	"crud-t/models"
)


type Repo interface{
	
	//usuario
	NewUser(ctx context.Context,  user *models.User) error
	GetByIdUser(ctx context.Context, id string) (*models.User, error)
	GetByEmailUser(ctx context.Context, email string) (*models.User, error)
	ListUser(ctx context.Context, page uint64) ([]*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (error)
	DeleteUser(ctx context.Context, id string) error
	//producto
	NewProduct(ctx context.Context, product *models.Product) error
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

func GetByIdUser(ctx context.Context, id string) (*models.User, error){
	return implementation.GetByIdUser(ctx, id)
}

func GetByEmailUser(ctx context.Context, email string) (*models.User, error){
	return implementation.GetByEmailUser(ctx, email)
}

func ListUser(ctx context.Context, page uint64) ([]*models.User, error){
	return implementation.ListUser(ctx, page)
}

func UpdateUser(ctx context.Context,  user *models.User) (error)  {
	return implementation.UpdateUser(ctx, user)
}

func DeleteUser(ctx context.Context, id string) error{
	return implementation.DeleteUser(ctx, id)
}

func NewProduct(ctx context.Context, product *models.Product) (error){
	return implementation.NewProduct(ctx, product)
}
