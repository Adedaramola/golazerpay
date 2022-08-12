package lazerpay

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const BaseURL string = "https://api.lazerpay.engineering"

type service struct {
	client *LazerpayClient
}

type LazerpayClient struct {
	common        service
	Client        *http.Client
	BaseURL       *url.URL
	SecretKey     string
	PublicKey     string
	Payment       *PaymentService
	Miscellaneous *MiscellaneousService
}

func Setup(secretKey string, publicKey string) *LazerpayClient {
	url, _ := url.Parse(BaseURL)

	c := &LazerpayClient{
		Client:    http.DefaultClient,
		SecretKey: secretKey,
		PublicKey: publicKey,
		BaseURL:   url,
	}

	c.common.client = c
	c.Payment = (*PaymentService)(&c.common)
	c.Miscellaneous = (*MiscellaneousService)(&c.common)

	return c
}

func (c *LazerpayClient) Send(method string, endpoint string, body, v interface{}) error {
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(body)

		if err != nil {
			return err
		}
	}

	url, _ := c.BaseURL.Parse(endpoint)

	req, err := http.NewRequest(method, url.String(), buf)

	if err != nil {
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Authorization", "Bearer "+c.SecretKey)
	req.Header.Set("x-api-key", c.PublicKey)

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(v)
}
