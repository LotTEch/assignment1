package services

import (
	"encoding/json"
	"errors"
	"math"
	"strconv"
	"strings"

	"assignment-1/models"
	"assignment-1/utils"
)

// FetchPopulationData fetches population data for a given country and filters by year range if specified.
// Henter befolkningsdata for et gitt land og filtrerer etter årstall hvis nødvendig.
func FetchPopulationData(code string, limit string) (*models.PopulationResponse, error) {
	url := "http://129.241.150.113:3500/api/v0.1/countries/population?country=" + code
	response, err := utils.FetchAPI(url)
	if err != nil {
		return nil, errors.New("error fetching population data")
	}

	// Decode JSON response
	// Dekoder JSON-svaret
	var data struct {
		Error bool                    `json:"error"`
		Data  []models.PopulationData `json:"data"`
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, errors.New("error decoding JSON")
	}
	if len(data.Data) == 0 {
		return nil, errors.New("no population data found")
	}

	// Filter the correct country based on the provided code
	// Filtrer riktig land basert på den angitte koden
	var selectedCountry *models.PopulationData
	for _, country := range data.Data {
		if country.Code == code {
			selectedCountry = &country
			break
		}
	}
	if selectedCountry == nil {
		return nil, errors.New("did not find population data for this country")
	}

	// Handle limit as a year interval (e.g., "2010-2015")
	// Håndter limit som et datointervall (f.eks. "2010-2015")
	if limit != "" {
		years := strings.Split(limit, "-")
		if len(years) == 2 {
			startYear, err1 := strconv.Atoi(years[0])
			endYear, err2 := strconv.Atoi(years[1])
			if err1 != nil || err2 != nil || startYear > endYear {
				return nil, errors.New("invalid limit value, use the format: ?limit=startYear-endYear")
			}

			// Filter out entries within the specified interval
			// Filtrer ut årstall innenfor intervallet
			var filteredCounts []models.PopulationValue
			for _, entry := range selectedCountry.PopulationCounts {
				if entry.Year >= startYear && entry.Year <= endYear {
					filteredCounts = append(filteredCounts, entry)
				}
			}
			selectedCountry.PopulationCounts = filteredCounts
		} else {
			return nil, errors.New("invalid limit value, use the format: ?limit=startYear-endYear")
		}
	}

	// Calculate the average population value
	// Kalkuler gjennomsnittet av befolkningstallene
	var sum int
	var count int
	for _, entry := range selectedCountry.PopulationCounts {
		sum += entry.Value
		count++
	}

	mean := 0
	if count > 0 {
		mean = int(math.Round(float64(sum) / float64(count)))
	}

	responseData := &models.PopulationResponse{
		Mean:   mean,
		Values: selectedCountry.PopulationCounts,
	}

	return responseData, nil
}
