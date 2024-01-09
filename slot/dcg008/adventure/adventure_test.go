package adventure

import (
	"fmt"
	"testing"

	"github.com/death12358/digitalopen/games"
	"github.com/shopspring/decimal"
)

var (
	r_default = games.Rounds{
		Id:         "1234567890",
		GameCode:   "DCG008",
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
	for i := 0; i < 1000000; i++ {
		m := NewAdventureGames(BgAdvPOTable, AdvPWTable, AdventurePayTable)
		process := m.PickProcess()
		//fmt.Printf("process: %+v\n ", process)

		Totalpoint = Totalpoint.Add(m.CalcAdventure(process, decimal.NewFromFloat(1)))

	}
	fmt.Printf("Totalpoint: %+v\n ", Totalpoint.Div(decimal.NewFromInt(1000000)))
}
