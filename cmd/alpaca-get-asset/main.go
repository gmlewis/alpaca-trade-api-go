// alpaca-get-asset gets asset information for the authenticated user.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gmlewis/alpaca-trade-api-go/client"
	jsoniter "github.com/json-iterator/go"
)

var (
	symbol = flag.String("symbol", "", "Asset symbol to get")
)

func main() {
	flag.Parse()

	if *symbol == "" {
		log.Fatalf("Must supply asset -symbol flag.")
	}

	c, err := client.New()
	if err != nil {
		log.Fatalf("client.New: %v", err)
	}
	defer c.Close()

	asset, err := c.GetAsset(*symbol)
	if err != nil {
		log.Fatalf("GetAsset: %v", err)
	}

	jsonOutput(asset)

	log.Printf("Done.")
}

func jsonOutput(in interface{}) {
	j, err := jsoniter.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(j))
}
