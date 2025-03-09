package utils

import (
    "errors"
    "io/ioutil"
    "net/http"
    "time"
)

var client = &http.Client{
    Timeout: 10 * time.Second,
}

// FetchAPI henter data fra en gitt URL.
func FetchAPI(url string) ([]byte, error) {
    resp, err := client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("failed to fetch data: " + resp.Status)
    }

    return ioutil.ReadAll(resp.Body)
}

// CheckAPIHealth sjekker helsestatusen til en gitt URL.
func CheckAPIHealth(url string) int {
    resp, err := client.Get(url)
    if err != nil {
        return 500
    }
    defer resp.Body.Close()
    return resp.StatusCode
}