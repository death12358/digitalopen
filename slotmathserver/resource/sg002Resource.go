package resource

import (
	"slotmathserver/resource/args"

	"github.com/Bofry/host-fasthttp/response"
	"github.com/Bofry/httparg"
	"github.com/shopspring/decimal"
	"github.com/valyala/fasthttp"
	"gitlab.com/gaas_math/slotmachine/sg002"
	"gitlab.com/gaas_module/games"
)

func init() {
	decimal.DivisionPrecision = 16
}

var (
	sg002_clone = games.NewGames(sg002.New())
	sg002_rtp   = "98"
	sg002_rtps  = map[string]string{
		"GCN": "98",
		"SCN": "98",
	}
	sg002_gamecode = sg002_clone.Name()

	sg002_games = New(
		sg002_rtps,
		sg002_clone,
		sg002_rtp,
	)
)

type SG002Resource struct{}

func (r *SG002Resource) Ping(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("PONG"))
}

func (r *SG002Resource) SPIN(ctx *fasthttp.RequestCtx) {
	argv := args.SpinArgs{}
	httparg.Args(&argv).
		ProcessContent(ctx.PostBody(), MIMETYPE_JSON).
		Validate()

	round, err := sg002_games.ExecuteSpin(argv)
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
