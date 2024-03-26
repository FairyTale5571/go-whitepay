package whitepay

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fairytale5571/go-whitepay/internal/api"
)

// CreateCryptoTopUpRequest body for CreateCryptoTopUp
type CreateCryptoTopUpRequest struct {
	CreateNewOrderRequest
}

// CreateCryptoTopUpResponse response for CreateCryptoTopUp
type CreateCryptoTopUpResponse struct {
	Order `json:"order"`
}

// CreateCryptoTopUp Create Crypto Top Up
func (wp *WhitePay) CreateCryptoTopUp(ctx context.Context, req *CreateCryptoTopUpRequest) (*CreateCryptoTopUpResponse, error) {
	var resp CreateCryptoTopUpResponse
	if err := wp.client.SendRequest(ctx, api.Request{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/private-api/crypto-topups/%s", wp.slug),
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

// ShowInvoiceCurrenciesForCryptoTopUpOrderCreation Get Show Crypto Top Up Details
func (wp *WhitePay) ShowInvoiceCurrenciesForCryptoTopUpOrderCreation(ctx context.Context) *ShowInvoiceCurrenciesForCryptoOrderCreationResponse {
	var resp ShowInvoiceCurrenciesForCryptoOrderCreationResponse
	_ = wp.client.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   "/private-api/currencies/crypto-topup-target-currencies",
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", wp.token),
		},
		Response: &resp,
	})
	return &resp
}
