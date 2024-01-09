package sg006

import (
	"strconv"

	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/slot/sg006/vs"
	"github.com/shopspring/decimal"
)

// SpinJackpot - spin VS
func (s *SG006) SpinVSFlow(rtp string, bet decimal.Decimal, round *games.Rounds, times int) (*games.Rounds, error) {
	unitbet := round.Result["0"].Pickem[0]
	vs_game := sg006_gameplay[unitbet].vs_game[rtp]
	bet = round.Result["0"].Bet

	// default 8 times
	count := times
	stage := int64(count) // 1,2,3,......
	FGinVSTimes := 0
	for i := 0; i < times; i++ {
		round.Stages++
		stage--
		id := strconv.Itoa(int(round.Stages))
		var vs_sym games.Symbol
		if FGinVSTimes == 0 {
			vs_sym = vs_game.PickReward()
		} else if FGinVSTimes > 0 {
			vs_sym = sg006vs_NoBattle[rtp].PickReward()
		}
		point := vs_game.GetVSPay(vs_sym)
		if vs.TriggerFGinVS(vs_sym) {
			FGinVSTimes++
		}
		res := &games.Records{
			Id:         id,
			Brand:      round.Brand,
			Username:   round.Username,
			Case:       games.NotStartedYet,
			Pickem:     []string{""},
			Symbols:    []string{vs.SymbolString[vs_sym]},
			Multiplier: point,
			Point:      point.Mul(bet),
		}
		if FGinVSTimes > 0 {
			res.Extra = append(res.Extra, "FGTriggers")
		}

		res.Stages = stage

		//var ShowNumbers games.ExtraSG006

		round.Result[id] = res
		round.TotalPoint = round.TotalPoint.Add(res.Point)
	}
	/*
		// round.Stages = 1
		round.Result["1"].Case = games.Lose
		if !round.Result["1"].Point.IsZero() {
			round.Result["1"].Case = games.Win
		}
	*/
	if FGinVSTimes >= 1 {
		// RemainRetrigger > 0 時，再次暫時更改Rest，待玩家選擇
		round.Status = round.Status.Push(games.BonusGame2)
		//round.Position = games.BonusGame2
		return s.FGFlow(rtp, bet, round)
	}
	return round, nil
}
