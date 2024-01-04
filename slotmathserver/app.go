package main

import (
	"encoding/json"
	"log"
	"slotmathserver/failure"
	. "slotmathserver/internal"
	. "slotmathserver/resource"

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
	*SG002Resource       `url:"/sg002"`
	*SG003Resource       `url:"/sg003"`
	*SG005Resource       `url:"/sg005"`
	*SG006Resource       `url:"/sg006"`
	*SG008Resource       `url:"/sg008"`
	*SG009Resource       `url:"/sg009"`
	*DCG001Resource      `url:"/dcg001"`
	*DCG002Resource      `url:"/dcg002"`
	*DCG003Resource      `url:"/dcg003"`
	*DCG006Resource      `url:"/dcg006"`
	*DCG008Resource      `url:"/dcg008"`
	*DCG014Resource      `url:"/dcg014"`
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
			fasthttp.UseResourceManager(&ResourceManager{}),
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
