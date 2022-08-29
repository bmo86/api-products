package server

import (
	"context"
	"errors"

	"github.com/gorilla/mux"
	"crud-t/websocket"
)


type Config struct {
	Port       	string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface{
	Config() *Config
	hub() *websocket.Hub 
}

type Broker struct{
	config *Config
	router *mux.Router
	hub *websocket.Hub 
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

















