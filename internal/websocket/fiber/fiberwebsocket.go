package fiberwebsocket

import (
	"ClientWorkerService/internal/websocket/fiber/worker"
	"ClientWorkerService/pkg/helper"
	"ClientWorkerService/pkg/kafka/kafkago"
	"ClientWorkerService/pkg/logger"
	"ClientWorkerService/pkg/logger/logrus_logstash_hook"
	redisCache "ClientWorkerService/pkg/redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"sync"
)

type fiberWs struct {
	Log logger.ILog
}

var FiberWebSocket = &fiberWs{
	Log: &logrus_logstash_hook.Logrus,
}

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
	go worker.Work(&wgGroup, app, "AdvEventDataModel", &redisCache.RedisCache{Client: redisCache.GetClient(), Log: f.Log},
		&kafkago.KafkaGo{Log: f.Log}, f.Log)
	go worker.Work(&wgGroup, app, "BuyingEventDataModel", &redisCache.RedisCache{Client: redisCache.GetClient(), Log: f.Log},
		&kafkago.KafkaGo{Log: f.Log}, f.Log)
	go worker.Work(&wgGroup, app, "HardwareInformationModel", &redisCache.RedisCache{Client: redisCache.GetClient(), Log: f.Log},
		&kafkago.KafkaGo{Log: f.Log}, f.Log)
	go worker.Work(&wgGroup, app, "LocationDataModel", &redisCache.RedisCache{Client: redisCache.GetClient(), Log: f.Log},
		&kafkago.KafkaGo{Log: f.Log}, f.Log)
	go worker.Work(&wgGroup, app, "ScreenSwipeDataModel", &redisCache.RedisCache{Client: redisCache.GetClient(), Log: f.Log},
		&kafkago.KafkaGo{Log: f.Log}, f.Log)
	go worker.Work(&wgGroup, app, "ScreenClickDataModel", &redisCache.RedisCache{Client: redisCache.GetClient(), Log: f.Log},
		&kafkago.KafkaGo{Log: f.Log}, f.Log)
	go worker.Work(&wgGroup, app, "LevelBaseSessionDataModel", &redisCache.RedisCache{Client: redisCache.GetClient(), Log: f.Log},
		&kafkago.KafkaGo{Log: f.Log}, f.Log)
	go worker.Work(&wgGroup, app, "GameSessionEveryLoginDataModel", &redisCache.RedisCache{Client: redisCache.GetClient(), Log: f.Log},
		&kafkago.KafkaGo{Log: f.Log}, f.Log)
	go worker.Work(&wgGroup, app, "InventoryDataModel", &redisCache.RedisCache{Client: redisCache.GetClient(), Log: f.Log},
		&kafkago.KafkaGo{Log: f.Log}, f.Log)
	go worker.Work(&wgGroup, app, "OfferBehaviorModel", &redisCache.RedisCache{Client: redisCache.GetClient(), Log: f.Log},
		&kafkago.KafkaGo{Log: f.Log}, f.Log)
	go worker.Work(&wgGroup, app, "ChurnPredictionResultModel", &redisCache.RedisCache{Client: redisCache.GetClient(), Log: f.Log},
		&kafkago.KafkaGo{Log: f.Log}, f.Log)
	err := app.Listen(helper.ResolvePath("WEBSOCKET_HOST", "WEBSOCKET_PORT"))
	if err != nil {
		f.Log.SendFatalLog("fiberWs", "ListenServer", err.Error())
		return
	}
	wgGroup.Wait()

}
