package sg006

import (
	"encoding/json"
	"strconv"

	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/random"
	"github.com/death12358/digitalopen/sg006/battle"
	"github.com/shopspring/decimal"
)

// games.BonusFreeGame1
var RemainRetrigger int

// FGFlow - Bonus  pay Symbol Free Game flow
func (s *SG006) FGFlow(rtp string, bet decimal.Decimal, round *games.Rounds) (*games.Rounds, error) {
	unitbet := round.Result["0"].Pickem[0]
	var MonsterNumberOnRing map[string]int = map[string]int{
		"H2": 1,
		"H3": 1,
		"H4": 1,
		"H5": 1,
	}
	// default 8 times
	count := 0
	stage := int64(count) // 1,2,3,......
	ring := MonsterNumberOnRing
	for {
		fg_reel, show := s.SpinFGReelStripts(unitbet, rtp)
		round.Stages++
		stage++

		//MonsterNumberOnBoard := battle.CountMonsters(fg_reel)
		//TotalMonster := battle.HowManyMonsters(MonsterNumberOnBoard) - MonsterNumberOnBoard["H1"] + battle.HowManyMonsters(ring)

		id := strconv.Itoa(int(round.Stages))
		res, ring, BattleIsLose := s.CalcFGResult(unitbet, fg_reel, id, round.Brand, round.Username, show, ring, bet)

		res.Stages = stage

		//var ShowNumbers games.ExtraSG006
		//ShowNumbers.WWmultiplier = WWmulti
		//if s.IsBonusWin(unitbet, fg_reel) {
		//	RemainRetrigger++
		//}
		//howNumbers.RemainRetrigger = RemainRetrigger
		//res.AppendExtraSG006(ShowNumbers)
		round.Result[id] = res
		round.TotalPoint = round.TotalPoint.Add(res.Point)
		if battle.BattleIsWin(ring) || BattleIsLose {
			break
		}
	}
	/*
		// round.Stages = 1
		round.Result["1"].Case = games.Lose
		if !round.Result["1"].Point.IsZero() {
			round.Result["1"].Case = games.Win
		}
	*/
	return round, nil
}

// SpinFGReelStripts - spin Pay Free Game reel strips
func (s *SG006) SpinFGReelStripts(unitbet, rtp string) (*games.ReelStrips, []string) {
	fg_game := sg006_gameplay[unitbet].fg_game
	rtps := games.RTPs(rtp)

	fglen := fg_game.GetReelsLen(rtps)
	pos := random.Intsn(fglen)

	reelStrips := fg_game.ContiguousReelStrips(rtps, pos)
	// show_reelStrips := ng_game.SpinNormal(rtps, pos, 1, 1).InvertRegularXYAxis().Strings()
	show_reelStrips := fg_game.ShowReelStrips(rtps, pos, 1, 1).InvertRegularXYAxis().Strings()

	return &reelStrips, show_reelStrips
}

// CalcFGResult - calc Pay Free Game  result
func (s *SG006) CalcFGResult(unitbet string, reel *games.ReelStrips, id, brand, user string, symbols []string, ring map[string]int, bet decimal.Decimal) (*games.Records, map[string]int, bool) {
	record := games.Records{
		Id:       id,
		Brand:    brand,
		Username: user,
		Case:     games.NotStartedYet,
		Pickem:   []string{""},
		Symbols:  symbols,
	}
	SFNumber := [5]int{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			if (*reel)[i][j] == games.ToSymbol("0") {
				SFNumber[i]++
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
		SF_Point = decimal.NewFromInt(100).Mul(bet).Div(unitbet_decimal).Mul(decimal.NewFromInt(int64(SFLines)))
	} else if SFLengh == 3 {
		SF_Point = decimal.NewFromInt(550).Mul(bet).Div(unitbet_decimal).Mul(decimal.NewFromInt(int64(SFLines)))
	}

	MonsterNumberOnBoard := battle.CountMonsters(reel)
	fg_game := sg006_gameplay[unitbet].fg_game
	fg_result := fg_game.CalcWayReel(*reel, bet)
	win_detail, _ := json.Marshal(fg_result.Wins)
	record.Extra = append(record.Extra, string(win_detail))
	ring, BattlePoint, BattleIsLose := battle.StartBattle(MonsterNumberOnBoard, ring)
	//fmt.Printf("\n\n\nBattlePoint:%v\n\n\n", BattlePoint)
	record.Extra = append(record.Extra, "BattleRing")
	BattleRing, _ := json.Marshal(ring)
	record.Extra = append(record.Extra, string(BattleRing))

	fg_result.TotalWin = fg_result.TotalWin.Add((BattlePoint.Mul(bet).Div(unitbet_decimal))).Add(SF_Point)
	//fmt.Printf("\n\n\nfg_result.TotalWint:%v\n\n\n", fg_result.TotalWin)
	//fmt.Printf("\n\n\nffg_result.IsWinGreaterThanZero() :%v\n\n\n", fg_result.IsWinGreaterThanZero())
	if fg_result.IsWinGreaterThanZero() {
		record.Point = fg_result.TotalWin
		//fmt.Printf("\n\n\nfg_result.TotalWin+BattlePoint:%v +%v\n\n\n", fg_result.TotalWin, BattlePoint)
		record.Multiplier = record.Point.Div(bet)
	}

	return &record, ring, BattleIsLose
}
