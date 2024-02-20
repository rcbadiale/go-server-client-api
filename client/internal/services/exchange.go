package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rcbadiale/go-server-client-api/client/internal/models"
)

/*
	Solicita para o server a cotação

Caso a resposta seja erro (status code >= 400) gera um erro personalizado
indicando o status code e a mensagem de erro obtida se ela existir.
*/
func GetExchangeRate(ctx context.Context, url string) (*models.ExchangeRate, error) {
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
	var rate models.ExchangeRate
	err = json.Unmarshal(body, &rate)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		err = fmt.Errorf("%d: %s", res.StatusCode, rate.Error)
		return nil, err
	}
	return &rate, nil
}
