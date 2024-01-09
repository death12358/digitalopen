package sg006

import (
	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/slots"
	"github.com/death12358/digitalopen/slot/sg006/vs"
	"github.com/shopspring/decimal"
)

var (
	// 牌面定義
	reelDef = &slots.ReelStripsDef{3, 3, 3, 3, 3}

	// unit bet
	uint_bet100 = decimal.NewFromInt(100)

	// 轉輪表
	ngTable_100 = games.ReelStripList{
		"98": &ngReelStrips98_100,
		"97": &ngReelStrips98_100,
		"96": &ngReelStrips98_100,
		"92": &ngReelStrips98_100,
		"88": &ngReelStrips98_100,
	}
	fgTable_100 = games.ReelStripList{
		"98": &fgReelStrips98_100,
		"97": &fgReelStrips98_100,
		"96": &fgReelStrips98_100,
		"92": &fgReelStrips98_100,
		"88": &fgReelStrips98_100,
	}
)

var (
	// ng
	sg006ng_100 = slots.NewWayGames(
		ngTable_100, reelDef, payTable_unitbet100, slots.SymbolList, slots.ScatterPosition, uint_bet100)

	// fg
	sg006fg_100 = slots.NewWayGames(
		fgTable_100, reelDef, payTable_unitbet100, slots.SymbolList, slots.ScatterPosition, uint_bet100)

	// vs
	sg006vs = map[string]*vs.Games{
		"98": sg006vs_98,
		"97": sg006vs_97,
		"96": sg006vs_96,
		"92": sg006vs_92,
		"88": sg006vs_88,
	}
	sg006vs_NoBattle = map[string]*vs.Games{
		"98": sg006vs_NoBattle_98,
		"97": sg006vs_NoBattle_97,
		"96": sg006vs_NoBattle_96,
		"92": sg006vs_NoBattle_92,
		"88": sg006vs_NoBattle_88,
	}
	sg006vs_98          = vs.NewGame(vs.VS_reel_98, vs.VSPayTable)
	sg006vs_NoBattle_98 = vs.NewGame(vs.VS_reel_NoBattle_98, vs.VSPayTable)

	sg006vs_97          = vs.NewGame(vs.VS_reel_97, vs.VSPayTable)
	sg006vs_NoBattle_97 = vs.NewGame(vs.VS_reel_NoBattle_97, vs.VSPayTable)

	sg006vs_96          = vs.NewGame(vs.VS_reel_96, vs.VSPayTable)
	sg006vs_NoBattle_96 = vs.NewGame(vs.VS_reel_NoBattle_96, vs.VSPayTable)

	sg006vs_92          = vs.NewGame(vs.VS_reel_92, vs.VSPayTable)
	sg006vs_NoBattle_92 = vs.NewGame(vs.VS_reel_NoBattle_92, vs.VSPayTable)

	sg006vs_88          = vs.NewGame(vs.VS_reel_88, vs.VSPayTable)
	sg006vs_NoBattle_88 = vs.NewGame(vs.VS_reel_NoBattle_88, vs.VSPayTable)
)
