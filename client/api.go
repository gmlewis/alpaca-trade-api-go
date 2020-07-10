package client

import (
	"time"

	"github.com/gmlewis/alpaca-trade-api-go/alpaca"
	"github.com/gmlewis/alpaca-trade-api-go/polygon"
)

// Alpaca v2
func (c *Client) GetAccount() (*alpaca.Account, error) {
	return c.AClient.GetAccount()
}
func (c *Client) GetAccountConfigurations() (*alpaca.AccountConfigurations, error) {
	return c.AClient.GetAccountConfigurations()
}
func (c *Client) UpdateAccountConfigurations(newConfigs alpaca.AccountConfigurationsRequest) (*alpaca.AccountConfigurations, error) {
	return c.AClient.UpdateAccountConfigurations(newConfigs)
}
func (c *Client) GetAccountActivities(activityType *string, opts *alpaca.AccountActivitiesRequest) ([]alpaca.AccountActvity, error) {
	return c.AClient.GetAccountActivities(activityType, opts)
}
func (c *Client) GetPortfolioHistory(period *string, timeframe *alpaca.RangeFreq, dateEnd *time.Time, extendedHours bool) (*alpaca.PortfolioHistory, error) {
	return c.AClient.GetPortfolioHistory(period, timeframe, dateEnd, extendedHours)
}
func (c *Client) ListPositions() ([]alpaca.Position, error) {
	return c.AClient.ListPositions()
}
func (c *Client) GetPosition(symbol string) (*alpaca.Position, error) {
	return c.AClient.GetPosition(symbol)
}
func (c *Client) GetAggregates(symbol, timespan, from, to string) (*alpaca.Aggregates, error) {
	return c.AClient.GetAggregates(symbol, timespan, from, to)
}
func (c *Client) GetLastQuote(symbol string) (*alpaca.LastQuoteResponse, error) {
	return c.AClient.GetLastQuote(symbol)
}
func (c *Client) GetLastTrade(symbol string) (*alpaca.LastTradeResponse, error) {
	return c.AClient.GetLastTrade(symbol)
}
func (c *Client) CloseAllPositions() error {
	return c.AClient.CloseAllPositions()
}
func (c *Client) ClosePosition(symbol string) error {
	return c.AClient.ClosePosition(symbol)
}
func (c *Client) GetClock() (*alpaca.Clock, error) {
	return c.AClient.GetClock()
}
func (c *Client) GetCalendar(start, end *string) ([]alpaca.CalendarDay, error) {
	return c.AClient.GetCalendar(start, end)
}
func (c *Client) ListOrders(status *string, until *time.Time, limit *int, nested *bool) ([]alpaca.Order, error) {
	return c.AClient.ListOrders(status, until, limit, nested)
}
func (c *Client) PlaceOrder(req alpaca.PlaceOrderRequest) (*alpaca.Order, error) {
	return c.AClient.PlaceOrder(req)
}
func (c *Client) GetOrder(orderID string) (*alpaca.Order, error) {
	return c.AClient.GetOrder(orderID)
}
func (c *Client) ReplaceOrder(orderID string, req alpaca.ReplaceOrderRequest) (*alpaca.Order, error) {
	return c.AClient.ReplaceOrder(orderID, req)
}
func (c *Client) CancelOrder(orderID string) error {
	return c.AClient.CancelOrder(orderID)
}
func (c *Client) CancelAllOrders() error {
	return c.AClient.CancelAllOrders()
}
func (c *Client) ListAssets(status *string) ([]alpaca.Asset, error) {
	return c.AClient.ListAssets(status)
}
func (c *Client) GetAsset(symbol string) (*alpaca.Asset, error) {
	return c.AClient.GetAsset(symbol)
}
func (c *Client) ListBars(symbols []string, opts alpaca.ListBarParams) (map[string][]alpaca.Bar, error) {
	return c.AClient.ListBars(symbols, opts)
}
func (c *Client) GetSymbolBars(symbol string, opts alpaca.ListBarParams) ([]alpaca.Bar, error) {
	return c.AClient.GetSymbolBars(symbol, opts)
}

// Polygon v2
func (c *Client) GetHistoricAggregatesV2(
	symbol string,
	multiplier int,
	resolution polygon.AggType,
	from, to *time.Time,
	unadjusted *bool) (*polygon.HistoricAggregatesV2, error) {
	return c.PClient.GetHistoricAggregatesV2(symbol,
		multiplier,
		resolution,
		from, to,
		unadjusted)
}
func (c *Client) GetHistoricTradesV2(ticker string, date string, opts *polygon.HistoricTicksV2Params) (*polygon.HistoricTradesV2, error) {
	return c.PClient.GetHistoricTradesV2(ticker, date, opts)
}
func (c *Client) GetHistoricQuotesV2(ticker string, date string, opts *polygon.HistoricTicksV2Params) (*polygon.HistoricQuotesV2, error) {
	return c.PClient.GetHistoricQuotesV2(ticker, date, opts)
}
func (c *Client) GetStockExchanges() ([]polygon.StockExchange, error) {
	return c.PClient.GetStockExchanges()
}
func (c *Client) GetPreviousClose(symbol string) (*polygon.PreviousCloseV2, error) {
	return c.PClient.GetPreviousClose(symbol)
}
