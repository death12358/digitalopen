package sg001

import (
	"errors"

	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/random"

	"github.com/shopspring/decimal"
)

// SpinNormalGame
func (s *SG001) SpinNormalGame(unitbet, rtp string, position []int, round *games.Rounds) (games.Rounds, *games.ReelStrips) {
	// 取得獎圖長度以已檢查
	// // 取得 unitbet 遊戲
	// game_math, ok := sg001_gameplay[unitbet]
	// if !ok {
	// 	return nil, nil, errors.New("unitbet not found: " + unitbet)
	// }
	game_math, _ := sg001_gameplay[unitbet]

	rtps := games.RTPs(rtp)

	// 取得盤面
	reelStrips := game_math.ng_game.ContiguousReelStrips(rtps, position)
	show_reelStrips := game_math.ng_game.ShowReelStrips(rtps, position, 1, 1).InvertRegularXYAxis().Strings()

	// 初始化 games.Records 結構
	record := &games.Records{
		Id:       "0",
		Brand:    round.Brand,
		Username: round.Username,
		Case:     games.Lose,
		Pickem:   []string{unitbet},
		Symbols:  show_reelStrips,
		Bet:      round.TotalBet,
	}

	// 換至 NG 遊戲 Lightning 獎圖
	round.Result["0"] = game_math.bgLightning_game.ReplaceNGSymbol(record)

	// 除以 unitbet
	// unit_bet, _ := decimal.NewFromString(unitbet)
	// bet := record.Bet.Mul(unit_bet)

	// 計算盤面
	ng_result := game_math.ng_game.CalcWayReel(reelStrips, round.TotalBet)
	if ng_result.IsWinGreaterThanZero() {
		record.Case = games.Win
		record.Multiplier = ng_result.TotalWin.Div(round.TotalBet)
		record.Point = ng_result.TotalWin
	}

	round.Status = round.Status.Push(record.Case)
	round.Result["0"] = record
	round.TotalPoint = round.TotalPoint.Add(record.Point)

	return *round, &reelStrips
}

// GetNGReelsLen - 取得NG滾輪長度
func (s *SG001) GetNGReelsLen(unitbet, rtp string) ([]int, error) {
	// 取得 unitbet 遊戲
	game_math, ok := sg001_gameplay[unitbet]
	if !ok {
		return nil, errors.New("unitbet not found: " + unitbet)
	}
	rtps := games.RTPs(rtp)
	return game_math.ng_game.GetReelsLen(rtps), nil
}

// SpinNGReelStripts - spin NG reel strips
func (s *SG001) SpinNGReelStripts(unitbet, rtp string) (*games.ReelStrips, []string, error) {
	game_math, ok := sg001_gameplay[unitbet]
	if !ok {
		return nil, nil, errors.New("unitbet not found: " + unitbet)
	}
	rtps := games.RTPs(rtp)
	nglen := game_math.ng_game.GetReelsLen(rtps)
	pos := random.Intsn(nglen)
	// // FG Test
	// pos = []int{0, 5, 60, 23, 0}
	// // BG Test
	// pos = []int{1, 25, 22, 4, 33}

	reelStrips := game_math.ng_game.ContiguousReelStrips(rtps, pos)
	show_reelStrips := game_math.ng_game.ShowReelStrips(rtps, pos, 1, 1).InvertRegularXYAxis().Strings()
	// fmt.Printf("pos: %v", show_reelStrips)
	// show_reelStrips := ng_game.SpinNormal(rtps, pos, 1, 1).Strings()

	return &reelStrips, show_reelStrips, nil
}

// CalcNGResult - calc NG result
func (s *SG001) CalcNGResult(unitbet string, reel *games.ReelStrips, id, brand, user string, pickem, symbols []string, bet decimal.Decimal) *games.Records {
	record := games.Records{
		Id:       "0",
		Brand:    brand,
		Username: user,
		Case:     games.Lose,
		Pickem:   pickem,
		Symbols:  symbols,
		Bet:      bet,
	}

	game_math := sg001_gameplay[unitbet]

	ng_result := game_math.ng_game.CalcWayReel(*reel, bet)
	if ng_result.IsWinGreaterThanZero() {
		record.Case = games.Win
		record.Multiplier = ng_result.TotalWin.Div(bet)
		record.Point = ng_result.TotalWin
	}

	return &record
}
