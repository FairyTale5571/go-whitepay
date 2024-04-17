package whitepay

type Status string

const (
	StatusComplete           Status = "COMPLETE"
	StatusOpen               Status = "OPEN"
	StatusDeclined           Status = "DECLINED"
	StatusInit               Status = "INIT"
	StatusPartiallyFulfilled Status = "PARTIALLY_FULFILLED"
)
