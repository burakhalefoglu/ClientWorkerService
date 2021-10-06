package websocketAdapter

type WebsocketAdapter interface {
	ListenServer()
}

func ListenServer(d WebsocketAdapter) {
	d.ListenServer()

}