package whitepay

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/fairytale5571/go-whitepay/internal/api"
)

// CreateNewOrderRequest body for CreateNewOrder
type CreateNewOrderRequest struct {
	// Amount Min value depends on currency. Becomes required in case if “currency“ was specified.
	Amount string `json:"amount"`

	// Currency Becomes required in case if “amount“ was specified.
	Currency string `json:"currency"`

	// ExternalOrderId Unique order identifier in client Database (usually Internal Payment or Order ID used). Number of characters: min 1, max 255. Allowed symbols: aA-zZ, 0-9. Special characters: “-“, “_“.
	ExternalOrderID string `json:"external_order_id"`

	// SuccessfulLink Link used to redirect user from payment form after order completion. Has to be URL
	SuccessfulLink string `json:"successful_link"`

	// FailureLink Link used to redirect user from payment form after order failure. Has to be URL
	FailureLink string `json:"failure_link"`
}

// CreateNewOrderResponse body for CreateNewOrder
type CreateNewOrderResponse struct {
	Order `json:"order"`
}

// CreateNewOrder Create New Crypto Order
func (wp *WhitePay) CreateNewOrder(ctx context.Context, req *CreateNewOrderRequest) (*CreateNewOrderResponse, error) {
	var resp CreateNewOrderResponse
	if err := wp.client.SendRequest(ctx, api.Request{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/private-api/crypto-orders/%s", wp.slug),
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

// GetCryptoOrderDetailsResponse body for GetCryptoOrderDetails
type GetCryptoOrderDetailsResponse struct {
	Order `json:"order"`
}

// ShowCryptoOrderDetails Get Show Crypto Order Details
func (wp *WhitePay) ShowCryptoOrderDetails(ctx context.Context, orderID string) (*GetCryptoOrderDetailsResponse, error) {
	var resp GetCryptoOrderDetailsResponse
	if err := wp.client.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("/private-api/crypto-orders/%s/%s", wp.slug, orderID),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", wp.token),
		},
		Response: &resp,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ShowCryptoOrdersListResponse body for ShowCryptoOrdersList
type ShowCryptoOrdersListResponse struct {
	Orders []*Order   `json:"orders"`
	Meta   Pagination `json:"meta"`
}

func (s *ShowCryptoOrdersListResponse) Next() bool {
	return s.Meta.Next()
}

// ShowCryptoOrdersList Get Show Crypto Orders List
func (wp *WhitePay) ShowCryptoOrdersList(ctx context.Context, filters ...Filters) (*ShowCryptoOrdersListResponse, error) {
	values := url.Values{}
	for _, filter := range filters {
		for key, value := range filter() {
			values[key] = value
		}
	}

	rawURL, err := url.Parse(fmt.Sprintf("/private-api/crypto-orders/%s", wp.slug))
	if err != nil {
		return nil, err
	}
	rawURL.RawQuery = values.Encode()

	var resp ShowCryptoOrdersListResponse
	if err := wp.client.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   rawURL.String(),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", wp.token),
		},
		Response: &resp,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ShowInvoiceCurrenciesForCryptoOrderCreationResponse struct {
	Currencies []*Currency `json:"currencies"`
}

// ShowInvoiceCurrenciesForCryptoOrderCreation Get Show Invoice Currencies For Crypto Order Creation
func (wp *WhitePay) ShowInvoiceCurrenciesForCryptoOrderCreation(ctx context.Context) (*ShowInvoiceCurrenciesForCryptoOrderCreationResponse, error) {
	var resp ShowInvoiceCurrenciesForCryptoOrderCreationResponse
	if err := wp.client.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   "/currencies/crypto-order-target-currencies",
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", wp.token),
		},
		Response: &resp,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CompleteCryptoOrderResponse body for CompleteCryptoOrder
type CompleteCryptoOrderResponse struct {
	Order `json:"order"`
}

// CompleteCryptoOrder Complete Crypto Order
func (wp *WhitePay) CompleteCryptoOrder(ctx context.Context, orderID string) (*CompleteCryptoOrderResponse, error) {
	var resp CompleteCryptoOrderResponse
	if err := wp.client.SendRequest(ctx, api.Request{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/private-api/crypto-orders/%s/%s/complete", wp.slug, orderID),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", wp.token),
		},
		Response: &resp,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}
