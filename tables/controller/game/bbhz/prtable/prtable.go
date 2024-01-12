package prtable

import (
	"github.com/death12358/digitalopen/tables/controller/internal"
)

type PRTableMap map[int]PRTable

type PRTable struct {
	Limit internal.ProbabilityLimit
	mgr   [][]string
}

// func (pr *PRTable) InitTable(inData internal.DataMap) error {

// 	pr.Limit = internal.CreateProbabilityLimit(inData)

// 	for i := 0; i < constants.ReelsLength; i++ {
// 		key := fmt.Sprintf("MGR1_%d", i+1)
// 		reel, ok := inData.GetDataByKey(key)
// 		if !ok {
// 			return errors.New(key + ": 載入失敗")
// 		}
// 		pr.mgr = append(pr.mgr, reel)
// 	}

// 	return nil

// }

// GetGameReels 取遊戲輪帶
func (pr PRTable) GetGameReels() [][]string {
	return pr.mgr
}
