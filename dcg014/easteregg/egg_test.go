package easteregg

import (
	"fmt"
	"testing"

	"github.com/death12358/digitalopen/games"
	"github.com/shopspring/decimal"
)

var (
	r_default = games.Rounds{
		Id:         "1234567890",
		GameCode:   "DCG011",
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

func Test_adventure(t *testing.T) {
	var Totalpoint decimal.Decimal
	freq := make(map[float64]int)
	freq2 := make(map[int]int)

	for i := 0; i < 1000000; i++ {
		eggcounts := 0
		m := NewEggGames(Bg_EggNameTable, Bg_EggWeightTable, PayTable, Bg_BonusNameTable, Bg_BonusWeightTable)
		process := m.PickProcess()
		for _, egg := range process {
			if egg == bgegg_200 {
				eggcounts++
			}
		}
		flag := eggcounts*10 + process[len(process)-1].Int()
		point := m.CalcAllPoints(process, decimal.NewFromFloat(1))
		freq2[flag]++
		P, _ := point.Float64()
		freq[P]++
		Totalpoint = Totalpoint.Add(point)
	}
	for c, f := range freq {
		fmt.Printf("%+v: %+v\n", c, f)
	}
	fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	for c, f := range freq2 {
		fmt.Printf("%+v: %+v\n", c, f)
	}
	fmt.Printf("Totalpoint: %+v\n ", Totalpoint.Div(decimal.NewFromInt(1000000)))
}
