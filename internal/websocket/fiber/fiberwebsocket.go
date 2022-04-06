package fiberwebsocket

import (
	"ClientWorkerService/internal/websocket/fiber/worker"
	"ClientWorkerService/pkg/helper"
	"ClientWorkerService/pkg/kafka/kafkago"
	redisCache "ClientWorkerService/pkg/redis/redis"
	"log"
	"sync"
)

type fiberWs struct {
}

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
	helper.CreateHealthFile()
	go worker.Work(&wgGroup, app, "AdvEventDataModel", &redisCache.RedisCache{Client: redisCache.GetClient()},
		&kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "BuyingEventDataModel", &redisCache.RedisCache{Client: redisCache.GetClient()},
		&kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "HardwareInformationModel", &redisCache.RedisCache{Client: redisCache.GetClient()},
		&kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "LocationDataModel", &redisCache.RedisCache{Client: redisCache.GetClient()},
		&kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "ScreenSwipeDataModel", &redisCache.RedisCache{Client: redisCache.GetClient()},
		&kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "ScreenClickDataModel", &redisCache.RedisCache{Client: redisCache.GetClient()},
		&kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "LevelBaseSessionDataModel", &redisCache.RedisCache{Client: redisCache.GetClient()},
		&kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "GameSessionEveryLoginDataModel", &redisCache.RedisCache{Client: redisCache.GetClient()},
		&kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "InventoryDataModel", &redisCache.RedisCache{Client: redisCache.GetClient()},
		&kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "OfferBehaviorModel", &redisCache.RedisCache{Client: redisCache.GetClient()},
		&kafkago.KafkaGo{})
	go worker.Work(&wgGroup, app, "ChurnPredictionResultModel", &redisCache.RedisCache{Client: redisCache.GetClient()},
		&kafkago.KafkaGo{})
	err := app.Listen("localhost:8080")
	if err != nil {
		log.Fatal("fiberWs", "ListenServer", err.Error())
		return
	}
	log.Print("main", "client web server started")
	wgGroup.Wait()

}
