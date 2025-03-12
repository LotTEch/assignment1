package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/LotTEch/assignment1/api"      // Oppdater med riktig modulbane
	"github.com/LotTEch/assignment1/services" // Importerer status-service for å sette startTime
)

// Variabel for å spore oppstartstidspunkt (for /status-endepunktet)
var startTime time.Time

func main() {
	// Laster environment-variabler fra .env (hvis finnes)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error reading .env file.")
	}

	// Setter starttid
	startTime = time.Now()
	// Setter oppstartstid i status-servicen slik at den kan beregne oppetid
	services.SetStartTime(startTime)

	// Henter port fra .env, eller bruker 8080 som default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Oppretter router (gorilla/mux)
	r := mux.NewRouter()

	// Registrer endepunkter fra våre API-handlere
	api.RegisterCountryInfoRoutes(r)
	api.RegisterPopulationRoutes(r)
	api.RegisterStatusRoutes(r)

	// Kjør serveren
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
