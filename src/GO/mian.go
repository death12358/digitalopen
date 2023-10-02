package main

import "fmt"

func main() {
	fmt.Println(isMatch([]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 7}))
}

func isMatch(NineBox []int) bool {
	if (NineBox[0] == 7 && NineBox[1] == 7 && NineBox[2] == 7) ||
		(NineBox[3] == 7 && NineBox[4] == 7 && NineBox[5] == 7) ||
		(NineBox[6] == 7 && NineBox[7] == 7 && NineBox[8] == 7) ||
		(NineBox[0] == 7 && NineBox[3] == 7 && NineBox[6] == 7) ||
		(NineBox[1] == 7 && NineBox[4] == 7 && NineBox[7] == 7) ||
		(NineBox[2] == 7 && NineBox[5] == 7 && NineBox[8] == 7) ||
		(NineBox[0] == 7 && NineBox[4] == 7 && NineBox[8] == 7) ||
		(NineBox[2] == 7 && NineBox[4] == 7 && NineBox[6] == 7) {
		return true
	}
	return false
}
