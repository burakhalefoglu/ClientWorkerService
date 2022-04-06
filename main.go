package main

import (
	IWebSocket "ClientWorkerService/internal/websocket"
	fiberwebsocket "ClientWorkerService/internal/websocket/fiber"
	"ClientWorkerService/pkg/helper"
	"fmt"
	"log"
	"runtime"

	"github.com/joho/godotenv"

	logger "github.com/appneuroncompany/light-logger"
)

func main() {
	defer helper.DeleteHealthFile()
	runtime.MemProfileRate = 0

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	logger.Log.App = "ClientWorkerService"

	fmt.Println("Starting listen!")
	IWebSocket.ListenServer(fiberwebsocket.FiberWebSocket)

}
