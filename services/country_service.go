package services

import (
	"assignment-1/models"
	"assignment-1/utils"
	"encoding/json"
	"errors"
	"sort"
)

// FetchCountryInfo fetches country information based on the country code and limits the number of cities if a limit is provided.
// Henter informasjon om et land basert på landkoden og begrenser antall byer med limit.
func FetchCountryInfo(code string, limit int) (*models.Country, error) {
	url := "http://129.241.150.113:8080/v3.1/alpha/" + code
	response, err := utils.FetchAPI(url)
	if err != nil {
		return nil, errors.New("error fetching country information")
	}

	// REST Countries API returns a list of countries
	// REST Countries API returnerer en liste med land
	var countries []map[string]interface{}
	if err := json.Unmarshal(response, &countries); err != nil {
		return nil, errors.New("error decoding JSON")
	}

	if len(countries) == 0 {
		return nil, errors.New("no country information found")
	}

	rawCountry := countries[0]

	// Retrieve name information
	// Hent ut navneinformasjon
	nameData, ok := rawCountry["name"].(map[string]interface{})
	if !ok {
		return nil, errors.New("error fetching name data")
	}
	commonName, ok := nameData["common"].(string)
	if !ok {
		return nil, errors.New("error fetching common name")
	}

	// Build the Country structure using helper functions from utils
	// Bygg Country-strukturen ved hjelp av hjelpefunksjoner fra utils
	countryInfo := models.Country{
		Name:       commonName,
		Continents: utils.ExtractStringArray(rawCountry["continents"].([]interface{})),
		Population: int(rawCountry["population"].(float64)),
		Languages:  utils.ExtractLanguages(rawCountry["languages"].(map[string]interface{})),
		Borders:    utils.ExtractStringArray(rawCountry["borders"].([]interface{})),
		Flag:       rawCountry["flags"].(map[string]interface{})["png"].(string),
		Capital:    utils.ExtractFirstString(rawCountry["capital"].([]interface{})),
	}

	// Fetch cities from the CountriesNow API
	// Hent byer fra CountriesNow API
	cities, err := utils.FetchCountriesNowCities(code)
	if err != nil {
		return nil, errors.New("error fetching city data: " + err.Error())
	}
	// Sort cities alphabetically
	// Sorter byene alfabetisk
	sort.Strings(cities)
	// Limit the number of cities if limit is provided (> 0)
	// Begrens antallet byer dersom limit er satt (> 0)
	if limit > 0 && limit < len(cities) {
		cities = cities[:limit]
	}
	countryInfo.Cities = cities

	return &countryInfo, nil
}
