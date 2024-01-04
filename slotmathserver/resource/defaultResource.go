package resource

import (
	"encoding/json"

	"github.com/Bofry/host-fasthttp/response"
	"github.com/valyala/fasthttp"
)

var (
	gameList = map[string]string{
		"dcg001": "/dcg001",
		"dcg002": "/dcg002",
		"dcg003": "/dcg003",
		"dcg006": "/dcg006",
		"dcg008": "/dcg008",
		"dcg014": "/dcg014",
		"sg001":  "/sg001",
		"sg002":  "/sg002",
		"sg003":  "/sg003",
		"sg005":  "/sg005",
		"sg006":  "/sg006",
		"sg008":  "/sg008",
		"sg009":  "/sg009",
	}
)

type DefaultResource struct{}

func (r *DefaultResource) GET(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("OK"))
}

func (r *DefaultResource) Ping(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("PONG"))
}

func (r *DefaultResource) LIST(ctx *fasthttp.RequestCtx) {
	body, _ := json.Marshal(gameList)
	response.Success(ctx, MIMETYPE_JSON, body)
}
