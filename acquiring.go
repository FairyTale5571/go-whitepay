package whitepay

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fairytale5571/go-whitepay/internal/api"
)

/*
	DO NOT USE THIS METHODS RIGHT NOW
*/

// ListCurrencies represents a list of cryptocurrencies.
// Deprecated: DO NOT USE THIS METHODS RIGHT NOW
type ListCurrencies struct {
	CryptoCurrency []CryptoCurrency `json:"crypto_currency"`
}

// GetListCurrencies gets a list of currencies.
// Deprecated: DO NOT USE THIS METHODS RIGHT NOW
func (wp *WhitePay) GetListCurrencies(ctx context.Context) (*ListCurrencies, error) {
	var listCurrencies ListCurrencies
	err := wp.client.SendRequest(ctx, api.Request{
		Method:   http.MethodGet,
		Path:     fmt.Sprintf("/acquiring/%s/currencies", wp.slug),
		Response: &listCurrencies,
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", wp.token),
		},
	})
	if err != nil {
		return nil, err
	}
	return &listCurrencies, nil
}

// CreatePaymentRequest represents a request to create a payment.
// Deprecated: DO NOT USE THIS METHODS RIGHT NOW
type CreatePaymentRequest struct {
	Amount     int    `json:"amount"`
	CurrencyID string `json:"currency_id"`
	Method     string `json:"method"`
	Network    string `json:"network"`
}

// CreatePaymentResponse represents a response to create a payment.
// Deprecated: DO NOT USE THIS METHODS RIGHT NOW
type CreatePaymentResponse struct {
	Order struct {
		ID string `json:"id"`
	} `json:"order"`
}

// CreatePayment creates a payment.
// Deprecated: DO NOT USE THIS METHODS RIGHT NOW
func (wp *WhitePay) CreatePayment(ctx context.Context, req *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	var resp CreatePaymentResponse
	if err := wp.client.SendRequest(ctx, api.Request{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/acquiring/%s/pay", wp.slug),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", wp.token),
		},
		Body:     req,
		Response: &resp,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// WalletOrderDetailResponse represents a response to get wallet order details.
// Deprecated: DO NOT USE THIS METHODS RIGHT NOW
type WalletOrderDetailResponse struct {
	Address string `json:"address"`
	Network string `json:"network"`
	Memo    string `json:"memo,omitempty"`
}

// GetWalletOrderDetail gets wallet order details.
// Deprecated: DO NOT USE THIS METHODS RIGHT NOW
func (wp *WhitePay) GetWalletOrderDetail(ctx context.Context, orderID string) (*WalletOrderDetailResponse, error) {
	var resp WalletOrderDetailResponse
	if err := wp.client.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("/acquiring/%s/order/%s", wp.slug, orderID),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", wp.token),
		},
		Response: &resp,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}
