package fastHttpServer

import (
	worker "ClientWorkerService/internal/websocket/fasthttp/worker"
	"log"

	"github.com/valyala/fasthttp"
)

type FasthttpConn struct {
	ConnString string
}

func (f FasthttpConn) ListenServer() { 
	h := requestHandler
	h = fasthttp.CompressHandler(h)
	
	fastHttpServer := fasthttp.Server{
		Name:    "ClientWorkerService",
		Handler: h,
	}
	log.Fatal(fastHttpServer.ListenAndServe(f.ConnString))
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/AdvEventDataModel":
		worker.Work(ctx,"AdvEventDataModel")
	case "/BuyingEventDataModel":
		worker.Work(ctx,"BuyingEventDataModel")
	case "/HardwareIndormationModel":
		worker.Work(ctx,"HardwareIndormationModel")
	case "/InventoryDataModel":
		worker.Work(ctx,"InventoryDataModel")
	case "/LocationModel":
		worker.Work(ctx,"LocationModel")
	case "/ClickDataModel":
		worker.Work(ctx,"ClickDataModel")
	case "/SwipeDataModel":
		worker.Work(ctx,"SwipeDataModel")
	case "/GameSessionEveryLoginDataModel":
		worker.Work(ctx,"GameSessionEveryLoginDataModel")
	case "/LevelBaseSessionDataModel":
		worker.Work(ctx,"LevelBaseSessionDataModel")
	case "/EnemyBaseEveryLoginLevelDatasModel":
		worker.Work(ctx,"EnemyBaseEveryLoginLevelDatasModel")	
	case "/EnemyBaseWithLevelFailDataModel":
		worker.Work(ctx,"EnemyBaseWithLevelFailDataModel")
	case "/ManuelFlowModel":
		worker.Work(ctx,"ManuelFlowModel")
	case "/OfferBehaviorModel":
		worker.Work(ctx,"OfferBehaviorModel")		
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}