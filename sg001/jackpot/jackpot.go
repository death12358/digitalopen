package jackpot

import (
	"github.com/death12358/digitalopn/games"
	"github.com/death12358/digitalopn/games/random"
	"github.com/death12358/digitalopn/games/slots"
	weights "github.com/death12358/digitalopn/games/weight"

	"github.com/shopspring/decimal"
)

// Games Game Structure
type Games struct {
	trigger   *weights.Games
	jp_pools  *weights.Games
	pays      *games.Pays
	symbols   games.Symbol
	roll_rate []decimal.Decimal
	basePool  []decimal.Decimal
}

// NewJackpotGame - 建立 Jackpot Game
func NewJackpotGame(tr, jppool *weights.Games, pays *games.Pays, roll []decimal.Decimal, basepool []decimal.Decimal) *Games {
	return &Games{
		trigger:   tr,
		jp_pools:  jppool,
		pays:      pays,
		symbols:   slots.WW,
		roll_rate: roll,
		basePool:  basepool,
	}
}

// TriggerJackpot - 是否觸發 Jackpot
func (j *Games) TriggerJackpot() bool {
	// dice := random.Intn(j.trigger.Sum())
	// pick, _ := j.trigger.Pick(dice)

	// 1/70
	//	pick := random.Intn(70)
	pick := random.Intn(70)

	if pick == 0 {
		return true
	}
	return false
}

// PickJackpot - 選擇 Jackpot
func (j *Games) PickJackpot() games.Symbol {
	dice := random.Intn(j.jp_pools.Sum())
	pick, _ := j.jp_pools.Pick(dice)
	return games.Symbol(pick)
}

// GetJackpotPay - 取得 Jackpot Pay
func (j *Games) GetJackpotPay(jp games.Symbol) decimal.Decimal {
	return (*j.pays)[jp]
}

// CalcJackpot - 計算 Jackpot
func (j *Games) CalcJackpotMuti(jp games.Symbol, totalbet decimal.Decimal) decimal.Decimal {
	// point = (*j.pays)[jp].Mul(totalbet)
	// return point
	return (*j.pays)[jp].Mul(totalbet)
}

// CalcJackpot - 計算 Jackpot
func (j *Games) CalcJackpot(jp games.Symbol, totalbet, pool decimal.Decimal) decimal.Decimal {
	point := (*j.pays)[jp].Mul(totalbet)
	// point = point.Add(pool)
	// return point
	return point.Add(pool)
}

// CalcRollRate - 計算 Roll Rate
func (j *Games) CalcRollRate(totalbet decimal.Decimal) []decimal.Decimal {
	roll_rate := make([]decimal.Decimal, len(j.roll_rate))
	for i, v := range j.roll_rate {
		if !v.IsZero() {
			roll_rate[i] = v.Mul(totalbet)
		}
	}

	return roll_rate
}

// GetBasePool - 取得 Base Pool
func (j *Games) GetBasePool(totalbet decimal.Decimal) []decimal.Decimal {
	basePool := make([]decimal.Decimal, len(j.basePool))
	for i, v := range j.basePool {
		if !v.IsZero() {
			basePool[i] = v.Mul(totalbet)
		}
	}

	return basePool
}
