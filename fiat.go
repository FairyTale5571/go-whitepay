package whitepay

import (
	"context"
	"fmt"
	"github.com/fairytale5571/go-whitepay/pkg/api"
	"net/http"
)

type CreateNewFiatOrder struct {
	Amount          string `json:"amount"`
	Currency        string `json:"currency"`
	ExternalOrderID string `json:"external_order_id"`
}

type CreateNewFiatOrderResponse struct {
	Order `json:"order"`
}

// CreateNewFiatOrder Create New Fiat Order
func (wp *WhitePay) CreateNewFiatOrder(ctx context.Context, req *CreateNewFiatOrder) (*CreateNewFiatOrderResponse, error) {
	var resp CreateNewFiatOrderResponse
	if err := wp.client.SendRequest(ctx, api.Request{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/private-api/orders/%s", wp.slug),
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

type ShowInvoiceCurrenciesForFiatOrderCreationResponse struct {
	Currencies []*Currency `json:"currencies"`
}

// ShowInvoiceCurrenciesForFiatOrderCreation Get Show Invoice Currencies For Fiat Order Creation
func (wp *WhitePay) ShowInvoiceCurrenciesForFiatOrderCreation(ctx context.Context) (*ShowInvoiceCurrenciesForFiatOrderCreationResponse, error) {
	var resp ShowInvoiceCurrenciesForFiatOrderCreationResponse
	if err := wp.client.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   "/currencies/fiat-order-target-currencies",
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", wp.token),
		},
		Response: &resp,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}
