package fastHttpServer

import (
	"ClientWorkerService/internal/websocket/fasthttp/worker"
	redisCache "ClientWorkerService/pkg/redis/redis"
	"log"
	"sync"

	"github.com/valyala/fasthttp"
)
var wg sync.WaitGroup

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

	v,err := trimFirstRune(string(ctx.Path()))
	if err {
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		return
	}
	worker.Work(ctx,&redisCache.RedisCache{
		Client: redisCache.GetClient(),
	}, v)
}

func trimFirstRune(s string) (string, bool) {
    for i := range s {
        if i > 0 {
            return s[i:], false
        }
    }
    return "", true
}