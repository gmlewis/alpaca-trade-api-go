// alpaca-list-bars lists bar lists for the named symbol(s).
package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/gmlewis/alpaca-trade-api-go/alpaca"
	"github.com/gmlewis/alpaca-trade-api-go/client"
	jsoniter "github.com/json-iterator/go"
	"github.com/tj/go-naturaldate"
)

var (
	end         = flag.String("end", "", "End time using natural date parser")
	limitFlag   = flag.Int("limit", 0, "Optional limit for request")
	start       = flag.String("start", "", "Start time using natural date parser")
	symbolsFlag = flag.String("symbols", "", "Comma-separated list of symbols to list")
	timeframe   = flag.String("timeframe", "day", "Timeframe: 'day', 'minute', '1Min', '5Min', '15Min', '1H', '1D'")
)

func main() {
	flag.Parse()

	if *symbolsFlag == "" {
		log.Fatalf("Must supply -symbols flag.")
	}

	c, err := client.New()
	if err != nil {
		log.Fatalf("client.New: %v", err)
	}
	defer c.Close()

	symbols := client.GetSymbols(*symbolsFlag)

	var startDt, endDt *time.Time
	if *start != "" {
		now := time.Now().Local()
		t, err := naturaldate.Parse(*start, now)
		if err != nil {
			log.Fatalf("Unable to parse start time: %v", err)
		}
		startDt = &t
	}
	if *end != "" {
		now := time.Now().Local()
		t, err := naturaldate.Parse(*end, now)
		if err != nil {
			log.Fatalf("Unable to parse end time: %v", err)
		}
		endDt = &t
	}
	var limit *int
	if *limitFlag != 0 {
		limit = limitFlag
	}

	opts := alpaca.ListBarParams{
		Timeframe: *timeframe,
		StartDt:   startDt,
		EndDt:     endDt,
		Limit:     limit,
	}
	barLists, err := c.ListBars(symbols, opts)
	if err != nil {
		log.Fatalf("ListBars: %v", err)
	}

	jsonOutput(barLists)

	log.Printf("Done.")
}

func jsonOutput(in interface{}) {
	j, err := jsoniter.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(j))
}
