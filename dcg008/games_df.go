package dcg008

import (
	"github.com/death12358/digitalopen/dcg008/adventure"
	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/slots"
	"github.com/shopspring/decimal"
)

var (
	// ng
	dcg008ng = slots.NewWayGames(
		dcg008ngTable, ngreelDef, ngpayTable_unitbet100, SymbolList, ScatterPosition, uint_bet_100)
	//fg
	dcg008fg = slots.NewWayGames(
		dcg008fgTable, fgreelDef, fgpayTable_unitbet100, SymbolList, ScatterPosition, uint_bet_100)
	dcg008bg = adventure.NewAdventureGames(adventure.BgAdvPOTable, adventure.AdvPWTable, adventure.AdventurePayTable)
)

// 这里是NG相关
var (
	// 定義轉輪為 3x3x3x3x3
	ngreelDef = &slots.ReelStripsDef{3, 3, 3, 3, 3}
	fgreelDef = &slots.ReelStripsDef{3, 3, 3, 3, 3}
	// unit bet
	uint_bet_100 = decimal.NewFromInt(100)
	////转轮长度
	//ng_len = DCG008ngReelStrips.Lengths()
	//fg_len = DCG008fgReelStrips.Lengths()

)

var (
	dcg008ngTable = games.ReelStripList{
		//rtp:87.878848%   100,000,000
		"88": &dcg008ngReelStrips_88,
		//rtp:91.9379529%  50,000,000
		"92": &dcg008ngReelStrips_92,
		//rtp:96.2795448  50,000,000
		"96": &dcg008ngReelStrips_96,
		//97.9523697%  50,000,000
		"98": &dcg008ngReelStrips_98,
		//more fgbg.
		"40": &dcg008testing,
	}
	dcg008fgTable = games.ReelStripList{
		"88": &dcg008fgReelStrips_88,
		"92": &dcg008fgReelStrips_92,
		"96": &dcg008fgReelStrips_96,
		"98": &dcg008fgReelStrips_98,
		"40": &dcg008fgReelStrips_88,
	}
)
