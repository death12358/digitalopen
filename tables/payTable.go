package tables

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"github.com/adimax2953/bftrtpmodel/bft302prob/path"
// 	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
// 	LogTool "github.com/adimax2953/log-tool"
// )

// func GetFishPayTable() (err error) {
// 	// Excel_PayTable_random, err := GetExcelData("C:/tables/payTable_randomFlow.xlsx")
// 	Excel_PayTable_random, err := GetExcelData(path.PayTableRandomFlowPath)
// 	if err != nil {
// 		return err
// 	}
// 	PayTableMap_random, err := Excel_PayTable_random.GetPayTableMap()
// 	if err != nil {
// 		return err
// 	}
// 	// Excel_PayTable_sysWin, err := GetExcelData("C:/tables/payTable_sysWin.xlsx")
// 	Excel_PayTable_sysWin, err := GetExcelData(path.PayTableSysWinPath)
// 	if err != nil {
// 		return err
// 	}
// 	PayTableMap_sysWin, err := Excel_PayTable_sysWin.GetPayTableMap()
// 	if err != nil {
// 		return err
// 	}
// 	FishPayTable = PayTableMap_flow{
// 		config.SystemWinMonthlyRTP:          &PayTableMap_sysWin,
// 		config.SystemWinDailySysLoss:        &PayTableMap_sysWin,
// 		config.SystemWinDailyPlayerProfit:   &PayTableMap_sysWin,
// 		config.SystemWinMonthlyPlayerProfit: &PayTableMap_sysWin,
// 		config.RandomFlowProfitLimit:        &PayTableMap_random,
// 	}
// 	return
// }

// // Definition
// // 遊戲流程 --> 該流程的表
// type PayTableMap_flow map[config.RTPFlowTypeID]*PayTableMap

// // 魚編號 --> 該隻魚的PayTable
// type PayTableMap map[int32]*PayTable

// // 遊戲流程 --> 該流程的表
// var FishPayTable PayTableMap_flow

// // 單隻魚的PayTable
// type PayTable struct {
// 	FishID int32 `yaml:"fish_id"`
// 	FixPay int64 `yaml:"fix_pay"`

// 	TableWeight    map[int32]int32            `yaml:"table_weight"`    // 第幾項 --> 該項的權重
// 	IntervalWeight map[string]map[int32]int32 `yaml:"interval_weight"` // key(魚編號_表編號)  --> weight(map: 第幾項 --> 該項的權重)
// 	PayIntervals   map[string][2]int64        `yaml:"pay_intervals"`   // key(魚編號_表編號_區間編號)  --> 分數區間

// 	FGTimesObject map[int32]int             `yaml:"fgtimes_object"` // 第幾項 --> 該項的FG次數
// 	FGTimesWeight map[int32]int32           `yaml:"fgtimes_weight"` // 第幾項 --> 該項的權重
// 	Roulette      map[int32]map[int32]int64 `yaml:"roulette"`       // 第幾個盤眠 --> 第幾項 --> 該項的分數
// }

// func (E ExcelData) GetPayTableMap() (payTableMap PayTableMap, err error) {
// 	payTableMap = make(PayTableMap)
// 	for _, idx := range config.AllFishID {
// 		payTableMap[int32(idx)] = new(PayTable)
// 		payTableMap[int32(idx)].TableWeight = make(map[int32]int32)
// 		payTableMap[int32(idx)].IntervalWeight = make(map[string]map[int32]int32)
// 		payTableMap[int32(idx)].PayIntervals = map[string][2]int64{}
// 		payTableMap[int32(idx)].FGTimesObject = make(map[int32]int)
// 		payTableMap[int32(idx)].FGTimesWeight = make(map[int32]int32)
// 		payTableMap[int32(idx)].Roulette = make(map[int32]map[int32]int64)
// 	}

