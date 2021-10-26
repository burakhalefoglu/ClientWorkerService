package main

import (
	IWebSocket "ClientWorkerService/internal/websocket"
	fastHttpServer "ClientWorkerService/internal/websocket/fasthttp"
	"runtime"
	"sync"
)

func main() {
	_ = make([]byte, 10<<30) 
	runtime.MemProfileRate = 0

	var wg *sync.WaitGroup

	IWebSocket.ListenServer(fastHttpServer.FastHttpServer, wg)
	wg.Wait()
}

