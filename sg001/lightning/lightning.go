package lightning

import (
	"errors"

	"games"
	"games/random"
	"games/slots"
	weights "games/weight"

	"github.com/shopspring/decimal"
)

// Games is a game of Lightning
// games.BonusFreeGame2
type Games struct {
	bg_reel_game *weights.Games
	fg_reel_game *weights.Games
	reelDef      int
}

// NewLightningGames - create new LightningGames
//
//	@param bg_reel	- bonus game reel weight
//	@param fg		- 中了FG後權重配置
//	@param def		- 盤面大小
func NewLightningGames(bg_reel, fg *weights.Games, def uint) *Games {
	return &Games{
		bg_reel_game: bg_reel,
		fg_reel_game: fg,
		reelDef:      int(def),
	}
}

// GetReelDef
func (g *Games) GetReelDef() int {
	return g.reelDef
}

// HoldNGReel - hold NG reel
func (g *Games) HoldNGReel(reels *games.Reels, record_last *games.Records) *games.Records {
	record := &games.Records{
		Id:         "1",
		Brand:      record_last.Brand,
		Username:   record_last.Username,
		Case:       games.NotStartedYet,
		Stages:     2, // default 3
		Pickem:     []string{""},
		Symbols:    []string{},
		Multiplier: decimal.Zero,
		Bet:        decimal.Zero,
		Point:      decimal.Zero,
	}

	idx := 0
	for i, v := range record_last.Symbols {

		if real_reels_pos_InvertRegularXYAxis[i] {
			sf, ok := bgSymbolMapping[v]
			if ok {
				record.Symbols = append(record.Symbols, v)
				(*reels)[idx] = sf
				record.Multiplier = record.Multiplier.Add(bgLightningPayTable[sf])
			} else {
				record.Symbols = append(record.Symbols, "0")
				(*reels)[idx] = 0
			}
			idx++
		}

		// // 取得真實轉輪位置
		// 		real_idx := i % 5
		// // show_reelStrips.InvertRegularXYAxis real is 1 2 3
		// if real_idx == 1 || real_idx == 2 || real_idx == 3 {
		// }
	}

	//log.Printf("LightningGames HoldNGReel: %+v", record)
	return record
}

// replaceNGSymbol -
func (g *Games) ReplaceNGSymbol(record *games.Records) *games.Records {
	for i, v := range record.Symbols {
		if v == slots.SF.String() {

			dice := random.Intn(g.fg_reel_game.Sum())

			pick, _ := g.fg_reel_game.Pick(dice)
			record.Symbols[i] = BGName(games.Symbol(pick))
		}
	}
	return record
}

// Respin - FG Respin
func (g *Games) Respin(reels *games.Reels, rc *games.Records, id string, bet decimal.Decimal) (*games.Reels, *games.Records, error) {
	bi, ok := g.CheckBlankSpaces(reels)
	if !ok {
		return nil, nil, errors.New("no blank spaces found or reel length is incorrect or reels is nil")
	}
	record := &games.Records{
		Id:         id,
		Brand:      rc.Brand,
		Username:   rc.Username,
		Case:       games.NotStartedYet,
		Stages:     rc.Stages - 1,
		Pickem:     rc.Pickem,
		Multiplier: decimal.Zero, /////////////////////////
		Point:      decimal.Zero,
	}

	// 待優化
	// bg 個數統計
	bg_stat := int64(0)
	for _, v := range bi {
		// 產生新的reel
		(*reels)[v] = g.SpinReel()
		if (*reels)[v] != bg_default {
			// 計算fg獎金
			record.Multiplier = record.Multiplier.Add(bgLightningPayTable[(*reels)[v]])
			if (*reels)[v] == bg_bg {
				bg_stat++
				record.Stages = 3
			}
		}
	}
	if bg_stat > 0 {
		record.Case = games.NotStartedYet
		record.Stages = 3
		// record.Case = games.NotStartedYet
		multi := g.CalcBonusPoint(reels)
		// bg 獎金等於所有fg加總乘上bg個數
		calc_bg := multi.Mul(decimal.NewFromInt(bg_stat))
		record.Multiplier = record.Multiplier.Add(calc_bg)
	}
	// 待優化
	// 倍率大於0，表示有中獎，並計算獎金
	if record.Multiplier.GreaterThan(decimal.Zero) {
		record.Case = games.NotStartedYet
		record.Stages = 3

		record.Point = record.Multiplier.Mul(bet)
	}
	// 有中BG，並計算獎金

	// 替換FG
	record.Symbols = BGReel(*reels)

	return reels, record, nil
}

// CalcBonusPoint - 計算Bonus獎金
func (g *Games) CalcBonusPoint(reels *games.Reels) decimal.Decimal {
	multi := decimal.Zero
	for _, v := range *reels {
		multi = multi.Add(bgLightningPayTable[v])
	}
	return multi
}

// SpinReel -
func (g *Games) SpinReel() games.Symbol {
	dice := random.Intn(g.bg_reel_game.Sum())
	pick, _ := g.bg_reel_game.Pick(dice) //RANDOM位置是否中獎

	// BG 是否中獎
	if pick == bg_fg.Int() {
		// FG 權重選擇
		dice := random.Intn(g.fg_reel_game.Sum())
		pick, _ := g.fg_reel_game.Pick(dice)

		return games.Symbol(pick)
	}

	return games.Symbol(pick)
}

// CheckBlankSpaces -
// checks the reels to find the indices of all blank spaces and returns a slice containing the indices and a boolean indicating whether any blank spaces were found.
func (g *Games) CheckBlankSpaces(reels *games.Reels) ([]int, bool) {
	// check reels length
	if reels.Length() != g.reelDef {
		return nil, false
	}
	// check reels is nil
	if reels == nil {
		return nil, false
	}

	// Initialize an empty slice to store the indices of the blank spaces
	blankIndices := make([]int, 0)

	// Iterate over the reels and check for blank spaces
	for i := 0; i < reels.Length(); i++ {
		if (*reels)[i] == bg_default {
			blankIndices = append(blankIndices, i)
		}
	}

	// Return the slice of indices and a boolean indicating whether any blank spaces were found
	return blankIndices, len(blankIndices) > 0
}
