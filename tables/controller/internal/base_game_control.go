package internal

import (
	"fmt"
)

type BaseGameControl[P ProbabilityTableMaker] struct {
	Pay                 *PayTable //賠付格式固定
	ProbabilitySettings map[int]string
	ProbabilityTable    map[string]P //機率表格式依遊戲設定而有所不同
	LimitControl        int
}

type ProbabilityTableMaker interface {
	InitTable(inData DataMap) error
}

// // InitPay 初始化賠付
// func (bgc *BaseGameControl[P]) InitPay(dir Folder, payStand int) (err error) {
// 	bgc.Pay, err = dir.ReadPayData(payStand)
// 	return
// }

// // InitProbabilitySetting 初始化機率表設定檔
// func (bgc *BaseGameControl[P]) InitProbabilitySetting(dir Folder) (err error) {
// 	bgc.ProbabilitySettings, err = dir.ReadMappingData()
// 	return
// }

// // SetProbabilityTable 載入機率表
// func (bgc *BaseGameControl[P]) InsertProbabilityTable(dir Folder, file string, pr P) (err error) {
// 	probData, err := dir.ReadProbabilityData(file)
// 	if err != nil {
// 		return
// 	}
// 	probabilityTable, err := bgc.serializeData(probData, pr)
// 	if err != nil {
// 		err = fmt.Errorf("file:[%s], %v", file, err)
// 		return
// 	}

// 	bgc.setProbabilityTable(file, probabilityTable)
// 	return
// }

func NewBGC[P ProbabilityTableMaker]() BaseGameControl[P] {
	base := BaseGameControl[P]{
		ProbabilityTable: make(map[string]P),
	}

	return base
}

// serializeData 將機率表內容轉至結構
func (bgc *BaseGameControl[P]) serializeData(data DataMap, pr P) (rslt P, err error) {
	if err = pr.InitTable(data); err == nil {
		rslt = pr
	}
	return
}

func (bgc *BaseGameControl[P]) setProbabilityTable(key string, value P) {
	bgc.ProbabilityTable[key] = value
}

// GetProbabilityTable 取得機率表
func (bgc *BaseGameControl[P]) GetProbabilityTable(inPRID int) (file string, id int, table P, err error) {
	id = inPRID

	file, ok := bgc.ProbabilitySettings[id]

	if !ok {
		err = fmt.Errorf("in PRID:[%d] 不存在", id)
	}

	rslt, ok := bgc.ProbabilityTable[file]

	if ok {
		table = rslt
	} else {
		err = fmt.Errorf("in File:[%s] 不存在", file)
	}

	return
}
