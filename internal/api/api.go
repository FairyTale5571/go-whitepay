package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL string
	Client  *http.Client
	ErrorParserFunc
}

// Request for API
type Request struct {
	Method   string
	Path     string
	Params   map[string]string
	Headers  map[string]string
	Body     any
	Response any
}

// prepareRequestBody prepare request body
func prepareRequestBody(req Request) (io.ReadCloser, error) {
	if req.Body == nil {
		return http.NoBody, nil
	}
	bodyBytes, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}
	return io.NopCloser(bytes.NewReader(bodyBytes)), nil
}

// isNotOkStatus check if status is not ok
func isNotOkStatus(status int) bool {
	return status < http.StatusOK || status > http.StatusMultipleChoices
}

// SendRequest send request to API
func (a *Client) SendRequest(ctx context.Context, r Request) error {
	reqURL, err := url.Parse(fmt.Sprintf("%s%s", a.BaseURL, r.Path))
	if err != nil {
		return err
	}
	reqURL.Path = r.Path
	values := url.Values{}
	for k, v := range r.Params {
		values.Add(k, v)
	}
	reqURL.RawQuery = values.Encode()

	body, err := prepareRequestBody(r)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, r.Method, reqURL.String(), body)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}

	resp, err := a.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if isNotOkStatus(resp.StatusCode) {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return a.ErrorParserFunc(bodyBytes)
	}

	if r.Response == nil {
		return nil
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyBytes, &r.Response)
}
