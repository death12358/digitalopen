package sg006

import "github.com/death12358/digitalopen/games"

var (
	// RTP: 95.722%

	ngReelStrips98_100 = games.ReelStrips{
		{1, 8, 11, 6, 7, 10, 5, 11, 6, 8, 2, 9, 11, 5, 10, 7, 2, 9, 10, 8, 11, 3, 9, 7, 4, 10, 9, 1, 8, 11, 6, 7, 10, 5, 11, 6, 8, 2, 9, 11, 5, 10, 7, 2, 6, 10, 8, 11, 3, 6, 7, 4, 10, 9, 1, 8, 11, 6, 7, 10, 5, 11, 9, 8, 2, 9, 11, 5, 10, 7, 2, 9, 10, 8, 11, 3, 9, 7, 4, 10, 9},
		{13, 13, 9, 13, 13, 11, 13, 9, 1, 11, 10, 3, 11, 6, 4, 6, 11, 5, 9, 6, 11, 2, 6, 8, 2, 7, 11, 7, 5, 8, 6, 3, 10, 9, 1, 11, 10, 3, 11, 6, 4, 6, 10, 5, 9, 6, 10, 2, 6, 8, 2, 7, 10, 7, 5, 8, 6, 3, 11, 9, 1, 11, 10, 3, 11, 6, 4, 6, 11, 5, 9, 6, 11, 2, 6, 8, 2, 7, 11, 7, 5, 8, 6, 3, 11, 9, 1, 11, 10, 3, 11, 6, 4, 6, 9, 5, 11, 6, 9, 2, 6, 8, 2, 7, 11, 7, 5, 8, 6, 3, 11, 9, 1, 11, 10, 3, 11, 6, 4, 6, 11, 5, 9, 6, 11, 2, 6, 8, 2, 7, 11, 7, 5, 8, 6, 3},
		{13, 13, 13, 10, 13, 1, 6, 8, 3, 11, 8, 10, 5, 7, 6, 3, 7, 10, 5, 9, 11, 4, 7, 10, 2, 8, 9, 1, 8, 10, 1, 6, 8, 3, 9, 8, 10, 5, 7, 6, 3, 7, 10, 5, 9, 11, 4, 7, 10, 2, 8, 11, 1, 8, 10, 1, 6, 8, 3, 11, 8, 10, 5, 7, 6, 3, 7, 10, 5, 9, 11, 4, 7, 10, 2, 8, 11, 1, 8, 10, 1, 6, 8, 3, 11, 8, 10, 5, 7, 6, 3, 7, 10, 5, 9, 11, 4, 7, 10, 2, 8, 9, 1, 8, 10, 1, 6, 8, 3, 9, 8, 10, 5, 7, 6, 3, 7, 10, 5, 9, 11, 4, 7, 10, 2, 8, 9, 1, 8, 10},
		{13, 11, 13, 13, 13, 11, 13, 9, 2, 11, 9, 3, 10, 7, 1, 10, 11, 5, 6, 8, 3, 10, 7, 2, 6, 7, 4, 6, 9, 5, 11, 9, 2, 11, 9, 3, 10, 7, 1, 10, 11, 5, 6, 8, 3, 10, 7, 2, 6, 7, 4, 6, 9, 5, 11, 9, 2, 11, 9, 3, 10, 7, 1, 10, 11, 5, 6, 8, 3, 10, 7, 2, 6, 7, 4, 6, 9, 5},
		{1, 6, 9, 3, 7, 6, 4, 6, 8, 2, 9, 11, 5, 10, 7, 4, 11, 8, 3, 10, 11, 5, 10, 6, 5, 8, 6, 2, 10, 9, 3, 6, 7, 1, 6, 9, 3, 7, 6, 4, 6, 8, 2, 9, 11, 5, 10, 7, 4, 11, 8, 3, 9, 11, 5, 10, 7, 5, 8, 6, 2, 10, 9, 3, 8, 7, 1, 6, 9, 3, 7, 10, 4, 6, 8, 2, 9, 11, 5, 10, 7, 4, 11, 8, 3, 9, 11, 5, 10, 7, 5, 8, 6, 2, 10, 9, 3, 8, 7},
	}
	fgReelStrips98_100 = games.ReelStrips{
		{1, 8, 11, 6, 7, 10, 5, 11, 6, 8, 2, 9, 11, 5, 10, 7, 2, 9, 10, 8, 11, 3, 9, 7, 4, 10, 9, 1, 8, 11, 6, 7, 10, 5, 11, 6, 8, 2, 9, 11, 5, 10, 7, 2, 6, 10, 8, 11, 3, 6, 7, 4, 10, 9, 1, 8, 11, 6, 7, 10, 5, 11, 9, 8, 2, 9, 11, 5, 10, 7, 2, 9, 10, 8, 11, 3, 9, 7, 4, 10, 9},
		{0, 0, 9, 0, 0, 11, 0, 9, 1, 11, 10, 3, 11, 6, 4, 6, 11, 5, 9, 6, 11, 2, 6, 8, 2, 7, 11, 7, 5, 8, 6, 3, 10, 9, 1, 11, 10, 3, 11, 6, 4, 6, 10, 5, 9, 6, 10, 2, 6, 8, 2, 7, 10, 7, 5, 8, 6, 3, 11, 9, 1, 11, 10, 3, 11, 6, 4, 6, 11, 5, 9, 6, 11, 2, 6, 8, 2, 7, 11, 7, 5, 8, 6, 3, 11, 9, 1, 11, 10, 3, 11, 6, 4, 6, 9, 5, 11, 6, 9, 2, 6, 8, 2, 7, 11, 7, 5, 8, 6, 3, 11, 9, 1, 11, 10, 3, 11, 6, 4, 6, 11, 5, 9, 6, 11, 2, 6, 8, 2, 7, 11, 7, 5, 8, 6, 3},
		{0, 0, 0, 10, 0, 1, 6, 8, 3, 11, 8, 10, 5, 7, 6, 3, 7, 10, 5, 9, 11, 4, 7, 10, 2, 8, 9, 1, 8, 10, 1, 6, 8, 3, 9, 8, 10, 5, 7, 6, 3, 7, 10, 5, 9, 11, 4, 7, 10, 2, 8, 11, 1, 8, 10, 1, 6, 8, 3, 11, 8, 10, 5, 7, 6, 3, 7, 10, 5, 9, 11, 4, 7, 10, 2, 8, 11, 1, 8, 10, 1, 6, 8, 3, 11, 8, 10, 5, 7, 6, 3, 7, 10, 5, 9, 11, 4, 7, 10, 2, 8, 9, 1, 8, 10, 1, 6, 8, 3, 9, 8, 10, 5, 7, 6, 3, 7, 10, 5, 9, 11, 4, 7, 10, 2, 8, 9, 1, 8, 10},
		{0, 11, 0, 0, 0, 11, 0, 9, 2, 11, 9, 3, 10, 7, 1, 10, 11, 5, 6, 8, 3, 10, 7, 2, 6, 7, 4, 6, 9, 5, 11, 9, 2, 11, 9, 3, 10, 7, 1, 10, 11, 5, 6, 8, 3, 10, 7, 2, 6, 7, 4, 6, 9, 5, 11, 9, 2, 11, 9, 3, 10, 7, 1, 10, 11, 5, 6, 8, 3, 10, 7, 2, 6, 7, 4, 6, 9, 5},
		{1, 6, 9, 3, 7, 6, 4, 6, 8, 2, 9, 11, 5, 10, 7, 4, 11, 8, 3, 10, 11, 5, 10, 6, 5, 8, 6, 2, 10, 9, 3, 6, 7, 1, 6, 9, 3, 7, 6, 4, 6, 8, 2, 9, 11, 5, 10, 7, 4, 11, 8, 3, 9, 11, 5, 10, 7, 5, 8, 6, 2, 10, 9, 3, 8, 7, 1, 6, 9, 3, 7, 10, 4, 6, 8, 2, 9, 11, 5, 10, 7, 4, 11, 8, 3, 9, 11, 5, 10, 7, 5, 8, 6, 2, 10, 9, 3, 8, 7},
	}
)
