package games

// // Definition
// // 遊戲流程 --> 該流程的表
// type PayTableMap_flow map[config.RTPFlowTypeID]*PayTableMap

// // 魚編號 --> 該隻魚的PayTable
// type PayTableMap map[int32]*PayTable

// // 遊戲流程 --> 該流程的表
// var FishPayTable PayTableMap_flow

func GetProbability(targetDir string) (map[string]Reels, error) {
	//rtp ->[]reels
	probTable := make(map[string]Reels)
	rtpType, err := GetRTPType(targetDir)
	if err != nil {
		return nil, err
	}
	for id, file := range rtpType {
		if id == "ID" {
			continue
		}
		rtp := file[0]
		reelConfig, err := GetReelConfig(targetDir, rtp)
		if err != nil {
			return nil, err
		}

		for k, v := range reelConfig {
			probTable[rtp+"_"+k] = GetSymbolReelFromStringReel(v)
		}

	}

	// LogTool.LogInfof("SlotReel", "%#v", probTable)
	return probTable, nil
}

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
