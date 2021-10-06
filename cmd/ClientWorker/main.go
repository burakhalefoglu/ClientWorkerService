package main

import (
	websocketAdapter "ClientWorkerService/internal/websocket"
	fastHttpServer "ClientWorkerService/internal/websocket/fasthttp"
	"runtime"
)

func main() {
	_ = make([]byte, 10<<30) 
	runtime.MemProfileRate = 0 
	
	conn := fastHttpServer.FasthttpConn{
		ConnString : "localhost:8080",
	}
	websocketAdapter.ListenServer(conn)
	//TODO: kafka ile redis direk bağlanabiliyor mu? kontrol etmelisin! Eğer mümkünse bu bağlantı aracılığı ile
	//TODO: kaçan veriler kayıtedilmeli.
}

