package resource

import "strings"

const (
	MIMETYPE_JSON = "application/json; charset=utf8"
	MIMETYPE_TEXT = "text/plain; charset=utf8"
)

// getRTP retrieves the RTP based on the provided currency. If the currency
// is not recognized, it returns the default RTP.
// currency: a string representing the currency for which to retrieve the RTP
// defaultValue: a string representing the default RTP to return if the currency is not recognized
// returns: a string representing the RTP for the specified currency or the default RTP
func getRTP(GameName string, currency string, defaultValue string) string {
	var rtp string

	switch strings.ToUpper(currency) {

	// Gold Coin
	case "GCN":
		rtp, _ = ChooseRTP(GameName)

	// Silver Coin
	case "SCN":
		_, rtp = ChooseRTP(GameName)

	default:
		rtp = defaultValue
		//rtp = "40"
	}

	return rtp
}

func ChooseRTP(GameName string) (string, string) {
	var GRTP, SRTP string
	switch GameName {
	// Gold Coin
	case "SG001":
		GRTP = "98"
		SRTP = "98"
	case "SG002":
		GRTP = "98"
		SRTP = "98"
	case "SG003":
		GRTP = "98"
		SRTP = "98"
	case "SG006":
		GRTP = "98"
		SRTP = "98"
	case "SG008":
		GRTP = "98"
		SRTP = "98"
	case "SG009":
		GRTP = "98"
		SRTP = "98"
	case "DCG001":
		GRTP = "98"
		SRTP = "98"
	case "DCG008": //瑪雅
		GRTP = "96"
		SRTP = "96"
	default:
		GRTP = "40"
		SRTP = "40"

	}
	return GRTP, SRTP
}
