package slots_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/shopspring/decimal"
	"gitlab.com/gaas_module/games"
	"gitlab.com/gaas_module/games/slots"
)

var sg001_ng = slots.NewWayGames(ngTable, reelDef, sg001paytable, slots.SymbolList, slots.ScatterPosition, decimal.NewFromInt(8))
var sg002_ng = slots.NewLineGames(ngTable, reelDef, sg001paytable, payline, kWinLines, slots.SymbolList, slots.ScatterPosition, decimal.NewFromInt(1))

var (
	ngTable = games.ReelStripList{
		"98": ngReelStrips,
	}
	fgTable = games.ReelStripList{
		"98": fgReelStrips,
	}
	reelDef = &slots.ReelStripsDef{3, 3, 3, 3, 3}
)

var ngReelStrips = &games.ReelStrips{
	{2, 2, 6, 13, 13, 5, 7, 4, 1, 5, 2, 2, 5, 4, 10, 2, 4, 6, 5, 3, 3, 4, 4, 4, 2, 1, 1, 3, 4, 5, 3, 4, 2, 9, 13, 13, 13, 11, 5, 4, 4, 3, 3, 2, 2, 2, 3, 3, 5, 5, 5, 8, 3, 3, 3, 7, 8, 9, 1, 1, 1, 3, 3},
	{3, 3, 2, 2, 2, 0, 1, 3, 3, 2, 2, 5, 3, 5, 2, 2, 2, 3, 3, 6, 2, 2, 10, 11, 8, 13, 13, 13, 7, 2, 2, 5, 3, 3, 9, 0, 7, 5, 4, 6, 5, 0, 2, 2, 3, 6, 7, 2, 3, 11, 13, 13, 4, 6, 3, 7, 4, 13, 13, 1, 1, 1, 3},
	{11, 5, 9, 8, 1, 1, 1, 4, 4, 0, 5, 9, 2, 13, 13, 1, 1, 5, 4, 8, 1, 10, 13, 13, 13, 4, 5, 1, 1, 4, 5, 1, 1, 4, 4, 4, 5, 1, 1, 0, 5, 7, 8, 5, 5, 4, 4, 1, 5, 5, 5, 2, 9, 3, 10, 4, 11, 3, 4, 5, 6, 4, 0},
	{5, 10, 4, 4, 13, 13, 13, 10, 2, 2, 6, 7, 10, 4, 4, 4, 6, 1, 2, 2, 2, 6, 9, 0, 10, 8, 5, 7, 4, 4, 4, 0, 4, 4, 10, 6, 2, 7, 0, 5, 7, 3, 10, 2, 2, 2, 11, 6, 1, 1, 1, 6, 7, 3, 7, 13, 13, 11, 2, 6, 7, 3, 3},
	{9, 2, 4, 9, 13, 13, 13, 5, 5, 9, 11, 3, 8, 9, 3, 11, 8, 1, 8, 5, 4, 3, 8, 9, 4, 8, 3, 3, 3, 11, 10, 5, 5, 1, 1, 1, 5, 5, 10, 8, 3, 11, 6, 2, 9, 6, 1, 10, 8, 5, 5, 5, 11, 6, 1, 7, 10, 2, 8, 13, 13, 7, 1, 1, 1, 6, 3, 3, 3, 6, 5, 7, 3},
}

