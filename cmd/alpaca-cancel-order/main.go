// alpaca-cancel-order cancels an order for the authenticated user.
package main

import (
	"flag"
	"log"

	"github.com/gmlewis/alpaca-trade-api-go/client"
)

var (
	id = flag.String("id", "", "Order ID to cancel")
)

func main() {
	flag.Parse()

	if *id == "" {
		log.Fatalf("Must supply order -id flag.")
	}

	c, err := client.New()
	if err != nil {
		log.Fatalf("client.New: %v", err)
	}
	defer c.Close()

	if err := c.CancelOrder(*id); err != nil {
		log.Fatalf("CancelOrder: %v", err)
	}

	log.Printf("Done.")
}
