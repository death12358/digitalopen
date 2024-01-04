package main

import (
	"digitalopen/slotmathserver/failure"
	. "digitalopen/slotmathserver/internal"
	. "digitalopen/slotmathserver/resource"
	"encoding/json"
	"log"

	"github.com/Bofry/config"
	fasthttp "github.com/Bofry/host-fasthttp"
	"github.com/Bofry/host-fasthttp/response"
	"github.com/Bofry/httparg"
	"github.com/shopspring/decimal"
)

//go:generate gen-host-fasthttp-resource
type ResourceManager struct {
	*DefaultResource     `url:"/"`
	*HeathycheckResource `url:"/heathycheck"`
	*SG001Resource       `url:"/sg001"`
}

func main() {
	ctx := AppContext{}
	decimal.DivisionPrecision = 16
	httparg.RegistryService.SetupErrorHandler(func(err error) {
		// log.Printf("err: %v", err.Error())
		failure.ThrowFailureMessage(failure.INVALID_ARGUMENT, err.Error())
		return
	})
	fasthttp.Startup(&ctx).
		Middlewares(
			// fasthttp.UseResourceManager(&ResourceManager{}),
			fasthttp.UseRequestManager(&ResourceManager{}),
			fasthttp.UseXHttpMethodHeader(),
			fasthttp.UseErrorHandler(func(ctx *fasthttp.RequestCtx, err interface{}) {
				fail, ok := err.(*failure.Failure)
				if ok {
					content, _ := json.Marshal(fail)

					log.Printf("uri: %+v err: %v, body: \n%+v", string(ctx.Request.RequestURI()), fail.Message, string(ctx.Request.Body()))
					if content != nil {
						response.Failure(ctx, "application/json", content, fasthttp.StatusBadRequest)
					}
				}
			}),
		).
		ConfigureConfiguration(func(service *config.ConfigurationService) {
			service.
				LoadYamlFile("config.yaml").
				LoadYamlFile("config.${Environment}.yaml").
				LoadEnvironmentVariables("").
				LoadResource(".").
				LoadResource(".conf/${Environment}").
				LoadCommandArguments()
		}).
		Run()
}
