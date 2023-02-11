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

	// body, err := network.GetBody("https://api.github.com/users/denis-kilchichakov")
	// if err != nil {
	// 	logger.Fatal(err)
	// }
	// logger.Println(string(body))

	// log that application is shutting down
	logger.Println("Shutting down application")
}

func getBody(s string) {
	panic("unimplemented")
}
