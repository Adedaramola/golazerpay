package lazerpay

import (
	"fmt"
	"net/http"
)

type PaymentService service

// func (p *PaymentService) Initialize(reference, customer_name, customer_email, coin, currency string, amount uint64, accept_partial_payment bool) (Response, error) {
// 	url := fmt.Sprintf("/transaction/initialize")

// 	payload := &InitializePaymentData{
// 		Reference:            reference,
// 		CustomerName:         customer_name,
// 		CustomerEmail:        customer_email,
// 		Coin:                 coin,
// 		Currency:             currency,
// 		AcceptPartialPayment: accept_partial_payment,
// 	}

// 	res := Response{}

// 	err := p.client.Send(http.MethodPost, url, payload, res)

// 	return res, err
// }

func (p *PaymentService) Verify(address_or_reference string) (Payment, error) {
	url := fmt.Sprintf("/transaction/verify/%s", address_or_reference)
	payment := Payment{}

	err := p.client.Send(http.MethodGet, url, nil, payment)

	return payment, err
}

func (p *PaymentService) GenerateLink(title, description string, amount uint64, link_type, logo, currency, redirect_url string) (PaymentLinkResponse, error) {
	url := fmt.Sprintf("/payment-links")

	payload := &PaymentLinkRequest{
		Title:       title,
		Description: description,
		Amount:      amount,
		Type:        link_type,
		LogoUrl:     logo,
		Currency:    currency,
		RedirectUrl: redirect_url,
	}

	res := PaymentLinkResponse{}

	err := p.client.Send(http.MethodPost, url, payload, &res)

	return res, err
}

func (p *PaymentService) GetPaymentLinks() (PaymentLinksResponse, error) {
	url := fmt.Sprintf("/payment-links")
	res := PaymentLinksResponse{}

	err := p.client.Send(http.MethodGet, url, nil, &res)

	return res, err
}

func (p *PaymentService) GetPaymentLink(id_or_reference string) (PaymentLinkResponse, error) {
	url := fmt.Sprintf("/payment-links/%s", id_or_reference)
	res := PaymentLinkResponse{}

	err := p.client.Send(http.MethodGet, url, nil, &res)

	return res, err
}

// func (p *PaymentService) UpdatePaymentLink(id_or_reference string, status string) (Response, error) {
// 	url := fmt.Sprintf("/payment-links/%s", id_or_reference)
// 	res := Response{}

// 	err := p.client.Send(http.MethodGet, url, status, res)

// 	return res, err
// }
