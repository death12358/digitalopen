package easteregg

import (
	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/random"
	weights "github.com/death12358/digitalopen/games/weight"
	"github.com/shopspring/decimal"
)

type Games struct {
	EggTable         []int
	EggWeightTable   [][]int
	Pays             *games.Pays
	OtherTable       []int
	OtherWeightTable [][]int
}

func WheelWeight([]games.Symbol) {

}

// @param bg_reel	- bonus game reel weight
// @param fg		- 中了FG後權重配置
// @param def		- 盤面大小
func NewEggGames(bgobjTable []int, bgpointTable [][]int, P *games.Pays, othertable []int, otherweighttable [][]int) *Games {
	return &Games{
		EggTable:         bgobjTable,
		EggWeightTable:   bgpointTable,
		Pays:             P,
		OtherTable:       othertable,
		OtherWeightTable: otherweighttable,
	}
}

func (m *Games) GetEggFloor(i int) *weights.Games {
	return weights.NewGames(m.EggWeightTable[i], m.EggTable)
}

func (m *Games) GetBonusWheel(eggcount int) *weights.Games {
	return weights.NewGames(m.OtherWeightTable[eggcount-2], m.OtherTable)
}
func (m *Games) PickEgg(floor int) int {
	// if floor < 2 {
	// 	return 1
	// }
	dice := random.Intn(m.GetEggFloor(floor).Sum())
	pick, _ := m.GetEggFloor(floor).Pick(dice)
	return pick
}

func (m *Games) PickBonus(eggcount int) int {
	var pick int
	// if eggcount == 2 {
	// 	if random.Intn(1) == 1 {
	// 		pick = 5
	// 	} else {
	// 		pick = 8
	// 	}
	// 	return pick
	// }
	dice := random.Intn(m.GetBonusWheel(eggcount).Sum())
	pick, _ = m.GetBonusWheel(eggcount).Pick(dice)

	return pick
}

func (m *Games) PickProcess() []games.Symbol {

	process := make([]games.Symbol, 0)
	hp := 1
	eggcounts := 0
	//pick egg process
	for floor := 0; floor < 5 && hp > 0; floor++ {
		egg := Symbol(m.PickEgg(floor))
		eggcounts++
		if egg == bgegg_default {
			hp -= 1
			eggcounts--
		}
		process = append(process, egg)
	}
	//play bonuswheel process
	bonus := Symbol(m.PickBonus(eggcounts))
	process = append(process, bonus)
	//fmt.Println(process)
	return process
}

func (m *Games) GetBonus(process []games.Symbol) games.Symbol {
	return process[len(process)-1]
}

func (m *Games) GetSymbolPoints(symbol games.Symbol) decimal.Decimal {
	return (*m.Pays)[symbol]
}

func (m *Games) CalcAllPoints(process []games.Symbol, bet decimal.Decimal) decimal.Decimal {
	points := decimal.Zero
	n := len(process)
	eggcounts := 0
	//egg process
	for _, egg := range process {
		if egg == bgegg_200 {
			eggcounts++
		}
	}

	points = m.GetSymbolPoints(bgegg_200).Mul(decimal.NewFromInt(int64(eggcounts)))
	//bonus wheel
	bonus := process[n-1]
	switch bonus {
	case bgMult_150:
		points = points.Mul(m.GetSymbolPoints(bonus))
	case bgMult_200:
		points = points.Mul(m.GetSymbolPoints(bonus))
	case bgMult_300:
		points = points.Mul(m.GetSymbolPoints(bonus))
	case bgMult_500:
		points = points.Mul(m.GetSymbolPoints(bonus))
	case bgAdd_5:
		points = points.Add(m.GetSymbolPoints(bonus))
	case bgAdd_10:
		points = points.Add(m.GetSymbolPoints(bonus))
	case bgAdd_15:
		points = points.Add(m.GetSymbolPoints(bonus))

	}
	return points.Div(unitbet).Mul(bet)

}
