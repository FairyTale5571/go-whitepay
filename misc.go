package whitepay

import (
	"context"
	"fmt"
	"github.com/fairytale5571/go-whitepay/pkg/api"
	"net/http"
	"net/url"
)

type ShorOrderDetailsResponse struct {
	Order `json:"order"`
}

// ShowOrderDetails Show Order Details
func (wp *WhitePay) ShowOrderDetails(ctx context.Context, orderID string) (*ShorOrderDetailsResponse, error) {
	var resp ShorOrderDetailsResponse
	if err := wp.client.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("/private-api/orders/%s/%s", wp.slug, orderID),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", wp.token),
		},
		Response: &resp,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ShowOrdersListResponse struct {
	Orders []*Order   `json:"orders"`
	Meta   Pagination `json:"meta"`
}

func (s *ShowOrdersListResponse) Next() bool {
	return s.Meta.Next()
}

// ShowOrdersList Get Show Orders List
func (wp *WhitePay) ShowOrdersList(ctx context.Context, filters ...Filters) (*ShowOrdersListResponse, error) {
	values := url.Values{}
	for _, filter := range filters {
		for key, value := range filter() {
			values[key] = value
		}
	}

	rawURL, err := url.Parse(fmt.Sprintf("/private-api/orders/%s", wp.slug))
	if err != nil {
		return nil, err
	}
	rawURL.RawQuery = values.Encode()

	var resp ShowOrdersListResponse
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

type OrderStatusesListResponse struct {
	Statuses []*Status `json:"statuses"`
}

// ShowOrderStatusesList Get Show Statuses List
func (wp *WhitePay) ShowOrderStatusesList(ctx context.Context) (*OrderStatusesListResponse, error) {
	var resp OrderStatusesListResponse
	if err := wp.client.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   "/private-api/order-statuses",
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", wp.token),
		},
		Response: &resp,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

type AccountBalancesResponse struct {
	Balances []*Balance `json:"balances"`
}

// ShowAccountBalances Get Show Account Balances
func (wp *WhitePay) ShowAccountBalances(ctx context.Context) (*AccountBalancesResponse, error) {
	var resp AccountBalancesResponse
	if err := wp.client.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   "/private-api/balances",
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", wp.token),
		},
		Response: &resp,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}
