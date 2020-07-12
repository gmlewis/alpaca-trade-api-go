// alpaca-cancel-all-orders cancels all orders for the authenticated user.
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

	if err := c.CancelAllOrders(); err != nil {
		log.Fatalf("CancelAllOrders: %v", err)
	}

	log.Printf("Done.")
}
