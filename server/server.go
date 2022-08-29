package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"crud-t/database"
	"crud-t/repository"
	"crud-t/websocket"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
)


type Config struct {
	Port       	string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface{
	Config() *Config
	Hub() *websocket.Hub 
}

type Broker struct{
	config *Config
	router *mux.Router
	hub    *websocket.Hub 
}

func (b *Broker) Config() *Config{
	return b.config
}


func (b *Broker) Hub() *websocket.Hub  {
	return b.hub
}

//ctx - encontrar posibles problemas en el codigo 


func NewServer(ctx context.Context, config *Config) (*Broker, error)  {
	if config.Port == ""{
		return nil, errors.New("port is required")
	}

	if config.JWTSecret == ""{
		return nil, errors.New("secretToken is required")
	}

	if config.DatabaseUrl == ""{
		return nil, errors.New("Data base is required") 
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		hub: websocket.NewHub(),
	} 

	return broker, nil


}


func (b *Broker) Start(binder func(s Server, r *mux.Router))  {
	b.router =  mux.NewRouter()
	binder(b, b.router)

	handler := cors.Default().Handler(b.router)
	
	repo, err := database.NewPostgresRepo(b.config.DatabaseUrl)
	if err != nil{
		log.Fatal(err)
	}

	go b.Hub().Run()

	repository.SetRepo(repo)

	log.Println("Starting server on port ",b.config.Port)
	if err := http.ListenAndServe(b.config.Port, handler); err != nil{
		log.Fatal("ListAndServer: ",  err)
	}
	


}
















