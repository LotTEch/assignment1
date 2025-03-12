package models

// PopulationResponse er datastrukturen for /population-endepunktets respons
type PopulationResponse struct {
	Mean   int                     `json:"mean"`
	Values []PopulationYearValue   `json:"values"`
}

type PopulationYearValue struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}
