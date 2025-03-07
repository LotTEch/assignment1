package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"assignment-1/services"
)

// GetCountryInfo henter informasjon om et land basert på koden
func GetCountryInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	country, err := services.FetchCountryInfo(code)
	if err != nil {
		http.Error(w, "Feil ved henting av landinformasjon", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(country)
}
