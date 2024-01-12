package sg001

import (
	"github.com/death12358/digitalopen/games"
)

var (
	ngReelStrips97_8 = games.ReelStrips{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		// {7, 2, 4, 13, 13, 6, 8, 2, 4, 9, 7, 11, 1, 11, 1, 1, 4, 2, 3, 7, 4, 10, 1, 2, 3, 9, 8, 2, 1, 3, 5, 4, 6, 4, 13, 13, 13, 4, 2, 6, 4, 2, 1, 3, 4, 2, 5, 8, 1, 1, 1, 2, 1, 4, 1, 1, 3, 2, 5, 1, 11, 10, 5},
		// {7, 5, 3, 7, 6, 0, 3, 1, 4, 7, 8, 1, 3, 6, 13, 13, 13, 7, 5, 2, 3, 1, 6, 3, 6, 13, 13, 13, 6, 1, 0, 8, 2, 3, 7, 5, 4, 1, 3, 6, 11, 9, 1, 3, 7, 2, 5, 3, 1, 10, 4, 1, 3, 6, 1, 0, 7, 2, 9, 5, 11, 10, 2},
		// {1, 2, 3, 6, 4, 8, 5, 4, 6, 1, 2, 5, 7, 3, 8, 4, 5, 3, 10, 5, 4, 11, 13, 13, 13, 4, 5, 3, 5, 1, 7, 4, 2, 13, 13, 13, 11, 5, 2, 0, 3, 11, 1, 5, 7, 1, 4, 8, 1, 9, 5, 2, 9, 0, 10, 4, 2, 10, 9, 2, 4, 5, 4},
		// {5, 3, 4, 7, 13, 13, 13, 11, 9, 7, 0, 9, 10, 1, 5, 6, 3, 6, 13, 13, 10, 6, 3, 4, 11, 7, 2, 9, 10, 5, 8, 0, 6, 4, 8, 3, 7, 5, 4, 1, 8, 4, 6, 3, 11, 4, 7, 1, 7, 4, 10, 4, 6, 0, 11, 5, 7, 4, 6, 2, 5, 8, 9},
		// {6, 2, 11, 7, 13, 13, 13, 8, 6, 1, 4, 1, 5, 1, 9, 6, 1, 3, 8, 5, 9, 3, 4, 5, 9, 7, 13, 13, 10, 9, 5, 4, 11, 1, 10, 2, 11, 7, 1, 9, 2, 1, 5, 8, 3, 11, 5, 1, 7, 1, 10, 5, 2, 8, 5, 2, 7, 3, 1, 10, 4, 8, 2, 11, 2, 5, 6, 2, 4, 5, 2, 10, 2},
	}
	hifgReelStrips97_8 = games.ReelStrips{
		{1, 3, 1, 4, 1, 5, 1, 4, 1, 1, 1, 2, 2, 5, 3, 4, 3, 5, 1, 1, 1, 1, 1, 1, 1, 2, 4, 2, 4, 3, 4, 5, 4, 3, 2, 1, 5, 3, 4, 1, 5, 2, 5, 1, 3, 4, 2, 1, 1, 1, 5, 3, 2, 3, 5, 2, 4, 4, 4, 4, 4, 2, 4, 4, 4, 3, 4, 4, 4, 5},
		{1, 1, 1, 3, 2, 4, 0, 4, 2, 1, 4, 4, 1, 5, 1, 4, 1, 4, 4, 3, 5, 2, 4, 3, 4, 4, 1, 1, 1, 1, 1, 4, 4, 4, 0, 2, 4, 4, 4, 0, 5, 4, 3, 1, 4, 4, 4, 1, 5, 1, 4, 5, 4, 5, 0, 5, 5, 4, 3, 2, 4, 4, 5, 5},
		{0, 1, 5, 3, 4, 2, 3, 2, 5, 4, 3, 2, 5, 1, 4, 1, 3, 0, 1, 5, 3, 4, 2, 3, 2, 5, 4, 3, 2, 5, 1, 4, 1, 3, 1, 1, 1, 1, 1, 1, 4, 5, 2, 4, 2, 2, 5, 1, 1, 5, 2, 4, 1, 3, 3, 4, 1, 3, 4, 2, 2, 0, 3, 5, 4, 5, 3, 4, 5, 2, 3, 4, 2, 5, 1, 1, 5, 2, 4, 1, 3, 3, 5, 1, 4, 5, 2, 2, 0, 3, 5, 4, 5, 3},
		{3, 2, 3, 1, 4, 2, 2, 2, 3, 1, 3, 4, 2, 0, 3, 1, 4, 3, 3, 3, 2, 5, 4, 1, 5, 3, 3, 1, 4, 5, 2, 5, 0, 1, 5, 3, 4, 2, 3, 5, 1, 2, 5, 3, 2, 0, 2, 5, 1, 2, 5, 1, 4, 5, 5, 2, 2, 2, 4, 5, 1, 4, 3, 5, 0, 3, 3, 5, 3, 5, 4, 3, 3, 3},
		{1, 1, 1, 1, 1, 1, 1, 2, 3, 4, 2, 3, 1, 4, 2, 2, 3, 4, 1, 5, 4, 2, 1, 3, 4, 5, 2, 1, 2, 4, 5, 1, 5, 1, 3, 3, 1, 1, 3, 4, 1, 5, 2, 5, 1, 5, 2, 4, 5, 1, 3, 5, 4, 1, 3, 4, 3, 1, 5, 2},
	}

	// 中間 2 3 4 大圖，同一座標去算
	bigfgReelStrips97_8 = games.ReelStrips{
		{8, 2, 6, 1, 11, 2, 3, 1, 7, 2, 1, 4, 10, 2, 1, 6, 2, 3, 4, 1, 2, 3, 11, 1, 2, 3, 10, 1, 3, 4, 2, 9, 3, 5, 9, 1, 3, 1, 2, 10, 2, 3, 11, 5, 2, 1, 8, 5, 1, 3, 7, 3, 2, 1},
		{3, 4, 3, 4, 8, 5, 6, 3, 5, 9, 4, 5, 2, 5, 6, 4, 10, 5, 8, 7, 4, 8, 5, 10, 4, 9, 6, 5, 4, 6, 5, 0, 6, 7, 4, 11, 1, 6, 11, 2, 4, 6, 4, 3, 5, 4, 1, 2},
		{3, 4, 3, 4, 8, 5, 6, 3, 5, 9, 4, 5, 2, 5, 6, 4, 10, 5, 8, 7, 4, 8, 5, 10, 4, 9, 6, 5, 4, 6, 5, 0, 6, 7, 4, 11, 1, 6, 11, 2, 4, 6, 4, 3, 5, 4, 1, 2},
		{3, 4, 3, 4, 8, 5, 6, 3, 5, 9, 4, 5, 2, 5, 6, 4, 10, 5, 8, 7, 4, 8, 5, 10, 4, 9, 6, 5, 4, 6, 5, 0, 6, 7, 4, 11, 1, 6, 11, 2, 4, 6, 4, 3, 5, 4, 1, 2},
		{9, 5, 3, 9, 11, 5, 9, 7, 4, 7, 9, 3, 11, 7, 2, 8, 5, 11, 4, 8, 6, 4, 8, 2, 4, 1, 11, 10, 5, 5, 1, 6, 1, 9, 5, 10, 8, 3, 11, 6, 2, 9, 6, 1, 10, 8, 5, 10, 7, 11, 6, 1, 7, 10, 2, 8, 7, 4, 10, 2, 6, 3, 4, 2, 6, 5, 7, 3},
	}
)

