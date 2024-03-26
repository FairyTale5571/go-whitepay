package whitepay

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// https://docs.whitepay.com/docs/webhooks/events#webhooks-settings

type EventType string

const (
	EventWithdrawalCompleted EventType = "withdrawal::completed"
	EventWithdrawalDeclined  EventType = "withdrawal::declined"
	EventRollbackToMerchant  EventType = "rollback::to_merchant"
	EventRollbackToClient    EventType = "rollback::to_client"

	EventOrderCompleted              EventType = "order::completed"
	EventOrderDeclined               EventType = "order::declined"
	EventOrderPartiallyFulfilled     EventType = "order::partially_fulfilled"
	EventOrderFinalAmountWasReceived EventType = "order::final_amount_was_received"

	EventTransactionCompleted        EventType = "transaction::complete"
	EventTransactionDeclined         EventType = "transaction::decline"
	EventTransactionWasFinalExchange EventType = "transaction::was_final_exchange"
)

type EventTransaction struct {
	Transaction `json:"transaction"`
	EventType   `json:"event_type"`
}

func (wp *WhitePay) SignEvent(event any) string {
	payload, err := json.Marshal(event)
	if err != nil {
		return ""
	}
	h := hmac.New(sha256.New, []byte(wp.token))
	h.Write(payload)
	return hex.EncodeToString(h.Sum(nil))
}
