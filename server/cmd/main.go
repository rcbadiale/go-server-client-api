package main

import (
	"log"
	"net/http"

	"github.com/rcbadiale/go-server-client-api/server/internal/databases"
	"github.com/rcbadiale/go-server-client-api/server/internal/routes"
)

func main() {
	log.Println("starting server...")
	log.Println("preparing database...")
	db, err := databases.Setup()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Println("migrating database...")
	err = databases.Migrate(db)
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/cotacao", routes.ExchangeRouteHandler)
	http.HandleFunc("/exchange/history", routes.ExchangeHistoryHandler)
	log.Println("server started...")
	http.ListenAndServe(":8080", nil)
}
