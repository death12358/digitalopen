package resource

import (
	"slotmathserver/resource/args"

	"github.com/Bofry/host-fasthttp/response"
	"github.com/Bofry/httparg"
	"github.com/shopspring/decimal"
	"github.com/valyala/fasthttp"
	"gitlab.com/gaas_math/slotmachine/sg006"
	"gitlab.com/gaas_module/games"
)

func init() {
	decimal.DivisionPrecision = 16
}

var (
	sg006_clone = games.NewGames(sg006.New())
	sg006_rtp   = "98"
	sg006_rtps  = map[string]string{
		"GCN": "98",
		"SCN": "98",
	}
	sg006_gamecode = sg006_clone.Name()

	sg006_games = New(
		sg006_rtps,
		sg006_clone,
		sg006_rtp,
	)
)

type SG006Resource struct{}

func (r *SG006Resource) Ping(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("PONG"))
}

func (r *SG006Resource) SPIN(ctx *fasthttp.RequestCtx) {
	argv := args.SpinArgs{}
	httparg.Args(&argv).
		ProcessContent(ctx.PostBody(), MIMETYPE_JSON).
		Validate()

	round, err := sg006_games.ExecuteSpin(argv)
	if err != nil {
		response.Failure(ctx, MIMETYPE_TEXT, []byte(err.Error()), fasthttp.StatusInternalServerError)
		return
	}

	resp, err := round.Marshal()
	if err != nil {
		response.Failure(ctx, MIMETYPE_TEXT, []byte(err.Error()), fasthttp.StatusInternalServerError)
		return
	}
	response.Success(ctx, MIMETYPE_JSON, resp)
	return
}
