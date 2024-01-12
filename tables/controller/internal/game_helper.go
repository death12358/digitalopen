package internal

import (
	"errors"
)

func CopyReels(inReels [][]string) [][]string {

	var result [][]string

	for r := 0; r < len(inReels); r++ {
		reel := make([]string, len(inReels[r]))
		copy(reel, inReels[r])
		result = append(result, reel)
	}
	return result
}

type Angle int

const (
	angle90 Angle = iota
	angle180
	angle270
)

// 旋转切片
func RotateSlice[T any](slice [][]T, angle Angle) ([][]T, error) {

	for i := 1; i < len(slice); i++ {
		if len(slice[i-1]) != len(slice[i]) {
			return slice, errors.New("slice asymmetrical")
		}
	}

	rows := len(slice)
	cols := len(slice[0])
	result := make([][]T, cols)

	for i := 0; i < cols; i++ {
		result[i] = make([]T, rows)
		for j := 0; j < rows; j++ {
			switch angle {
			case angle90:
				result[i][j] = slice[rows-j-1][i]
			case angle180:
				result[i][j] = slice[rows-i-1][cols-j-1]
			case angle270:
				result[i][j] = slice[j][cols-i-1]
			default:
				result[i][j] = slice[i][j]
				return slice, errors.New("Angle is invalid parameter")
			}
		}
	}
	return result, nil

}
