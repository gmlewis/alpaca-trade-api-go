// alpaca-list-assets lists assets information for the authenticated user.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gmlewis/alpaca-trade-api-go/client"
	jsoniter "github.com/json-iterator/go"
)

type Server struct {
	c *client.Client
}

func main() {
	flag.Parse()

	c, err := client.New()
	if err != nil {
		log.Fatalf("client.New: %v", err)
	}
	defer c.Close()

	status := "active"
	assets, err := c.ListAssets(&status)
	if err != nil {
		log.Fatalf("ListAssets: %v", err)
	}

	jsonOutput(assets)

	log.Printf("Found %v %v assets.", len(assets), status)
	log.Printf("Done.")
}

func jsonOutput(in interface{}) {
	j, err := jsoniter.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(j))
}
