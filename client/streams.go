package client

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gmlewis/alpaca-trade-api-go/api"
)

// GetSymbols splits the comma-separated symbol list and upcases the names.
func GetSymbols(allSymbols string) []string {
	allSymbols = strings.ToUpper(allSymbols)
	return strings.Split(allSymbols, ",")
}

// StartStreams starts the data streams using the specified source.
//
// symbols is a slice of upper-case symbol names to stream.
//
// handlerGetter gets the handlers for the various streams.
//
// The provided channel is sent an empty struct when the stream ends.
// For "paper" or "live", the stream ends at 4pm ET. For a filename source,
// the stream ends at the end of the file.
func (c *Client) StartStreams(symbols []string, handlerGetter api.HandlerGetter) (<-chan struct{}, error) {
	// switch c.source {
	// case "live":
	// 	return nil, fmt.Errorf("source %v not yet supported", c.source)
	// case "paper":
	return c.startPaperStreams(symbols, handlerGetter)
	// default:
	// 	return c.simClient.StartLogfileStreams(c.source, symbols, handlerGetter)
	// }
}

func (c *Client) startPaperStreams(symbols []string, hg api.HandlerGetter) (<-chan struct{}, error) {
	if ah := hg.GetAccountHandler(); ah != nil {
		if err := c.NewAccountHandler(ah); err != nil {
			return nil, fmt.Errorf("NewAccountHandler: %v", err)
		}
	} else {
		log.Printf("No AccountHandler created.")
	}

	if oh := hg.GetOrderHandler(); oh != nil {
		if err := c.NewOrderHandler(oh); err != nil {
			return nil, fmt.Errorf("NewOrderHandler: %v", err)
		}
	} else {
		log.Printf("No OrderHandler created.")
	}

	for _, symbol := range symbols {
		if qh := hg.GetQuoteHandler(symbol); qh != nil {
			if err := c.NewQuoteHandler(symbol, qh); err != nil {
				return nil, fmt.Errorf("NewQuoteHandler(%q): %v", symbol, err)
			}
		} else {
			log.Printf("No QuoteHandler created for symbol %q.", symbol)
		}

		if th := hg.GetTradeHandler(symbol); th != nil {
			if err := c.NewTradeHandler(symbol, th); err != nil {
				return nil, fmt.Errorf("NewTradeHandler(%q): %v", symbol, err)
			}
		} else {
			log.Printf("No TradeHandler created for symbol %q.", symbol)
		}

		if mbh := hg.GetMinuteBarHandler(symbol); mbh != nil {
			if err := c.NewMinuteBarHandler(symbol, mbh); err != nil {
				return nil, fmt.Errorf("NewMinuteBarHandler(%q): %v", symbol, err)
			}
		} else {
			log.Printf("No MinuteBarHandler created for symbol %q.", symbol)
		}

		if sbh := hg.GetSecondBarHandler(symbol); sbh != nil {
			if err := c.NewSecondBarHandler(symbol, sbh); err != nil {
				return nil, fmt.Errorf("NewSecondBarHandler(%q): %v", symbol, err)
			}
		} else {
			log.Printf("No SecondBarHandler created for symbol %q.", symbol)
		}
	}

	// Fire the doneCh at 4pm ET.
	doneCh := make(chan struct{})
	now := time.Now().Local()
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		return nil, err
	}
	marketClose := time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, location)
	dur := marketClose.Sub(now)
	timer := time.NewTimer(dur)
	go func() {
		<-timer.C
		doneCh <- struct{}{}
	}()

	return doneCh, nil
}