var (
	ngReelStrips97_18 = games.ReelStrips{
		{7, 3, 4, 13, 13, 6, 8, 1, 4, 9, 7, 1, 1, 11, 1, 1, 2, 2, 2, 7, 4, 10, 4, 4, 3, 9, 8, 3, 3, 3, 5, 4, 6, 1, 13, 13, 13, 4, 2, 6, 4, 2, 1, 3, 4, 2, 5, 8, 1, 1, 1, 2, 1, 4, 1, 1, 3, 2, 5, 1, 11, 10, 5},
		{4, 5, 3, 7, 6, 0, 3, 1, 7, 3, 8, 1, 4, 6, 13, 13, 13, 7, 5, 2, 4, 3, 6, 1, 6, 13, 13, 13, 6, 3, 0, 8, 2, 4, 1, 5, 11, 1, 4, 6, 4, 9, 1, 4, 7, 1, 5, 4, 1, 10, 3, 2, 3, 6, 1, 0, 7, 2, 9, 5, 11, 10, 2},
		{4, 4, 1, 6, 3, 8, 5, 2, 6, 1, 2, 5, 7, 3, 8, 1, 5, 3, 10, 5, 3, 11, 13, 13, 13, 4, 4, 4, 5, 1, 0, 5, 2, 13, 13, 13, 11, 5, 2, 0, 1, 11, 4, 4, 7, 1, 4, 8, 1, 9, 5, 2, 9, 0, 10, 5, 2, 10, 9, 2, 5, 4, 4},
		{5, 3, 4, 7, 13, 13, 13, 11, 9, 7, 0, 9, 10, 1, 4, 5, 3, 6, 13, 13, 10, 6, 3, 4, 11, 7, 2, 9, 10, 5, 8, 0, 6, 4, 8, 3, 7, 5, 4, 1, 8, 4, 6, 3, 11, 4, 7, 1, 7, 4, 10, 4, 6, 0, 11, 5, 7, 4, 6, 2, 5, 8, 9},
		{6, 2, 11, 7, 13, 13, 13, 8, 6, 1, 4, 1, 5, 1, 9, 6, 1, 3, 8, 5, 9, 3, 4, 5, 9, 7, 13, 13, 10, 9, 5, 4, 11, 1, 10, 2, 11, 7, 1, 9, 2, 1, 5, 8, 3, 11, 5, 1, 7, 1, 10, 5, 2, 8, 5, 2, 7, 3, 1, 10, 4, 8, 2, 11, 2, 5, 6, 2, 4, 5, 2, 10, 2},
	}

	hifgReelStrips97_18 = games.ReelStrips{
		{1, 3, 1, 4, 1, 5, 1, 3, 1, 1, 1, 3, 2, 5, 3, 4, 3, 5, 1, 3, 1, 1, 3, 1, 1, 2, 4, 2, 4, 3, 4, 5, 4, 3, 2, 1, 5, 3, 4, 1, 5, 2, 5, 1, 3, 4, 2, 3, 1, 1, 5, 3, 2, 3, 5, 2, 4, 4, 4, 4, 4, 2, 4, 4, 4, 3, 4, 4, 4, 5},
		{1, 1, 1, 3, 2, 3, 0, 3, 2, 1, 4, 3, 1, 5, 0, 4, 1, 3, 4, 3, 5, 2, 4, 3, 4, 3, 1, 1, 3, 1, 1, 4, 4, 4, 0, 2, 4, 4, 3, 0, 5, 4, 3, 1, 4, 3, 1, 4, 5, 1, 3, 5, 4, 5, 0, 5, 5, 4, 3, 2, 4, 3, 5, 5},
		{4, 1, 5, 3, 4, 2, 3, 2, 5, 4, 3, 2, 5, 1, 4, 1, 3, 0, 1, 5, 3, 4, 2, 3, 0, 5, 3, 3, 2, 5, 1, 4, 1, 3, 1, 1, 1, 1, 1, 1, 4, 5, 2, 4, 2, 2, 5, 1, 1, 5, 2, 4, 1, 3, 3, 4, 1, 3, 4, 2, 2, 0, 3, 5, 4, 5, 3, 4, 5, 2, 3, 4, 2, 5, 1, 4, 5, 2, 4, 1, 3, 3, 5, 4, 4, 5, 2, 2, 0, 3, 5, 4, 5, 3},
		{3, 4, 3, 1, 4, 2, 2, 2, 3, 1, 3, 4, 2, 4, 3, 1, 4, 3, 3, 3, 2, 5, 4, 1, 5, 3, 3, 1, 4, 5, 2, 5, 0, 1, 5, 3, 4, 2, 3, 0, 1, 2, 5, 3, 2, 0, 2, 5, 1, 2, 5, 1, 4, 5, 5, 4, 2, 4, 4, 4, 4, 4, 3, 5, 0, 3, 3, 5, 3, 5, 4, 3, 3, 3},
		{1, 1, 1, 1, 1, 1, 4, 4, 4, 4, 4, 3, 4, 4, 2, 2, 3, 4, 1, 5, 4, 2, 1, 3, 4, 5, 2, 1, 2, 4, 5, 1, 5, 1, 3, 3, 1, 1, 3, 4, 1, 5, 2, 5, 1, 5, 2, 4, 5, 1, 3, 5, 4, 1, 3, 4, 3, 1, 5, 2},
	}

	// 中間 2 3 4 大圖，同一座標去算
	bigfgReelStrips97_18 = games.ReelStrips{
		{8, 2, 6, 1, 11, 2, 3, 1, 7, 2, 1, 4, 10, 2, 1, 6, 2, 3, 4, 1, 2, 3, 11, 1, 2, 3, 10, 1, 3, 4, 2, 9, 3, 5, 9, 1, 3, 1, 2, 10, 2, 3, 11, 5, 2, 1, 8, 5, 1, 3, 7, 3, 2, 1},
		{3, 4, 3, 4, 8, 5, 6, 3, 5, 9, 4, 5, 2, 5, 6, 4, 10, 5, 8, 7, 4, 8, 5, 10, 4, 9, 6, 5, 4, 6, 5, 0, 6, 7, 4, 11, 1, 6, 11, 2, 4, 6, 4, 3, 5, 4, 1, 2},
		{3, 4, 3, 4, 8, 5, 6, 3, 5, 9, 4, 5, 2, 5, 6, 4, 10, 5, 8, 7, 4, 8, 5, 10, 4, 9, 6, 5, 4, 6, 5, 0, 6, 7, 4, 11, 1, 6, 11, 2, 4, 6, 4, 3, 5, 4, 1, 2},
		{3, 4, 3, 4, 8, 5, 6, 3, 5, 9, 4, 5, 2, 5, 6, 4, 10, 5, 8, 7, 4, 8, 5, 10, 4, 9, 6, 5, 4, 6, 5, 0, 6, 7, 4, 11, 1, 6, 11, 2, 4, 6, 4, 3, 5, 4, 1, 2},
		{9, 5, 3, 9, 11, 5, 9, 7, 4, 7, 9, 3, 11, 7, 2, 8, 5, 11, 4, 8, 6, 4, 8, 2, 4, 1, 11, 10, 5, 5, 1, 6, 1, 9, 5, 10, 8, 3, 11, 6, 2, 9, 6, 1, 10, 8, 5, 10, 7, 11, 6, 1, 7, 10, 2, 8, 7, 4, 10, 2, 6, 3, 4, 2, 6, 5, 7, 3},
	}
)
var (
	ngReelStrips97_38 = games.ReelStrips{
		{2, 1, 6, 1, 4, 4, 4, 3, 9, 2, 4, 11, 1, 3, 2, 2, 9, 1, 3, 11, 2, 2, 4, 7, 1, 2, 4, 5, 11, 3, 3, 3, 5, 11, 2, 2, 9, 2, 4, 10, 2, 1, 10, 2, 1, 6, 13, 13, 13, 8, 3, 3, 3, 7, 5, 2, 6, 13, 13, 13, 8, 5, 2, 4, 8, 2, 9, 2, 4, 10, 2, 7, 1, 2, 1, 6, 13, 13, 1, 7, 2, 4, 8, 4, 4, 5, 10},
		{0, 6, 1, 1, 6, 1, 1, 8, 2, 1, 10, 3, 1, 9, 2, 4, 11, 2, 4, 8, 2, 2, 6, 3, 3, 11, 3, 3, 3, 2, 2, 5, 1, 10, 3, 5, 7, 13, 13, 13, 5, 7, 13, 13, 9, 4, 2, 5, 3, 4, 8, 0, 2, 3, 9, 1, 2, 11, 2, 3, 5, 4, 8, 3, 3, 7, 3, 3, 0, 11, 4, 4, 2, 9, 0, 10, 2, 2, 2, 7, 13, 13, 13, 6, 1, 2, 4, 10, 4, 4, 1},
		{0, 1, 8, 13, 13, 2, 2, 2, 1, 6, 2, 1, 6, 2, 2, 2, 1, 8, 13, 13, 13, 6, 3, 1, 4, 3, 9, 4, 1, 2, 1, 1, 1, 11, 4, 1, 8, 13, 13, 13, 7, 2, 5, 2, 2, 7, 2, 3, 5, 6, 0, 5, 7, 3, 3, 3, 9, 4, 3, 4, 9, 3, 4, 3, 11, 1, 2, 9, 0, 11, 2, 3, 7, 4, 4, 4, 8, 0, 10, 5, 2, 4, 5, 3, 10, 2, 3, 11, 3},
		{0, 8, 1, 2, 11, 3, 3, 10, 1, 1, 1, 8, 2, 2, 2, 9, 1, 3, 4, 11, 0, 10, 1, 2, 7, 3, 5, 6, 13, 13, 10, 2, 1, 5, 7, 4, 4, 10, 1, 2, 9, 0, 11, 1, 4, 4, 7, 5, 2, 8, 0, 11, 2, 2, 2, 2, 6, 13, 13, 13, 9, 4, 5, 2, 6, 13, 13, 9, 3, 3, 3, 4, 7, 3, 4, 5, 4, 6, 1, 3, 4, 8, 4},
		{3, 1, 1, 11, 1, 3, 5, 3, 1, 4, 6, 1, 3, 5, 1, 10, 1, 5, 3, 1, 3, 7, 1, 3, 4, 6, 13, 13, 5, 3, 3, 7, 3, 2, 10, 2, 3, 7, 2, 2, 8, 2, 4, 3, 2, 4, 9, 1, 1, 1, 9, 13, 13, 13, 10, 2, 4, 3, 3, 3, 8, 2, 3, 3, 2, 6, 13, 13, 5, 11, 1, 4, 8, 1, 4, 9, 4, 3, 4, 3, 1, 4, 2, 3},
	}

	hifgReelStrips97_38 = games.ReelStrips{
		{1, 1, 1, 1, 4, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 5, 4, 1, 4, 2, 2, 2, 2, 3, 3, 3, 3, 2, 4, 3, 5, 1, 2, 5, 1, 3, 1, 2, 5, 4, 3, 2, 4, 3, 5, 4, 1, 3, 4, 1},
		{0, 5, 1, 2, 0, 4, 5, 2, 2, 1, 1, 1, 1, 3, 1, 2, 3, 2, 1, 5, 0, 2, 2, 2, 0, 3, 3, 3, 1, 5, 0, 4, 1, 2, 4, 3, 3, 3, 2, 0, 3, 1, 4, 3, 3, 3, 3, 1, 4, 3, 4, 4},
		{0, 1, 1, 1, 0, 5, 1, 2, 3, 3, 1, 3, 3, 2, 2, 4, 1, 3, 2, 1, 0, 5, 2, 4, 0, 1, 3, 2, 5, 4, 0, 2, 5, 4, 4, 4, 1, 3, 4, 0, 2, 2, 1, 3, 4, 2, 5, 3, 1},
		{0, 2, 5, 4, 0, 2, 5, 1, 1, 1, 1, 1, 1, 4, 2, 2, 4, 1, 2, 5, 0, 2, 1, 4, 0, 2, 2, 1, 5, 4, 0, 4, 1, 2, 4, 1, 3, 3, 3, 0, 5, 3, 3, 4, 3, 3, 4, 3, 3},
		{1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 2, 1, 1, 2, 1, 2, 3, 2, 4, 2, 3, 3, 3, 2, 3, 3, 2, 4, 3, 2, 5, 1, 1, 1, 2, 5, 2, 4, 3, 2, 5, 4, 1, 3, 5, 4, 2, 3, 3, 5, 3, 2, 4, 2, 4, 4},
	}

	// 中間 2 3 4 大圖，同一座標去算
	bigfgReelStrips97_38 = games.ReelStrips{
		{8, 2, 6, 1, 11, 2, 3, 1, 7, 2, 1, 4, 10, 2, 1, 6, 2, 3, 4, 1, 2, 3, 11, 1, 2, 3, 10, 1, 3, 4, 2, 9, 3, 5, 9, 1, 3, 1, 2, 10, 2, 3, 11, 5, 2, 1, 8, 5, 1, 3, 7, 3, 2, 1},
		{3, 4, 3, 4, 8, 3, 1, 3, 2, 9, 2, 5, 2, 1, 1, 2, 10, 3, 8, 3, 4, 8, 5, 10, 2, 9, 1, 5, 2, 6, 5, 0, 6, 7, 4, 11, 1, 6, 11, 2, 4, 6, 4, 3, 5, 4, 1, 2},
		{3, 4, 3, 4, 8, 3, 1, 3, 2, 9, 2, 5, 2, 1, 1, 2, 10, 3, 8, 3, 4, 8, 5, 10, 2, 9, 1, 5, 2, 6, 5, 0, 6, 7, 4, 11, 1, 6, 11, 2, 4, 6, 4, 3, 5, 4, 1, 2},
		{3, 4, 3, 4, 8, 3, 1, 3, 2, 9, 2, 5, 2, 1, 1, 2, 10, 3, 8, 3, 4, 8, 5, 10, 2, 9, 1, 5, 2, 6, 5, 0, 6, 7, 4, 11, 1, 6, 11, 2, 4, 6, 4, 3, 5, 4, 1, 2},
		{9, 5, 3, 9, 11, 5, 9, 7, 4, 7, 9, 3, 11, 7, 2, 8, 5, 11, 4, 8, 6, 4, 8, 2, 4, 1, 11, 10, 5, 5, 1, 6, 1, 9, 5, 10, 8, 3, 11, 6, 2, 9, 6, 1, 10, 8, 5, 10, 7, 11, 6, 1, 7, 10, 2, 8, 7, 4, 10, 2, 6, 3, 4, 2, 6, 5, 7, 3},
	}
)

