package whitepay

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// EventType https://docs.whitepay.com/docs/webhooks/events#webhooks-settings
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

type Events struct {
	Transaction *Transaction `json:"transaction,omitempty"`
	Order       *Order       `json:"order,omitempty"`
	EventType   EventType    `json:"event_type"`
}

func (wp *WhitePay) SignTransaction(payload []byte) string {
	h := hmac.New(sha256.New, []byte(wp.merchantSignatureToken))
	h.Write(payload)
	return hex.EncodeToString(h.Sum(nil))
}

func (wp *WhitePay) SignOrder(payload []byte) string {
	h := hmac.New(sha256.New, []byte(wp.slugSignatureToken))
	h.Write(payload)
	return hex.EncodeToString(h.Sum(nil))
}
