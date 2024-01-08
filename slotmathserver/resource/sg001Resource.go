package resource

import (
	"encoding/json"
	"errors"

	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/sg001"
	"github.com/death12358/digitalopen/slotmathserver/resource/args"

	"github.com/Bofry/host-fasthttp/response"
	"github.com/Bofry/httparg"
	"github.com/shopspring/decimal"
	"github.com/valyala/fasthttp"
)

func init() {
	decimal.DivisionPrecision = 16
}

var (
	sg001_clone = games.NewGames(sg001.New())
	sg001_rtp   = "98"
	sg001_rtps  = map[string]string{
		"GCN": "98",
		"SCN": "98",
	}
	sg001_gamecode = sg001_clone.Name()

	sg001_games = NewJP(
		sg001_rtps,
		sg001_clone,
		sg001_rtp,
		&sg001_jp_def,
		sg001_jppools_sorts,
		sg001_jp_mapping,
	)
)

// jp def
var (
	sg001_jp_def = []string{"POOLCODE.Mini", "POOLCODE.Minor", "POOLCODE.Major", "POOLCODE.Grand"}

	sg001_jppools_sorts = []string{"POOLCODE.Mini", "POOLCODE.Minor", "POOLCODE.Major", "POOLCODE.Grand"}

	sg001_jp_mapping = map[string]string{"0": "Mini", "1": "Minor", "2": "Major", "3": "Grand"}
)

type SG001Resource struct{}

func (r *SG001Resource) Ping(ctx *fasthttp.RequestCtx) {
	response.Success(ctx, "text/plain", []byte("PONG"))
}

func (r *SG001Resource) SPIN(ctx *fasthttp.RequestCtx) {
	argv := args.SpinArgs{}
	httparg.Args(&argv).
		ProcessContent(ctx.PostBody(), MIMETYPE_JSON).
		Validate()

	round, err := sg001_games.ExecuteSpin(argv)
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

func (r *SG001Resource) POOL(ctx *fasthttp.RequestCtx) {
	argv := args.JackpotArgs{}
	httparg.Args(&argv).
		ProcessContent(ctx.PostBody(), MIMETYPE_JSON).
		Validate()
	rtp := getRTP("SG001", argv.Currency, sg001_rtp)

	bet, err := decimal.NewFromString(argv.Bet.String())
	if err != nil {
		response.Failure(ctx, MIMETYPE_TEXT, []byte(err.Error()), fasthttp.StatusInternalServerError)
		return
	}

	basepool := sg001_clone.BaseJackpot("8", rtp, bet)

	cache, err := slot_Scriptor.JackpotPeeks("0", argv.Brand+sg001_gamecode, argv.Currency)
	if err != nil {
		response.Failure(ctx, MIMETYPE_TEXT, []byte(err.Error()), fasthttp.StatusInternalServerError)
		return
	}
	if len(cache.Pools) == 0 {
		resp, err := json.Marshal(basepool)
		if err != nil {
			response.Failure(ctx, MIMETYPE_TEXT, []byte(err.Error()), fasthttp.StatusInternalServerError)
			return
		}
		// log.Printf("PICKEM: %s", string(resp))
		response.Success(ctx, MIMETYPE_JSON, resp)
		return
	}

	for i, v := range basepool {
		pool, err := decimal.NewFromString(cache.Pools[sg001_jp_def[i]].String())
		if err != nil {
			response.Failure(ctx, MIMETYPE_TEXT, []byte(err.Error()), fasthttp.StatusInternalServerError)
			return
		}
		basepool[i] = v.Add(pool)
	}

	resp, err := json.Marshal(basepool)
	if err != nil {
		response.Failure(ctx, MIMETYPE_TEXT, []byte(err.Error()), fasthttp.StatusInternalServerError)
		return
	}
	// log.Printf("PICKEM: %s", string(resp))
	response.Success(ctx, MIMETYPE_JSON, resp)
	return
}

func (r *SG001Resource) JPPOOLS(ctx *fasthttp.RequestCtx) {
	argv := args.PoolArgs{}
	httparg.Args(&argv).
		ProcessContent(ctx.PostBody(), MIMETYPE_JSON).
		Validate()

	jp, err := slot_Scriptor.JackpotCurrencyPeeks("0", argv.Brand+sg001_gamecode)
	if err != nil {
		response.Failure(ctx, MIMETYPE_TEXT, []byte(err.Error()), fasthttp.StatusInternalServerError)
		return
	}

	pools := jp.Pools
	poolResp := map[string][]json.Number{}
	if len(pools) != 0 {
		for i, v := range pools {
			poolResp[i] = make([]json.Number, 4)
			for v_i, v_v := range sg001_jppools_sorts {
				poolResp[i][v_i] = v[v_v]
			}
		}

	}

	resp, err := json.Marshal(poolResp)
	if err != nil {
		response.Failure(ctx, MIMETYPE_TEXT, []byte(err.Error()), fasthttp.StatusInternalServerError)
		return
	}

	response.Success(ctx, MIMETYPE_JSON, resp)
	return
}

func (r *SG001Resource) JPPoolPush(unitbet, jpname, currency string, totalBet decimal.Decimal) error {
	rtp := getRTP("SG001", currency, sg001_rtp)
	jp_roll_rate := sg001_clone.CalcRollRate(unitbet, rtp, totalBet)
	_, err := slot_Scriptor.JackpotPush("0", jpname, currency, "Mini", jp_roll_rate[0])
	_, err = slot_Scriptor.JackpotPush("0", jpname, currency, "Minor", jp_roll_rate[1])
	_, err = slot_Scriptor.JackpotPush("0", jpname, currency, "Major", jp_roll_rate[2])
	_, err = slot_Scriptor.JackpotPush("0", jpname, currency, "Grand", jp_roll_rate[3])

	if err != nil {
		return err
	}

	return nil
}

func (r *SG001Resource) JPCalc(rounds *games.Rounds) (*games.Rounds, error) {
	jp_result, ok := rounds.Result["1"]
	if !ok {
		return nil, errors.New("JACKPOT_RESULT_NOT_FOUND")
	}
	symbol := jp_result.Symbols[0]
	if symbol == "0" || symbol == "1" {
		return rounds, nil
	}

	symbol_key := sg001_jp_mapping[symbol]
	pool, err := slot_Scriptor.JackpotSweep("0", rounds.Brand+sg001_gamecode, rounds.Currency, symbol_key, rounds.TotalBet)
	if err != nil {
		return nil, err
	}
	if pool.Jackpot_Pool.IsZero() {
		return rounds, nil
	}

	rounds.Result["1"].Point = rounds.Result["1"].Point.Add(pool.Jackpot_Pool)
	rounds.Result["1"].Multiplier = rounds.Result["1"].Point.Div(rounds.TotalBet)

	return rounds, nil
}
