package lazerpay

import (
	"fmt"
	"net/http"
	"time"
)

type MiscellaneousService service

type CoinsResponse struct {
	Message    string `json:"message"`
	Data       []Coin `json:"data"`
	StatusCode uint   `json:"statusCode"`
	Status     string `json:"status"`
}

type Coin struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Symbol     string     `json:"symbol"`
	Logo       string     `json:"logo"`
	Address    string     `json:"address"`
	Network    string     `json:"network"`
	Blockchain string     `json:"blockchain"`
	Status     string     `json:"status"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

func (m *MiscellaneousService) GetCoins() (CoinsResponse, error) {
	url := fmt.Sprintf("/api/v1/coins")
	coins := CoinsResponse{}

	err := m.client.Send(http.MethodGet, url, nil, &coins)

	return coins, err
}