var (
	ngReelStrips97_68 = games.ReelStrips{
		{1, 1, 1, 3, 1, 3, 1, 6, 1, 4, 6, 3, 1, 7, 3, 2, 7, 4, 1, 8, 3, 1, 9, 4, 4, 10, 3, 1, 11, 2, 2, 10, 1, 1, 1, 5, 1, 4, 4, 6, 1, 4, 6, 1, 4, 7, 13, 13, 13, 7, 1, 1, 8, 2, 2, 2, 8, 13, 13, 13, 9, 4, 4, 11, 2, 2, 2, 9, 2, 4, 10, 5, 4, 8, 3, 2, 10, 3, 4, 11, 1, 3, 6, 13, 13, 7, 2, 2, 2, 6, 3, 3, 7, 4, 5, 9, 5, 1, 11, 5, 4},
		{0, 6, 2, 2, 6, 3, 3, 6, 7, 5, 4, 7, 5, 3, 4, 7, 5, 4, 10, 8, 1, 1, 1, 1, 1, 2, 1, 1, 8, 1, 1, 1, 1, 9, 3, 5, 9, 13, 13, 13, 8, 3, 2, 9, 13, 13, 8, 4, 3, 11, 10, 0, 2, 10, 4, 4, 4, 11, 3, 4, 10, 0, 3, 6, 2, 1, 4, 11, 0, 4, 6, 2, 2, 6, 2, 2, 2, 2, 2, 7, 13, 13, 13, 7, 3, 1, 1, 3, 0, 7, 5, 1, 9, 0, 1, 11, 2},
		{0, 1, 7, 13, 13, 8, 1, 1, 9, 2, 3, 11, 1, 2, 10, 0, 6, 3, 7, 13, 13, 13, 8, 3, 1, 11, 3, 1, 6, 3, 7, 2, 3, 8, 4, 0, 1, 7, 13, 13, 13, 8, 4, 4, 10, 4, 4, 6, 3, 2, 1, 0, 10, 1, 2, 2, 9, 4, 5, 10, 4, 5, 11, 3, 4, 11, 2, 3, 6, 0, 7, 3, 0, 9, 2, 5, 9, 5, 0, 8, 1, 1, 1, 9, 1, 2, 11, 2, 2, 2, 2, 9, 1, 5, 10, 4, 4, 11, 4, 1, 10},
		{0, 8, 3, 1, 4, 3, 1, 1, 1, 1, 1, 1, 3, 10, 3, 7, 0, 6, 5, 7, 0, 11, 4, 2, 4, 4, 8, 2, 5, 10, 3, 3, 10, 5, 2, 11, 4, 4, 1, 2, 7, 0, 4, 3, 11, 1, 3, 3, 10, 1, 0, 11, 6, 1, 2, 4, 9, 13, 13, 13, 9, 1, 8, 5, 9, 13, 13, 9, 5, 9, 5, 6, 7, 0, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 6, 5, 6, 8, 13, 13, 8, 2, 2, 3, 7},
		{1, 1, 1, 1, 1, 1, 3, 6, 7, 3, 11, 10, 5, 7, 9, 4, 7, 6, 4, 3, 8, 10, 1, 1, 1, 1, 1, 10, 11, 5, 5, 6, 11, 1, 1, 10, 8, 2, 2, 2, 2, 2, 6, 7, 1, 6, 7, 3, 8, 6, 3, 7, 9, 1, 7, 6, 1, 1, 8, 10, 4, 4, 4, 4, 4, 4, 4, 11, 3, 6, 11, 13, 13, 10, 7, 5, 3, 8, 9, 1, 8, 10, 5, 9, 11, 13, 13, 13, 11, 6, 3, 5, 10, 7, 4, 4, 9, 5, 5, 9, 11, 4, 4, 4, 9, 13, 13, 11, 3, 2, 5, 2, 8, 3, 2, 2, 2, 8, 2, 2, 2, 8, 2, 2, 2, 2, 5, 9, 8, 2, 2, 2, 10, 9},
	}
	hifgReelStrips97_68 = games.ReelStrips{
		{1, 1, 1, 1, 3, 5, 4, 4, 1, 1, 1, 1, 1, 1, 1, 3, 5, 2, 2, 2, 1, 5, 4, 4, 4, 2, 2, 2, 3, 1, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 1, 5, 3, 2, 2, 5, 1, 4, 4, 4},
		{1, 1, 1, 1, 5, 4, 4, 1, 1, 2, 3, 0, 4, 4, 4, 1, 3, 5, 0, 3, 5, 0, 4, 3, 1, 2, 2, 2, 1, 4, 4, 0, 2, 2, 1, 2, 2, 0, 2, 5, 0, 2, 2, 2, 2, 2, 2, 4, 4, 5, 3},
		{1, 1, 1, 1, 5, 4, 4, 1, 2, 4, 3, 0, 4, 4, 3, 1, 3, 5, 0, 3, 5, 0, 4, 3, 1, 2, 2, 2, 1, 3, 4, 0, 2, 2, 1, 2, 2, 0, 2, 5, 0, 3, 4, 4, 3, 2, 4, 3, 4, 5, 3},
		{1, 1, 1, 1, 1, 1, 2, 2, 2, 1, 0, 1, 2, 2, 2, 1, 3, 4, 5, 0, 3, 3, 5, 0, 4, 3, 3, 0, 5, 4, 4, 1, 4, 3, 2, 4, 5, 0, 4, 4, 4, 0, 5, 3, 2, 4},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 4, 4, 2, 2, 2, 5, 2, 2, 2, 5, 3, 1, 1, 3, 3, 2, 2, 4, 4, 5, 1, 1, 1, 5, 4, 3, 1, 5, 3, 3, 1, 3, 3, 2, 2, 2, 4, 4, 4},
	}

	// 中間 2 3 4 大圖，同一座標去算
	bigfgReelStrips97_68 = games.ReelStrips{
		{1, 9, 2, 6, 1, 3, 6, 8, 1, 3, 9, 10, 4, 9, 5, 2, 1, 8, 4, 11, 5, 4, 5, 11, 4, 1, 6, 1, 8, 6, 1, 11, 6, 4, 2, 7, 10, 1, 7, 11, 3, 1, 7, 1, 10, 8, 2, 1, 10, 3, 9, 7, 2, 5, 3, 1, 5, 2},
		{3, 9, 0, 10, 2, 9, 1, 10, 8, 6, 1, 2, 5, 6, 1, 2, 1, 11, 1, 7, 5, 1, 7, 1, 5, 1, 8, 2, 7, 5, 11, 2, 1, 3, 4, 2, 6, 3, 4, 3, 4, 2, 1, 4, 1, 1, 1},
		{3, 9, 0, 10, 2, 9, 1, 10, 8, 6, 1, 2, 5, 6, 1, 2, 1, 11, 1, 7, 5, 1, 7, 1, 5, 1, 8, 2, 7, 5, 11, 2, 1, 3, 4, 2, 6, 3, 4, 3, 4, 2, 1, 4, 1, 1, 1},
		{3, 9, 0, 10, 2, 9, 1, 10, 8, 6, 1, 2, 5, 6, 1, 2, 1, 11, 1, 7, 5, 1, 7, 1, 5, 1, 8, 2, 7, 5, 11, 2, 1, 3, 4, 2, 6, 3, 4, 3, 4, 2, 1, 4, 1, 1, 1},
		{5, 6, 1, 5, 7, 2, 8, 10, 1, 7, 5, 10, 2, 6, 2, 5, 11, 6, 5, 1, 9, 8, 2, 11, 5, 3, 10, 1, 1, 8, 5, 7, 2, 8, 5, 11, 1, 9, 11, 2, 6, 5, 3, 1, 10, 4, 6, 3, 9, 4, 5, 9, 1, 10, 1, 3, 5, 7, 4},
	}
)

