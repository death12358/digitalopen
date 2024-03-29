package args

import (
	"encoding/json"
	"fmt"
)

type SpinArgs struct {
	Id       string      `json:"id"`
	Brand    string      `json:"brand"`
	Username string      `json:"username"`
	Pickem   []string    `json:"pickem"`
	Currency string      `json:"currency"`
	Bet      json.Number `json:"bet"`
}

// Validate validates the SpinArgs
func (v *SpinArgs) Validate() error {
	// return assert.Assert(
	// 	assert.NonEmptyString(v.Id, "id"),
	// 	assert.NonEmptyString(v.Brand, "brand"),
	// 	assert.NonEmptyString(v.Username, "username"),
	// 	assert.JsonNumber(v.Bet, "bet",
	// 		assert.NonNanNorInf,
	// 		assert.NonNegativeNumber,
	// 	),
	// 	NonEmptyArray(v.Pickem, "pickem"),
	// 	assert.NonEmptyString(v.Currency, "currency"),
	// )
	return nil
}

const (
	ERR_EMPTY_ARRAY = "%T cannot be an empty array elements"
	ERR_NIL_ARRAY   = "cannot be an nil array"
	ERR_NON_ARRAY   = "is %T, %s should be an array"
)

func NonEmptyArray(v interface{}, name string) error {
	switch x := v.(type) {
	case nil:
		return fmt.Errorf(name, "cannot be an nil array")
	case []string:
		if len(x) == 0 {
			return fmt.Errorf(name, "cannot be an nil array")
		}
	case []int:
		if len(x) == 0 {
			return fmt.Errorf(name, "cannot be an nil array")
		}
	case []int32:
		if len(x) == 0 {
			return fmt.Errorf(name, "cannot be an nil array")
		}
	case []int64:
		if len(x) == 0 {
			return fmt.Errorf(name, "cannot be an nil array")
		}
	case []float32:
		if len(x) == 0 {
			return fmt.Errorf(name, "cannot be an nil array")
		}
	case []float64:
		if len(x) == 0 {
			return fmt.Errorf(name, "cannot be an nil array")
		}
	default:
		return fmt.Errorf(name, "cannot be an nil array")
	}

	return nil
}
