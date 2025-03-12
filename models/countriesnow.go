package models

// CountriesNowCitiesResponse for "POST /countries/cities"
type CountriesNowCitiesResponse struct {
	Error bool     `json:"error"`
	Msg   string   `json:"msg"`
	Data  []string `json:"data"`
}

// CountriesNowPopulationResponse for "POST /countries/population"
type CountriesNowPopulationResponse struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  struct {
		Country          string                       `json:"country"`
		Code             string                       `json:"code"`
		PopulationCounts []CountriesNowPopulationCount `json:"populationCounts"`
	} `json:"data"`
}

type CountriesNowPopulationCount struct {
    Year  int `json:"year"`
    Value int `json:"value"`
}
