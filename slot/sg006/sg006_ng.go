package sg006

import (
	"encoding/json"
	"errors"

	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/random"
	"github.com/shopspring/decimal"
)

// SpinNormalGame
func (s *SG006) SpinNormalGame(unitbet, rtp string, position []int, round *games.Rounds) (games.Rounds, *games.ReelStrips) {
	// 取得獎圖長度以已檢查
	// // 取得 unitbet 遊戲
	// game_math, ok := sg006_gameplay[unitbet]
	// if !ok {
	// 	return nil, nil, errors.New("unitbet not found: " + unitbet)
	// }
	game_math, _ := sg006_gameplay[unitbet]
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
		Point:    decimal.Zero,
	}

	// 除以 unitbet
	// unit_bet, _ := decimal.NewFromString(unitbet)
	// bet := record.Bet.Mul(unit_bet)

	// 計算盤面
	//新令一個輪以避開ARRAY賦值的問題(待優化)
	var A games.Symbol
	var reelStrips_SFtoH1 games.ReelStrips

	R1 := games.Reels{A, A, A}
	R2 := games.Reels{A, A, A}
	R3 := games.Reels{A, A, A}
	R4 := games.Reels{A, A, A}
	R5 := games.Reels{A, A, A}
	reelStrips_SFtoH1 = append(reelStrips_SFtoH1, R1)
	reelStrips_SFtoH1 = append(reelStrips_SFtoH1, R2)
	reelStrips_SFtoH1 = append(reelStrips_SFtoH1, R3)
	reelStrips_SFtoH1 = append(reelStrips_SFtoH1, R4)
	reelStrips_SFtoH1 = append(reelStrips_SFtoH1, R5)
	SF := games.ToSymbol("13")
	H1 := games.ToSymbol("1")
	SFNumber := [5]int{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			if reelStrips[i][j] == SF {
				SFNumber[i]++
				reelStrips_SFtoH1[i][j] = H1
			} else {
				reelStrips_SFtoH1[i][j] = reelStrips[i][j]
			}
		}
	}
	//計算有多少ＳＦ連線
	SFLengh := 0
	for i := 1; i < 4; i++ {
		if SFNumber[i] == 0 {
			SFNumber[i] = 1
		} else {
			SFLengh++
		}
	}
	SFLines := SFNumber[1] * SFNumber[2] * SFNumber[3]
	SF_Point := decimal.Zero
	unitbet_decimal, _ := decimal.NewFromString(unitbet)

	if SFLengh == 2 {
		SF_Point = decimal.NewFromInt(100).Mul(round.TotalBet).Div(unitbet_decimal).Mul(decimal.NewFromInt(int64(SFLines)))
	} else if SFLengh == 3 {
		SF_Point = decimal.NewFromInt(550).Mul(round.TotalBet).Div(unitbet_decimal).Mul(decimal.NewFromInt(int64(SFLines)))
	}

	ng_result := game_math.ng_game.CalcWayReel(reelStrips_SFtoH1, round.TotalBet)
	win_detail, _ := json.Marshal(ng_result.Wins)
	record.Extra = append(record.Extra, string(win_detail))

	ng_result.TotalWin = ng_result.TotalWin.Add(SF_Point)
	if ng_result.IsWinGreaterThanZero() {
		record.Case = games.Win
		record.Multiplier = ng_result.TotalWin.Div(round.TotalBet)
		record.Point = ng_result.TotalWin
	}

	round.Status = round.Status.Push(record.Case)
	round.Result["0"] = record
	round.TotalPoint = round.TotalPoint.Add(record.Point)
	/*j_round, err := json.Marshal(round)
	if err != nil {
		fmt.Printf("Marshal error: %s", err.Error())
	}
	fmt.Printf("Round: %s", string(j_round))*/

	return *round, &reelStrips
}

// GetNGReelsLen - 取得NG滾輪長度
func (s *SG006) GetNGReelsLen(unitbet, rtp string) ([]int, error) {
	// 取得 unitbet 遊戲
	game_math, ok := sg006_gameplay[unitbet]
	if !ok {
		return nil, errors.New("unitbet not found: " + unitbet)
	}
	rtps := games.RTPs(rtp)
	return game_math.ng_game.GetReelsLen(rtps), nil
}

// SpinNGReelStripts - spin NG reel strips
func (s *SG006) SpinNGReelStripts(unitbet, rtp string) (*games.ReelStrips, []string, error) {
	game_math, ok := sg006_gameplay[unitbet]
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
func (s *SG006) CalcNGResult(unitbet string, reel *games.ReelStrips, id, brand, user string, pickem, symbols []string, bet decimal.Decimal) *games.Records {
	record := games.Records{
		Id:       "0",
		Brand:    brand,
		Username: user,
		Case:     games.Lose,
		Pickem:   pickem,
		Symbols:  symbols,
		Bet:      bet,
	}

	game_math := sg006_gameplay[unitbet]

	ng_result := game_math.ng_game.CalcWayReel(*reel, bet)
	if ng_result.IsWinGreaterThanZero() {
		record.Case = games.Win
		record.Multiplier = ng_result.TotalWin.Div(bet)
		record.Point = ng_result.TotalWin
	}

	return &record
}