// 	for _, sheet := range E {
// 		fishID_int64, err := strconv.ParseInt(sheet["FishID"][0], 0, 32)
// 		if err != nil {
// 			LogTool.LogFatalf("GetPayTableMap", "型態轉換失敗:%v", err)
// 			break
// 		}

// 		fishID := int32(fishID_int64)
// 		payTableMap[fishID].FishID = fishID
// 		for k, v := range sheet {
// 			if k == "fix_pay" {
// 				val, err := strconv.ParseInt(v[0], 0, 32)
// 				if err != nil {
// 					LogTool.LogFatalf("GetPayTableMap", "型態轉換失敗:%v", err)
// 				}
// 				payTableMap[fishID].FixPay = val
// 			}

// 			if k == "tableWeight" {
// 				for idx := 0; idx < len(v); idx++ {
// 					val, err := strconv.ParseInt(v[idx], 0, 32)
// 					if err != nil {
// 						LogTool.LogFatalf("GetPayTableMap", "型態轉換失敗:%v", err)
// 					}
// 					payTableMap[fishID].TableWeight[int32(idx)] = int32(val)
// 				}
// 			}

// 			if strings.HasPrefix(k, "intervalWeight_") {
// 				key := k[len("intervalWeight_"):]
// 				for idx := 0; idx < len(v); idx++ {
// 					val, err := strconv.ParseInt(v[idx], 0, 32)
// 					if err != nil {
// 						LogTool.LogFatalf("GetPayTableMap", "型態轉換失敗:%v", err)
// 					}
// 					// map不存在時先初始化
// 					if len(payTableMap[fishID].IntervalWeight[key]) == 0 {
// 						intervalWeight := make(map[int32]int32)
// 						payTableMap[fishID].IntervalWeight[key] = intervalWeight
// 					}
// 					payTableMap[fishID].IntervalWeight[key][int32(idx)] = int32(val)
// 				}
// 			}

// 			if strings.HasPrefix(k, "interval_") {
// 				key := k[len("interval_"):]
// 				val1, err := strconv.ParseInt(v[0], 0, 32)
// 				if err != nil {
// 					LogTool.LogFatalf("GetPayTableMap", "型態轉換失敗:%v", err)
// 				}
// 				val2, err := strconv.ParseInt(v[1], 0, 32)
// 				if err != nil {
// 					LogTool.LogFatalf("GetPayTableMap", "型態轉換失敗:%v", err)
// 				}
// 				payTableMap[fishID].PayIntervals[key] = [2]int64{val1, val2}
// 			}

// 			if k == "FGTimesObject" {
// 				for idx := 0; idx < len(v); idx++ {
// 					val, err := strconv.ParseInt(v[idx], 0, 32)
// 					if err != nil {
// 						LogTool.LogFatalf("GetPayTableMap", "型態轉換失敗:%v", err)
// 					}
// 					payTableMap[fishID].FGTimesObject[int32(idx)] = int(val)
// 				}
// 			}

// 			if k == "FGTimesWeight" {
// 				for idx := 0; idx < len(v); idx++ {
// 					val, err := strconv.ParseInt(v[idx], 0, 32)
// 					if err != nil {
// 						LogTool.LogFatalf("GetPayTableMap", "型態轉換失敗:%v", err)
// 					}
// 					payTableMap[fishID].FGTimesWeight[int32(idx)] = int32(val)
// 				}
// 			}
// 			if strings.HasPrefix(k, "roulette_") {
// 				k, err := strconv.Atoi(k[len("roulette_"):])
// 				key := int32(k)
// 				if err != nil {
// 					LogTool.LogFatalf("GetPayTableMap", "型態轉換失敗:%v", err)
// 				}
// 				for idx := 0; idx < len(v); idx++ {
// 					val, err := strconv.ParseInt(v[idx], 0, 32)
// 					if err != nil {
// 						LogTool.LogFatalf("GetPayTableMap", "型態轉換失敗:%v", err)
// 					}
// 					// map不存在時先初始化
// 					if len(payTableMap[fishID].Roulette[key]) == 0 {
// 						rouletteTable := make(map[int32]int64)
// 						payTableMap[fishID].Roulette[key] = rouletteTable
// 					}
// 					payTableMap[fishID].Roulette[key][int32(idx)] = int64(val)
// 				}
// 			}
// 		}
// 	}
// 	return
// }

