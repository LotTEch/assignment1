package services

import (
	
	"time"

	"github.com/LotTEch/assignment1/utils"
)

// StatusInfo representerer JSON-strukturen vi sender tilbake på /status
type StatusInfo struct {
	CountriesNowAPI string  `json:"countriesnowapi"`
	RestCountriesAPI string  `json:"restcountriesapi"`
	Version          string  `json:"version"`
	Uptime           float64 `json:"uptime"`
}

// globalStartTime er en package-variabel som settes fra main eller lignende.
// For enkelthet, kan du hente fra main. Her viser vi bare konseptuelt.
var globalStartTime time.Time

// SetStartTime settes fra main.go (om du ønsker). Ellers kan du bare referere globalt.
func SetStartTime(t time.Time) {
	globalStartTime = t
}

// GetStatusInfo sjekker status for eksterne tjenester og regner ut oppetid.
func GetStatusInfo() StatusInfo {
	// Sjekk CountriesNow
	cnStatus := checkCountriesNow()

	// Sjekk RestCountries
	rcStatus := checkRestCountries()

	// Oppetid
	uptime := time.Since(globalStartTime).Seconds()

	return StatusInfo{
		CountriesNowAPI: cnStatus,
		RestCountriesAPI: rcStatus,
		Version:         "v1",
		Uptime:          uptime,
	}
}

// checkCountriesNow gjør et kjapt kall for å sjekke status
func checkCountriesNow() string {
	testURL := "http://129.241.150.113:3500/api/v0.1/countries/codes" // Kan være et greit testendepunkt
	resp, err := utils.HttpClient.Get(testURL)
	if err != nil {
		return "ERROR"
	}
	defer resp.Body.Close()

	return resp.Status
}

// checkRestCountries gjør et kjapt kall for å sjekke status
func checkRestCountries() string {
	testURL := "http://129.241.150.113:8080/v3.1/all?fields=ccn3" // Minimal data
	resp, err := utils.HttpClient.Get(testURL)
	if err != nil {
		return "ERROR"
	}
	defer resp.Body.Close()

	return resp.Status
}