var (
	ngReelStrips97_88 = games.ReelStrips{
		{3, 3, 1, 1, 1, 1, 1, 6, 1, 1, 6, 3, 2, 7, 3, 2, 7, 4, 4, 8, 3, 3, 9, 4, 4, 10, 3, 1, 11, 2, 2, 10, 1, 1, 1, 1, 1, 4, 4, 6, 5, 4, 6, 1, 4, 7, 13, 13, 13, 7, 1, 1, 8, 2, 2, 2, 8, 13, 13, 13, 9, 4, 4, 11, 2, 2, 2, 9, 2, 4, 10, 5, 4, 8, 3, 2, 10, 3, 4, 11, 1, 3, 6, 13, 13, 7, 2, 2, 2, 6, 3, 3, 7, 4, 5, 9, 5, 1, 11, 5, 4},
		{0, 6, 2, 0, 6, 3, 3, 6, 2, 2, 7, 5, 4, 7, 5, 3, 4, 7, 5, 4, 10, 8, 1, 1, 1, 1, 1, 2, 1, 1, 8, 1, 1, 1, 1, 9, 3, 5, 9, 13, 13, 13, 8, 3, 2, 9, 13, 13, 8, 4, 3, 11, 10, 0, 2, 10, 4, 4, 4, 11, 3, 4, 10, 0, 3, 6, 2, 1, 4, 11, 0, 4, 6, 2, 2, 6, 4, 4, 4, 4, 4, 7, 13, 13, 13, 7, 3, 1, 1, 3, 0, 7, 5, 1, 9, 0, 1, 11, 2},
		{0, 6, 7, 13, 13, 8, 2, 1, 9, 2, 3, 11, 1, 2, 10, 0, 6, 3, 7, 13, 13, 13, 8, 3, 1, 11, 3, 1, 6, 3, 7, 2, 3, 8, 4, 0, 6, 7, 13, 13, 13, 8, 4, 4, 10, 4, 4, 6, 3, 2, 7, 0, 10, 8, 1, 2, 9, 4, 5, 10, 4, 5, 11, 3, 4, 11, 2, 3, 6, 0, 7, 3, 0, 9, 2, 5, 9, 5, 0, 8, 1, 1, 1, 9, 1, 2, 11, 4, 4, 4, 4, 9, 1, 5, 10, 4, 4, 11, 4, 2, 10},
		{0, 8, 10, 1, 4, 8, 10, 1, 1, 1, 7, 11, 3, 10, 3, 6, 7, 0, 6, 10, 5, 7, 0, 11, 10, 4, 4, 4, 8, 2, 5, 9, 3, 3, 10, 5, 2, 11, 4, 4, 5, 8, 7, 0, 7, 6, 3, 11, 8, 3, 3, 10, 9, 0, 11, 6, 1, 2, 11, 9, 13, 13, 13, 9, 5, 8, 5, 9, 13, 13, 9, 5, 9, 5, 6, 7, 0, 10, 7, 2, 11, 5, 11, 6, 4, 4, 4, 4, 4, 3, 6, 5, 6, 8, 13, 13, 2, 8, 9, 3, 7},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 6, 7, 3, 11, 10, 5, 7, 9, 4, 7, 6, 4, 3, 8, 10, 1, 1, 1, 1, 1, 10, 11, 5, 5, 6, 11, 1, 1, 10, 8, 4, 4, 4, 4, 4, 6, 7, 1, 6, 7, 3, 8, 6, 3, 7, 9, 1, 7, 6, 1, 1, 8, 10, 4, 4, 4, 4, 4, 4, 4, 11, 3, 6, 11, 13, 13, 10, 7, 5, 3, 8, 9, 1, 8, 10, 5, 9, 11, 13, 13, 13, 11, 6, 3, 5, 10, 7, 4, 4, 9, 5, 5, 9, 11, 4, 4, 4, 9, 13, 13, 11, 3, 2, 5, 2, 8, 3, 2, 2, 2, 8, 2, 2, 2, 8, 2, 2, 2, 2, 5, 9, 8, 2, 2, 2, 10, 9},
	}

	hifgReelStrips97_88 = games.ReelStrips{
		{1, 1, 1, 1, 3, 5, 4, 4, 1, 1, 1, 1, 1, 1, 1, 3, 5, 2, 2, 2, 1, 5, 4, 4, 4, 2, 2, 2, 3, 1, 4, 4, 4, 4, 4, 4, 4, 2, 2, 3, 1, 5, 3, 2, 2, 5, 1, 4, 4, 4},
		{1, 1, 1, 1, 5, 4, 4, 1, 1, 2, 3, 0, 4, 4, 4, 1, 3, 5, 0, 3, 5, 0, 4, 3, 1, 2, 2, 2, 1, 4, 4, 0, 2, 2, 1, 2, 2, 0, 2, 5, 0, 4, 4, 4, 4, 2, 4, 4, 4, 5, 3},
		{1, 1, 1, 1, 5, 4, 4, 1, 2, 4, 3, 0, 4, 4, 3, 1, 3, 5, 0, 3, 5, 0, 4, 3, 1, 2, 2, 2, 1, 3, 4, 0, 2, 2, 1, 2, 2, 0, 2, 5, 0, 3, 4, 4, 3, 2, 4, 3, 4, 5, 3},
		{1, 1, 2, 1, 1, 1, 2, 1, 1, 2, 0, 1, 3, 3, 3, 1, 3, 4, 5, 0, 3, 3, 5, 0, 4, 3, 3, 0, 5, 4, 4, 1, 4, 3, 2, 4, 5, 0, 4, 4, 4, 0, 5, 3, 2, 4},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 4, 4, 2, 3, 4, 4, 1, 1, 1, 5, 2, 2, 2, 5, 3, 1, 1, 3, 3, 2, 2, 4, 4, 5, 1, 1, 1, 5, 4, 3, 1, 5, 3, 3, 1, 3, 3, 2, 2, 2, 4, 4, 4},
	}

	// 中間 2 3 4 大圖，同一座標去算
	bigfgReelStrips97_88 = games.ReelStrips{
		{6, 9, 2, 6, 7, 3, 6, 8, 1, 3, 9, 10, 4, 9, 5, 2, 10, 8, 4, 11, 5, 4, 5, 11, 4, 11, 6, 1, 8, 6, 1, 11, 6, 4, 2, 7, 10, 1, 7, 11, 3, 8, 7, 1, 10, 8, 2, 9, 10, 3, 9, 7, 2, 5, 3, 1, 5, 2},
		{3, 9, 0, 10, 2, 9, 1, 10, 3, 6, 1, 4, 3, 6, 1, 3, 4, 11, 1, 7, 5, 3, 7, 1, 5, 1, 8, 2, 7, 5, 11, 2, 8, 3, 4, 2, 6, 3, 4, 3, 4, 4, 1, 4, 4},
		{3, 9, 0, 10, 2, 9, 1, 10, 3, 6, 1, 4, 3, 6, 1, 3, 4, 11, 1, 7, 5, 3, 7, 1, 5, 1, 8, 2, 7, 5, 11, 2, 8, 3, 4, 2, 6, 3, 4, 3, 4, 4, 1, 4, 4},
		{3, 9, 0, 10, 2, 9, 1, 10, 3, 6, 1, 4, 3, 6, 1, 3, 4, 11, 1, 7, 5, 3, 7, 1, 5, 1, 8, 2, 7, 5, 11, 2, 8, 3, 4, 2, 6, 3, 4, 3, 4, 4, 1, 4, 4},
		{5, 6, 1, 5, 7, 2, 8, 10, 1, 7, 5, 10, 2, 6, 8, 9, 1, 6, 7, 1, 5, 8, 2, 11, 5, 3, 10, 11, 1, 8, 5, 7, 2, 8, 5, 11, 1, 9, 11, 2, 6, 5, 3, 9, 10, 4, 6, 3, 9, 4, 5, 9, 1, 10, 11, 3, 5, 7, 4},
	}
)
