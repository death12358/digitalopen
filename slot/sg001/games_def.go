package sg001

import (
	"fmt"

	"github.com/death12358/digitalopen/games/slots"
	"github.com/death12358/digitalopen/slot/sg001/jackpot"
	"github.com/death12358/digitalopen/slot/sg001/lightning"

	"github.com/death12358/digitalopen/games"

	"github.com/shopspring/decimal"
)

// var (
// 	ngReelStrips97_8    games.ReelStrips
// 	hifgReelStrips97_8  games.ReelStrips
// 	bigfgReelStrips97_8 games.ReelStrips

// 	ngReelStrips97_18    games.ReelStrips
// 	hifgReelStrips97_18  games.ReelStrips
// 	bigfgReelStrips97_18 games.ReelStrips
// )

func init() {
	fmt.Println(2)
	var err error
	probTable, err = games.GetProbability("D:\\Golang\\src\\github.com\\death12358\\digitalopen\\slot\\sg001\\bbhz")
	if err != nil {
		fmt.Printf("GetProbability ERR: %#v", err)
	}
	ngReelStrips97_8 = games.ReelStrips{probTable["97_MGR1_1"], probTable["97_MGR1_2"], probTable["97_MGR1_3"], probTable["97_MGR1_4"], probTable["97_MGR1_5"]}
	hifgReelStrips97_8 = games.ReelStrips{probTable["97_HIGHFGR1_1"], probTable["97_HIGHFGR1_2"], probTable["97_HIGHFGR1_1"], probTable["97_HIGHFGR1_1"], probTable["97_BIGFGR1_1"]}
	bigfgReelStrips97_8 = games.ReelStrips{probTable["97_BIGFGR1_1"], probTable["97_BIGFGR1_2"], probTable["97_BIGFGR1_1"], probTable["97_BIGFGR1_1"], probTable["97_BIGFGR1_1"]}

	ngReelStrips97_18 = games.ReelStrips{probTable["97_MGR2_1"], probTable["97_MGR2_2"], probTable["97_MGR2_1"], probTable["97_MGR2_1"], probTable["97_MGR2_1"]}
	hifgReelStrips97_18 = games.ReelStrips{probTable["97_HIGHFGR2_1"], probTable["97_HIGHFGR2_2"], probTable["97_HIGHFGR2_1"], probTable["97_HIGHFGR2_1"], probTable["97_HIGHFGR2_1"]}
	bigfgReelStrips97_18 = games.ReelStrips{probTable["97_BIGFGR1_1"], probTable["97_BIGFGR1_2"], probTable["97_BIGFGR1_1"], probTable["97_BIGFGR1_1"], probTable["97_BIGFGR1_1"]}

	// ng
	*sg001ng_8 = *slots.NewWayGames(
		ngTable_8, reelDef, payTable_unitbet8, slots.SymbolList, slots.ScatterPosition, uint_bet8)
	sg001ng_18 = slots.NewWayGames(
		ngTable_18, reelDef, payTable_unitbet18, slots.SymbolList, slots.ScatterPosition, uint_bet18)
	sg001ng_38 = slots.NewWayGames(
		ngTable_38, reelDef, payTable_unitbet38, slots.SymbolList, slots.ScatterPosition, uint_bet38)
	sg001ng_68 = slots.NewWayGames(
		ngTable_68, reelDef, payTable_unitbet68, slots.SymbolList, slots.ScatterPosition, uint_bet68)
	sg001ng_88 = slots.NewWayGames(
		ngTable_88, reelDef, payTable_unitbet88, slots.SymbolList, slots.ScatterPosition, uint_bet88)
	fmt.Println(2.5)

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
	sg001lightning_8 = lightning.NewLightningGames(lightning.LightningGame_8, lightning.LightningFreeGame_8, 15)

	sg001lightning_18 = lightning.NewLightningGames(lightning.LightningGame_18, lightning.LightningFreeGame_18, 15)

	sg001lightning_38 = lightning.NewLightningGames(lightning.LightningGame_38, lightning.LightningFreeGame_38, 15)
	sg001lightning_68 = lightning.NewLightningGames(lightning.LightningGame_68, lightning.LightningFreeGame_68, 15)

	sg001lightning_88 = lightning.NewLightningGames(lightning.LightningGame_88, lightning.LightningFreeGame_88, 15)

	// jackpot
	sg001jackpot = jackpot.NewJackpotGame(JPTrigger, JPGame, jpPayTable, roll_rate, jpBasePool)
}

