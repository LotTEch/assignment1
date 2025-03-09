package utils

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

// LoadEnv laster miljøvariabler fra .env-filen.
func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }
}

// GetEnvString henter en strengverdi fra miljøvariabler eller returnerer en default-verdi.
func GetEnvString(key, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        return defaultValue
    }
    return value
}