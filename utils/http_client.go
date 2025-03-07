package utils

import (
	"io/ioutil"
	"net/http"
)

// FetchAPI henter data fra en gitt URL og returnerer svaret som bytes.
func FetchAPI(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// CheckAPIHealth sjekker om en ekstern API er tilgjengelig.
func CheckAPIHealth(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		return 500 // Returnerer 500 hvis API-et ikke er tilgjengelig.
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
