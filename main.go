package main

import (
	"log"
	"net/http"
	"os"

	"assignment-1/api"
	"assignment-1/utils"

	"github.com/gorilla/mux"
)

func main() {
	// Last inn miljøvariabler fra .env-filen
	utils.LoadConfig()

	// Hent port fra miljøvariabler
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	r := mux.NewRouter()

	// Definerer API-endepunkter
	r.HandleFunc("/countryinfo/v1/info/{code}", api.GetCountryInfo).Methods("GET")
	r.HandleFunc("/countryinfo/v1/population/{code}", api.GetPopulationData).Methods("GET")
	r.HandleFunc("/countryinfo/v1/status", api.GetAPIStatus).Methods("GET")

	log.Printf("Server kjører på port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
