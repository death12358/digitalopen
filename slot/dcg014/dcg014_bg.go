package dcg014

import (
	"strconv"

	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/slots"
	"github.com/shopspring/decimal"
)

const bg_SBmatch_def = 3

func (s *DCG014) IsBonusWin(unitbet string, reel *games.ReelStrips) bool {
	_, count := reel.CalcSymbolMatches(slots.SB)
	return count >= bg_SFmatch_def
}

func SymbolSliceToStringslice(symbol []games.Symbol) []string {
	result := make([]string, len(symbol))
	for i, val := range symbol {
		result[i] = val.String()
	}
	return result
}

func (s *DCG014) EggFlow(bet decimal.Decimal, unitbet string, round *games.Rounds) (games.Rounds, error) {
	bonus_game := dcg014_gameplay[unitbet].bg_game
	round.Status = round.Status.Push(games.Bonus)
	round.Position = round.Position.Push(games.Bonus)
	process := bonus_game.PickProcess()
	score := bonus_game.CalcAllPoints(process, bet)
	ID := strconv.Itoa(len(round.Result))
	Process := SymbolSliceToStringslice(process)
	record := &games.Records{
		Id:         ID,
		Brand:      round.Brand,
		Username:   round.Username,
		Case:       games.NotStartedYet,
		Stages:     round.Stages,
		Extra:      []string{""},
		Pickem:     []string{unitbet},
		Symbols:    Process,
		Multiplier: score.Div(bet),
		Bet:        bet,
		Point:      score,
	}
	round.Result[ID] = record
	round.TotalPoint = round.TotalPoint.Add(record.Point)

	return *round, nil
}
