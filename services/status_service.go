package services

import (
    "time"
	"net/http"
    "github.com/LotTEch/assignment1/utils"
)

// StatusService definerer funksjonalitet for å hente statusdata.
type StatusService interface {
    GetStatus() (map[string]interface{}, error)
}

// statusService er en konkret implementasjon av StatusService.
type statusService struct {
    httpClient        *http.Client
    startTime         time.Time
    restCountriesBase string
    countriesNowBase  string
}

// NewStatusService oppretter en ny instans av statusService.
func NewStatusService(startTime time.Time, restCountriesBase, countriesNowBase string) StatusService {
    return &statusService{
        httpClient:        &http.Client{},
        startTime:         startTime,
        restCountriesBase: restCountriesBase,
        countriesNowBase:  countriesNowBase,
    }
}

// GetStatus henter statusdata fra eksterne API-er og returnerer et statuskart.
func (ss *statusService) GetStatus() (map[string]interface{}, error) {
    restCountriesStatus := utils.CheckAPIHealth(ss.restCountriesBase)
    countriesNowStatus := utils.CheckAPIHealth(ss.countriesNowBase)

    status := map[string]interface{}{
        "restcountriesapi": restCountriesStatus,
        "countriesnowapi":  countriesNowStatus,
        "version":          "v1",
        "uptime":           int(time.Since(ss.startTime).Seconds()),
    }

    return status, nil
}