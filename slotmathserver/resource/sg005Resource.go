package resource

import (
	"slotmathserver/resource/args"

	"github.com/Bofry/host-fasthttp/response"
	"github.com/Bofry/httparg"
	"github.com/shopspring/decimal"
	"github.com/valyala/fasthttp"
	"gitlab.com/gaas_math/slotmachine/sg005"
	"gitlab.com/gaas_module/games"
)

func init() {
	decimal.DivisionPrecision = 16
}

var (
	sg005_clone = games.NewGames(sg005.New())
	sg005_rtp   = "98"
	sg005_rtps  = map[string]string{
		"GCN": "98",
		"SCN": "98",
	}
	sg005_gamecode = sg005_clone.Name()

	sg005_games = New(
		sg005_rtps,
		sg005_clone,
		sg005_rtp,
	)
)

type SG005Resource struct{}

func (r *SG005Resource) Ping(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("PONG"))
}

func (r *SG005Resource) SPIN(ctx *fasthttp.RequestCtx) {
	argv := args.SpinArgs{}
	httparg.Args(&argv).
		ProcessContent(ctx.PostBody(), MIMETYPE_JSON).
		Validate()

	round, err := sg005_games.ExecuteSpin(argv)
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
