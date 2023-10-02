package sg001

import (
	"github.com/shopspring/decimal"
	"gitlab.com/gaas_math/slotmachine/sg001/jackpot"
	"gitlab.com/gaas_math/slotmachine/sg001/lightning"
	"gitlab.com/gaas_module/games"
	"gitlab.com/gaas_module/games/slots"
)

var (
	// 牌面定義
	reelDef = &slots.ReelStripsDef{3, 3, 3, 3, 3}

	// unit bet
	uint_bet8  = decimal.NewFromInt(8)
	uint_bet18 = decimal.NewFromInt(18)
	uint_bet38 = decimal.NewFromInt(38)
	uint_bet68 = decimal.NewFromInt(68)
	uint_bet88 = decimal.NewFromInt(88)

	// 轉輪表
	ngTable_8 = games.ReelStripList{
		"98": &ngReelStrips98_8,
	}
	hifgTable_8 = games.ReelStripList{
		"98": &hifgReelStrips98_8,
	}
	bigfgTable_8 = games.ReelStripList{
		"98": &bigfgReelStrips98_8,
	}

	ngTable_18 = games.ReelStripList{
		"98": &ngReelStrips98_18,
	}
	hifgTable_18 = games.ReelStripList{
		"98": &hifgReelStrips98_18,
	}
	bigfgTable_18 = games.ReelStripList{
		"98": &bigfgReelStrips98_18,
	}

	ngTable_38 = games.ReelStripList{
		"98": &ngReelStrips98_38,
	}
	hifgTable_38 = games.ReelStripList{
		"98": &hifgReelStrips98_38,
	}
	bigfgTable_38 = games.ReelStripList{
		"98": &bigfgReelStrips98_38,
	}

	ngTable_68 = games.ReelStripList{
		"98": &ngReelStrips98_68,
	}
	hifgTable_68 = games.ReelStripList{
		"98": &hifgReelStrips98_68,
	}
	bigfgTable_68 = games.ReelStripList{
		"98": &bigfgReelStrips98_68,
	}

	ngTable_88 = games.ReelStripList{
		"98": &ngReelStrips98_88,
	}
	hifgTable_88 = games.ReelStripList{
		"98": &hifgReelStrips98_88,
	}
	bigfgTable_88 = games.ReelStripList{
		"98": &bigfgReelStrips98_88,
	}
)

var (
	// ng
	sg001ng_8 = slots.NewWayGames(
		ngTable_8, reelDef, payTable_unitbet8, slots.SymbolList, slots.ScatterPosition, uint_bet8)
	sg001ng_18 = slots.NewWayGames(
		ngTable_18, reelDef, payTable_unitbet18, slots.SymbolList, slots.ScatterPosition, uint_bet18)
	sg001ng_38 = slots.NewWayGames(
		ngTable_38, reelDef, payTable_unitbet38, slots.SymbolList, slots.ScatterPosition, uint_bet38)
	sg001ng_68 = slots.NewWayGames(
		ngTable_68, reelDef, payTable_unitbet68, slots.SymbolList, slots.ScatterPosition, uint_bet68)
	sg001ng_88 = slots.NewWayGames(
		ngTable_88, reelDef, payTable_unitbet88, slots.SymbolList, slots.ScatterPosition, uint_bet88)

	// hifg
	sg001hifg_8 = slots.NewWayGames(
		hifgTable_8, reelDef, payTable_unitbet8, slots.SymbolList, slots.ScatterPosition, uint_bet8)
	sg001hifg_18 = slots.NewWayGames(
		hifgTable_18, reelDef, payTable_unitbet18, slots.SymbolList, slots.ScatterPosition, uint_bet18)
	sg001hifg_38 = slots.NewWayGames(
		hifgTable_38, reelDef, payTable_unitbet38, slots.SymbolList, slots.ScatterPosition, uint_bet38)
	sg001hifg_68 = slots.NewWayGames(
		hifgTable_68, reelDef, payTable_unitbet68, slots.SymbolList, slots.ScatterPosition, uint_bet68)
	sg001hifg_88 = slots.NewWayGames(
		hifgTable_88, reelDef, payTable_unitbet88, slots.SymbolList, slots.ScatterPosition, uint_bet88)

	// bigfg
	sg001bigfg_8 = slots.NewWayGames(
		bigfgTable_8, reelDef, payTable_unitbet8, slots.SymbolList, slots.ScatterPosition, uint_bet8)
	sg001bigfg_18 = slots.NewWayGames(
		bigfgTable_18, reelDef, payTable_unitbet18, slots.SymbolList, slots.ScatterPosition, uint_bet18)
	sg001bigfg_38 = slots.NewWayGames(
		bigfgTable_38, reelDef, payTable_unitbet38, slots.SymbolList, slots.ScatterPosition, uint_bet38)
	sg001bigfg_68 = slots.NewWayGames(
		bigfgTable_68, reelDef, payTable_unitbet68, slots.SymbolList, slots.ScatterPosition, uint_bet68)
	sg001bigfg_88 = slots.NewWayGames(
		bigfgTable_88, reelDef, payTable_unitbet88, slots.SymbolList, slots.ScatterPosition, uint_bet88)

	// lightning
	sg001lightning = lightning.NewLightningGames(lightning.LightningGame, lightning.LightningFreeGame, 15)

	// jackpot
	sg001jackpot = jackpot.NewJackpotGame(JPTrigger, JPGame, jpPayTable, roll_rate, jpBasePool)
)
