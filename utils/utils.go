package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

// ParseYear forsøker å parse en streng til et heltall (år). Returnerer feil hvis mislykket.
func ParseYear(yearStr string) (int, error) {
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return 0, fmt.Errorf("invalid year: %s", yearStr)
	}
	return year, nil
}

// WriteJSONResponse hjelper med å sette riktig Content-Type og skrive JSON-respons
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(data)
}
