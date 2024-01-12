package tables

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"os"
// )

// func TableInit() (err error) {
// 	// DeadProb會用payTable中的值去算, 所以Get的順序不能變
// 	// if err = GetFishPayTable(); err != nil {
// 	// 	return err
// 	// }
// 	// if err = GetFishDeadProb(); err != nil {
// 	// 	return err
// 	// }
// 	// for k := range *FishPayTable[5] {
// 	// 	LogTool.LogInfof("FishPayTable:", "%#v", (*FishPayTable[1])[k])
// 	// }
// 	// LogTool.LogInfof("FishDeadProb1:", "%#v", FishDeadProb[1])

// 	// LogTool.LogInfof("FishDeadProb5:", "%#v", FishDeadProb[5])
// 	FishPayTable = dummy_FishPayTable
// 	FishDeadProb = dummy_FishDeadProb
// 	return nil
// }

// func TableInit2() (err error) {
// 	// // DeadProb會用payTable中的值去算, 所以Get的順序不能變
// 	if err = GetFishPayTable(); err != nil {
// 		return err
// 	}
// 	if err = GetFishDeadProb(); err != nil {
// 		return err
// 	}

// 	// for k := range *FishPayTable[5] {
// 	// 	LogTool.LogInfof("FishPayTable:", "%v:%#v\n", k, (*FishPayTable[5])[k])
// 	// }
// 	// LogTool.LogInfof("FishDeadProb1:", "%#v\n", FishDeadProb[1])
// 	// LogTool.LogInfof("FishDeadProb2:", "%#v\n", FishDeadProb[2])

// 	// LogTool.LogInfof("FishDeadProb3:", "%#v\n", FishDeadProb[3])
// 	// LogTool.LogInfof("FishDeadProb4:", "%#v\n", FishDeadProb[4])

// 	return nil
// }

// // Sheet --> 資料內容(map)
// type ExcelData map[string]DataMap

// // 數據名稱(first col) --> 數據內容([]string)
// type DataMap map[string][]string

// // fileName: "path/XXX.xlsx"
// // map: Sheet名稱 --> 檔案內容(map)
// func GetExcelData(fileName string) (excelData ExcelData, err error) {
// 	excelData = map[string]DataMap{}
// 	// f, err := excelize.OpenFile(fileName)
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("Excel讀檔失敗:%v", err)
// 	// }

// 	file, err := os.Open("example.csv")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// 創建一個CSV Reader
// 	reader := csv.NewReader(file)

// 	// 讀取所有的記錄
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		fmt.Println("Error reading CSV:", err)
// 		return
// 	}

// 	// 輸出CSV中的每條記錄
// 	for _, record := range records {
// 		fmt.Println(record)
// 	}

// 	// for _, sheet := range f.GetSheetList() {
// 	// 	dataMap := make(map[string][]string)
// 	// 	if _, ok := excelData[sheet]; !ok {
// 	// 		excelData[sheet] = make(map[string][]string)
// 	// 	}

// 	// 	rows, err := f.GetRows(sheet)
// 	// 	if err != nil {
// 	// 		return nil, fmt.Errorf("Excel讀取資料失敗:%v", err)
// 	// 	}

// 	// 	for _, row := range rows {
// 	// 		// 将第一column的值作为Key
// 	// 		key := row[0]
// 	// 		dataMap[key] = make([]string, 0)
// 	// 		// 将Key对应的后续列的数据存储到Map中
// 	// 		for _, value := range row[1:] {
// 	// 			dataMap[key] = append(dataMap[key], value)
// 	// 		}
// 	// 	}
// 	// 	excelData[sheet] = dataMap
// 	// }
// 	return
// }
