package battle

import (
	"github.com/death12358/digitalopen/games"
	"github.com/shopspring/decimal"
)

func StartBattle(MonsterNumberOnBoard, MonsterNumberOnRing map[string]int) (map[string]int, decimal.Decimal, bool) {
	MonsterNumberOnRing = GotoRing(MonsterNumberOnBoard, MonsterNumberOnRing)
	if BattleIsLose(MonsterNumberOnRing) {
		return MonsterNumberOnRing, decimal.Zero, true
	}
	MonsterNumberOnRing, BattlePoint := KillMonsters(MonsterNumberOnBoard["H1"], MonsterNumberOnRing)
	return MonsterNumberOnRing, BattlePoint, false
}

func GotoRing(MonsterNumberOnBoard, MonsterNumberOnRing map[string]int) map[string]int {
	if MonsterNumberOnBoard["H1"] > 0 {
		MonsterNumberOnRing["H2"] += MonsterNumberOnBoard["H2"]
		MonsterNumberOnRing["H3"] += MonsterNumberOnBoard["H3"]
		MonsterNumberOnRing["H4"] += MonsterNumberOnBoard["H4"]
		MonsterNumberOnRing["H5"] += MonsterNumberOnBoard["H5"]
	}
	return MonsterNumberOnRing
}

func HowManyMonsters(MonsterNumberOnRing map[string]int) int {
	OnRingNumber := 0
	for _, v := range MonsterNumberOnRing {
		OnRingNumber += v
	}
	return OnRingNumber
}

func BattleIsLose(MonsterNumberOnRing map[string]int) bool {
	return HowManyMonsters(MonsterNumberOnRing) >= 16
}

func BattleIsWin(MonsterNumberOnRing map[string]int) bool {
	return HowManyMonsters(MonsterNumberOnRing) == 0
}

func KillMonsters(H1Number int, MonsterNumberOnRing map[string]int) (map[string]int, decimal.Decimal) {
	BattlePoint := decimal.Zero
	for i := 0; i < H1Number; i++ {
		if MonsterNumberOnRing["H5"] != 0 {
			BattlePoint = BattlePoint.Add(CountBattlePoint("H5", MonsterNumberOnRing["H5"]))
			MonsterNumberOnRing["H5"] = 0
			continue
		}
		if MonsterNumberOnRing["H4"] != 0 {
			BattlePoint = BattlePoint.Add(CountBattlePoint("H4", MonsterNumberOnRing["H4"]))
			MonsterNumberOnRing["H4"] = 0
			continue
		}
		if MonsterNumberOnRing["H3"] != 0 {
			BattlePoint = BattlePoint.Add(CountBattlePoint("H3", MonsterNumberOnRing["H3"]))
			MonsterNumberOnRing["H3"] = 0
			continue
		}
		if MonsterNumberOnRing["H2"] != 0 {
			BattlePoint = BattlePoint.Add(CountBattlePoint("H2", MonsterNumberOnRing["H2"]))
			MonsterNumberOnRing["H2"] = 0
			continue
		}
	}
	return MonsterNumberOnRing, BattlePoint
}

func CountBattlePoint(WhichMonster string, MonsterNumber int) decimal.Decimal {
	if WhichMonster == "H2" {
		return MonsterH2Point[MonsterNumber]
	}
	if WhichMonster == "H3" {
		return MonsterH3Point[MonsterNumber]
	}
	if WhichMonster == "H4" {
		return MonsterH4Point[MonsterNumber]

	}
	if WhichMonster == "H5" {
		return MonsterH5Point[MonsterNumber]
	}
	return decimal.Zero
}

func CountMonsters(reel *games.ReelStrips) map[string]int { //int { //
	MonstersNumber := make(map[string]int)
	for _, r := range *reel {
		for _, v := range r {
			switch v {
			case 1:
				MonstersNumber["H1"]++
			case 2:
				MonstersNumber["H2"]++
			case 3:
				MonstersNumber["H3"]++
			case 4:
				MonstersNumber["H4"]++
			case 5:
				MonstersNumber["H5"]++
			default:
			}
		}
	}

	return MonstersNumber
}
