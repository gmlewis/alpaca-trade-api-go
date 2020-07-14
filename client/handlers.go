package client

import (
	"fmt"
	"log"

	"github.com/gmlewis/alpaca-trade-api-go/alpaca"
	"github.com/gmlewis/alpaca-trade-api-go/api"
	"github.com/gmlewis/alpaca-trade-api-go/polygon"
)

// NewAccountHandler creates a new account handler.
func (c *Client) NewAccountHandler(handler api.AccountHandler) error {
	f := func(msg interface{}) {
		account, ok := msg.(alpaca.AccountUpdate)
		if !ok {
			log.Printf("accountHandler: unknown msg type %T = %#v", msg, msg)
			return
		}
		handler(account)
	}

	if err := c.AStream.Subscribe(alpaca.AccountUpdates, f); err != nil {
		return fmt.Errorf("NewAccountHandler: %v", err)
	}

	return nil
}

// NewOrderHandler creates a new order handler.
func (c *Client) NewOrderHandler(handler api.OrderHandler) error {
	f := func(msg interface{}) {
		order, ok := msg.(alpaca.TradeUpdate)
		if !ok {
			log.Printf("orderHandler: unknown msg type %T = %#v", msg, msg)
			return
		}
		handler(order)
	}

	if err := c.AStream.Subscribe(alpaca.TradeUpdates, f); err != nil {
		return fmt.Errorf("NewOrderHandler: %v", err)
	}

	return nil
}

// NewQuoteHandler creates a new quote handler.
func (c *Client) NewQuoteHandler(symbol string, handler api.QuoteHandler) error {
	f := func(msg interface{}) {
		quote, ok := msg.(polygon.StreamQuote)
		if !ok {
			log.Printf("quoteHandler: unknown msg type %T = %#v", msg, msg)
			return
		}
		handler(quote)
	}

	if err := c.PStream.Subscribe("Q."+symbol, f); err != nil {
		return fmt.Errorf("NewQuoteHandler(%q): %v", symbol, err)
	}

	return nil
}

// NewTradeHandler creates a new trade handler.
func (c *Client) NewTradeHandler(symbol string, handler api.TradeHandler) error {
	f := func(msg interface{}) {
		trade, ok := msg.(polygon.StreamTrade)
		if !ok {
			log.Printf("tradeHandler: unknown msg type %T = %#v", msg, msg)
			return
		}
		handler(trade)
	}

	if err := c.PStream.Subscribe("T."+symbol, f); err != nil {
		return fmt.Errorf("NewTradeHandler(%q): %v", symbol, err)
	}

	return nil
}

// NewMinuteBarHandler creates a new minuteBar handler.
func (c *Client) NewMinuteBarHandler(symbol string, handler api.MinuteBarHandler) error {
	f := func(msg interface{}) {
		minuteBar, ok := msg.(polygon.StreamAggregate)
		if !ok {
			log.Printf("minuteBarHandler: unknown msg type %T = %#v", msg, msg)
			return
		}
		handler(minuteBar)
	}

	if err := c.PStream.Subscribe("AM."+symbol, f); err != nil {
		return fmt.Errorf("NewMinuteBarHandler(%q): %v", symbol, err)
	}

	return nil
}

// NewSecondBarHandler creates a new secondBar handler.
func (c *Client) NewSecondBarHandler(symbol string, handler api.SecondBarHandler) error {
	f := func(msg interface{}) {
		secondBar, ok := msg.(polygon.StreamAggregate)
		if !ok {
			log.Printf("secondBarHandler: unknown msg type %T = %#v", msg, msg)
			return
		}
		handler(secondBar)
	}

	if err := c.PStream.Subscribe("A."+symbol, f); err != nil {
		return fmt.Errorf("NewSecondBarHandler(%q): %v", symbol, err)
	}

	return nil
}
