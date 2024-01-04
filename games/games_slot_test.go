package games_test

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"games/slots"

	"games"
	"games/random"

	"github.com/shopspring/decimal"
)

var (
	r_default = games.Rounds{
		Id:         "1234567890",
		GameCode:   "SG001",
		Brand:      "brand_test",
		Username:   "user_test",
		Status:     games.State(0),
		Position:   games.State(0),
		Stages:     0,
		Result:     games.NewResults(),
		Currency:   "TestCoin",
		Start:      1669596839071688000,
		Fisish:     1669596839071688000,
		TotalBet:   decimal.Zero,
		TotalPoint: decimal.Zero,
	}
)

func BenchmarkSG001(b *testing.B) {
	mock_game := NewMockSlotSG001()
	for i := 0; i < b.N; i++ {
		mock_game.Spin("98", decimal.NewFromFloat(1.0), []string{"8"}, r_default)

	}
}

func BenchmarkSG001_Safe(b *testing.B) {
	mock_game := NewMockSlotSG001()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mock_game.Spin("98", decimal.NewFromFloat(1.0), []string{"8"}, r_default)
		}
	})
}

// example of a test that uses a slot machine sg001.
func TestSlotSG001(t *testing.T) {
	// create a new slot machine
	mock_game := games.NewGames(NewMockSlotSG001())
	t.Logf("GameName: %s", mock_game.Name())
	t.Logf("GameInfo: %s", mock_game.Info())

	// spin
	round, err := mock_game.Spin("98", decimal.NewFromFloat(1.0), []string{"8"}, r_default)
	if err != nil {
		t.Errorf("Spin error: %s", err.Error())
	}
	j_round, err := json.Marshal(round)
	if err != nil {
		t.Errorf("Marshal error: %s", err.Error())
	}
	t.Logf("Round: %s", string(j_round))

	// t.Logf("Round: %v", round)
}

type MockSlotSG001 struct {
	name string
	info string
}

// NewMockSlotSG001 - 建立遊戲定義
func NewMockSlotSG001() *MockSlotSG001 {
	return &MockSlotSG001{
		name: "SG001",
		info: `
		GameName: SG001\n
		Type: Slot\n
		RTP: 0.9800\n
		`,
	}
}

// Name - 遊戲名稱
func (g *MockSlotSG001) Name() string {
	return g.name
}

// Info - 遊戲資訊
func (g *MockSlotSG001) Info() string {
	return g.info
}

// Spin - 遊戲開始
func (g *MockSlotSG001) Spin(rtp string, bet decimal.Decimal, pickem []string, round games.Rounds) (*games.Rounds, error) {
	reels, display := g.SpinNGReelStrips(rtp)

	ng := g.CalcNGResult(reels, "0", round.Brand, round.Username, pickem, display, bet)
	round.Result["0"] = ng
	round.Status = round.Status.Push(ng.Case)
	round.TotalBet = round.TotalBet.Add(bet)
	round.TotalPoint = round.TotalPoint.Add(ng.Point)
	round.Start = time.Now().Unix()

	if g.IsFG(reels) {
		round = *g.FGFlow(rtp, bet, round)
	}
	round.Fisish = time.Now().Unix()

	return &round, nil
}

// SpinNGReelStrips - 滾動 NG Reel Strips
//
//	@param rtp - RTP
//	@return reel - 滾動後的 NG Reel Strips
//	@return client display reel - 滾動後的 NG Reel Strips 的位置
func (g *MockSlotSG001) SpinNGReelStrips(rtp string) (*games.ReelStrips, []string) {
	rtps := games.RTPs(rtp)
	nglen := sg001_ng.GetReelsLen(rtps)
	pos := random.Intsn(nglen)
	// // FG Test
	// pos = []int{0, 5, 9, 23, 0}

	reelStrips := sg001_ng.ContiguousReelStrips(rtps, pos)
	show_reelStrips := sg001_ng.ShowReelStrips(rtps, pos, 1, 1).InvertRegularXYAxis().Strings()

	return &reelStrips, show_reelStrips
}

