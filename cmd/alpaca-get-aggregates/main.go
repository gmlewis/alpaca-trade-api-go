// alpaca-get-aggregate gets aggregate information for the authenticated user.
package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/gmlewis/alpaca-trade-api-go/client"
	jsoniter "github.com/json-iterator/go"
)

var (
	symbol   = flag.String("symbol", "SPY", "Symbol to query")
	timespan = flag.String("timespan", "day", "Timespan to query: 'minute', '1Min', '5Min', '15Min', 'day' or '1D'")
	from     = flag.String("from", "2020-01-01", "From date")
	to       = flag.String("to", "", "To date")
)

func main() {
	flag.Parse()

	if *to == "" {
		now := time.Now().Local()
		*to = fmt.Sprintf("%v-%02d-%02d", now.Year(), int(now.Month()), now.Day())
	}

	c, err := client.New()
	if err != nil {
		log.Fatalf("client.New: %v", err)
	}
	defer c.Close()

	agg, err := c.GetAggregates(*symbol, *timespan, *from, *to)
	if err != nil {
		log.Fatalf("GetAggregates: %v", err)
	}

	jsonOutput(agg)

	log.Printf("Done.")
}

func jsonOutput(in interface{}) {
	j, err := jsoniter.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(j))
}
