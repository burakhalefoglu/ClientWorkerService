package main

import (
	IWebSocket "ClientWorkerService/internal/websocket"
	fiberwebsocket "ClientWorkerService/internal/websocket/fiber"
	"ClientWorkerService/pkg/helper"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"runtime"
)

func main() {
	defer helper.DeleteHealthFile()
	runtime.MemProfileRate = 0

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	fmt.Println("Starting listen!")
	IWebSocket.ListenServer(fiberwebsocket.FiberWebSocket)

}
