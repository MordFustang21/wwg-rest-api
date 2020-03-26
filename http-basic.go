package main

import (
	"log"
	"net/http"
)

func main() {
	// calling our run function and handling the error
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// run is used so we can just return errors and handle a single exit point in main.
func run() error {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}
