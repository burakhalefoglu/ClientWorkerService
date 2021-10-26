package IWebSocket

import "sync"

type IWebSocket interface {
	ListenServer(wg *sync.WaitGroup)
}

func ListenServer(d IWebSocket, wg *sync.WaitGroup) {
	d.ListenServer(wg)

}