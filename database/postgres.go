package database

import (
	"context"
	"crud-t/models"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepo(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}

func (repo *PostgresRepository) NewUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id, email, pass, name, lastname, date_brithday) VALUES ($1, $2, $3, $4, $5, $6)", user.Id, user.Email, user.Pass, user.Name, user.LastName, user.DateBrithday)
	return err
}

func (repo *PostgresRepository) GetByIdUser(ctx context.Context, id string) (*models.User, error){
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

func (repo *PostgresRepository) GetByEmailUser(ctx context.Context, email string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email, pass FROM users WHERE email= $1", email)

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
		if err = rows.Scan(&user.Id, &user.Email, &user.Pass); err == nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return &user, nil
	}
	return &user, nil
}

func (repo *PostgresRepository) ListUser(ctx context.Context, page uint64) ([]*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email, pass, created_at, name, lastname, date_brithday FROM users LIMIT $1 OFFSET $2", 5, page*5)

	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var users []*models.User

	for rows.Next() {
		var user = models.User{}
		
		if err = rows.Scan(&user.Id, &user.Email, &user.Pass, &user.CreatedAt, &user.Name, &user.LastName, &user.DateBrithday); err == nil {
			users = append(users, &user)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}


func (repo *PostgresRepository) UpdateUser(ctx context.Context,  user *models.User) (error) {
	var query string
	query = "UPDATE users SET  email = $1, pass = $2, name = $3, lastname = $4, date_brithday = $5 WHERE id = $6"
	_, err := repo.db.ExecContext(ctx, query , user.Email, user.Pass, user.Name, user.LastName, user.DateBrithday, user.Id)
	return err
}

func (repo *PostgresRepository) DeleteUser(ctx context.Context, id string) (error){
	_, err := repo.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	return err
}

//product
func (repo *PostgresRepository) NewProduct(ctx context.Context, product *models.Product) error{
	_, err := repo.db.ExecContext(ctx, "INSERT INTO product (id, name, price, stock, stockMin, description, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)", product.Id, product.Name, product.Price, product.Stock, product.StockMin, product.Description, product.UserId)
	return err
}

func (repo *PostgresRepository) UpdateProduct(ctx context.Context, product *models.Product) (error){
	var query string
	query = "UPDATE product SET name = $1, price = $2, stock = $3, stockMin = $4,description = $5 WHERE id = $6"
	_, err := repo.db.ExecContext(ctx, query, product.Name, product.Price, product.Stock, product.StockMin, product.Description)
	return err
}

func (repo *PostgresRepository) DeleteProduct(ctx context.Context, id string) (error){
	_, err := repo.db.ExecContext(ctx, "DELETE FROM product WHERE id = $1", id)
	return err
}

func (repo *PostgresRepository) ListProduct(ctx context.Context, page uint64) ([]*models.Product, error){
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, price, stock, stockMin, description, user_id FROM product LIMIT $1 OFFSET $2",  5, page*5)

	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var products []*models.Product
	
	for rows.Next() {
		var pro = models.Product{}
		
		err = rows.Scan(&pro.Id ,&pro.Name ,&pro.Price ,&pro.Stock ,&pro.StockMin ,&pro.Description ,&pro.UserId)

		if err == nil {
			products = append(products, &pro)
		} 
		
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}








