package api

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "github.com/gorilla/mux"

    "github.com/LotTEch/assignment1/services"
)

// PopulationHandler håndterer /countryinfo/v1/population/{:two_letter_country_code}
type PopulationHandler struct {
    PopulationService services.PopulationService
}

func NewPopulationHandler(ps services.PopulationService) *PopulationHandler {
    return &PopulationHandler{
        PopulationService: ps,
    }
}

// GetPopulationData håndterer GET /countryinfo/v1/population/{:two_letter_country_code}?limit=YYYY-YYYY
func (h *PopulationHandler) GetPopulationData(w http.ResponseWriter, r *http.Request) {
    // Hent landkoden
    vars := mux.Vars(r)
    countryCode := vars["two_letter_country_code"]
    if len(countryCode) == 0 {
        http.Error(w, "Country code missing path", http.StatusBadRequest)
        return
    }

    // Hent ev. limit-parameter for tidsrom (startYear-endYear)
    limitQuery := r.URL.Query().Get("limit")
    var startYear, endYear int
    var err error

    if limitQuery != "" {
        parts := strings.Split(limitQuery, "-")
        if len(parts) == 2 {
            startYear, err = strconv.Atoi(parts[0])
            if err != nil {
                http.Error(w, "Invalid start year in limit parameter", http.StatusBadRequest)
                return
            }
            endYear, err = strconv.Atoi(parts[1])
            if err != nil {
                http.Error(w, "Invalid end year in limit parameter", http.StatusBadRequest)
                return
            }
        } else {
            http.Error(w, "Invalid format of limit parameter. Expected YYYY-YYYY", http.StatusBadRequest)
            return
        }
    }

    // Kall service-laget
    popData, err := h.PopulationService.GetPopulationHistory(countryCode, startYear, endYear)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Returner JSON i henhold til spesifikasjon:
    // {
    //   "mean": <int>,
    //   "values": [
    //       {"year": 2010, "value": 4889252},
    //       ...
    //   ]
    // }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(popData)
}