var (
	// 牌面定義
	reelDef = &slots.ReelStripsDef{3, 3, 3, 3, 3}

	// unit bet
	uint_bet8  = decimal.NewFromInt(8)
	uint_bet18 = decimal.NewFromInt(18)
	uint_bet38 = decimal.NewFromInt(38)
	uint_bet68 = decimal.NewFromInt(68)
	uint_bet88 = decimal.NewFromInt(88)

	probTable = make(map[string]games.Reels)

	// 轉輪表
	ngTable_8 = games.ReelStripList{
		// "98": &ngReelStrips98_8,
		// "88": &ngReelStrips88_8,
		"97": &ngReelStrips97_8,
	}
	hifgTable_8 = games.ReelStripList{
		// "98": &hifgReelStrips98_8,
		"97": &hifgReelStrips97_8,
		// "88": &hifgReelStrips88_8,
	}
	bigfgTable_8 = games.ReelStripList{
		// "98": &bigfgReelStrips98_8,
		"97": &bigfgReelStrips97_8,
		// "88": &bigfgReelStrips88_8,
	}

	ngTable_18 = games.ReelStripList{
		"98": &ngReelStrips98_18,
		"97": &ngReelStrips97_18,
		"88": &ngReelStrips88_18,
	}
	hifgTable_18 = games.ReelStripList{
		"98": &hifgReelStrips98_18,
		"97": &hifgReelStrips97_18,
		"88": &hifgReelStrips88_18,
	}
	bigfgTable_18 = games.ReelStripList{
		"98": &bigfgReelStrips98_18,
		"97": &bigfgReelStrips97_18,
		"88": &bigfgReelStrips88_18,
	}

	ngTable_38 = games.ReelStripList{
		"98": &ngReelStrips98_38,
		"88": &ngReelStrips88_38,
	}
	hifgTable_38 = games.ReelStripList{
		"98": &hifgReelStrips98_38,
		"88": &hifgReelStrips88_38,
	}
	bigfgTable_38 = games.ReelStripList{
		"98": &bigfgReelStrips98_38,
		"88": &bigfgReelStrips88_38,
	}

	ngTable_68 = games.ReelStripList{
		"98": &ngReelStrips98_68,
		"88": &ngReelStrips88_68,
	}
	hifgTable_68 = games.ReelStripList{
		"98": &hifgReelStrips98_68,
		"88": &hifgReelStrips88_68,
	}
	bigfgTable_68 = games.ReelStripList{
		"98": &bigfgReelStrips98_68,
		"88": &bigfgReelStrips88_68,
	}

	ngTable_88 = games.ReelStripList{
		"98": &ngReelStrips98_88,
		"88": &ngReelStrips88_88,
	}
	hifgTable_88 = games.ReelStripList{
		"98": &hifgReelStrips98_88,
		"88": &hifgReelStrips88_88,
	}
	bigfgTable_88 = games.ReelStripList{
		"98": &bigfgReelStrips98_88,
		"88": &bigfgReelStrips88_88,
	}
)

var (
	// ng
	// sg001ng_8 = new(slots.WayGames)
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
	sg001lightning_8 = lightning.NewLightningGames(lightning.LightningGame_8, lightning.LightningFreeGame_8, 15)

	sg001lightning_18 = lightning.NewLightningGames(lightning.LightningGame_18, lightning.LightningFreeGame_18, 15)

	sg001lightning_38 = lightning.NewLightningGames(lightning.LightningGame_38, lightning.LightningFreeGame_38, 15)
	sg001lightning_68 = lightning.NewLightningGames(lightning.LightningGame_68, lightning.LightningFreeGame_68, 15)

	sg001lightning_88 = lightning.NewLightningGames(lightning.LightningGame_88, lightning.LightningFreeGame_88, 15)

	// jackpot
	sg001jackpot = jackpot.NewJackpotGame(JPTrigger, JPGame, jpPayTable, roll_rate, jpBasePool)
)
