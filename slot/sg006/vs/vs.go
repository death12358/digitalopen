package vs

import (
	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/random"
	weights "github.com/death12358/digitalopen/games/weight"
	"github.com/shopspring/decimal"
)

// Games Game Structure
type Games struct {
	vs_reel *weights.Games
	pays    *games.Pays
}

// NewVSGame - 建立 VS Game
func NewGame(vsreel *weights.Games, pays *games.Pays) *Games {
	return &Games{
		vs_reel: vsreel,
		pays:    pays,
	}
}

// PickReward - 轉轉輪決定獎勵
func (vs *Games) PickReward() games.Symbol {
	dice := random.Intn(vs.vs_reel.Sum())
	pick, _ := vs.vs_reel.Pick(dice)
	return games.Symbol(pick)
}

// GetJackpotPay - 取得 Jackpot Pay
func (v *Games) GetVSPay(jp games.Symbol) decimal.Decimal {
	return (*v.pays)[jp]
}

// CalcJackpot - 計算 Jackpot
func (j *Games) CalcJackpot(jp games.Symbol, totalbet, pool decimal.Decimal) decimal.Decimal {
	point := (*j.pays)[jp].Mul(totalbet)
	// point = point.Add(pool)
	// return point
	return point.Add(pool)
}

func CreateVSReel(rtp string, reel *games.ReelStrips) [5]int {
	return VSPosition(rtp, HiSymbolsOnEachReel(reel))
}
func HiSymbolsOnEachReel(reel *games.ReelStrips) map[string]string { //int { //
	HiSymbolOnReel := make(map[string]string)
	for i, s := range *reel {
		switch i {
		case 1:
			for _, sym := range s {
				switch sym {
				case 2:
					HiSymbolOnReel["R2"] = "H2"
				case 3:
					HiSymbolOnReel["R2"] = "H3"
				case 4:
					HiSymbolOnReel["R2"] = "H4"
				case 5:
					HiSymbolOnReel["R2"] = "H5"
				default:

				}
			}
		case 2:
			for _, sym := range s {
				switch sym {
				case 2:
					HiSymbolOnReel["R3"] = "H2"
				case 3:
					HiSymbolOnReel["R3"] = "H3"
				case 4:
					HiSymbolOnReel["R3"] = "H4"
				case 5:
					HiSymbolOnReel["R3"] = "H5"
				default:
				}
			}
		case 3:
			for _, sym := range s {
				switch sym {
				case 2:
					HiSymbolOnReel["R4"] = "H2"
				case 3:
					HiSymbolOnReel["R4"] = "H3"
				case 4:
					HiSymbolOnReel["R4"] = "H4"
				case 5:
					HiSymbolOnReel["R4"] = "H5"
				default:
				}
			}

		}

	}

	return HiSymbolOnReel
}

func VSPosition(rtp string, HiSymbolOnReel map[string]string) [5]int {
	var VSReel [5]int
	if HiSymbolOnReel["R2"] != "" && HiToVSHi(rtp) == 1 {
		VSReel[1] = 1
	}
	if HiSymbolOnReel["R3"] != "" && HiToVSHi(rtp) == 1 {
		VSReel[2] = 1
	}
	if HiSymbolOnReel["R4"] != "" && HiToVSHi(rtp) == 1 {
		VSReel[3] = 1
	}
	return VSReel
}

func HiToVSHi(rtp string) int {
	dice := random.Intn(vs_HiToVSHiGame[rtp].Sum())
	pick, _ := vs_HiToVSHiGame[rtp].Pick(dice)
	return pick
}

func TriggerFGinVS(jp games.Symbol) bool {
	return jp == vs_Battle || jp == vs_JPMiniPlusBattle
}
