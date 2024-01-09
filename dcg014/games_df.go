package dcg014

import (
	"github.com/shopspring/decimal"

	"github.com/death12358/digitalopen/dcg014/easteregg"
	"github.com/death12358/digitalopen/dcg014/grow"
	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/slots"
)

var (
	// ng
	dcg014ng = slots.NewWayGames(
		dcg014ngTable, ngreelDef, ngpayTable_unitbet100, SymbolList, ScatterPosition, uint_bet_100)
	//fg
	dcg014fg1 = slots.NewWayGames(
		dcg014fg1Table, fgreel1Def, fgpayTable_unitbet100, SymbolList, ScatterPosition, uint_bet_100)
	dcg014fg2 = slots.NewWayGames(
		dcg014fg2Table, fgreel2Def, fgpayTable_unitbet100, SymbolList, ScatterPosition, uint_bet_100)
	dcg014fg3 = slots.NewWayGames(
		dcg014fg3Table, fgreel3Def, fgpayTable_unitbet100, SymbolList, ScatterPosition, uint_bet_100)
	dcg014fg4 = slots.NewWayGames(
		dcg014fg4Table, fgreel4Def, fgpayTable_unitbet100, SymbolList, ScatterPosition, uint_bet_100)
	dcg014fg = grow.NewGrowGames(map[int]*slots.WayGames{1: dcg014fg1, 2: dcg014fg2, 3: dcg014fg3, 4: dcg014fg4})
	dcg014bg = easteregg.NewEggGames(easteregg.Bg_EggNameTable, easteregg.Bg_EggWeightTable, easteregg.PayTable, easteregg.Bg_BonusNameTable, easteregg.Bg_BonusWeightTable)
	Dcg014fg = grow.NewGrowGames(map[int]*slots.WayGames{1: dcg014fg1, 2: dcg014fg2, 3: dcg014fg3, 4: dcg014fg4})
)

// NG
var (
	// reels define 3x3x3x3x3
	ngreelDef  = &slots.ReelStripsDef{6, 6, 6, 6, 6, 6}
	fgreel1Def = &slots.ReelStripsDef{6, 6, 6, 6, 6, 6}
	fgreel2Def = &slots.ReelStripsDef{7, 7, 7, 7, 7, 7}
	fgreel3Def = &slots.ReelStripsDef{8, 8, 8, 8, 8, 8}
	fgreel4Def = &slots.ReelStripsDef{9, 9, 9, 9, 9, 9}
	// unit bet
	uint_bet_100 = decimal.NewFromInt(100)
	////reels
	//ng_len = DCG014ngReelStrips.Lengths()
	//fg_len = DCG014fgReelStrips.Lengths()

)

var (
	dcg014ngTable = games.ReelStripList{
		"98": &dcg014ngReelStrips_98,
		"96": &dcg014ngReelStrips_96,
		"92": &dcg014ngReelStrips_92,
		"88": &dcg014ngReelStrips_88,
		"40": &dcg014testing,
	}
	dcg014fg1Table = games.ReelStripList{
		"98": &dcg014fg1ReelStrips,
		"96": &dcg014fg1ReelStrips,
		"92": &dcg014fg1ReelStrips,
		"88": &dcg014fg1ReelStrips,
		"40": &dcg014fg1ReelStrips,
	}
	dcg014fg2Table = games.ReelStripList{
		"98": &dcg014fg2ReelStrips,
		"96": &dcg014fg2ReelStrips,
		"92": &dcg014fg2ReelStrips,
		"88": &dcg014fg2ReelStrips,
		"40": &dcg014fg2ReelStrips,
	}
	dcg014fg3Table = games.ReelStripList{
		"98": &dcg014fg3ReelStrips,
		"96": &dcg014fg3ReelStrips,
		"92": &dcg014fg3ReelStrips,
		"88": &dcg014fg3ReelStrips,
		"40": &dcg014fg3ReelStrips,
	}
	dcg014fg4Table = games.ReelStripList{
		"98": &dcg014fg4ReelStrips,
		"96": &dcg014fg4ReelStrips,
		"92": &dcg014fg4ReelStrips,
		"88": &dcg014fg4ReelStrips,
		"40": &dcg014fg4ReelStrips,
	}
)
