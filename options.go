package whitepay

import "net/http"

type Option func(*WhitePay)

// WithClient sets the HTTP client for the WhitePay client.
func WithClient(client *http.Client) Option {
	return func(wp *WhitePay) {
		wp.client.Client = client
	}
}

// WithChannelSize sets the size of the updates channel.
func WithChannelSize(size int) Option {
	return func(wp *WhitePay) {
		wp.updates = make(chan *Events, size)
	}
}

// WithSlugSignatureToken sets the slug signature token for the WhitePay client.
func WithSlugSignatureToken(token string) Option {
	return func(wp *WhitePay) {
		wp.slugSignatureToken = token
	}
}

// WithMerchantSignatureToken sets the merchant signature token for the WhitePay client.
func WithMerchantSignatureToken(token string) Option {
	return func(wp *WhitePay) {
		wp.merchantSignatureToken = token
	}
}
