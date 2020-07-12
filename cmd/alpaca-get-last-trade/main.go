// alpaca-get-last-trade gets last trade information for the authenticated user.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gmlewis/alpaca-trade-api-go/client"
	jsoniter "github.com/json-iterator/go"
)

var (
	symbol = flag.String("symbol", "", "Get the last trade for this symbol")
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

	lastTrade, err := c.GetLastTrade(*symbol)
	if err != nil {
		log.Fatalf("GetLastTrade: %v", err)
	}

	jsonOutput(lastTrade)

	log.Printf("Done.")
}

func jsonOutput(in interface{}) {
	j, err := jsoniter.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(j))
}
