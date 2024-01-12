package internal

// func ParseToReelsPB(inReels [][]string) []*pb_game.Reel {
// 	var result []*pb_game.Reel
// 	for col := 0; col < len(inReels); col++ {
// 		var reel pb_game.Reel

// 		reel.R = make([]string, len(inReels[col]))

// 		copy(reel.R, inReels[col])

// 		result = append(result, &reel)
// 	}
// 	return result
// }

// func ParseToPayDetailPB(inData []*PayDetail) []*pb_game.PayDetail {
// 	var result []*pb_game.PayDetail
// 	for i := 0; i < len(inData); i++ {
// 		result = append(result, &pb_game.PayDetail{
// 			Symbol:   inData[i].Symbol,
// 			NumOrWay: int32(inData[i].NumOrWay),
// 			Lines:    int32(inData[i].Lines),
// 			Rate:     int32(inData[i].Rate),
// 		})
// 	}
// 	return result
// }
