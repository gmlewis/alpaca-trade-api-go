// alpaca-get-account-configurations gets account configuration information for the authenticated user.
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

	cfgs, err := c.GetAccountConfigurations()
	if err != nil {
		log.Fatalf("GetAccountConfigurations: %v", err)
	}

	jsonOutput(cfgs)

	log.Printf("Done.")
}

func jsonOutput(in interface{}) {
	j, err := jsoniter.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(j))
}
