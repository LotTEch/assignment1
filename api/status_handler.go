package api

import (
	"encoding/json"
	"net/http"
	

	"github.com/gorilla/mux"

	"github.com/LotTEch/assignment1/services"
	"github.com/LotTEch/assignment1/utils"
)

// Du trenger en referanse til startTime. For enkelhets skyld kan du enten
// - Hente den fra main via en global variabel, eller
// - Legge i en package-variabel.
// Her bruker vi en service-funksjon for å hente oppetid.

func RegisterStatusRoutes(r *mux.Router) {
	r.HandleFunc("/countryinfo/v1/status/", statusHandler).Methods(http.MethodGet)
}

// statusHandler håndterer kall til /countryinfo/v1/status/
func statusHandler(w http.ResponseWriter, r *http.Request) {
	// Hent statusinformasjon fra service
	statusInfo := services.GetStatusInfo()

	// Konverterer struct til JSON
	respBytes, err := json.Marshal(statusInfo)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, []byte(`{"error":"failed to marshal status info"}`))
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, respBytes)
}
