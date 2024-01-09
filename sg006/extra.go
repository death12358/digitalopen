package sg006

import (
	"encoding/json"

	"github.com/death12358/digitalopen/games"
)

type ExtraReels struct {
	Reels []string `json:"reels"`
}

// Marshal
func (r ExtraReels) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Unmarshal
func (r *ExtraReels) Unmarshal(roundsBytes []byte) error {
	return json.Unmarshal(roundsBytes, r)
}

type ExtraSG006 struct {
	BattleRing map[string]int `json:"battlering"`
	VSReel     [5]int         `json:"vsreel"` //標注哪幾輪上有ＶＳ
}

// Marshal
func (r ExtraSG006) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Unmarshal
func (r *ExtraSG006) Unmarshal(roundsBytes []byte) error {
	return json.Unmarshal(roundsBytes, r)
}

// AppendExtraReels - 新增額外資訊
func AppendExtraReels(r *games.Records, extraReels ExtraReels) error {
	s, err := extraReels.Marshal()
	if err != nil {
		return err
	}
	r.Symbols = append(r.Symbols, string(s))
	return nil
}

// AppendExtraSG006 - 新增 SG006 額外資訊
func AppendExtraSG006(r *games.Records, extrasg006 ExtraSG006) error {
	s, err := extrasg006.Marshal()
	if err != nil {
		return err
	}
	r.Symbols = append(r.Symbols, string(s))
	return nil
}
