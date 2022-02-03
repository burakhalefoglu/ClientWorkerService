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
				log.Fatal("Work", "Work", channel, err)
				break
			}
			log.Print("Work", "Work", channel, "Message received")
			var id = []byte(c.Query("clientId"))

			kafkaErr := kafka.Produce(&id, &msg, channel)
			if kafkaErr != nil {
				log.Fatal("Work", "Work_Kafka_error", channel, kafkaErr)
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
				log.Fatal("Work", "Work_Cache_error", channel, err, "veri kaybı uyarısı!!!", val)
				return
			}

			if len(val) > 0 {
				for k, v := range val {
					message := []byte(v)
					kafkaErr := kafka.Produce(&id, &message, channel)
					if kafkaErr == nil {
						_, err := cache.Delete(channel, k)
						if err != nil {
							log.Fatal("Work", "Work_kafkaErr_error",
								channel, err)
						}
					}
				}
			}
		}
		wgGroup.Done()
	}))
}