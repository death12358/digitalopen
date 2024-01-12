package tables

// import (
// 	"encoding/json"
// 	"fmt"
// 	"testing"

// 	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
// )

// var (
// 	host, password, port = "0000.00.00", "", 6379
// 	scriptDefinition     = "Bft|0.0.1"
// )
// var err error

// func TestSetTable2(t *testing.T) {
// 	TableInit2()
// 	m, _ := NewTableManager(host, password, port)
// 	m.SetFishPayTable(FishPayTable)
// 	m.SetFishDeadTable(FishDeadProb, DeadTableMaps)
// }
// func TestAutoWriteRedisData(t *testing.T) {
// 	AutoWriteRedisData(host, password, port)
// }

// func TestSetTable(t *testing.T) {
// 	TableInit2()
// 	m, _ := NewTableManager(host, password, port)
// 	m.SetFishPayTable(FishPayTable)
// 	m.SetFishDeadTable(FishDeadProb, DeadTableMaps)
// }

// func TestGetlastNumberFromKey(t *testing.T) {
// 	Test_GetlastNumberFromKey()
// }

// func TestGetFishPayTable(t *testing.T) {
// 	GetFishPayTable()
// 	js, err := json.Marshal(FishPayTable)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("FishPayTable:%#v", string(js))
// }

// func TestGetFitTable(t *testing.T) {
// 	Test_GetFitTable()
// }

// func TestGetFishDeadProb(t *testing.T) {
// 	GetFishDeadProb()

// 	js, err := json.Marshal(FishDeadProb)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for i := 1; i <= 5; i++ {
// 		fmt.Printf("FishDeadTable_Flow(%d):%+v\n", i, FishDeadProb[config.RTPFlowTypeID(i)])
// 	}
// 	fmt.Printf("js_FishDeadProb:%+v\n", string(js))
// }

// //	func TestGetFitTable(t *testing.T) {
// //		TestPayTable.GetFitTable(25)
// //	}

// func TestDeadTableMap_GetExpectPay(t *testing.T) {
// 	deadTableMap := DeadTableMap{
// 		13: &DeadTable{
// 			ExpectedRTP: 97.0,
// 		},
// 	}

// 	payTableMap := PayTableMap{
// 		13: &PayTable{
// 			TableWeight: map[int32]int32{
// 				1: 100,
// 				2: 100,
// 			},
// 			IntervalWeight: map[string]map[int32]int32{
// 				"13_1": {
// 					1: 100,
// 					2: 100,
// 				},
// 				"13_2": {
// 					1: 100,
// 					2: 100,
// 				},
// 			},
// 			PayIntervals: map[string][2]int64{
// 				"13_1_1": {1, 9},
// 				"13_1_2": {10, 20},
// 				"13_2_1": {1, 5},
// 				"13_2_2": {6, 10},
// 			},
// 			FGTimesWeight: map[int32]int32{
// 				1: 100,
// 				2: 100,
// 			},
// 			FGTimesObject: map[int32]int{
// 				1: 10,
// 				3: 10,
// 			},
// 		},
// 	}

// 	(&deadTableMap).getExpectPay(payTableMap)
// 	js, _ := json.Marshal(deadTableMap)
// 	fmt.Printf("deadTableMap:%+v\n", string(js))
// }
