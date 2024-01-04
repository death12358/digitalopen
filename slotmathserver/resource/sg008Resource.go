package resource

import (
	"slotmathserver/resource/args"

	"github.com/Bofry/host-fasthttp/response"
	"github.com/Bofry/httparg"
	"github.com/shopspring/decimal"
	"github.com/valyala/fasthttp"
	"gitlab.com/gaas_math/slotmachine/sg008"
	"gitlab.com/gaas_module/games"
)

func init() {
	decimal.DivisionPrecision = 16
}

var (
	sg008_clone = games.NewGames(sg008.New())
	sg008_rtp   = "98"
	sg008_rtps  = map[string]string{
		"GCN": "98",
		"SCN": "98",
	}
	sg008_gamecode = sg008_clone.Name()

	sg008_games = New(
		sg008_rtps,
		sg008_clone,
		sg008_rtp,
	)
)

type SG008Resource struct{}

func (r *SG008Resource) Ping(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("PONG"))
}

func (r *SG008Resource) SPIN(ctx *fasthttp.RequestCtx) {
	argv := args.SpinArgs{}
	httparg.Args(&argv).
		ProcessContent(ctx.PostBody(), MIMETYPE_JSON).
		Validate()

	round, err := sg008_games.ExecuteSpin(argv)
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
