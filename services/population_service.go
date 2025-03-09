package services

import (
    "encoding/json"
    
    "fmt"
    "net/http"
    

    "github.com/LotTEch/assignment1/models"
)

// PopulationService definerer funksjonalitet for å hente befolkningsdata.
type PopulationService interface {
    GetPopulationHistory(countryCode string, startYear, endYear int) (models.PopulationResponse, error)
}

// populationService er en konkret implementasjon av PopulationService.
type populationService struct {
    httpClient        *http.Client
    restCountriesBase string
    countriesNowBase  string
}

// NewPopulationService oppretter en ny instans av populationService.
func NewPopulationService(restCountriesBase, countriesNowBase string) PopulationService {
    return &populationService{
        httpClient:        &http.Client{},
        restCountriesBase: restCountriesBase,
        countriesNowBase:  countriesNowBase,
    }
}

// GetPopulationHistory henter befolkningsdata fra eksterne API-er og returnerer et "models.PopulationResponse".
func (ps *populationService) GetPopulationHistory(countryCode string, startYear, endYear int) (models.PopulationResponse, error) {
    url := fmt.Sprintf("%s/alpha/%s", ps.restCountriesBase, countryCode)

    resp, err := ps.httpClient.Get(url)
    if err != nil {
        return models.PopulationResponse{}, fmt.Errorf("error retrieving population data from RestCountries: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return models.PopulationResponse{}, fmt.Errorf("RestCountries replied with status: %d", resp.StatusCode)
    }

    var restCountriesData struct {
        Population int `json:"population"`
    }

    err = json.NewDecoder(resp.Body).Decode(&restCountriesData)
    if err != nil {
        return models.PopulationResponse{}, fmt.Errorf("error when unmarshaling RestCountries data: %w", err)
    }

    url = fmt.Sprintf("%s/countries/population/historical", ps.countriesNowBase)
    req, err := http.NewRequest(http.MethodPost, url, nil)
    if err != nil {
        return models.PopulationResponse{}, fmt.Errorf("error creating request to CountriesNow: %w", err)
    }

    resp, err = ps.httpClient.Do(req)
    if err != nil {
        return models.PopulationResponse{}, fmt.Errorf("error retrieving population data from CountriesNow: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return models.PopulationResponse{}, fmt.Errorf("CountriesNow responded with status: %d", resp.StatusCode)
    }

    var countriesNowData struct {
        Data []struct {
            Year  int `json:"year"`
            Value int `json:"value"`
        } `json:"data"`
    }

    err = json.NewDecoder(resp.Body).Decode(&countriesNowData)
    if err != nil {
        return models.PopulationResponse{}, fmt.Errorf("error when unmarshaling CountriesNow data: %w", err)
    }

    var filteredData []models.PopulationValue
    for _, d := range countriesNowData.Data {
        if (startYear == 0 || d.Year >= startYear) && (endYear == 0 || d.Year <= endYear) {
            filteredData = append(filteredData, models.PopulationValue{
                Year:  d.Year,
                Value: d.Value,
            })
        }
    }

    sum := 0
    for _, d := range filteredData {
        sum += d.Value
    }

    mean := 0
    if len(filteredData) > 0 {
        mean = sum / len(filteredData)
    }

    return models.PopulationResponse{
        Mean:   mean,
        Values: filteredData,
    }, nil
}