// // 修改paytable以符合倍數上限需求
// func (payTable PayTable) GetFitTable(multipleLimit int64) (PayTable, error) {
// 	FitPayIntervals := make(map[string][2]int64)
// 	FitTableWeight := make(map[int32]int32)
// 	FitIntervalWeight := make(map[string]map[int32]int32)
// 	if len(payTable.PayIntervals) == 0 {
// 		return payTable, fmt.Errorf("%s", "GetFitTable: payTable.PayIntervals為空")
// 	}

// 	for k, v := range payTable.PayIntervals {
// 		if v[0] <= multipleLimit {
// 			if v[1] <= multipleLimit {
// 				FitPayIntervals[k] = [2]int64{v[0], v[1]}
// 			} else {
// 				FitPayIntervals[k] = [2]int64{v[0], multipleLimit}
// 			}
// 			//Number從0開始 key從1開始
// 			intervalNumber := GetlastNumberFromKey(k) - 1
// 			tableKey := GetTableKeyFromIntervalKey(k)

// 			if _, ok := FitIntervalWeight[tableKey]; !ok {
// 				FitIntervalWeight[tableKey] = make(map[int32]int32)
// 			}

// 			FitIntervalWeight[tableKey][intervalNumber] = payTable.IntervalWeight[tableKey][intervalNumber]
// 			//Number從0開始 key從1開始
// 			tableNumber := GetlastNumberFromKey(tableKey) - 1
// 			FitTableWeight[tableNumber] = payTable.TableWeight[tableNumber]
// 		}
// 	}

// 	FitPayTables := PayTable{
// 		FishID:         payTable.FishID,
// 		FixPay:         payTable.FixPay,
// 		TableWeight:    FitTableWeight,
// 		IntervalWeight: FitIntervalWeight,
// 		PayIntervals:   FitPayIntervals,
// 	}
// 	return FitPayTables, nil
// }

// func GetTableKeyFromIntervalKey(key string) string {
// 	return GetKeyContractLastNumber(key)
// }

// func GetKeyContractLastNumber(key string) string {
// 	lastIndex := strings.LastIndex(key, "_")
// 	return key[:lastIndex]
// }

// func GetlastNumberFromKey(key string) int32 {
// 	lastIndex := strings.LastIndex(key, "_")
// 	if lastIndex == -1 {
// 		// 如果没有下划线，将整个字符串作为第一个部分，最後的部分0
// 		firstPart := key
// 		fmt.Printf("%s不是可分割的key,", firstPart)
// 		return 0
// 	} else {
// 		secondPartStr := key[lastIndex+1:]
// 		secondPart, err := strconv.ParseInt(secondPartStr, 10, 32)
// 		if err != nil {
// 			fmt.Println("无法解析第二部分为int32:", err)
// 			return 0
// 		}
// 		return int32(secondPart)
// 	}
// }

// func paytableMapToMap(payTable PayTable) map[string]interface{} {
// 	fmt.Println("fishRTP/config/config/configToMap")
// 	data := make(map[string]interface{})
// 	data["FishID"] = payTable.FishID
// 	data["FixPay"] = payTable.FixPay
// 	data["TableWeight"] = payTable.TableWeight
// 	data["IntervalWeight"] = payTable.IntervalWeight
// 	data["PayIntervals"] = payTable.PayIntervals
// 	data["FGTimesObject"] = payTable.FGTimesObject
// 	data["FGTimesWeight"] = payTable.FGTimesWeight
// 	return data
// }
