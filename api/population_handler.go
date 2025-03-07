package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"assignment-1/services"
)

// GetPopulationData håndterer API-forespørselen for å hente befolkningsdata.
func GetPopulationData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	// Hent "limit" parameteren fra URL-en
	limit := r.URL.Query().Get("limit")

	// Send kode og limit videre til service-laget
	population, err := services.FetchPopulationData(code, limit)
	if err != nil {
		http.Error(w, "Feil ved henting av befolkningsdata: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(population)
}
