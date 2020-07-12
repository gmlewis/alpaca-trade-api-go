// alpaca-get-lastQuote gets last quote information for the authenticated user.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gmlewis/alpaca-trade-api-go/client"
	jsoniter "github.com/json-iterator/go"
)

var (
	symbol = flag.String("symbol", "", "Get the last quote for this symbol")
)

func main() {
	flag.Parse()

	if *symbol == "" {
		log.Fatalf("Must supply -symbol flag.")
	}

	c, err := client.New()
	if err != nil {
		log.Fatalf("client.New: %v", err)
	}
	defer c.Close()

	lastQuote, err := c.GetLastQuote(*symbol)
	if err != nil {
		log.Fatalf("GetLastQuote: %v", err)
	}

	jsonOutput(lastQuote)

	log.Printf("Done.")
}

func jsonOutput(in interface{}) {
	j, err := jsoniter.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(j))
}
