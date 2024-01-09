package adventure

import (
	"github.com/death12358/digitalopen/games"
	"github.com/shopspring/decimal"

	"github.com/death12358/digitalopen/games/random"
	weights "github.com/death12358/digitalopen/games/weight"
)

type Games struct {
	BgobjTable         []int
	BgPointWeightTable [][]int
	pays               *games.Pays
}

// NewLightningGames - create new LightningGames
//
//	@param bg_reel	- bonus game reel weight
//	@param fg		- 中了FG後權重配置
//	@param def		- 盤面大小
func NewAdventureGames(bgobjTable []int, bgpointTable [][]int, P *games.Pays) *Games {
	return &Games{
		BgobjTable:         bgobjTable,
		BgPointWeightTable: bgpointTable,
		pays:               P,
	}
}

func (m *Games) GetFloor(i int) *weights.Games {
	return weights.NewGames(m.BgPointWeightTable[i], m.BgobjTable)
}
func (m *Games) PickCrystal(floor int) int {
	dice := random.Intn(m.GetFloor(floor).Sum())
	pick, _ := m.GetFloor(floor).Pick(dice)
	return pick
}

func (m *Games) PickProcess() []games.Symbol {
	/*record := &games.Records{
		Id:         id,
		Brand:      rc.Brand,
		Username:   rc.Username,
		Username:   rc.Username,
		Username:   rc.Username,
		Case:       games.NotStardYet,
		Stages:     rc.Stages -
		Pickem:     rc.Pickem,,
		Pickem:     rc.Pickem,
		Muiplier: decimal.Zero,
	}*/
	//決幾輪
	process := make([]games.Symbol, 0)
	hp := 2
	for r := 0; r < 6 && hp > 0; r++ {
		n := Symbol(m.PickCrystal(r))
		if n == bgAdv_default {
			hp -= 1
		}
		process = append(process, n)
		//fmt.Println(process)
	}
	return process
}
func (m *Games) CalcAdventure(process []games.Symbol, bet decimal.Decimal) decimal.Decimal {
	var point decimal.Decimal
	for i := 0; i < len(process); i++ {
		point = point.Add((*AdventurePayTable)[process[i]].Mul(bet.Div(unitbet)))
	}
	return point
}
func (m *Games) CalcAdventureOneRound(a games.Symbol, bet decimal.Decimal) decimal.Decimal {
	return (*AdventurePayTable)[a].Mul(bet).Div(unitbet)
}
