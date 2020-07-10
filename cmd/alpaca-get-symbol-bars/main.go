// alpaca-get-symbol-bars gets symbol bars for the authorized account.
package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gmlewis/alpaca-trade-api-go/alpaca"
	"github.com/gmlewis/alpaca-trade-api-go/client"
	jsoniter "github.com/json-iterator/go"
)

const (
	timeFmt = "2006-01-02"
)

var (
	symbol    = flag.String("symbol", "GOOG", "Get this symbol's bars")
	timespan  = flag.String("timespan", "day", "Timespan to query: 'minute', '1Min', '5Min', '15Min', 'day' or '1D'")
	startDate = flag.String("startDate", "", "Start of query (e.g. '2010-01-01')")
	endDate   = flag.String("endDate", "", "End of query (e.g. '2020-01-01')")
	limit     = flag.Int("limit", 0, "Limit number of candles (0 = max)")
)

func main() {
	flag.Parse()

	c, err := client.New()
	if err != nil {
		log.Fatalf("client.New: %v", err)
	}
	defer c.Close()

	*symbol = strings.ToUpper(*symbol)

	maxLimit := 1000
	opts := alpaca.ListBarParams{Timeframe: *timespan, Limit: &maxLimit}
	if *startDate != "" {
		start, err := time.ParseInLocation(timeFmt, *startDate, time.Local)
		if err != nil {
			log.Fatalf("unable to parse startDate %q: %v", *startDate, err)
		}
		opts.StartDt = &start
	}
	if *endDate != "" {
		end, err := time.ParseInLocation(timeFmt, *endDate, time.Local)
		if err != nil {
			log.Fatalf("unable to parse endDate %q: %v", *endDate, err)
		}
		opts.EndDt = &end
	}
	if *limit > 0 {
		opts.Limit = limit
	}

	bars, err := c.GetSymbolBars(*symbol, opts)
	if err != nil {
		log.Fatalf("GetSymbolBars: %v", err)
	}

	jsonOutput(bars)

	if len(bars) > 0 {
		start := time.Unix(bars[0].Time, 0).Local()
		end := time.Unix(bars[len(bars)-1].Time, 0).Local()
		log.Printf("Got %v candles for %q from %v to %v", len(bars), *symbol, start, end)
	} else {
		log.Printf("No data returned for %q", *symbol)
	}
	log.Printf("Done.")
}

func jsonOutput(in interface{}) {
	j, err := jsoniter.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(j))
}
