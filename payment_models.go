package lazerpay

import "time"

type Payment struct {
}

type InitializePaymentData struct {
	Reference            string `json:"reference"`
	CustomerName         string `json:"customer_name"`
	CustomerEmail        string `json:"customer_email"`
	Coin                 string `json:"coin"`
	Currency             string `json:"currency"`
	Amount               uint64 `json:"amount"`
	AcceptPartialPayment bool   `json:"accept_partial_payment"`
}

type InitializePaymentResponse struct {
	Message    string  `json:"message"`
	Status     string  `json:"status"`
	Data       Payment `json:"data"`
	StatusCode uint    `json:"statusCode"`
}

type PaymentLinkRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      uint64 `json:"amount,omitempty"`
	Type        string `json:"type,omitempty"`
	LogoUrl     string `json:"logo,omitempty"`
	Currency    string `json:"currency,omitempty"`
	RedirectUrl string `json:"redirect_url,omitempty"`
}

type PaymentLinksResponse struct {
	Message    string        `json:"message"`
	Data       []PaymentLink `json:"data"`
	StatusCode string        `json:"statusCode"`
	Status     string        `json:"status"`
}

type PaymentLinkResponse struct {
	Message    string      `json:"message"`
	Data       PaymentLink `json:"data"`
	StatusCode string      `json:"statusCode"`
	Status     string      `json:"status"`
}

type PaymentLink struct {
	ID          string     `json:"id"`
	Reference   string     `json:"reference"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Amount      string     `json:"amount"`
	Currency    string     `json:"currency"`
	RedirectUrl string     `json:"redirect_url"`
	Logo        string     `json:"logo"`
	Type        string     `json:"type"`
	Network     string     `json:"network"`
	Status      string     `json:"status"`
	PaymentUrl  string     `json:"payment_url"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
