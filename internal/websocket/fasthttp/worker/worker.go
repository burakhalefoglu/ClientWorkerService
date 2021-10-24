package worker

import (
	Ikafka "ClientWorkerService/pkg/kafka"
	cache "ClientWorkerService/pkg/redis"
	"github.com/google/uuid"
	"log"

	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

var upgrade = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
} 


func Work(ctx *fasthttp.RequestCtx, cache cache.ICache, kafka Ikafka.IKafka , topic string) {
	err := upgrade.Upgrade(ctx, func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			req := &ctx.Request
			id := req.URI().QueryArgs().Peek("clientId")
			kafkaErr := kafka.Produce(&id, &message, topic)
			if kafkaErr != nil {
				log.Println(kafkaErr)
				cache.Add(topic, map[string]interface{}{
					uuid.New().String(): message,
				})
				return
			}

			val, err := cache.Get(topic)
			if err != nil {
				log.Fatal(err)
				//! veri kaybı olma ihtimali oluşuyor!!!
				return
			}
			if len(val) > 0 {
				for k, v := range val {
					message := []byte(v)
					kafkaErr := kafka.Produce(&id, &message, topic)
					if kafkaErr == nil {
						result, err := cache.Delete(topic, k)
						if err != nil {
							log.Println(result)
						}
						log.Println(result)
					}
				}
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
