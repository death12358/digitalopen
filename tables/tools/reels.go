package tools

import (
	"fmt"

	"github.com/death12358/digitalopen/tables/tools/random"
)

// BuildReels 建立輪帶資料
func BuildReels(inReelsData [][]string, length, height int) ([][]string, error) {
	rslt := make([][]string, length)
	if len(inReelsData) != length {
		return rslt, fmt.Errorf("in reels data len:[%d] not equal length:[%d]", len(inReelsData), length)
	}
	for x := 0; x < length; x++ {
		if oneReel, err := RandReelsData(inReelsData[x], height); err != nil {
			return rslt, err
		} else {
			rslt[x] = oneReel
		}
	}
	return rslt, nil
}

// RandReelsData 亂數產生輪帶資料 count=要連續幾顆
func RandReelsData(inReel []string, count int) ([]string, error) {
	beginIdx, ok := random.RandSliceIndex(inReel)
	if !ok || len(inReel) < count {
		return []string{}, fmt.Errorf("rand slice index issue. data:[%v]", inReel)
	}

	result := make([]string, count)
	for i := 0; i < count; i++ {
		idx := (beginIdx + i) % len(inReel)
		result[i] = inReel[idx]
	}

	return result, nil
}
