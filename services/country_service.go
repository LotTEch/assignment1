package services

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
    "sort"
    "strings"

    "github.com/LotTEch/assignment1/models"
)

// CountryService definerer funksjonalitet for å hente landinformasjon.
type CountryService interface {
    GetCountryInfo(countryCode string, limit int) (models.Country, error)
}

// countryService er en konkret implementasjon av CountryService.
type countryService struct {
    httpClient         *http.Client
    restCountriesBase  string
    countriesNowBase   string
}

// NewCountryService oppretter en ny instans av countryService.
func NewCountryService(restCountriesBase, countriesNowBase string) CountryService {
    return &countryService{
        httpClient:        &http.Client{},
        restCountriesBase: restCountriesBase,
        countriesNowBase:  countriesNowBase,
    }
}

// GetCountryInfo henter data fra eksterne API-er og returnerer et "models.Country".
func (cs *countryService) GetCountryInfo(countryCode string, limit int) (models.Country, error) {
    // 1. Bygg URL dynamisk fra restCountriesBase
    restCountriesURL := fmt.Sprintf("%s/alpha/%s", cs.restCountriesBase, countryCode)

    resp, err := cs.httpClient.Get(restCountriesURL)
    if err != nil {
        return models.Country{}, fmt.Errorf("feil ved henting av landinfo fra RestCountries: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return models.Country{}, fmt.Errorf("RestCountries svarte med status: %d", resp.StatusCode)
    }

    // RestCountries for v3.1/alpha returnerer ofte et array av land. Vi definerer en hjelpestruktur:
    var restCountriesData []struct {
        Name struct {
            Common string `json:"common"`
        } `json:"name"`
        Continents []string          `json:"continents"`
        Population int               `json:"population"`
        Languages  map[string]string `json:"languages"`
        Borders    []string          `json:"borders"`
        Capital    []string          `json:"capital"`
        Flags      struct {
            Png string `json:"png"`
            Svg string `json:"svg"`
        } `json:"flags"`
    }

    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return models.Country{}, fmt.Errorf("feil ved lesing av respons: %w", err)
    }

    err = json.Unmarshal(bodyBytes, &restCountriesData)
    if err != nil {
        return models.Country{}, fmt.Errorf("feil ved unmarshal av RestCountries-data: %w", err)
    }

    if len(restCountriesData) == 0 {
        return models.Country{}, errors.New("fant ingen data for gitt landkode i RestCountries")
    }

    rc := restCountriesData[0] // For enkelhets skyld tar vi bare første element
    countryName := rc.Name.Common

    // 2. Kall CountriesNow API for å hente liste over byer.
    cities, err := cs.fetchCitiesFromCountriesNow(countryName)
    if err != nil {
        return models.Country{}, fmt.Errorf("feil ved henting av byer: %w", err)
    }

    // Sorter byene alfabetisk
    sort.Strings(cities)

    // Begrens antall byer hvis limit > 0
    if limit > 0 && limit < len(cities) {
        cities = cities[:limit]
    }

    // 3. Sett sammen en "models.Country"
    result := models.Country{
        Name:       rc.Name.Common,
        Continents: rc.Continents,
        Population: rc.Population,
        Languages:  rc.Languages,
        Borders:    rc.Borders,
        Flag:       rc.Flags.Png, // du kan velge PNG eller SVG
        Capital:    "",
        Cities:     cities,
    }

    // RestCountries gir ofte capital som en liste. Vi tar første hvis den finnes
    if len(rc.Capital) > 0 {
        result.Capital = rc.Capital[0]
    }

    return result, nil
}

// fetchCitiesFromCountriesNow henter byliste fra CountriesNow for gitt landnavn.
func (cs *countryService) fetchCitiesFromCountriesNow(countryName string) ([]string, error) {
    // Bygg URL dynamisk fra countriesNowBase
    url := cs.countriesNowBase + "/countries/cities"

    payload := struct {
        Country string `json:"country"`
    }{
        Country: countryName,
    }

    jsonPayload, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPayload))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := cs.httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("CountriesNow responded with status: %d", resp.StatusCode)
    }

    var cnResponse struct {
        Error bool     `json:"error"`
        Msg   string   `json:"msg"`
        Data  []string `json:"data"` // antar at byene returneres som en liste av strenger
    }
    err = json.NewDecoder(resp.Body).Decode(&cnResponse)
    if err != nil {
        return nil, err
    }

    if cnResponse.Error {
        return nil, fmt.Errorf("CountriesNow error message: %s", cnResponse.Msg)
    }

    var cities []string
    for _, c := range cnResponse.Data {
        cTrim := strings.TrimSpace(c)
        if cTrim != "" {
            cities = append(cities, cTrim)
        }
    }

    return cities, nil
}