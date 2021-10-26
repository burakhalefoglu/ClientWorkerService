package fastHttpServer

import (
	"ClientWorkerService/internal/websocket/fasthttp/worker"
	"ClientWorkerService/pkg/kafka/confluent"
	redisCache "ClientWorkerService/pkg/redis/redis"
	"log"
	"sync"

	"github.com/valyala/fasthttp"
)

type FasthttpConn struct {
	ConnString string
}

var FastHttpServer = &FasthttpConn{
	ConnString : "localhost:8080",
}
var wgGroup *sync.WaitGroup

func assignAndGetWaitGroup(group *sync.WaitGroup) *sync.WaitGroup{
	wgGroup := group
	return wgGroup
}
func (f FasthttpConn) ListenServer(wg *sync.WaitGroup) {
	assignAndGetWaitGroup(wg)
	h := requestHandler
	h = fasthttp.CompressHandler(h)

	fastHttpServer := fasthttp.Server{
		Name:    "ClientWorkerService",
		Handler: h,
	}
	log.Fatal(fastHttpServer.ListenAndServe(f.ConnString))
}

func requestHandler(ctx *fasthttp.RequestCtx) {

	v,err := trimFirstRune(string(ctx.Path()))
	if err {
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		return
	}
	wgGroup.Add(1)
	worker.Work(ctx,&redisCache.RedisCache{
		Client: redisCache.GetClient(),
	},
	&confluent.Kafka{},
	v)
	wgGroup.Done()
}

func trimFirstRune(s string) (string, bool) {
    for i := range s {
        if i > 0 {
            return s[i:], false
        }
    }
    return "", true
}