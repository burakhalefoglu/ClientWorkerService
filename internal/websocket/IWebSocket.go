package IWebSocket

type IWebSocket interface {
	ListenServer()
}

func ListenServer(d IWebSocket) {
	d.ListenServer()

}