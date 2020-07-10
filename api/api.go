// Package API defines the Alpaca v2 and Polygon v2 APIs interface.
package api

import (
	"time"

	"github.com/gmlewis/alpaca-trade-api-go/alpaca"
	"github.com/gmlewis/alpaca-trade-api-go/polygon"
)

// API specifies the Alpaca v2 and Polygon Rest APIs that can be called.
type API interface {
	// Alpaca v2
	GetAccount() (*alpaca.Account, error)
	GetAccountConfigurations() (*alpaca.AccountConfigurations, error)
	UpdateAccountConfigurations(newConfigs alpaca.AccountConfigurationsRequest) (*alpaca.AccountConfigurations, error)
	GetAccountActivities(activityType *string, opts *alpaca.AccountActivitiesRequest) ([]alpaca.AccountActvity, error)
	GetPortfolioHistory(period *string, timeframe *alpaca.RangeFreq, dateEnd *time.Time, extendedHours bool) (*alpaca.PortfolioHistory, error)
	ListPositions() ([]alpaca.Position, error)
	GetPosition(symbol string) (*alpaca.Position, error)
	GetAggregates(symbol, timespan, from, to string) (*alpaca.Aggregates, error)
	GetLastQuote(symbol string) (*alpaca.LastQuoteResponse, error)
	GetLastTrade(symbol string) (*alpaca.LastTradeResponse, error)
	CloseAllPositions() error
	ClosePosition(symbol string) error
	GetClock() (*alpaca.Clock, error)
	GetCalendar(start, end *string) ([]alpaca.CalendarDay, error)
	ListOrders(status *string, until *time.Time, limit *int, nested *bool) ([]alpaca.Order, error)
	PlaceOrder(req alpaca.PlaceOrderRequest) (*alpaca.Order, error)
	GetOrder(orderID string) (*alpaca.Order, error)
	ReplaceOrder(orderID string, req alpaca.ReplaceOrderRequest) (*alpaca.Order, error)
	CancelOrder(orderID string) error
	CancelAllOrders() error
	ListAssets(status *string) ([]alpaca.Asset, error)
	GetAsset(symbol string) (*alpaca.Asset, error)
	ListBars(symbols []string, opts alpaca.ListBarParams) (map[string][]alpaca.Bar, error)
	GetSymbolBars(symbol string, opts alpaca.ListBarParams) ([]alpaca.Bar, error)

	// Polygon v2
	GetHistoricAggregatesV2(
		symbol string,
		multiplier int,
		resolution polygon.AggType,
		from, to *time.Time,
		unadjusted *bool) (*polygon.HistoricAggregatesV2, error)
	GetHistoricTradesV2(ticker string, date string, opts *polygon.HistoricTicksV2Params) (*polygon.HistoricTradesV2, error)
	GetHistoricQuotesV2(ticker string, date string, opts *polygon.HistoricTicksV2Params) (*polygon.HistoricQuotesV2, error)
	GetStockExchanges() ([]polygon.StockExchange, error)
}

// HandlerGetter gets the handlers for the various streams.
type HandlerGetter interface {
	GetAccountHandler() AccountHandler
	GetOrderHandler() OrderHandler
	GetQuoteHandler(symbol string) QuoteHandler
	GetTradeHandler(symbol string) TradeHandler
	GetMinuteBarHandler(symbol string) MinuteBarHandler
	GetSecondBarHandler(symbol string) SecondBarHandler
}

// AccountHandler is a handler that processes accounts.
type AccountHandler func(account interface{})

// OrderHandler is a handler that processes orders.
type OrderHandler func(order alpaca.TradeUpdate)

// QuoteHandler is a handler that processes quotes.
type QuoteHandler func(quote polygon.StreamQuote)

// TradeHandler is a handler that processes trades.
type TradeHandler func(trade polygon.StreamTrade)

// MinuteBarHandler is a handler that processes minuteBars.
type MinuteBarHandler func(minuteBar polygon.StreamAggregate)

// SecondBarHandler is a handler that processes secondBars.
type SecondBarHandler func(secondBar polygon.StreamAggregate)
