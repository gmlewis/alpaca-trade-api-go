// alpaca-get-calendar gets calendar information for the authenticated user.
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
	start = flag.String("start", "2020-01-01", "Start date")
	end   = flag.String("end", "", "End date")
)

func main() {
	flag.Parse()

	if *end == "" {
		now := time.Now().Local()
		*end = fmt.Sprintf("%v-%02d-%02d", now.Year(), int(now.Month()), now.Day())
	}

	c, err := client.New()
	if err != nil {
		log.Fatalf("client.New: %v", err)
	}
	defer c.Close()

	agg, err := c.GetCalendar(start, end)
	if err != nil {
		log.Fatalf("GetCalendar: %v", err)
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
