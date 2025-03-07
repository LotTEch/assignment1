package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadConfig laster miljøvariabler fra .env-filen.
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Ingen .env-fil funnet, bruker standard miljøvariabler")
	}
}
