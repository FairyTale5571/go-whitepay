package whitepay

import (
	"net/http"

	"github.com/fairytale5571/go-whitepay/internal/api"
)

const BaseURL = "https://api.whitepay.com"

// WhitePay represents a WhitePay client.
type WhitePay struct {
	client *api.Client
	token  string
	slug   string
}

// New creates a new WhitePay client with the given token.
func New(token, slug string, opts ...Option) *WhitePay {
	wp := &WhitePay{
		client: &api.Client{
			BaseURL:         BaseURL,
			Client:          http.DefaultClient,
			ErrorParserFunc: errorParser,
		},
		token: token,
		slug:  slug,
	}

	for _, opt := range opts {
		opt(wp)
	}

	return wp
}

func (wp *WhitePay) SetSlug(slug string) {
	wp.slug = slug
}
