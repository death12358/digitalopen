package util

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// fileName: "path/XXX.xlsx"
// map: Sheet名稱 --> 檔案內容(map)
func GetExcelData(fileName string) (excelData ExcelData, err error) {
	excelData = map[string]DataMap{}
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Excel讀檔失敗:%v", err)
	}

	for _, sheet := range f.GetSheetList() {
		dataMap := make(map[string][]string)
		if _, ok := excelData[sheet]; !ok {
			excelData[sheet] = make(map[string][]string)
		}

		rows, err := f.GetRows(sheet)
		if err != nil {
			return nil, fmt.Errorf("Excel讀取資料失敗:%v", err)
		}

		for _, row := range rows {
			// 将第一column的值作为Key
			key := row[0]
			dataMap[key] = make([]string, 0)
			// 将Key对应的后续列的数据存储到Map中
			for _, value := range row[1:] {
				dataMap[key] = append(dataMap[key], value)
			}
		}
		excelData[sheet] = dataMap
	}
	return
}

func main() {
	// 打開CSV檔案
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 創建一個CSV Reader
	reader := csv.NewReader(file)

	// 讀取所有的記錄
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// 輸出CSV中的每條記錄
	for _, record := range records {
		fmt.Println(record)
	}
}
