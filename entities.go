package whitepay

type Pagination struct {
	PerPage  int `json:"per_page"`
	LastPage int `json:"last_page"`
	Total    int `json:"total"`
}

func (p *Pagination) Next() bool {
	return p.LastPage > 1
}

// Order represents a WhitePay order.
type Order struct {
	// ID 	internal order identifier
	ID string `json:"id"`

	// Currency invoice currency (ticker)
	Currency string `json:"currency"`

	// OrderNumber 	unique 9-digit numeric order identifier in WhitePay CRM system.
	OrderNumber string `json:"order_number"`

	// Value (numeric)	invoice currency amount
	Value string `json:"value"`

	// ExpectedAmount (numeric)	expected amount to be deposited
	ExpectedAmount string `json:"expected_amount"`

	// ReceivedTotal (numeric)	received amount in balance currency
	ReceivedTotal string `json:"received_total"`

	// ExchangeRate (numeric)	fixated exchange rate of invoice
	ExchangeRate string `json:"exchange_rate"`

	// CleanExchangeRate -
	CleanExchangeRate interface{} `json:"clean_exchange_rate"`

	// IsInternal 	identifies whether payment was made via WhiteBIT platform or not where true = WhiteBIT platform payment and false = any other platform.
	IsInternal        bool        `json:"is_internal"`
	DepositedCurrency interface{} `json:"deposited_currency"`

	// ReceivedCurrency balance currency (ticker)
	ReceivedCurrency string `json:"received_currency"`

	// Status 	order status label
	Status string `json:"status"`

	// ExternalOrderID external order identifier
	ExternalOrderID string `json:"external_order_id"`

	// CreatedAt 	order creation date
	CreatedAt string `json:"created_at"`

	// CompletedAt 	order completion date
	CompletedAt string `json:"completed_at"`

	// AcquiringURL	unique payment form url. Should be used to redirect website customer to the payment form.
	AcquiringURL string `json:"acquiring_url"`

	// SuccessfulLink Link used to redirect user from payment form after order completion. Has to be URL
	SuccessfulLink string `json:"successful_link"`

	// FailureLink Link used to redirect user from payment form after order failure. Has to be URL
	FailureLink string `json:"failure_link"`

	Transactions []*Transaction `json:"transactions"`
}

// Transaction represents a WhitePay transaction.
type Transaction struct {
	// ID 	transaction identifier
	ID string `json:"id"`
	// Currency 	transaction currency (ticker)
	Currency string `json:"currency"`
	// Value 	transaction amount
	Value string `json:"value"`
	// Status 	transaction status label
	Status string `json:"status"`
	// Hash transaction hash
	Hash string `json:"hash"`
	// IsInternal 	identifies whether payment was made via WhiteBIT platform or not where true = WhiteBIT platform payment and false = any other platform.
	IsInternal bool `json:"is_internal"`
	// CreatedAt 	transaction creation date
	CreatedAt string `json:"created_at"`
	// CompletedAt 	transaction completion date
	CompletedAt string `json:"completed_at"`
	// StockOrders 	stock orders
	StockOrders []*StockOrder `json:"stock_orders,omitempty"`
}

// InvoiceCurrencies represents a WhitePay invoice currency.
type InvoiceCurrencies struct {
	// ID 	currency identifier
	ID string `json:"id"`
	// Ticker 	currency ticker
	Ticker string `json:"ticker"`
	// IsFiat a parameter that determines whether the currency is fiat
	IsFiat bool `json:"is_fiat"`
	// Title 	currency title
	Title string `json:"title"`
	// MinAmount 	minimum amount for the currency (numeric)
	MinAmount string `json:"min_amount"`
	// MaxAmount 	maximum amount for the currency (numeric)
	MaxAmount string `json:"max_amount"`
	// Precision (numeric) currency precision - number of characters after the decimal point
	Precision string `json:"precision"`
	// Icon 	currency icon for acquiring form
	Icon string `json:"icon"`
}

// Currency represents a WhitePay currency.
type Currency struct {
	// ID 	currency identifier
	ID string `json:"id"`

	// Ticker 	currency ticker
	Ticker string `json:"ticker"`

	// Title 	currency title
	Title string `json:"title"`

	// Value 	currency value
	Value string `json:"value"`
}

// Status represents a WhitePay status.
type Status struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// Balance represents a WhitePay balance.
type Balance struct {
	ID       string   `json:"id"`
	Value    string   `json:"value"`
	Currency Currency `json:"currency"`
}

// StockOrder represents a WhitePay stock order.
type StockOrder struct {
	Amount       string `json:"amount"`
	Status       string `json:"status"`
	Pair         string `json:"pair"`
	Date         string `json:"date"`
	Time         string `json:"time"`
	ExchangeRate string `json:"exchange_rate"`
	CreatedAt    string `json:"created_at"`
}

type CoinLimit struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

type Network struct {
	MinAmount string      `json:"min_amount"`
	MaxAmount string      `json:"max_amount"`
	Fixed     string      `json:"fixed"`
	Flex      interface{} `json:"flex"`
}

type CryptoCurrency struct {
	ID                   string               `json:"id"`
	Precision            int                  `json:"precision"`
	Icon                 string               `json:"icon"`
	Ticker               string               `json:"ticker"`
	IsFiat               bool                 `json:"is_fiat"`
	Limits               map[string]CoinLimit `json:"limits"`
	Min                  string               `json:"min"`
	Max                  string               `json:"max"`
	IsStable             bool                 `json:"is_stable"`
	Name                 string               `json:"name"`
	Networks             map[string]Network   `json:"networks"`
	ExchangeRate         string               `json:"exchange_rate"`
	PresetDonationValues []interface{}        `json:"preset_donation_values"`
}
