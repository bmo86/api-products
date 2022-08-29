package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)



var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {return true},
}

type Hub struct{
	Clients 	[]*Client
	Register 	chan *Client
	UnRegister 	chan *Client
	Mutex 		*sync.Mutex
}

func NewHub() *Hub{
	return &Hub{
		Clients: 	make([]*Client, 0),
		Register: 	make(chan *Client),
		UnRegister: make(chan *Client),
		Mutex: 		&sync.Mutex{},
	}
}


func (hub *Hub) Run() {

	for{
		select {
		case cli := <-hub.Register:
			hub.onConnect(cli)
		
		case cli := <-hub.UnRegister: 
			hub.onDisconnects(cli)
		}
	}
	
}

func (h *Hub) BroadCast(msg interface{}, ignore *Client){
	data, _ := json.Marshal(msg)
	for _, cli := range h.Clients{
		if cli != ignore{
			cli.OutBand <- data
		} 
	}
}  

func (hub *Hub) onConnect(c *Client){
	log.Println(" -> Client Connected! ", c.Socket.RemoteAddr())

	hub.Mutex.Lock()
	defer hub.Mutex.Unlock()
	c.Id = c.Socket.RemoteAddr().String()
	hub.Clients = append(hub.Clients, c)
}

func (hub *Hub) onDisconnects (cli *Client){
	log.Println(" -> Client Disconect ", cli.Socket.RemoteAddr())

	cli.Close();
	hub.Mutex.Lock()
	defer hub.Mutex.Unlock()
	i := -1
	for j, c := range hub.Clients{
		if c.Id == cli.Id {
			i = j
			break
		}
	}

	copy(hub.Clients[i:], hub.Clients[i+1:])
	hub.Clients[len(hub.Clients)-1] = nil
	hub.Clients = hub.Clients[:len(hub.Clients)-1]

}



func (h *Hub) handlerWs(w http.ResponseWriter, r *http.Request){

	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil { 
		log.Println(err)
		http.Error(w, "Error Upgrading connection", http.StatusInternalServerError)

		cli := NewClient(h, socket)
		h.Register <- cli
		go cli.Write()
	}
}

