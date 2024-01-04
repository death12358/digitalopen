package games

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

// Results -
type Results map[string]*Records

// NewResults - 建立回合結果
func NewResults() Results {
	return Results{}
}

// AddRecord - 新增回合結果
func (r Results) AddRecord(rec Records) error {
	// 檢查 ID 是否存在
	if _, ok := r[rec.Id]; ok {
		// ID 已存在，返回錯誤
		return fmt.Errorf("record with ID '%s' already exists", rec.Id)
	}

	// ID 不存在，新增記錄
	r[rec.Id] = &rec
	return nil
}

// Marshal
func (r Results) Marshal() ([]byte, error) {
	return json.Marshal(map[string]*Records(r))
}

// Unmarshal
func (r *Results) Unmarshal(resultsBytes []byte) error {
	return json.Unmarshal(resultsBytes, r)
}

// Split - 清除未中獎

// Records - 遊戲紀錄
//
//	Id			- 該筆紀錄識別碼
//	Brand		- 品牌
//	Username	- 使用者
//	Case		- 該回合狀態
//	Stages 		- 該回合狀態階段或局數，例如：FreeGame的1(第一局)、2(第二局)
//	Pickem		- 選擇項目
//	Symbol		- 該回合獎項
//	Extra		- 額外資訊
//	Multiplier	- 贏分倍數
//	Bet			- 下注
//	Point		- 贏分
type Records struct {
	Id         string          `json:"id"`
	Brand      string          `json:"brand"`
	Username   string          `json:"username"`
	Case       State           `json:"case"`
	Stages     int64           `json:"stages"`
	Pickem     []string        `json:"pickem"`
	Symbols    []string        `json:"symbols"`
	Extra      []string        `json:"extra"`
	Multiplier decimal.Decimal `json:"multiplier"`
	Bet        decimal.Decimal `json:"bet"`
	Point      decimal.Decimal `json:"point"`
}

// Marshal
func (r Records) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Unmarshal
func (r *Records) Unmarshal(recordsBytes []byte) error {
	return json.Unmarshal(recordsBytes, r)
}

// AppendExtraReels - 新增額外資訊
func (r *Records) AppendExtraReels(extraReels ExtraReels) error {
	s, err := extraReels.Marshal()
	if err != nil {
		return err
	}
	r.Symbols = append(r.Symbols, string(s))
	return nil
}

// AppendExtraSG008 - 新增 SG008 額外資訊
func (r *Records) AppendExtraSG008(extrasg008 ExtraSG008) error {
	s, err := extrasg008.Marshal()
	if err != nil {
		return err
	}
	r.Symbols = append(r.Symbols, string(s))
	return nil
}
