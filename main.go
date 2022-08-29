package main

import (
	"context"
	"crud-t/server"
	"log"
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


}