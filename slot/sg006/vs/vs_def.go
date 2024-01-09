package vs

import (
	"github.com/death12358/digitalopen/games"
	weights "github.com/death12358/digitalopen/games/weight"
	"github.com/shopspring/decimal"
)

var (
	//vs_HiToVSHiGameWeightTable_98 = []int{8, 1} //Test
	vs_HiToVSHiGameWeightTable_98 = []int{60, 1}
	vs_HiToVSHiGameWeightTable_97 = []int{60, 1}
	vs_HiToVSHiGameWeightTable_96 = []int{60, 1}
	vs_HiToVSHiGameWeightTable_92 = []int{68, 1}
	vs_HiToVSHiGameWeightTable_88 = []int{68, 1}
	vs_HiToVSHiGameObjectTable    = []int{0, 1}

	vs_HiToVSHiGame map[string]*weights.Games = map[string]*weights.Games{
		"98": vs_HiToVSHiGame_98,
		"97": vs_HiToVSHiGame_97,
		"96": vs_HiToVSHiGame_96,
		"92": vs_HiToVSHiGame_92,
		"88": vs_HiToVSHiGame_88,
	}

	vs_HiToVSHiGame_98 = weights.NewGames(
		vs_HiToVSHiGameWeightTable_98,
		vs_HiToVSHiGameObjectTable,
	)
	vs_HiToVSHiGame_97 = weights.NewGames(
		vs_HiToVSHiGameWeightTable_97,
		vs_HiToVSHiGameObjectTable,
	)
	vs_HiToVSHiGame_96 = weights.NewGames(
		vs_HiToVSHiGameWeightTable_96,
		vs_HiToVSHiGameObjectTable,
	)
	vs_HiToVSHiGame_92 = weights.NewGames(
		vs_HiToVSHiGameWeightTable_92,
		vs_HiToVSHiGameObjectTable,
	)
	vs_HiToVSHiGame_88 = weights.NewGames(
		vs_HiToVSHiGameWeightTable_88,
		vs_HiToVSHiGameObjectTable,
	)

	VS_reel_98 = weights.NewGames(
		vsReelWeightTable_98,
		vsReelObjectTable,
	)

	VS_reel_NoBattle_98 = weights.NewGames(
		vsReelNoBattleWeightTable_98,
		vsReelNoBattleObjectTable,
	)
	VS_reel_97 = weights.NewGames(
		vsReelWeightTable_97,
		vsReelObjectTable,
	)

	VS_reel_NoBattle_97 = weights.NewGames(
		vsReelNoBattleWeightTable_97,
		vsReelNoBattleObjectTable,
	)
	VS_reel_96 = weights.NewGames(
		vsReelWeightTable_96,
		vsReelObjectTable,
	)

	VS_reel_NoBattle_96 = weights.NewGames(
		vsReelNoBattleWeightTable_96,
		vsReelNoBattleObjectTable,
	)
	VS_reel_92 = weights.NewGames(
		vsReelWeightTable_92,
		vsReelObjectTable,
	)

	VS_reel_NoBattle_92 = weights.NewGames(
		vsReelNoBattleWeightTable_92,
		vsReelNoBattleObjectTable,
	)
	VS_reel_88 = weights.NewGames(
		vsReelWeightTable_88,
		vsReelObjectTable,
	)

	VS_reel_NoBattle_88 = weights.NewGames(
		vsReelNoBattleWeightTable_88,
		vsReelNoBattleObjectTable,
	)
)

/*
// CalcJackpot

	func (s *SG006) CalcJackpot(unitbet string, jp games.Symbol, totalbet, pool decimal.Decimal) decimal.Decimal {
		vs_game := sg006_gameplay[unitbet].vs_game
		return vs_game.CalcVS(jp, totalbet, pool)
	}
*/
const (
	vs_JPGrand = games.Symbol(0) + iota
	vs_JPMajor
	vs_JPMinor
	vs_Battle
	vs_JPMiniPlusBattle
	vs_JPMini
	vs_10x
	vs_5x
	vs_3x
	jp_match_def = 1
)

