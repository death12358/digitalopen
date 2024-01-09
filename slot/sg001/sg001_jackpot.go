package sg001

// 每次牌面上出現WW，有 1/70 的機率觸發 Jackpot。
// NG 才會有 Jackpot，且該款遊戲不會同時觸發。
// 起始倍率：
//  Mini:	10  倍，無彩池累積，權重： 4700
//  Minor:	25  倍，無彩池累積，權重： 5100
//  Major:	100 倍，無彩池累積，權重： 195
//  Grand:	500 倍，無彩池累積，權重： 5

import (
	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/slots"
	"github.com/death12358/digitalopen/games/weight"

	"github.com/shopspring/decimal"
)

// IsJackpotTrigger - 判斷是否觸發 Jackpot
func (s *SG001) IsJackpotTrigger(unitbet string, reel *games.ReelStrips) bool {
	_, count := reel.CalcSymbolMatches(slots.WW)

	if count >= jp_match_def {
		jp_game := sg001_gameplay[unitbet].jp_game

		return jp_game.TriggerJackpot()
	}
	return false
}

// SpinJackpot - spin Jackpot
func (s *SG001) SpinJackpot(unitbet string) games.Symbol {
	jp_game := sg001_gameplay[unitbet].jp_game
	return jp_game.PickJackpot()
}

// CalcJackpot
func (s *SG001) CalcJackpot(unitbet string, jp games.Symbol, totalbet, pool decimal.Decimal) decimal.Decimal {
	jp_game := sg001_gameplay[unitbet].jp_game
	return jp_game.CalcJackpot(jp, totalbet, pool)
}

const (
	JPMini = games.Symbol(0) + iota
	JPMinor
	JPMajor
	JPGrand

	jp_match_def = 1
)

var (
	jpTriggerWeightTable = []int{1, 69}
	jpTriggerObjectTable = []int{0, 1}

	jpGameWeightTable = []int{4700, 5100, 195, 5}
	jpGameObjectTable = []int{JP_Mini, JP_Minor, JP_Major, JP_Grand}

	jpBasePool = []decimal.Decimal{
		decimal.NewFromInt(10),
		decimal.NewFromInt(25),
		decimal.NewFromInt(100),
		decimal.NewFromInt(500),
	}

	jpSymbolList = []games.Symbol{
		JPMini,
		JPMinor,
		JPMajor,
		JPGrand,
	}

	jpSymbolString = []string{
		"mini",
		"minor",
		"major",
		"grand",
	}

	jpPayTable = &games.Pays{
		decimal.NewFromInt(10),
		decimal.NewFromInt(25),
		decimal.NewFromInt(100),
		decimal.NewFromInt(500),
	}

	JPTrigger = weight.NewGames(jpTriggerWeightTable, jpTriggerObjectTable)
	JPGame    = weight.NewGames(jpGameWeightTable, jpGameObjectTable)

	roll_rate = []decimal.Decimal{
		decimal.Zero,
		decimal.Zero,
		decimal.NewFromFloat(0.006),
		decimal.NewFromFloat(0.004),
	}
)

// JPSymbolDefine - 定義 JP symbol
// ------------------------------------------------------------
const (
	// JPSymbolDefine - 定義 JP symbol
	JP_Mini = int(0) + iota
	JP_Minor
	JP_Major
	JP_Grand

	// JPWeightDefine - 定義 JP weight
	JP_MiniWeight  = 4700
	JP_MinorWeight = 5100
	JP_MajorWeight = 195
	JP_GrandWeight = 5
)
