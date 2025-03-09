package api

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"

    "github.com/LotTEch/assignment1/services"
)

// CountryHandler håndterer /countryinfo/v1/info/{:two_letter_country_code}
type CountryHandler struct {
    CountryService services.CountryService
}

func NewCountryHandler(cs services.CountryService) *CountryHandler {
    return &CountryHandler{
        CountryService: cs,
    }
}

// GetCountryInfo håndterer GET /countryinfo/v1/info/{:two_letter_country_code}?limit=10
func (h *CountryHandler) GetCountryInfo(w http.ResponseWriter, r *http.Request) {
    // Hent landkoden
    vars := mux.Vars(r)
    countryCode := vars["two_letter_country_code"]
    if len(countryCode) == 0 {
        http.Error(w, "Country code missing path", http.StatusBadRequest)
        return
    }

    // Hent ev. limit-parameter
    limitQuery := r.URL.Query().Get("limit")
    var limit int
    var err error

    if limitQuery != "" {
        limit, err = strconv.Atoi(limitQuery)
        if err != nil {
            http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
            return
        }
    }

    // Kall service-laget
    countryInfo, err := h.CountryService.GetCountryInfo(countryCode, limit)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Returner JSON i henhold til spesifikasjon:
    // {
    //   "name": "Norway",
    //   "continents": ["Europe"],
    //   "population": 4700000,
    //   "languages": {"nno":"Norwegian Nynorsk","nob":"Norwegian Bokmål","smi":"Sami"},
    //   "borders": ["FIN","SWE","RUS"],
    //   "flag": "https://flagcdn.com/w320/no.png",
    //   "capital": "Oslo",
    //   "cities": ["Abelvaer","Adalsbruk",...]
    // }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(countryInfo)
}