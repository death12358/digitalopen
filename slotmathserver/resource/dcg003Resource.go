package resource

import (
	"slotmathserver/resource/args"

	"github.com/Bofry/host-fasthttp/response"
	"github.com/Bofry/httparg"
	"github.com/valyala/fasthttp"
	"gitlab.com/gaas_math/slotmachine/dcg003"
	"gitlab.com/gaas_module/games"
)

var (
	dcg003_clone = games.NewGames(dcg003.New())
	dcg003_rtp   = "98"
	dcg003_rtps  = map[string]string{
		"GCN": "98",
		"SCN": "98",
	}
	dcg003_gamecode = dcg003_clone.Name()

	dcg003_games = New(
		dcg003_rtps,
		dcg003_clone,
		dcg003_rtp,
	)
)

type DCG003Resource struct{}

func (r *DCG003Resource) Ping(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("PONG"))
}

func (r *DCG003Resource) SPIN(ctx *fasthttp.RequestCtx) {
	argv := args.SpinArgs{}
	httparg.Args(&argv).
		ProcessContent(ctx.PostBody(), MIMETYPE_JSON).
		Validate()

	round, err := dcg003_games.ExecuteSpin(argv)
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