// go mod edit -replace [old git package]@[version]=[new git package]@[version]

// SpinFGReelStrips - 滾動 FG Reel Strips
//
//	@param rtp - RTP
//	@return reel - 滾動後的 FG Reel Strips
//	@return client display reel - 滾動後的 FG Reel Strips 的位置
func (g *MockSlotSG001) SpinFGReelStrips(rtp string) (*games.ReelStrips, []string) {
	rtps := games.RTPs(rtp)
	fglen := sg001_fg.GetReelsLen(rtps)
	pos := random.Intsn(fglen)

	reelStrips := sg001_fg.ContiguousReelStrips(rtps, pos)
	show_reelStrips := sg001_fg.ShowReelStrips(rtps, pos, 1, 1).InvertRegularXYAxis().Strings()

	return &reelStrips, show_reelStrips
}

// CalcNGResult - 計算 NG 結果
//
//	@param reel	- 滾動後的 NG Reel Strips
//	@param id 		- 遊戲 ID
//	@param brand	- 廠商
//	@param user	- 使用者
//	@param bet		- 下注金額
//	@return result	- NG 結果
func (g *MockSlotSG001) CalcNGResult(reel *games.ReelStrips, id, brand, user string, pickem, symbols []string, bet decimal.Decimal) *games.Records {
	record := games.Records{
		Id:       id,
		Brand:    brand,
		Username: user,
		Pickem:   pickem,
		Symbols:  symbols,
		Bet:      bet,
	}

	ng_result := sg001_ng.CalcWayReel(*reel, bet)
	if ng_result.IsWinGreaterThanZero() {
		record.Case = games.Win
		record.Multiplier = ng_result.TotalWin.Div(bet)
		record.Point = ng_result.TotalWin
	}

	return &record
}

// FGFlow - FG 流程
//
//	@param rtp - RTP
//	@param bet - 下注金額
//	@param round - Round
//	@return round - Round
func (g *MockSlotSG001) FGFlow(rtp string, bet decimal.Decimal, round games.Rounds) *games.Rounds {
	count := 6
	round.Status = round.Status.Push(games.FreeGame)
	for i := 0; i < count; i++ {
		fg_reel, show := g.SpinFGReelStrips(rtp)
		round.Stages++
		id := strconv.Itoa(int(round.Stages))
		res := g.CalcFGResult(fg_reel, id, round.Brand, round.Username, show, bet)
		res.Stages = round.Stages
		round.Result[id] = res
		round.TotalPoint = round.TotalPoint.Add(res.Point)
		if g.IsFG(fg_reel) {
			count += 6
		}
	}

	return &round
}

// IsFG - 是否 FG
//
//	@param reel *slots.ReelStrips
//	@return bool
func (g *MockSlotSG001) IsFG(reel *games.ReelStrips) bool {
	sc := sg001_fg.CalcScatter(*reel, decimal.Zero, slots.WW)

	if sc.Wins[0].Match >= 3 {
		return true
	}
	return false
}

// CalcFGResult - 計算 FG 結果
//
//	@param reel	- 滾動後的 FG Reel Strips
//	@param id 		- 遊戲 ID
//	@param brand	- 廠商
//	@param user	- 使用者
//	@param bet		- 下注金額
//	@return result	- FG 結果
func (g *MockSlotSG001) CalcFGResult(reel *games.ReelStrips, id, brand, user string, symbols []string, bet decimal.Decimal) *games.Records {
	record := games.Records{
		Id:       id,
		Brand:    brand,
		Username: user,
		Case:     games.FreeGame,
		Symbols:  symbols,
	}

	fg_result := sg001_fg.CalcWayReel(*reel, bet)
	if fg_result.IsWinGreaterThanZero() {
		// record.Case = games.Win
		record.Multiplier = fg_result.TotalWin.Div(bet)
		record.Point = fg_result.TotalWin
	}

	return &record
}

