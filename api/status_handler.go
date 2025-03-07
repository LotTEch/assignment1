package api

import (
	"encoding/json"
	"net/http"
	"time"

	"assignment-1/utils"
)

var startTime = time.Now()

func GetAPIStatus(w http.ResponseWriter, r *http.Request) {
	countriesStatus := utils.CheckAPIHealth("http://129.241.150.113:8080/v3.1/")
	populationStatus := utils.CheckAPIHealth("http://129.241.150.113:3500/api/v0.1/")

	status := map[string]interface{}{
		"countriesnowapi":  countriesStatus,
		"restcountriesapi": populationStatus,
		"version":          "v1",
		"uptime":           int(time.Since(startTime).Seconds()),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
