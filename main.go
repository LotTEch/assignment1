package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/LotTEch/assignment1/api"
	"github.com/LotTEch/assignment1/services"
	"github.com/LotTEch/assignment1/utils"
)

func main() {
	// 1. Last .env-variabler (om filen finnes)
	utils.LoadEnv()

	// 2. Hent nødvendige konfigurasjoner
	port := utils.GetEnvString("PORT", "")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}
	restCountriesURL := utils.GetEnvString("RESTCOUNTRIES_URL", "")
	if restCountriesURL == "" {
		log.Fatal("RESTCOUNTRIES_URL environment variable is not set")
	}
	countriesNowURL := utils.GetEnvString("COUNTRIESNOW_URL", "")
	if countriesNowURL == "" {
		log.Fatal("COUNTRIESNOW_URL environment variable is not set")
	}

	// 3. Opprett services, og send inn base-URLer for eksterne API-er
	countryService := services.NewCountryService(restCountriesURL, countriesNowURL)
	populationService := services.NewPopulationService(restCountriesURL, countriesNowURL)
	statusService := services.NewStatusService(time.Now(), restCountriesURL, countriesNowURL)

	// 4. Opprett handlers
	countryHandler := api.NewCountryHandler(countryService)
	populationHandler := api.NewPopulationHandler(populationService)
	statusHandler := api.NewStatusHandler(statusService)

	// 5. Sett opp router
	r := mux.NewRouter()
	r.HandleFunc("/countryinfo/v1/info/{two_letter_country_code}", countryHandler.GetCountryInfo).Methods("GET")
	r.HandleFunc("/countryinfo/v1/population/{two_letter_country_code}", populationHandler.GetPopulationData).Methods("GET")
	r.HandleFunc("/countryinfo/v1/status", statusHandler.GetStatus).Methods("GET")

	// 6. Start serveren
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting server on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
