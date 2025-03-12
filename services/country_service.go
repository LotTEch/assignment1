package services

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sort"
	

	"github.com/LotTEch/assignment1/models"
	"github.com/LotTEch/assignment1/utils"
)

// GetCountryInfo henter generell informasjon om et land (2-bokstavskode)
// samt en liste over byer fra CountriesNow (valgfritt limit).
func GetCountryInfo(twoLetterCode string, limit int) (*models.CountryInfo, error) {
	// 1) Kall RestCountries for å hente detaljert info
	restCountriesURL := "http://129.241.150.113:8080/v3.1/alpha/" + twoLetterCode
	resp, err := utils.HttpClient.Get(restCountriesURL)
	if err != nil {
		log.Println("Error calling RestCountries API:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("RestCountries API returned non-200:", resp.StatusCode)
		return nil, errors.New("invalid country code or RestCountries error")
	}

	var restCountriesData []models.RestCountriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&restCountriesData); err != nil {
		log.Println("Error decoding RestCountries response:", err)
		return nil, err
	}
	if len(restCountriesData) == 0 {
		return nil, errors.New("no data returned from RestCountries for code " + twoLetterCode)
	}

	// For enkelhets skyld tar vi bare første element i arrayet
	rc := restCountriesData[0]

	// 2) Finn landets offisielle/vanlige navn
	countryName := rc.Name.Common

	// 3) Hent liste over byer fra CountriesNow
	cities, err := fetchCitiesFromCountriesNow(countryName)
	if err != nil {
		log.Println("Error fetching cities from CountriesNow:", err)
		return nil, err
	}

	// Sorter byer alfabetisk
	sort.Strings(cities)

	// Hvis limit > 0, avkort listen
	if limit > 0 && limit < len(cities) {
		cities = cities[:limit]
	}

	// 4) Bygg opp CountryInfo-struct
	countryInfo := &models.CountryInfo{
		Name:       rc.Name.Common,
		Continents: rc.Continents,
		Population: rc.Population,
		Languages:  rc.Languages,
		Borders:    rc.Borders,
		Flag:       rc.Flags.Png,
		Capital:    "",
		Cities:     cities,
	}

	// capital i RestCountries er en slice, vi tar første om den finnes
	if len(rc.Capital) > 0 {
		countryInfo.Capital = rc.Capital[0]
	}

	return countryInfo, nil
}

// fetchCitiesFromCountriesNow kaller CountriesNow API for å hente alle byer for et gitt land (navn).
func fetchCitiesFromCountriesNow(countryName string) ([]string, error) {
	url := "http://129.241.150.113:3500/api/v0.1/countries/cities"
	// Body må sendes som JSON
	reqBody := map[string]string{
		"country": countryName,
	}

	resp, err := utils.DoPostJSON(url, reqBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("CountriesNow returned status %d for country %s\n", resp.StatusCode, countryName)
		return nil, errors.New("could not fetch cities from CountriesNow")
	}

	var cnResp models.CountriesNowCitiesResponse
	if err := json.NewDecoder(resp.Body).Decode(&cnResp); err != nil {
		return nil, err
	}

	if cnResp.Error {
		return nil, errors.New("CountriesNow API error: " + cnResp.Msg)
	}

	// Returner by-listen
	return cnResp.Data, nil
}
