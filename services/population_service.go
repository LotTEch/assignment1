package services

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	

	"github.com/LotTEch/assignment1/models"
	"github.com/LotTEch/assignment1/utils"
)

// GetPopulationData henter befolkningshistorikk for et gitt land (2-bokstavskode).
// Kan begrenses til en tidsperiode (startYear, endYear). Hvis startYear/endYear == 0, ingen filter.
func GetPopulationData(twoLetterCode string, startYear, endYear int) (*models.PopulationResponse, error) {
	// 1) Først trenger vi landets *navn* (CountriesNow krever navn, ikke kode).
	countryName, err := getCountryNameFromCode(twoLetterCode)
	if err != nil {
		return nil, err
	}

	// 2) Kall CountriesNow for å få befolkningsdata
	popCounts, err := fetchPopulationFromCountriesNow(countryName)
	if err != nil {
		return nil, err
	}

	// 3) Filterer på startYear/endYear hvis spesifisert
	filteredCounts := make([]models.PopulationYearValue, 0)
	for _, pc := range popCounts {
		yearInt := pc.Year
        valueInt := pc.Value

		// Dersom startYear og endYear er 0, tar vi alle
		// Hvis ikke, sjekk om yearInt er innenfor [startYear, endYear]
		if (startYear == 0 && endYear == 0) || (yearInt >= startYear && yearInt <= endYear) {
			filteredCounts = append(filteredCounts, models.PopulationYearValue{
				Year:  yearInt,
				Value: valueInt,
			})
		}
	}

	// 4) Regn ut gjennomsnittet (avrundet)
	var sum, count int
	for _, fc := range filteredCounts {
		sum += fc.Value
		count++
	}
	var mean int
	if count > 0 {
		mean = sum / count
	}

	popResp := &models.PopulationResponse{
		Mean:   mean,
		Values: filteredCounts,
	}

	return popResp, nil
}

// getCountryNameFromCode bruker RestCountries for å mappe en ISO2-kode til landets navn
func getCountryNameFromCode(twoLetterCode string) (string, error) {
	restCountriesURL := "http://129.241.150.113:8080/v3.1/alpha/" + twoLetterCode
	resp, err := utils.HttpClient.Get(restCountriesURL)
	if err != nil {
		log.Println("Error calling RestCountries API:", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("RestCountries API returned non-200:", resp.StatusCode)
		return "", errors.New("invalid country code or RestCountries error")
	}

	var restCountriesData []models.RestCountriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&restCountriesData); err != nil {
		log.Println("Error decoding RestCountries response:", err)
		return "", err
	}
	if len(restCountriesData) == 0 {
		return "", errors.New("no data returned from RestCountries for code " + twoLetterCode)
	}

	return restCountriesData[0].Name.Common, nil
}

// fetchPopulationFromCountriesNow kaller /v0.1/countries/population med landets navn.
func fetchPopulationFromCountriesNow(countryName string) ([]models.CountriesNowPopulationCount, error) {
	url := "http://129.241.150.113:3500/api/v0.1/countries/population"
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
		return nil, errors.New("could not fetch population from CountriesNow")
	}

	var cnResp models.CountriesNowPopulationResponse
	if err := json.NewDecoder(resp.Body).Decode(&cnResp); err != nil {
		return nil, err
	}

	if cnResp.Error {
		return nil, errors.New("CountriesNow API error: " + cnResp.Msg)
	}

	return cnResp.Data.PopulationCounts, nil
}
