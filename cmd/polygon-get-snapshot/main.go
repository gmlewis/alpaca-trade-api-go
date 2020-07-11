// polygon-get-snapshot gets snapshot information for the authenticated user.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gmlewis/alpaca-trade-api-go/client"
	jsoniter "github.com/json-iterator/go"
)

var (
	symbol = flag.String("symbol", "GOOG", "Get this symbol's bars")
)

func main() {
	flag.Parse()

	c, err := client.New()
	if err != nil {
		log.Fatalf("client.New: %v", err)
	}
	defer c.Close()

	exch, err := c.PClient.GetSnapshot(*symbol)
	if err != nil {
		log.Fatalf("GetSnapshot: %v", err)
	}

	jsonOutput(exch)

	log.Printf("Done.")
}

func jsonOutput(in interface{}) {
	j, err := jsoniter.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(j))
}
