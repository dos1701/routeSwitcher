package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Only log the Debug severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("Starting Route Switcher Service")
}
