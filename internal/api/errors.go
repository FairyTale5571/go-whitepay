package api

// ErrorParser interface for parsing errors
type ErrorParser interface {
	Error([]byte) error
}

// ErrorParserFunc function for parsing errors
type ErrorParserFunc func([]byte) error

// Error parsing errors
func (f ErrorParserFunc) Error(body []byte) error {
	return f(body)
}
