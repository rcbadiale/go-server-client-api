package routes

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/rcbadiale/go-server-client-api/server/internal/databases"
	"github.com/rcbadiale/go-server-client-api/server/internal/models"
	"github.com/rcbadiale/go-server-client-api/server/internal/services"
)

/*
	Define o comportamento da rota de cotação

Em casos de erro será retornado um json no formato `{"error":
"{mensagem de erro}"}` com status code 500.
*/
func ExchangeRouteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	log.Println("request started")
	defer log.Println("request ended")

	rate, err := services.GetExchangeRate(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "context deadline exceeded") {
			err = errors.New("request timed out by server")
		}
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(models.ErrorResponseJson(err))
		return
	}

	data, err := json.Marshal(rate)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(models.ErrorResponseJson(err))
		return
	}

	err = databases.InsertRate(rate)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(models.ErrorResponseJson(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

/*
	Define o comportamento da rota de histórico de cotações

Para essa rota ter um funcionamento mais correto o DB deveria ter uma coluna
de data de inserção, mas por se tratar de uma rota apenas para visualização
de que os dados foram salvos essa alteração não será implementada.
*/
func ExchangeHistoryHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("request started")
	defer log.Println("request ended")

	rates, err := databases.RateHistory()
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(models.ErrorResponseJson(err))
		return
	}

	data, err := json.Marshal(rates)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(models.ErrorResponseJson(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
