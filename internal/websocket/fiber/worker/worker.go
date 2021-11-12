package worker

import (
	IKafka "ClientWorkerService/pkg/kafka"
	ICache "ClientWorkerService/pkg/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"log"
	"sync"
)


func Work(wgGroup *sync.WaitGroup,
	app *fiber.App,
	channel string,
	cache ICache.ICache,
	kafka IKafka.IKafka) {
	app.Get("/" + channel, websocket.New(func(c *websocket.Conn) {

		var (
			msg []byte
			err error
		)
		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)
			var id = []byte(c.Query("clientId"))

			kafkaErr := kafka.Produce(&id, &msg, channel)
			if kafkaErr != nil {
				log.Println(kafkaErr)
				_, err := cache.Add(channel, map[string]interface{}{
					uuid.New().String(): msg,
				})
				if err != nil {
					return
				}
				continue
			}

			val, err := cache.Get(channel)
			if err != nil {
				log.Fatal(err)
				//! veri kaybı olma ihtimali oluşuyor!!!
				return
			}

			if len(val) > 0 {
				for k, v := range val {
					message := []byte(v)
					kafkaErr := kafka.Produce(&id, &message, channel)
					if kafkaErr == nil {
						result, err := cache.Delete(channel, k)
						if err != nil {
							log.Println(result)
						}
						log.Println(result)
					}
				}
			}
		}
		wgGroup.Done()
	}))
}