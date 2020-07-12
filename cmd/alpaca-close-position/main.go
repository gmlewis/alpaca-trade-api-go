// alpaca-close-position closes a position for the authenticated user.
package main

import (
	"flag"
	"log"

	"github.com/gmlewis/alpaca-trade-api-go/client"
)

var (
	symbol = flag.String("symbol", "", "Close this account's position for this symbol")
)

func main() {
	flag.Parse()

	if *symbol == "" {
		log.Fatalf("Must supply position -symbol flag.")
	}

	c, err := client.New()
	if err != nil {
		log.Fatalf("client.New: %v", err)
	}
	defer c.Close()

	if err := c.ClosePosition(*symbol); err != nil {
		log.Fatalf("ClosePosition: %v", err)
	}

	log.Printf("Done.")
}
