package main

import (
	IWebSocket "ClientWorkerService/internal/websocket"
	fiberwebsocket "ClientWorkerService/internal/websocket/fiber"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"runtime"
)

func main() {
	_ = make([]byte, 10<<30)
	runtime.MemProfileRate = 0

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	fmt.Println("Starting listen!")
	IWebSocket.ListenServer(fiberwebsocket.FiberWebSocket)
	
}