var fgReelStrips = &games.ReelStrips{
	{3, 1, 5, 4, 3, 3, 4, 2, 3, 4, 4, 3, 3, 2, 2, 2, 1, 13, 13, 3, 2, 4, 3, 1, 1, 1, 13, 13, 13, 2, 3, 3, 1, 1, 1, 3, 2, 1, 1, 3, 4, 1, 3, 2, 1, 4, 3, 2, 4, 1, 3, 4, 2, 2, 2, 4, 3, 2, 1, 4, 3, 5, 4},
	{3, 1, 2, 2, 3, 3, 13, 13, 2, 2, 2, 4, 1, 1, 1, 5, 4, 2, 3, 1, 5, 4, 2, 3, 4, 5, 0, 4, 2, 3, 3, 2, 2, 4, 1, 2, 3, 4, 2, 5, 3, 4, 4, 5, 3, 1, 4, 13, 13, 13, 5, 5, 3, 4, 1, 5, 4, 3, 2, 2, 0, 1, 1},
	{2, 2, 1, 1, 1, 5, 3, 3, 2, 2, 5, 1, 2, 2, 1, 1, 1, 0, 4, 2, 5, 4, 1, 1, 2, 4, 5, 0, 2, 4, 5, 3, 3, 2, 5, 13, 13, 1, 1, 2, 2, 5, 3, 2, 0, 5, 5, 1, 1, 4, 5, 2, 2, 2, 5, 1, 1, 1, 5, 2, 13, 13, 13},
	{4, 3, 3, 4, 3, 3, 4, 4, 2, 2, 5, 3, 3, 3, 1, 4, 4, 3, 3, 3, 5, 5, 13, 13, 13, 1, 3, 1, 5, 4, 4, 0, 1, 5, 2, 3, 1, 5, 5, 3, 4, 4, 3, 2, 4, 4, 0, 3, 5, 5, 5, 2, 3, 4, 4, 3, 5, 4, 13, 13, 4, 4, 4},
	{2, 2, 1, 3, 5, 5, 2, 2, 3, 4, 5, 3, 4, 1, 1, 3, 2, 2, 2, 3, 4, 4, 4, 1, 3, 5, 5, 1, 3, 1, 5, 3, 1, 1, 3, 4, 4, 3, 5, 5, 2, 2, 3, 4, 5, 5, 4, 4, 2, 2, 2, 3, 1, 1, 1, 3, 3, 3, 1, 1, 1, 13, 13, 13, 1, 1, 5, 4, 4, 4, 1, 1, 5},
}

var sg001paytable = &games.PayTable{
	// WW
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(15), decimal.NewFromInt(25), decimal.NewFromInt(30)},
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

// Win Lines
var payline = &games.PayLine{
	{5, 6, 7, 8, 9}, {0, 1, 2, 3, 4}, {10, 11, 12, 13, 14}, {0, 6, 12, 8, 4}, {10, 6, 2, 8, 14},
	{0, 1, 7, 3, 4}, {10, 11, 7, 13, 14}, {5, 11, 12, 13, 9}, {5, 1, 2, 3, 9}, {0, 6, 7, 8, 4},
	{10, 6, 7, 8, 14}, {0, 6, 2, 8, 4}, {10, 6, 12, 8, 14}, {5, 1, 7, 3, 9}, {5, 11, 7, 13, 9},
	{5, 6, 2, 8, 9}, {5, 6, 12, 8, 9}, {0, 11, 2, 13, 4}, {10, 1, 12, 3, 14}, {5, 1, 12, 3, 9},
	{5, 11, 2, 13, 9}, {0, 1, 12, 3, 4}, {10, 11, 2, 13, 14}, {0, 11, 12, 13, 4}, {10, 1, 2, 3, 14},

	{0, 11, 7, 13, 4}, {10, 1, 7, 3, 14}, {5, 6, 7, 8, 14}, {0, 1, 7, 13, 14}, {10, 11, 7, 3, 4},
	{0, 6, 7, 8, 14}, {10, 6, 7, 8, 4}, {0, 6, 12, 8, 14}, {10, 6, 2, 8, 4}, {0, 1, 2, 3, 9},
	{10, 11, 12, 13, 9}, {0, 6, 2, 8, 14}, {10, 6, 12, 8, 4}, {5, 1, 7, 13, 9}, {5, 11, 7, 3, 9},
	{5, 6, 2, 3, 4}, {5, 6, 12, 13, 14}, {5, 1, 2, 8, 14}, {5, 11, 12, 8, 4}, {5, 1, 7, 13, 14},
	{5, 11, 7, 3, 4}, {10, 6, 2, 3, 9}, {0, 6, 12, 13, 9}, {0, 1, 7, 13, 9}, {10, 11, 7, 3, 9},
}

var kWinLines = games.KWinLines(50)

// Test Fields
func TestFields(t *testing.T) {
	s := strings.Fields("input output")
	t.Logf("Fields are: %q", s)
	createTable("input output output outputoutput")
}

func createTable(input string) {
	// 將文字按照空格分離成多個欄位
	columns := strings.Split(input, " ")

	// 計算每一個欄位的最大長度
	maxLength := 0
	for _, column := range columns {
		length := len(column)
		if length > maxLength {
			maxLength = length
		}
	}

	// 輸出等寬文字表格
	for _, column := range columns {
		fmt.Printf("%*s", -maxLength, column)
		fmt.Print(" | ")
	}
	fmt.Println()
}

