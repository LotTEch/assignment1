package models

// PopulationData represents the raw data fetched from external APIs.
// PopulationData representerer den rå dataen som hentes fra eksterne API-er.
type PopulationData struct {
	Country          string            `json:"country"`
	Code             string            `json:"code"`
	Iso3             string            `json:"iso3"`
	PopulationCounts []PopulationValue `json:"populationCounts"`
}

// PopulationValue represents the population count for a single year.
// PopulationValue representerer befolkningstall for et enkelt år.
type PopulationValue struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

// PopulationResponse is the format we want to return to the client.
// PopulationResponse er formatet vi ønsker å returnere til klienten.
type PopulationResponse struct {
	Mean   int               `json:"mean"`
	Values []PopulationValue `json:"values"`
}
