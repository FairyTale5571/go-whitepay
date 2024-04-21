package whitepay

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/fairytale5571/go-whitepay/internal/api"
)

const (
	defaultUpdatesChanCap = 1024
	BaseURL               = "https://api.whitepay.com"
)

// WhitePay represents a WhitePay client.
type WhitePay struct {
	client *api.Client
	token  string

	merchantSignatureToken string
	slugSignatureToken     string

	slug    string
	updates chan *Events
}

// New creates a new WhitePay client with the given token.
func New(token, slug string, opts ...Option) *WhitePay {
	wp := &WhitePay{
		client: &api.Client{
			BaseURL:         BaseURL,
			Client:          http.DefaultClient,
			ErrorParserFunc: errorParser,
		},
		token:   token,
		slug:    slug,
		updates: make(chan *Events, defaultUpdatesChanCap),
	}

	for _, opt := range opts {
		opt(wp)
	}

	return wp
}

// SetSlug sets the slug for the WhitePay client.
func (wp *WhitePay) SetSlug(slug string) {
	wp.slug = slug
}

// WebhookHandler returns a http.HandlerFunc that can be used to handle WhitePay webhooks.
func (wp *WhitePay) WebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, errReadBody := io.ReadAll(r.Body)
		if errReadBody != nil {
			log.Printf("[WHITEPAY]: error read request body, %s", errReadBody.Error())
			return
		}

		signature := r.Header.Get("Signature")

		update := &Events{}
		errDecode := json.Unmarshal(body, update)
		if errDecode != nil {
			log.Printf("[WHITEPAY]: error decode request body, %s, %s", body, errDecode.Error())
			return
		}

		if signature != "" {
			var compareSignature string
			switch {
			case update.Order != nil && wp.slugSignatureToken != "":
				compareSignature = wp.SignOrder(body)
			case update.Transaction != nil && wp.merchantSignatureToken != "":
				compareSignature = wp.SignTransaction(body)
			}
			if signature != compareSignature {
				log.Printf("[WHITEPAY]: error signature mismatch")
				return
			}
		}

		select {
		case wp.updates <- update:
		default:
			log.Printf("[WHITEPAY]: error send update to processing, channel is full")
		}
	}
}

// Updates returns a channel that can be used to receive WhitePay events.
func (wp *WhitePay) Updates() <-chan *Events {
	return wp.updates
}
