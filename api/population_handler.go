package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/LotTEch/assignment1/services"
	"github.com/LotTEch/assignment1/utils"
)

// RegisterPopulationRoutes registrerer ruter relatert til befolknings-endepunktet.
func RegisterPopulationRoutes(r *mux.Router) {
	r.HandleFunc("/countryinfo/v1/population/", populationRootHandler).Methods(http.MethodGet)
	r.HandleFunc("/countryinfo/v1/population/{twoLetterCode}", getPopulationHandler).Methods(http.MethodGet)
}

// populationRootHandler håndterer kall til /countryinfo/v1/population/ uten ekstra parametere
func populationRootHandler(w http.ResponseWriter, r *http.Request) {
	// Viser en enkel hjelpetekst
	message := `{"message": "Use /countryinfo/v1/population/{ISO2Code}?limit=YYYY-YYYY to get population data. Example: /countryinfo/v1/population/no?limit=2010-2015"}`
	utils.WriteJSONResponse(w, http.StatusOK, []byte(message))
}

// getPopulationHandler håndterer kall til /countryinfo/v1/population/{:two_letter_country_code}{?limit={:startYear-endYear}}
func getPopulationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["twoLetterCode"]

	// Sjekk om vi har limit-range i query
	rangeParam := r.URL.Query().Get("limit")
	var startYear, endYear int
	var err error
	if rangeParam != "" {
		parts := strings.Split(rangeParam, "-")
		if len(parts) == 2 {
			startYear, err = utils.ParseYear(parts[0])
			if err != nil {
				log.Println("Invalid startYear")
				utils.WriteJSONResponse(w, http.StatusBadRequest, []byte(`{"error":"Invalid start year"}`))
				return
			}
			endYear, err = utils.ParseYear(parts[1])
			if err != nil {
				log.Println("Invalid endYear")
				utils.WriteJSONResponse(w, http.StatusBadRequest, []byte(`{"error":"Invalid end year"}`))
				return
			}
		} else {
			log.Println("Invalid limit param format. Must be startYear-endYear.")
			utils.WriteJSONResponse(w, http.StatusBadRequest, []byte(`{"error":"limit param must be in the format startYear-endYear"}`))
			return
		}
	}

	// Kall service-laget for å hente populasjonsinfo
	populationData, serviceErr := services.GetPopulationData(code, startYear, endYear)
	if serviceErr != nil {
		log.Printf("Error fetching population data: %v\n", serviceErr)
		utils.WriteJSONResponse(w, http.StatusBadRequest, []byte(`{"error": "`+serviceErr.Error()+`"}`))
		return
	}

	// Konverterer struct til JSON
	respBytes, err := json.Marshal(populationData)
	if err != nil {
		log.Printf("Error marshalling population data: %v\n", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, []byte(`{"error": "failed to marshal population data"}`))
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, respBytes)
}