// /乘倍的是否要去掉
var (
	vsReelWeightTable_98 = []int{1, 53, 101, 293, 157, 211, 251, 400, 440}
	vsReelWeightTable_97 = []int{1, 53, 101, 293, 157, 211, 350, 400, 440}
	vsReelWeightTable_96 = []int{1, 50, 50, 293, 157, 211, 251, 455, 440}
	//vsReelWeightTable_96 = []int{1, 50, 50, 293, 157, 190, 172, 455, 530}
	vsReelWeightTable_92 = []int{1, 50, 50, 293, 157, 211, 251, 455, 440}
	vsReelWeightTable_88 = []int{1, 10, 30, 293, 157, 211, 151, 505, 550}
	vsReelObjectTable    = []int{
		int(vs_JPGrand),
		int(vs_JPMajor),
		int(vs_JPMinor),
		int(vs_Battle),
		int(vs_JPMiniPlusBattle),
		int(vs_JPMini),
		int(vs_10x),
		int(vs_5x),
		int(vs_3x),
	}

	vsReelNoBattleWeightTable_98 = []int{1, 123, 200, 0, 0, 350, 333, 488, 100}
	vsReelNoBattleWeightTable_97 = []int{1, 123, 200, 0, 0, 350, 416, 488, 100}
	vsReelNoBattleWeightTable_96 = []int{1, 120, 133, 0, 0, 350, 333, 488, 100}
	//vsReelNoBattleWeightTable_96 = []int{1, 120, 133, 0, 0, 322, 321, 478, 150}
	vsReelNoBattleWeightTable_92 = []int{1, 120, 133, 0, 0, 350, 333, 488, 100}
	vsReelNoBattleWeightTable_88 = []int{1, 80, 100, 0, 0, 400, 380, 453, 111}
	vsReelNoBattleObjectTable    = []int{
		int(vs_JPGrand),
		int(vs_JPMajor),
		int(vs_JPMinor),
		int(vs_Battle),
		int(vs_JPMiniPlusBattle),
		int(vs_JPMini),
		int(vs_10x),
		int(vs_5x),
		int(vs_3x),
	}
	vsPool = []decimal.Decimal{
		decimal.NewFromInt(300),
		decimal.NewFromInt(100),
		decimal.NewFromInt(50),
		decimal.NewFromInt(36),
		decimal.NewFromInt(56),
		decimal.NewFromInt(20),
		decimal.NewFromInt(10),
		decimal.NewFromInt(5),
		decimal.NewFromInt(3),
	}

	jpSymbolList = []games.Symbol{
		vs_JPGrand,
		vs_JPMajor,
		vs_JPMinor,
		vs_Battle,
		vs_JPMiniPlusBattle,
		vs_JPMini,
		vs_10x,
		vs_5x,
		vs_3x,
	}

	SymbolString = []string{
		"vs_JPGrand",
		"vs_JPMajor",
		"vs_JPMinor",
		"vs_Battle",
		"vs_JPMiniPlusBattle",
		"vs_JPMini",
		"vs_10x",
		"vs_5x",
		"vs_3x",
	}

	VSPayTable = &games.Pays{
		decimal.NewFromInt(300),
		decimal.NewFromInt(100),
		decimal.NewFromInt(50),
		decimal.NewFromInt(0),
		decimal.NewFromInt(20),
		decimal.NewFromInt(20),
		decimal.NewFromInt(10),
		decimal.NewFromInt(5),
		decimal.NewFromInt(3),
	}

	/*
	   JPTrigger = weights.NewGames(jpTriggerWeightTable, jpTriggerObjectTable)
	*/
	//vs_reel = weights.NewGames(vsReelWeightTable, vsReelObjectTable)
)
