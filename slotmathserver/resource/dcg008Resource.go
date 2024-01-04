package resource

import (
	"slotmathserver/resource/args"

	"github.com/Bofry/host-fasthttp/response"
	"github.com/Bofry/httparg"
	"github.com/valyala/fasthttp"
	"gitlab.com/gaas_math/slotmachine/dcg008"
	"gitlab.com/gaas_module/games"
)

var (
	dcg008_clone = games.NewGames(dcg008.New())
	dcg008_rtp   = "96"
	dcg008_rtps  = map[string]string{
		"GCN": "98",
		"SCN": "98",
	}
	dcg008_gamecode = dcg008_clone.Name()

	dcg008_games = New(
		dcg008_rtps,
		dcg008_clone,
		dcg008_rtp,
	)
)

type DCG008Resource struct{}

func (r *DCG008Resource) Ping(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("PONG"))
}

func (r *DCG008Resource) SPIN(ctx *fasthttp.RequestCtx) {
	argv := args.SpinArgs{}
	httparg.Args(&argv).
		ProcessContent(ctx.PostBody(), MIMETYPE_JSON).
		Validate()

	round, err := dcg008_games.ExecuteSpin(argv)
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
