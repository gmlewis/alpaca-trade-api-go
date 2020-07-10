// alpaca-get-portfolio-history gets portfolio history information for the authenticated user.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gmlewis/alpaca-trade-api-go/client"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	flag.Parse()

	c, err := client.New()
	if err != nil {
		log.Fatalf("client.New: %v", err)
	}
	defer c.Close()

	hist, err := c.GetPortfolioHistory(nil, nil, nil, true)
	if err != nil {
		log.Fatalf("GetPortfolioHistory: %v", err)
	}

	jsonOutput(hist)

	log.Printf("Done.")
}

func jsonOutput(in interface{}) {
	j, err := jsoniter.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(j))
}
