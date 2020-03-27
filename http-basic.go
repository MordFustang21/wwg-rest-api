package main

import (
	"fmt"
	"log"
	"net/http"
)

// START OMIT
func main() {
	// calling our run function and handling the error
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// run is used so we can just return errors and handle a single exit point in main.
func run() error {
	fmt.Println("Starting server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}
// END OMIT
