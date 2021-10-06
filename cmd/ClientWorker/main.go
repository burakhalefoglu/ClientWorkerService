package main

import (
	websocketAdapter "ClientWorkerService/internal/websocket"
	fastHttpServer "ClientWorkerService/internal/websocket/fasthttp"
)

func main() {

	conn := fastHttpServer.FasthttpConn{
		ConnString : "localhost:8080",
	}
	websocketAdapter.ListenServer(conn)
}

