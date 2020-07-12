// alpaca-close-all-positions closes all positions for the authenticated user.
package main

import (
	"flag"
	"log"

	"github.com/gmlewis/alpaca-trade-api-go/client"
)

func main() {
	flag.Parse()

	c, err := client.New()
	if err != nil {
		log.Fatalf("client.New: %v", err)
	}
	defer c.Close()

	if err := c.CloseAllPositions(); err != nil {
		log.Fatalf("CloseAllPositions: %v", err)
	}

	log.Printf("Done.")
}
