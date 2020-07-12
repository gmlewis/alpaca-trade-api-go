// alpaca-place-order places an order for the authenticated user.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gmlewis/alpaca-trade-api-go/alpaca"
	"github.com/gmlewis/alpaca-trade-api-go/client"
	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
)

var (
	limit    = flag.Float64("limit", 0, "Price for limit order")
	market   = flag.Bool("market", false, "Market order")
	qty      = flag.Float64("qty", 0, "Number of shares for order")
	side     = flag.String("side", "buy", "Side of order")
	stop     = flag.Float64("stop", 0, "Price for stop order (optionally add -limit)")
	symbol   = flag.String("symbol", "", "Symbol for order")
	timeFlag = flag.String("time", "day", "Time in force for order")
)

func main() {
	flag.Parse()

	if *symbol == "" {
		log.Fatalf("Must provide -symbol.")
	}

	var orderType alpaca.OrderType
	var limitPrice *decimal.Decimal
	var stopPrice *decimal.Decimal
	switch {
	case *stop != 0 && *limit != 0:
		orderType = alpaca.StopLimit
		limitPrice = Decimal(limit)
		stopPrice = Decimal(stop)
	case *stop != 0:
		orderType = alpaca.Stop
		stopPrice = Decimal(stop)
	case *limit != 0:
		orderType = alpaca.Limit
		limitPrice = Decimal(limit)
	case *market:
	default:
		log.Fatalf("Must provide -limit, -market, and/or -stop flag(s).")
	}

	var orderSide = alpaca.Buy
	switch *side {
	case "buy":
	case "sell":
		orderSide = alpaca.Sell
	default:
		log.Fatalf("-side must be 'buy' or 'sell'.")
	}

	var timeInForce alpaca.TimeInForce
	switch *timeFlag {
	case "day":
		timeInForce = alpaca.Day
	case "gtc":
		timeInForce = alpaca.GTC
	case "opg":
		timeInForce = alpaca.OPG
	case "ioc":
		timeInForce = alpaca.IOC
	case "fok":
		timeInForce = alpaca.FOK
	case "gtx":
		timeInForce = alpaca.GTX
	case "gtd":
		timeInForce = alpaca.GTD
	case "cls":
		timeInForce = alpaca.CLS
	default:
		log.Fatalf("-time must be one of: day, gtc, opg, ioc, fok, gtx, gtd, cls")
	}

	if *qty <= 0 {
		log.Fatalf("-qty must not be > 0.")
	}

	c, err := client.New()
	if err != nil {
		log.Fatalf("client.New: %v", err)
	}
	defer c.Close()

	req := alpaca.PlaceOrderRequest{
		// AccountID:     0, // string
		AssetKey:    symbol,
		Qty:         decimal.NewFromFloat(*qty),
		Side:        orderSide,
		Type:        orderType,
		TimeInForce: timeInForce,
		LimitPrice:  limitPrice,
		StopPrice:   stopPrice,
		// ClientOrderID: 0, // string
		// OrderClass:    0, // OrderClass
		// TakeProfit:    0, // *TakeProfit
		// StopLoss:      0, // *StopLoss
	}
	order, err := c.PlaceOrder(req)
	if err != nil {
		log.Fatalf("PlaceOrder: %v", err)
	}

	jsonOutput(order)

	log.Printf("Done.")
}

func Decimal(v *float64) *decimal.Decimal {
	tmp := decimal.NewFromFloat(*v)
	return &tmp
}

func jsonOutput(in interface{}) {
	j, err := jsoniter.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(j))
}
