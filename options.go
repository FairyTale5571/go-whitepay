package whitepay

import "net/http"

type Option func(*WhitePay)

// WithClient sets the HTTP client for the WhitePay client.
func WithClient(client *http.Client) Option {
	return func(wp *WhitePay) {
		wp.client.Client = client
	}
}
