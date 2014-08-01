// Sample program to show how to access an exported identifier.
package main

import (
	"fmt"

	"github.csv.comcast.com/VariousArtists/gotraining/gotraining/02-packaging/example1/counters"
)

// main is the entry point for the application.
func main() {
	// Create a variable of the exported type and
	// initialize the value to 10.
	counter := counters.AlertCounter(10)

	fmt.Printf("Counter: %d\n", counter)
}
