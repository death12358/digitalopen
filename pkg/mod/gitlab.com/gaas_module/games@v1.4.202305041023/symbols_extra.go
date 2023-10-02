package games

import (
	"encoding/json"

	"github.com/shopspring/decimal"
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

type ExtraSG008 struct {
	RemainRetrigger int             `json:"remain_retrigger"`
	WWmultiplier    decimal.Decimal `json:"wwmultiplier"`
}

// Marshal
func (r ExtraSG008) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Unmarshal
func (r *ExtraSG008) Unmarshal(roundsBytes []byte) error {
	return json.Unmarshal(roundsBytes, r)
}
