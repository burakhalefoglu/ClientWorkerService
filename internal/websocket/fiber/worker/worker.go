package worker

import (
	IKafka "ClientWorkerService/pkg/kafka"
	ICache "ClientWorkerService/pkg/redis"
	"sync"

	"github.com/appneuroncompany/light-logger/clogger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

func Work(wgGroup *sync.WaitGroup,
	app *fiber.App,
	channel string,
	cache ICache.ICache,
	kafka IKafka.IKafka) {
	app.Get("/"+channel, websocket.New(func(c *websocket.Conn) {

		var (
			msg []byte
			err error
		)
		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				clogger.Error(&map[string]interface{}{
					"error message: ": err,
					"channel: ":       channel,
				})
				break
			}
			clogger.Info(&map[string]interface{}{ // use it wherever you want
				"message: ": "Message received",
				"channel: ": channel,
			})
			var id = []byte(c.Query("clientId"))

			kafkaErr := kafka.Produce(&id, &msg, channel)
			if kafkaErr != nil {
				clogger.Error(&map[string]interface{}{
					"kafka error: ": kafkaErr,
					"channel: ":     channel,
				})
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
				clogger.Error(&map[string]interface{}{
					"cache error: ": err,
					"channel: ":     channel,
				})
				return
			}

			if len(val) > 0 {
				for k, v := range val {
					message := []byte(v)
					kafkaErr := kafka.Produce(&id, &message, channel)
					if kafkaErr == nil {
						_, err := cache.Delete(channel, k)
						if err != nil {
							clogger.Error(&map[string]interface{}{
								"kafkaErr error: ": err,
								"channel: ":        channel,
							})
						}
					}
				}
			}
		}
		wgGroup.Done()
	}))
}
