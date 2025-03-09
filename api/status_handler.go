package api

import (
    "encoding/json"
    "net/http"
    

    "github.com/LotTEch/assignment1/services"
)

// StatusHandler håndterer /countryinfo/v1/status
type StatusHandler struct {
    StatusService services.StatusService
}

func NewStatusHandler(ss services.StatusService) *StatusHandler {
    return &StatusHandler{
        StatusService: ss,
    }
}

// GetStatus håndterer GET /countryinfo/v1/status
func (h *StatusHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
    status, err := h.StatusService.GetStatus()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Returner JSON i henhold til spesifikasjon:
    // {
    //   "countriesnowapi": "<http status code for CountriesNow API>",
    //   "restcountriesapi": "<http status code for RestCountries API>",
    //   "version": "v1",
    //   "uptime": <time in seconds since the last re/start of your service>
    // }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(status)
}