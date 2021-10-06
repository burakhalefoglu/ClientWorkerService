package hub

import (
	"fmt"

	"github.com/fasthttp/websocket"
)

type Channel struct {
	Clients []Client
	Name string
}

type Client struct {
	Id         string
	Connection *websocket.Conn
}


func (ps *Channel) AddClient(client Client) {

	ps.Clients = append(ps.Clients, client)
	fmt.Println("adding new client to the list", client.Id, len(ps.Clients))
	payload := []byte("Hello Client ID:" +
		client.Id)

	client.Connection.WriteMessage(1, payload)
}


func (ps *Channel) RemoveClient(client Client) {

	for index, c := range ps.Clients {

		if c.Id == client.Id {
			ps.Clients = append(ps.Clients[:index], ps.Clients[index+1:]...)
		}

	}
}