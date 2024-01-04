package resource

import (
	"slotmathserver/resource/args"

	"github.com/Bofry/host-fasthttp/response"
	"github.com/Bofry/httparg"
	"github.com/valyala/fasthttp"
	"gitlab.com/gaas_math/slotmachine/dcg006"
	"gitlab.com/gaas_module/games"
)

var (
	dcg006_clone = games.NewGames(dcg006.New())
	dcg006_rtp   = "98"
	dcg006_rtps  = map[string]string{
		"GCN": "98",
		"SCN": "98",
	}
	dcg006_gamecode = dcg006_clone.Name()

	dcg006_games = New(
		dcg006_rtps,
		dcg006_clone,
		dcg006_rtp,
	)
)

type DCG006Resource struct{}

func (r *DCG006Resource) Ping(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("PONG"))
}

func (r *DCG006Resource) SPIN(ctx *fasthttp.RequestCtx) {
	argv := args.SpinArgs{}
	httparg.Args(&argv).
		ProcessContent(ctx.PostBody(), MIMETYPE_JSON).
		Validate()

	round, err := dcg006_games.ExecuteSpin(argv)
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
