// alpaca-stream-market streams Trades, Quotes, and Minute-Bars for the named market(s).
package main

import (
	"flag"
	"log"
	"time"

	"github.com/gmlewis/alpaca-trade-api-go/alpaca"
	"github.com/gmlewis/alpaca-trade-api-go/polygon"
	"github.com/gmlewis/alpaca-trade-api-go/api"
	"github.com/gmlewis/alpaca-trade-api-go/client"
)

var (
	symbolsFlag = flag.String("symbols", "SPY", "Comma-separated list of symbols to stream")
)

// Server keeps track of all trading during this run.
// Server implements the client.HandlerGetter interface.
type Server struct {
	c *client.Client
	m map[string]*Market
}

// Market keeps track of a single market.
type Market struct {
	s   *Server
	sym string

	dayAgg *polygon.HistoricAggregatesV2
	minAgg *polygon.HistoricAggregatesV2
}

func main() {
	flag.Parse()

	c, err := client.New()
	if err != nil {
		log.Fatalf("client.New: %v", err)
	}
	defer c.Close()

	srv := &Server{c: c, m: map[string]*Market{}}
	symbols := client.GetSymbols(*symbolsFlag)

	now := time.Now().Local()
	fromDay := now.Add(time.Duration(-1000) * 24 * time.Hour)
	fromMinute := now.Add(time.Duration(-1000) * time.Minute)

	for _, symbol := range symbols {
		m := srv.getMarket(symbol)
		m.dayAgg, err = c.GetHistoricAggregatesV2(symbol, 1, "day", &fromDay, &now, nil)
		if err != nil {
			log.Fatalf("GetHistoricAggregatesV2(%q, 1, 'day'): %v", symbol, err)
		}
		log.Printf("dayHistoricAgg(%q): %#v", symbol, m.dayAgg)

		m.minAgg, err = c.GetHistoricAggregatesV2(symbol, 1, "minute", &fromMinute, &now, nil)
		if err != nil {
			log.Fatalf("GetHistoricAggregatesV2(%q, 1, 'minute'): %v", symbol, err)
		}
		log.Printf("minuteHistoricAgg(%q): %#v", symbol, m.minAgg)
	}

	doneCh, err := c.StartStreams(symbols, srv)
	if err != nil {
		log.Fatalf("StartStreams: %v", err)
	}

	<-doneCh
	log.Printf("Done.")
}

func (s *Server) accountHandler(account interface{}) {
	log.Printf("accountHandler: %#v", account)
}

func (s *Server) orderHandler(order alpaca.TradeUpdate) {
	log.Printf("orderHandler: %#v", order)
}

func (m *Market) quoteHandler(quote polygon.StreamQuote) {
	log.Printf("quoteHandler: %#v", quote)
}

func (m *Market) tradeHandler(trade polygon.StreamTrade) {
	log.Printf("tradeHandler: %#v", trade)
}

func (m *Market) minuteBarHandler(minuteBar polygon.StreamAggregate) {
	log.Printf("minuteBarHandler: %#v", minuteBar)
}

func (m *Market) secondBarHandler(secondBar polygon.StreamAggregate) {
	log.Printf("secondBarHandler: %#v", secondBar)
}

func (s *Server) GetAccountHandler() api.AccountHandler { return s.accountHandler }
func (s *Server) GetOrderHandler() api.OrderHandler     { return s.orderHandler }

func (s *Server) GetQuoteHandler(symbol string) api.QuoteHandler {
	m := s.getMarket(symbol)
	return m.quoteHandler
}

func (s *Server) GetTradeHandler(symbol string) api.TradeHandler {
	m := s.getMarket(symbol)
	return m.tradeHandler
}

func (s *Server) GetMinuteBarHandler(symbol string) api.MinuteBarHandler {
	m := s.getMarket(symbol)
	return m.minuteBarHandler
}

func (s *Server) GetSecondBarHandler(symbol string) api.SecondBarHandler {
	m := s.getMarket(symbol)
	return m.secondBarHandler
}

func (s *Server) getMarket(symbol string) *Market {
	if m, ok := s.m[symbol]; ok {
		return m
	}

	m := &Market{s: s, sym: symbol}
	s.m[symbol] = m
	return m
}
