// alpaca-get-order gets order information for the authenticated user.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gmlewis/alpaca-trade-api-go/client"
	jsoniter "github.com/json-iterator/go"
)

var (
	id = flag.String("id", "", "Order ID to get")
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

	order, err := c.GetOrder(*id)
	if err != nil {
		log.Fatalf("GetOrder: %v", err)
	}

	jsonOutput(order)

	log.Printf("Done.")
}

func jsonOutput(in interface{}) {
	j, err := jsoniter.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(j))
}