// slot game flow
// 1. start round
// 2. spin reels
// 3. check paylines
// 4. pay out
// 5. end round

func TestWayGame(t *testing.T) {
	bet := decimal.NewFromFloat(1)
	reels := games.ReelStrips{
		{3, 5, 4},
		{3, 1, 5},
		{4, 5, 0},
		{5, 5, 3},
		{4, 4, 1},
	}
	normal := sg001_ng.CalcWayReel(reels, bet)
	t.Logf("big way: sg001 normal: %+v\n", normal)
}

func TestLineGame(t *testing.T) {

	/*
		a := decimal.NewFromInt(5)
		b := decimal.NewFromInt(15)
		if a.Cmp(b) == 1 {
			log.Printf("a大于b")
		} else if a.Cmp(b) == -1 {
			log.Printf("a小于b")
		}else if a.Cmp(b) == 0 {
			log.Printf("a等于b")
		}
	*/

	bet := decimal.NewFromFloat(1)
	reels := games.ReelStrips{
		{0, 7, 4},
		{0, 5, 2},
		{0, 3, 1},
		{13, 5, 6},
		{1, 8, 1},
	}
	normal := sg002_ng.CalcLineReel(reels, bet)
	t.Logf("big way: sg001 normal: %+v\n", normal)
}

func Test_Test(t *testing.T) {
	bet := decimal.NewFromFloat(8)
	reels_set := []int{3, 3, 3, 3, 3}
	pos := []int{1, 1, 1, 1, 1}
	spin_ts := ngReelStrips.ShowReelStrips(reels_set, pos, 1, 1)
	t.Logf("spin_ts reels: %+v\n", spin_ts)

	spin := ngReelStrips.ContiguousReelStrips(reels_set, pos)
	t.Logf("spin reels: %+v\n", spin)

	rep_spin_ts := ngReelStrips.ShowRepeatedReelStrips(reels_set, pos, []bool{false, true, true, true, false}, 1, 1)
	t.Logf("rep_spin_ts reels: %+v\n", rep_spin_ts)

	rep_spin := ngReelStrips.RepeatedReelStrips(reels_set, pos, []bool{false, true, true, true, false})
	t.Logf("spin RepeatedReelStrips: %+v\n", rep_spin)

	invSpin := spin.InvertRegularXYAxis()
	t.Logf("invert reels: %+v\n", invSpin)

	invSpinStr := spin.InvertRegularXYAxis().Strings()
	t.Logf("invert reels: %+v\n", invSpinStr)

	sc_match, sc_count := spin.CalcSymbolMatches(slots.H3)
	t.Logf("scatter match: %+v, count: %+v", sc_match, sc_count)

	dup := spin[0].RemoveDuplicates()
	t.Logf("dup reels: %+v\n", dup)

	sstrs := spin.Strings()
	t.Logf("sstrs reels: %+v\n", sstrs)

	str := spin[0].Strings()
	t.Logf("str reels: %+v\n", str)

	match, multi := spin.CalcSymbolsMatchFromLeft(slots.H3, slots.WW)
	t.Logf("match: %+v, multi: %+v\n", match, multi)

	spin = sg001_ng.ContiguousReelStrips("98", pos)
	t.Logf("way: spin reels: %+v\n", spin)

	normal := sg001_ng.CalcWayReel(spin, bet)
	t.Logf("way: sg001 normal: %+v\n", normal)

	sc := sg001_ng.CalcScatter(spin, bet, slots.H3)
	t.Logf("way: sg001 sc: %+v\n", sc)

	spin = sg001_ng.BigSymbolsReelStrips("98", pos, []bool{false, true, true, true, false})
	t.Logf("big way: spin reels: %+v\n", spin)

	normal = sg001_ng.CalcWayReel(spin, bet)
	t.Logf("big way: sg001 normal: %+v\n", normal)

	sc = sg001_ng.CalcScatter(spin, bet, slots.H3)
	t.Logf("big way: sg001 sc: %+v\n", sc)

	nglen := sg001_ng.GetReelsLen("98")
	t.Logf("sg001 nglen: %+v\n", nglen)
}
