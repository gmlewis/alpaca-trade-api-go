// alpaca-get-position gets position information for the authenticated user.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gmlewis/alpaca-trade-api-go/client"
	jsoniter "github.com/json-iterator/go"
)

var (
	symbol = flag.String("symbol", "", "Get this account's position for this symbol")
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

	position, err := c.GetPosition(*symbol)
	if err != nil {
		log.Fatalf("GetPosition: %v", err)
	}

	jsonOutput(position)

	log.Printf("Done.")
}

func jsonOutput(in interface{}) {
	j, err := jsoniter.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(j))
}
