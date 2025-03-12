package utils

import (
	"log"
	"os"
)

// GetEnv henter en environment-variabel eller returnerer en defaultverdi hvis den ikke finnes
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Env variable %s not set, using default '%s'\n", key, fallback)
		return fallback
	}
	return value
}