// NGSpin -
func (g *MockSlotSG001) NGSpin(rtp string, bet decimal.Decimal, pickem []string, round games.Rounds) (*games.Rounds, error) {
	rtps := games.RTPs(rtp)
	nglen := sg001_ng.GetReelsLen(rtps)
	pos := random.Intsn(nglen)

	record := games.Records{
		Id:       "0",
		Brand:    round.Brand,
		Username: round.Username,
		Pickem:   pickem,
	}

	reel := sg001_ng.ContiguousReelStrips(rtps, pos)
	// 取前後各一個
	show_reel := sg001_ng.ShowReelStrips(rtps, pos, 1, 1)
	record.Symbols = show_reel.InvertRegularXYAxis().Strings()
	normal := sg001_ng.CalcWayReel(reel, bet)

	if normal.IsWinGreaterThanZero() {
		round.Status = round.Status.Push(games.Win)
		record.Case = games.Win
		record.Multiplier = normal.TotalWin.Div(bet)
		record.Bet = bet
		record.Point = normal.TotalWin
		round.TotalPoint = round.TotalPoint.Add(normal.TotalWin)
	}

	round.Result["0"] = &record
	round.TotalBet = round.TotalBet.Add(bet)

	return &round, nil
}

var sg001_ng = slots.NewWayGames(ngTable, reelDef, sg001paytable, slots.SymbolList, slots.ScatterPosition, decimal.NewFromInt(8))

var sg001_fg = slots.NewWayGames(fgTable, reelDef, sg001paytable, slots.SymbolList, slots.ScatterPosition, decimal.NewFromInt(8))

var (
	ngTable = games.ReelStripList{
		"98": ngReelStrips98,
	}
	fgTable = games.ReelStripList{
		"98": fgReelStrips98,
	}
	reelDef = &slots.ReelStripsDef{3, 3, 3, 3, 3}
)

var ngReelStrips98 = &games.ReelStrips{
	{2, 2, 6, 13, 13, 5, 7, 4, 1, 5, 2, 2, 5, 4, 10, 2, 4, 6, 5, 3, 3, 4, 4, 4, 2, 1, 1, 3, 4, 5, 3, 4, 2, 9, 13, 13, 13, 11, 5, 4, 4, 3, 3, 2, 2, 2, 3, 3, 5, 5, 5, 8, 3, 3, 3, 7, 8, 9, 1, 1, 1, 3, 3},
	{3, 3, 2, 2, 2, 0, 1, 3, 3, 2, 2, 5, 3, 5, 2, 2, 2, 3, 3, 6, 2, 2, 10, 11, 8, 13, 13, 13, 7, 2, 2, 5, 3, 3, 9, 0, 7, 5, 4, 6, 5, 0, 2, 2, 3, 6, 7, 2, 3, 11, 13, 13, 4, 6, 3, 7, 4, 13, 13, 1, 1, 1, 3},
	{11, 5, 9, 8, 1, 1, 1, 4, 4, 0, 5, 9, 2, 13, 13, 1, 1, 5, 4, 8, 1, 10, 13, 13, 13, 4, 5, 1, 1, 4, 5, 1, 1, 4, 4, 4, 5, 1, 1, 0, 5, 7, 8, 5, 5, 4, 4, 1, 5, 5, 5, 2, 9, 3, 10, 4, 11, 3, 4, 5, 6, 4, 0},
	{5, 10, 4, 4, 13, 13, 13, 10, 2, 2, 6, 7, 10, 4, 4, 4, 6, 1, 2, 2, 2, 6, 9, 0, 10, 8, 5, 7, 4, 4, 4, 0, 4, 4, 10, 6, 2, 7, 0, 5, 7, 3, 10, 2, 2, 2, 11, 6, 1, 1, 1, 6, 7, 3, 7, 13, 13, 11, 2, 6, 7, 3, 3},
	{9, 2, 4, 9, 13, 13, 13, 5, 5, 9, 11, 3, 8, 9, 3, 11, 8, 1, 8, 5, 4, 3, 8, 9, 4, 8, 3, 3, 3, 11, 10, 5, 5, 1, 1, 1, 5, 5, 10, 8, 3, 11, 6, 2, 9, 6, 1, 10, 8, 5, 5, 5, 11, 6, 1, 7, 10, 2, 8, 13, 13, 7, 1, 1, 1, 6, 3, 3, 3, 6, 5, 7, 3},
}

