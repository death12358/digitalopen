package resource

import (
	"slotmathserver/resource/args"

	"github.com/Bofry/host-fasthttp/response"
	"github.com/Bofry/httparg"
	"github.com/shopspring/decimal"
	"github.com/valyala/fasthttp"
	"gitlab.com/gaas_math/slotmachine/sg003"
	"gitlab.com/gaas_module/games"
)

func init() {
	decimal.DivisionPrecision = 16
}

var (
	sg003_clone = games.NewGames(sg003.New())
	sg003_rtp   = "96"
	sg003_rtps  = map[string]string{
		"GCN": "98",
		"SCN": "98",
	}
	sg003_gamecode = sg003_clone.Name()

	sg003_games = New(
		sg003_rtps,
		sg003_clone,
		sg003_rtp,
	)
)

type SG003Resource struct{}

func (r *SG003Resource) Ping(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("PONG"))
}

func (r *SG003Resource) SPIN(ctx *fasthttp.RequestCtx) {
	argv := args.SpinArgs{}
	httparg.Args(&argv).
		ProcessContent(ctx.PostBody(), MIMETYPE_JSON).
		Validate()

	round, err := sg003_games.ExecuteSpin(argv)
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
