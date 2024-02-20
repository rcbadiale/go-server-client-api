package services

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/rcbadiale/go-server-client-api/server/internal/models"
)

/* Solicita a cotação para um serviço de terceiro */
func GetExchangeRate(ctx context.Context) (*models.ExchangeRate, error) {
	// Provavelmente deveria ser configurável 🤷‍♂️
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var conversion models.Conversion
	err = json.Unmarshal(body, &conversion)
	if err != nil {
		return nil, err
	}
	conversion.ExchangeRate.Id = uuid.New().String()
	return &conversion.ExchangeRate, nil
}
