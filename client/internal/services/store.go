package services

import (
	"fmt"
	"os"

	"github.com/rcbadiale/go-server-client-api/client/internal/models"
)

func StoreFile(filename string, exchange *models.ExchangeRate) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Fprintf(file, "Dólar: %s", exchange.Bid)
	return nil
}
