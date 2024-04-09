package whitepay

import (
	"encoding/json"
	"errors"

	"github.com/fairytale5571/go-whitepay/internal/api"
)

// ErrorBody represents an error response body.
type ErrorBody struct {
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors,omitempty"`
}

func (e ErrorBody) Error() string {
	return e.Message
}

var errorParser api.ErrorParserFunc = func(body []byte) error {
	var errBody ErrorBody
	if err := json.Unmarshal(body, &errBody); err != nil {
		return err
	}
	return errors.New(errBody.Error())
}
