package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/LotTEch/assignment1/services"
	"github.com/LotTEch/assignment1/utils"
)

// RegisterCountryInfoRoutes registrerer ruter relatert til land-info-endepunktet.
func RegisterCountryInfoRoutes(r *mux.Router) {
	r.HandleFunc("/countryinfo/v1/info/", infoRootHandler).Methods(http.MethodGet)
	r.HandleFunc("/countryinfo/v1/info/{twoLetterCode}", getCountryInfoHandler).Methods(http.MethodGet)
}

// infoRootHandler håndterer kall til /countryinfo/v1/info/ uten ekstra parametere
func infoRootHandler(w http.ResponseWriter, r *http.Request) {
	// Viser en enkel hjelpetekst
	message := `{"message": "Use /countryinfo/v1/info/{ISO2Code}?limit=10 to get country info. Example: /countryinfo/v1/info/no?limit=5"}`
	utils.WriteJSONResponse(w, http.StatusOK, []byte(message))
}

// getCountryInfoHandler håndterer kall til /countryinfo/v1/info/{:two_letter_country_code}{?limit=10}
func getCountryInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["twoLetterCode"]

	// Henter optional limit-query
	limitStr := r.URL.Query().Get("limit")
	var limit int
	var err error
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Println("Invalid limit query parameter")
			utils.WriteJSONResponse(w, http.StatusBadRequest, []byte(`{"error": "limit parameter must be a number"}`))
			return
		}
	}

	// Kall service-laget for å hente info
	countryInfo, serviceErr := services.GetCountryInfo(code, limit)
	if serviceErr != nil {
		log.Printf("Error fetching country info: %v\n", serviceErr)
		utils.WriteJSONResponse(w, http.StatusBadRequest, []byte(`{"error": "`+serviceErr.Error()+`"}`))
		return
	}

	// Konverterer struct til JSON
	respBytes, err := json.Marshal(countryInfo)
	if err != nil {
		log.Printf("Error marshalling country info: %v\n", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, []byte(`{"error": "failed to marshal country info"}`))
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, respBytes)
}
