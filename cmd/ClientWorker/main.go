package main

import (
	IWebSocket "ClientWorkerService/internal/websocket"
	fiberwebsocket "ClientWorkerService/internal/websocket/fiber"
	"fmt"
	"runtime"
)

func main() {
	_ = make([]byte, 10<<30) 
	runtime.MemProfileRate = 0

	fmt.Println("Starting listen!")
	IWebSocket.ListenServer(fiberwebsocket.FiberWebSocket)
}

