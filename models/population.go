package models

// PopulationValue representerer befolkningsdata for et gitt år.
type PopulationValue struct {
    Year  int `json:"year"`
    Value int `json:"value"`
}

// PopulationResponse representerer responsen for befolkningsdata.
type PopulationResponse struct {
    Mean   int              `json:"mean"`
    Values []PopulationValue `json:"values"`
}