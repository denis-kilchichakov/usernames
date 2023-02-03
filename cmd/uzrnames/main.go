package main

import (
	"log"
	"os"
)

func main() {
	// create logger
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// log that application is starting
	logger.Println("Starting application")

	// log that application is shutting down
	logger.Println("Shutting down application")
}
