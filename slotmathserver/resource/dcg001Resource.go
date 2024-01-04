package resource

import (
	"slotmathserver/resource/args"

	"games"

	"github.com/Bofry/host-fasthttp/response"
	"github.com/Bofry/httparg"
	"github.com/shopspring/decimal"
	"github.com/valyala/fasthttp"
	"gitlab.com/gaas_math/slotmachine/dcg001"
)

func init() {
	decimal.DivisionPrecision = 16
}

var (
	dcg001_clone = games.NewGames(dcg001.New())
	dcg001_rtp   = "98"
	dcg001_rtps  = map[string]string{
		"GCN": "98",
		"SCN": "98",
	}
	dcg001_gamecode = dcg001_clone.Name()

	dcg001_games = New(
		dcg001_rtps,
		dcg001_clone,
		dcg001_rtp,
	)
)

type DCG001Resource struct{}

func (r *DCG001Resource) Ping(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("PONG"))
}

func (r *DCG001Resource) SPIN(ctx *fasthttp.RequestCtx) {
	argv := args.SpinArgs{}
	httparg.Args(&argv).
		ProcessContent(ctx.PostBody(), MIMETYPE_JSON).
		Validate()

	round, err := dcg001_games.ExecuteSpin(argv)
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
