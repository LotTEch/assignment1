package models

// RestCountriesResponse representerer (forenklet) responsen fra REST Countries v3.1
type RestCountriesResponse struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	Capital    []string          `json:"capital"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Borders    []string          `json:"borders"`
	Flags      struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags"`
	Languages map[string]string `json:"languages"`
}
