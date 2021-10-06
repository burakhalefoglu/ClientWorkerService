package worker

import (
	kafka "ClientWorkerService/internal/kafka"
	redisAdapter "ClientWorkerService/internal/redis"
	"context"
	"log"

	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

var upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
} 


func Work(ctx *fasthttp.RequestCtx, topic string) {
	err := upgrader.Upgrade(ctx, func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			kafkaErr := kafka.Produce(context.Background(), nil, message, topic)
			if(kafkaErr != nil){
				log.Println(kafkaErr)
				redisAdapter.SetDict(topic, message)
			}
		}
	})

	if err != nil {
		if _, ok := err.(websocket.HandshakeError); ok {
			log.Println(err)
		}
		return
	}
}
