package fiberwebsocket

import (
	"ClientWorkerService/internal/websocket/fiber/worker"
	"ClientWorkerService/pkg/kafka/kafkago"
	redisCache "ClientWorkerService/pkg/redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
	"os"
	"sync"
)

type fiberWs struct {}

var FiberWebSocket = &fiberWs{}

func (f fiberWs) ListenServer() {
	app := fiber.New()
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	var wgGroup sync.WaitGroup
	wgGroup.Add(11)
	go worker.Work(&wgGroup, app, "AdvEventDataModel", &redisCache.RedisCache{ Client: redisCache.GetClient() }, &kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "BuyingEventDataModel", &redisCache.RedisCache{ Client: redisCache.GetClient() }, &kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "HardwareInformationModel", &redisCache.RedisCache{ Client: redisCache.GetClient() }, &kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "LocationDataModel", &redisCache.RedisCache{ Client: redisCache.GetClient() }, &kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "ScreenSwipeDataModel", &redisCache.RedisCache{ Client: redisCache.GetClient() }, &kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "ScreenClickDataModel", &redisCache.RedisCache{ Client: redisCache.GetClient() }, &kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "LevelBaseSessionDataModel", &redisCache.RedisCache{ Client: redisCache.GetClient() }, &kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "GameSessionEveryLoginDataModel", &redisCache.RedisCache{ Client: redisCache.GetClient() }, &kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "InventoryDataModel", &redisCache.RedisCache{ Client: redisCache.GetClient() }, &kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "OfferBehaviorModel", &redisCache.RedisCache{ Client: redisCache.GetClient() }, &kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "ChurnPredictionResultModel", &redisCache.RedisCache{ Client: redisCache.GetClient() }, &kafkago.KafkaGo{})
	log.Fatal(app.Listen(os.Getenv("WEBSOCKET_CONN")))
	wgGroup.Wait()

}