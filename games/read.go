package games

import (
	"strings"

	LogTool "github.com/adimax2953/log-tool"
	"github.com/death12358/digitalopen/tables/tools"
)

const (
	dataKeyIndex = 0
	gameDir      = "gametemplate"
	dataDir      = "datas"
)

type FilePathType int

const (
	Config FilePathType = iota
	Mapping
	Table
	PayLines
	Probability
)

var (
	mappingPath     string
	payTablePath    string
	payLinesPath    string
	probabilityPath string
)

type tableData map[string]interface{}

type DataMap map[string][]string

// func initReels(targetDir string) (dm DataMap, err error)    {}
// func InitPayTable(targetDir string) (dm DataMap, err error) {}
// func initLines(targetDir string) (dm DataMap, err error)    {}
func GetRTPType(targetDir string) (dm DataMap, err error) {
	mappingPath = strings.ReplaceAll(targetDir+"\\config\\mapping.csv", "\\", "/")
	dm, err = CreateDataMap(mappingPath)
	if err != nil {
		LogTool.LogErrorf("CreateDataMap", "%v", err)
	}
	// LogTool.LogInfof("CreateDataMap", "%v", dm)
	return
}
func GetPayTable(targetDir string) (dm DataMap, err error) {
	payTablePath = strings.ReplaceAll(targetDir+"\\pay\\table.csv", "\\", "/")
	dm, err = CreateDataMap(payTablePath)
	if err != nil {
		LogTool.LogErrorf("CreateDataMap", "%v", err)
	}
	// LogTool.LogInfof("CreateDataMap", "%v", dm)
	return
}
func GetPayLine(targetDir string) (dm DataMap, err error) {
	payLinesPath = strings.ReplaceAll(targetDir+"\\pay\\lines.csv", "\\", "/")
	dm, err = CreateDataMap(payLinesPath)
	if err != nil {
		LogTool.LogErrorf("CreateDataMap", "%v", err)
	}
	// LogTool.LogInfof("CreateDataMap", "%v", dm)
	return
}
func GetReelConfig(targetDir, rtpType string) (dm DataMap, err error) {
	probabilityPath = strings.ReplaceAll(targetDir+"\\probability\\"+rtpType, "\\", "/")
	dm, err = CreateDataMap(probabilityPath)
	if err != nil {
		LogTool.LogErrorf("CreateDataMap", "%v", err)
	}
	// LogTool.LogInfof("CreateDataMap", "%v", dm)
	return
}

// CreateDataMap 建一個讀檔格式(橫式的csv檔,機率表...)
func CreateDataMap(path string) (DataMap, error) {
	records, err := tools.OpenCSV(path)

	if err != nil {
		return nil, err
	}
	dataMap := make(DataMap)
	//逐行處理記錄
	for _, record := range records {
		// 逐欄位讀取記錄中的值,整理資料
		var k string
		if len(record) > dataKeyIndex {
			k = record[dataKeyIndex]
		}
		for i := dataKeyIndex + 1; i < len(record); i++ {
			if record[i] == "" {
				dataMap[k] = record[dataKeyIndex+1 : i]
				break
			} else if i == len(record)-1 {
				dataMap[k] = record[dataKeyIndex+1:]
				break
			}
		}
	}
	return dataMap, nil
}

// GetDataByKey 取該row
func (rm DataMap) GetDataByKey(inKey string) ([]string, bool) {
	data, ok := rm[inKey]
	return data, ok
}

// GetDataByIndex 取該row index裡的資料
func (rm DataMap) GetDataByIndex(inKey string, index int) (string, bool) {
	data, ok := rm[inKey]

	if ok && len(data)-1 >= index {
		return data[index], ok
	}
	return "", ok
}

// // ParseRowDataToWeight 轉為權重結構
// func (rm DataMap) ParseRowDataToWeight(inKey string) (random.ProbWeight, bool) {
// 	var rslt random.ProbWeight
// 	data, ok := rm.GetDataByKey(inKey)
// 	if ok {
// 		var probSlice []int64
// 		for i := 0; i < len(data); i++ {
// 			val, err := strconv.Atoi(data[i])
// 			if err != nil {
// 				return rslt, false
// 			}
// 			probSlice = append(probSlice, int64(val))
// 		}
// 		rslt = random.CreateProb(probSlice...)
// 	}
// 	return rslt, ok
// }

// // ParseRowDataToInt 轉為int型態
// func (rm DataMap) ParseRowDataToInt(inKey string) ([]int, bool) {
// 	data, ok := rm.GetDataByKey(inKey)
// 	rslt := make([]int, len(data))
// 	if ok {
// 		for i := 0; i < len(data); i++ {
// 			v, _ := strconv.Atoi(data[i])
// 			rslt[i] = v
// 		}
// 	}
// 	return rslt, ok
// }

// // ReadPayData 讀取
// func (f Folder) ReadPayData(calcStand int) (*PayTable, error) {
// 	payTable := NewPayTable(calcStand)
// 	payTableData, err := CreateDataMap(f.getPath(Table))
// 	if err != nil {
// 		return payTable, err
// 	}
// 	payTable.InitPayInfo(payTableData)

// 	if calcStand == int(Line) {
// 		payLinesData, err := CreateDataMap(f.getPath(PayLines))
// 		if err != nil {
// 			return payTable, err
// 		}
// 		err = payTable.setPayLine(payLinesData)
// 		if err != nil {
// 			return payTable, err
// 		}
// 	}
// 	return payTable, err
// }

// func (f Folder) ReadPayLinesData() (DataMap, error) {
// 	return CreateDataMap(f.getPath(PayLines))
// }

// type ProbabilitySetting struct {
// 	ID   int
// 	File string
// }

// // ReadMappingData 讀機率表設定檔
// func (f Folder) ReadMappingData() (map[int]string, error) {

// 	path := f.getPath(Mapping)

// 	records, err := tools.OpenCSV(path)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var (
// 		field      []string
// 		settingMap map[int]string = make(map[int]string)
// 	)

// 	field = records[0]
// 	for i := 1; i < len(records); i++ {
// 		setting := ProbabilitySetting{}
// 		if err := tools.SerializeStructData(field, records[i], &setting); err != nil {
// 			log.Println(err)
// 			continue
// 		} else {
// 			settingMap[setting.ID] = setting.File
// 		}
// 	}

// 	return settingMap, nil
// }

// // ReadProbabilityData 讀機率表
// func (f Folder) ReadProbabilityData(file string) (DataMap, error) {
// 	path := f.getPath(Probability) + file
// 	return CreateDataMap(path)
// }
