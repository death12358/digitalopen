package resource

import (
	"slotmathserver/resource/args"

	"github.com/Bofry/host-fasthttp/response"
	"github.com/Bofry/httparg"
	"github.com/valyala/fasthttp"
	"gitlab.com/gaas_math/slotmachine/dcg014"
	"gitlab.com/gaas_module/games"
)

var (
	dcg014_clone = games.NewGames(dcg014.New())
	dcg014_rtp   = "98"
	dcg014_rtps  = map[string]string{
		"GCN": "98",
		"SCN": "98",
	}
	dcg014_gamecode = dcg014_clone.Name()

	dcg014_games = New(
		dcg014_rtps,
		dcg014_clone,
		dcg014_rtp,
	)
)

type DCG014Resource struct{}

func (r *DCG014Resource) Ping(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("PONG"))
}

func (r *DCG014Resource) SPIN(ctx *fasthttp.RequestCtx) {
	argv := args.SpinArgs{}
	httparg.Args(&argv).
		ProcessContent(ctx.PostBody(), MIMETYPE_JSON).
		Validate()

	round, err := dcg014_games.ExecuteSpin(argv)
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
