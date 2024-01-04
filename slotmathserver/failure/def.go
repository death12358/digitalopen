package failure

import "time"

const (
	ACCOUNT_FROZEN           = "ACCOUNT_FROZEN"
	DUPLICATED_OPERATION     = "DUPLICATED_OPERATION"
	DUPLICATED_STAGE         = "DUPLICATED_STAGE"
	INSUFFICIENT_BALANCE     = "INSUFFICIENT_BALANCE"
	INVALID_ARGUMENT         = "INVALID_ARGUMENT"
	INVALID_OPERATION        = "INVALID_OPERATION"
	MISSING_STAGE            = "MISSING_STAGE"
	MISSING_TICKET           = "MISSING_TICKET"
	MISSING_WALLET           = "MISSING_WALLET"
	NOT_AFFECTED             = "NOT_AFFECTED"
	UNKNOWN_ERROR            = "UNKNOWN_ERROR"
	DUPLICATE_WALLET_ADDRESS = "DUPLICATE_WALLET_ADDRESS"
)

func IsKnownErrorCode(message string) bool {
	if size := len(message); (size == 0) || (size > 64) {
		return false
	}

	for i := 0; i < len(message); i++ {
		ch := message[i]

		if ch == '_' ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') {
			continue
		} else {
			return false
		}
	}
	return true
}

func ThrowFailure(err error) {
	var failure *Failure
	if IsKnownErrorCode(err.Error()) {
		failure = &Failure{
			Message:   err.Error(),
			Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
		}
	} else {
		failure = &Failure{
			Message:     INVALID_OPERATION,
			Description: err.Error(),
			Timestamp:   time.Now().UnixNano() / int64(time.Millisecond),
		}
	}
	panic(failure)
}

func ThrowFailureMessage(message string, reason string) {
	var failure *Failure
	if IsKnownErrorCode(message) {
		failure = &Failure{
			Message:     message,
			Description: reason,
			Timestamp:   time.Now().UnixNano() / int64(time.Millisecond),
		}
	} else {
		var desc string = message
		if len(reason) > 0 {
			desc += " " + reason
		}
		failure = &Failure{
			Message:     INVALID_OPERATION,
			Description: desc,
			Timestamp:   time.Now().UnixNano() / int64(time.Millisecond),
		}
	}
	panic(failure)
}
