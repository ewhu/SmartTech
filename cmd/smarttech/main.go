// cmd/smarttech/main.go
package main

import (
	"flag"
	"log"
	"os"

	"smarttech/internal/smarttech"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := smarttech.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
