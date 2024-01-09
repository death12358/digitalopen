package grow

import (
	"encoding/json"

	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/slots"
	"github.com/shopspring/decimal"
)

type Games struct {
	grow map[int]*slots.WayGames
}

// NewLightningGames - create new LightningGames
//
//	@param bg_reel	- bonus game reel weight
//	@param fg		- 中了FG後權重配置
//	@param def		- 盤面大小
func NewGrowGames(fg map[int]*slots.WayGames) *Games {
	return &Games{
		grow: fg,
	}
}

func (s *Games) GetGrowgames(i int) *slots.WayGames {
	// game_math, ok := dcg011_gameplay[unitbet]
	// if !ok {
	// 	return nil, errors.New("unitbet not found: " + unitbet)
	// }
	return s.grow[i]
}

// GetNGReelsLen - get the reels length
func (s *Games) GetReelsLen(i int, unitbet, rtp string) []int {
	//get the unitbet playsource.

	rtps := games.RTPs(rtp)
	return s.GetGrowgames(i).GetReelsLen(rtps)
}

// SpinNGReelStripts - spin NG reel strips
//
//	@param i level.
func (s *Games) FirstSpinReelStripts(i int, unitbet string, rtp string, pos []int) (*games.ReelStrips, [][]int, []string, error) {

	rtps := games.RTPs(rtp)
	nglen := s.GetReelsLen(i, unitbet, rtp)
	reelStrips := s.GetGrowgames(i).ContiguousReelStrips(rtps, pos)
	//fmt.Printf("reelStrips: %+v\n", reelStrips)
	contiguousPos := make([][]int, len(nglen))
	for i := 0; i < len(reelStrips); i++ {
		contiguousPos[i] = make([]int, len(reelStrips[i]))
		contiguousPos[i][0] = pos[i]
		for j := 1; j < len(reelStrips[i]); j++ {
			contiguousPos[i][j] = (contiguousPos[i][j-1] + 1) % nglen[i]
		}
	}
	//fmt.Printf("🔵contiguousPos: %+v\n", contiguousPos)
	show_reelStrips := s.GetGrowgames(i).ShowReelStrips(rtps, pos, 1, 1).InvertRegularXYAxis().Strings()
	//fmt.Printf("🔴show_reelStrips: %+v\n", show_reelStrips)
	return &reelStrips, contiguousPos, show_reelStrips, nil
}

// CalcNGResult - calc FG result
//
//	@return &record
//	@return fg_result_newpos
//	@return &fg_result_newstrips
//	@return extra.NextStripsDef
//	@return respin
func (s *Games) CalcFGResult(level int, unitbet, rtp string, reel games.ReelStrips, pos [][]int, id, brand, user string, symbols []string, bet decimal.Decimal) (*games.Records, [][]int, *games.ReelStrips, int, int, int) {
	//define the record.Extra with json.
	respin := 0
	type ExtraCascading struct {
		CurrentStripsDef int      `json:"currentfloor"`
		Cascadingpoint   []string `json:"cascadingpoints"`
		NextFloor        int      `json:"nextfloor"`
		Respin           int      `json:"respin"`
	}
	extra := ExtraCascading{
		Cascadingpoint:   make([]string, 0),
		CurrentStripsDef: level,
		NextFloor:        level,
		Respin:           respin,
	}
	record := games.Records{
		Id:       id,
		Brand:    brand,
		Username: user,
		Case:     games.FreeGame,
		Extra:    make([]string, 0),
		Pickem:   []string{unitbet},
		Symbols:  symbols,
		Bet:      bet,
		Point:    decimal.NewFromFloat(0.0),
	}
	rtps := games.RTPs(rtp)
	fg_game := s.GetGrowgames(level)
	fg_result_removedstrips, fg_result_newpos, fg_result_newstrips, fg_result := fg_game.CalcCascadingWayReel(rtps, pos, reel, bet)
	if !fg_result.IsWinGreaterThanZero() {
		extra.Cascadingpoint = []string{"0"}
	}
	//fmt.Printf("ng_result_removedstrips: %+v\n, ng_result_newpos: %+v\n ,ng_result_newstrips: %+v\n,ng_result: %+v\n", ng_result_removedstrips, ng_result_newpos, ng_result_newstrips, ng_result)

	if fg_result.IsWinGreaterThanZero() {
		//i := 1
		if extra.NextFloor < 4 {
			extra.NextFloor++
		}
	}
	for fg_result.IsWinGreaterThanZero() {

		//record.Case = games.Win
		record.Multiplier = record.Multiplier.Add(fg_result.TotalWin.Div(bet))
		extra.Cascadingpoint = append(extra.Cascadingpoint, fg_result.TotalWin.String())

		record.Point = record.Point.Add(fg_result.TotalWin)
		record.Symbols = append(append(record.Symbols, fg_result_removedstrips.InvertRegularXYAxis().Strings()...), fg_result_newstrips.InvertRegularXYAxis().Strings()...)
		fg_result_removedstrips, fg_result_newpos, fg_result_newstrips, fg_result = fg_game.CalcCascadingWayReel(rtps, fg_result_newpos, fg_result_newstrips, bet)
	}
	_, respin = fg_result_newstrips.CalcSymbolMatches(slots.SF)
	extra.Respin = respin
	c, _ := json.Marshal(extra)
	record.Extra = []string{string(c)}
	//Cul the re-spin counts in this round.

	return &record, fg_result_newpos, &fg_result_newstrips, extra.CurrentStripsDef, extra.NextFloor, respin
}
