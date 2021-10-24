package main

import (
	"ClientWorkerService/internal/websocket"
	"ClientWorkerService/internal/websocket/fasthttp"
	"runtime"
)

func main() {
	_ = make([]byte, 10<<30) 
	runtime.MemProfileRate = 0 
	
	conn := fastHttpServer.FasthttpConn{
		ConnString: "localhost:8080",
	}
	websocketAdapter.ListenServer(conn)
}

