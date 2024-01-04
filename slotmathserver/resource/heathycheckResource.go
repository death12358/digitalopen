package resource

import (
	"github.com/Bofry/host-fasthttp/response"
	"github.com/valyala/fasthttp"
)

type HeathycheckResource struct{}

func (r *HeathycheckResource) HEAD(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("OK"))
}

func (r *HeathycheckResource) GET(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("HEAD"))
}

func (r *HeathycheckResource) PING(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("PONG"))
}
