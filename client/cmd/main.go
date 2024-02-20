package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/rcbadiale/go-server-client-api/client/internal/services"
)

func main() {
	log.Println("running client...")
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	url := "http://localhost:8080/cotacao"
	exchange_rate, err := services.GetExchangeRate(ctx, url)
	if err != nil {
		log.Fatalln(err)
	}
	data, err := json.Marshal(exchange_rate)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(data))
	err = services.StoreFile("cotacao.txt", exchange_rate)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("client done!")
}
