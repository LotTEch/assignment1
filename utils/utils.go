package utils

import (
	"log"
	"time"
)

// UptimeTracker kan brukes til å spore hvor lenge applikasjonen har kjørt.
type UptimeTracker struct {
	StartTime time.Time
}

// NewUptimeTracker initialiserer en ny UptimeTracker.
func NewUptimeTracker() *UptimeTracker {
	return &UptimeTracker{
		StartTime: time.Now(),
	}
}

// GetUptimeSeconds returnerer hvor mange sekunder som har gått siden StartTime.
func (u *UptimeTracker) GetUptimeSeconds() int64 {
	return int64(time.Since(u.StartTime).Seconds())
}

// LogInfo er et eksempel på en hjelpefunksjon for logging.
func LogInfo(message string) {
	log.Printf("[INFO] %s\n", message)
}

// LogError er et eksempel på en hjelpefunksjon for feillogging.
func LogError(err error) {
	log.Printf("[ERROR] %v\n", err)
}
