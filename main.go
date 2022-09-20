package main

import (
	"context"
	"crud-t/handlers"
	"crud-t/middleware"
	"crud-t/server"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)




func main(){
	err := godotenv.Load(".env")//loading file
	
	if err != nil{
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret: JWT_SECRET,
		Port: PORT,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)

}

func BindRoutes(s server.Server, r *mux.Router){
	
	api := r.PathPrefix("/api/v1").Subrouter()

	api.Use(middleware.CheckAuthMiddleware(s))

	//login && singup
	r.HandleFunc("/", handlers.HandlerHome(s)).Methods(http.MethodGet)
	r.HandleFunc("/singup", handlers.SingUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.HandlerLogin(s)).Methods(http.MethodPost)
	api.HandleFunc("/me", handlers.HandlerMe(s)).Methods(http.MethodGet)
	//user
	api.HandleFunc("/users/{userID}", handlers.UpdateUserHandler(s)).Methods(http.MethodPut)
	api.HandleFunc("/users/{userID}", handlers.DeleteUserHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/users", handlers.ListUserHandler(s)).Methods(http.MethodGet)
	//product
	api.HandleFunc("/product", handlers.NewProductHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/product/{productID}", handlers.UpdateProductHandler(s)).Methods(http.MethodPut)
	api.HandleFunc("/product/{productID}", handlers.DeleteProductHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("//api/v1/products", handlers.ListProductHandler(s)).Methods(http.MethodGet) 

	r.HandleFunc("/ws", s.Hub().HandlerWs)
}