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
	

	r.Use(middleware.CheckAuthMiddleware(s))

	r.HandleFunc("/", handlers.HandlerHome(s)).Methods(http.MethodGet)
	r.HandleFunc("/singup", handlers.SingUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.HandlerLogin(s)).Methods(http.MethodPost)
	r.HandleFunc("/me", handlers.HandlerMe(s)).Methods(http.MethodGet)
	r.HandleFunc("/updateUsers", handlers.UpdateUserHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/users", handlers.ListUserHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/ws", s.Hub().HandlerWs)
}