var fgReelStrips98 = &games.ReelStrips{
	{3, 1, 5, 4, 3, 3, 4, 2, 3, 4, 4, 3, 3, 2, 2, 2, 1, 13, 13, 3, 2, 4, 3, 1, 1, 1, 13, 13, 13, 2, 3, 3, 1, 1, 1, 3, 2, 1, 1, 3, 4, 1, 3, 2, 1, 4, 3, 2, 4, 1, 3, 4, 2, 2, 2, 4, 3, 2, 1, 4, 3, 5, 4},
	{3, 1, 2, 2, 3, 3, 13, 13, 2, 2, 2, 4, 1, 1, 1, 5, 4, 2, 3, 1, 5, 4, 2, 3, 4, 5, 0, 4, 2, 3, 3, 2, 2, 4, 1, 2, 3, 4, 2, 5, 3, 4, 4, 5, 3, 1, 4, 13, 13, 13, 5, 5, 3, 4, 1, 5, 4, 3, 2, 2, 0, 1, 1},
	{2, 2, 1, 1, 1, 5, 3, 3, 2, 2, 5, 1, 2, 2, 1, 1, 1, 0, 4, 2, 5, 4, 1, 1, 2, 4, 5, 0, 2, 4, 5, 3, 3, 2, 5, 13, 13, 1, 1, 2, 2, 5, 3, 2, 0, 5, 5, 1, 1, 4, 5, 2, 2, 2, 5, 1, 1, 1, 5, 2, 13, 13, 13},
	{4, 3, 3, 4, 3, 3, 4, 4, 2, 2, 5, 3, 3, 3, 1, 4, 4, 3, 3, 3, 5, 5, 13, 13, 13, 1, 3, 1, 5, 4, 4, 0, 1, 5, 2, 3, 1, 5, 5, 3, 4, 4, 3, 2, 4, 4, 0, 3, 5, 5, 5, 2, 3, 4, 4, 3, 5, 4, 13, 13, 4, 4, 4},
	{2, 2, 1, 3, 5, 5, 2, 2, 3, 4, 5, 3, 4, 1, 1, 3, 2, 2, 2, 3, 4, 4, 4, 1, 3, 5, 5, 1, 3, 1, 5, 3, 1, 1, 3, 4, 4, 3, 5, 5, 2, 2, 3, 4, 5, 5, 4, 4, 2, 2, 2, 3, 1, 1, 1, 3, 3, 3, 1, 1, 1, 13, 13, 13, 1, 1, 5, 4, 4, 4, 1, 1, 5},
}

var sg001paytable = &games.PayTable{
	// WW
	games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
	// H1
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20)},
	// H2
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20)},
	// H3
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20)},
	// H4
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20)},
	// H5
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(8), decimal.NewFromInt(10)},
	// LA
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(8), decimal.NewFromInt(10)},
	// LK
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(8), decimal.NewFromInt(10)},
	// LQ
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(8), decimal.NewFromInt(10)},
	// LJ
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(8), decimal.NewFromInt(10)},
	// LT
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(8), decimal.NewFromInt(10)},
	// LN
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(8), decimal.NewFromInt(10)},
	// SE
	games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
	// SF
	games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
	// SB
	games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
}
