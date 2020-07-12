// alpaca-replace-order replaces an order for the authenticated user.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gmlewis/alpaca-trade-api-go/alpaca"
	"github.com/gmlewis/alpaca-trade-api-go/client"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
)

var (
	id       = flag.String("id", "", "Order ID to replace")
	limit    = flag.Float64("limit", 0, "Price for limit order")
	market   = flag.Bool("market", false, "Market order")
	qtyFlag  = flag.Float64("qty", 0, "Number of shares for order")
	stop     = flag.Float64("stop", 0, "Price for stop order (optionally add -limit)")
	timeFlag = flag.String("time", "day", "Time in force for order")
)

func main() {
	flag.Parse()

	if *id == "" {
		log.Fatalf("Must supply order -id flag.")
	}

	var limitPrice *decimal.Decimal
	var stopPrice *decimal.Decimal
	switch {
	case *stop != 0 && *limit != 0:
		limitPrice = Decimal(limit)
		stopPrice = Decimal(stop)
	case *stop != 0:
		stopPrice = Decimal(stop)
	case *limit != 0:
		limitPrice = Decimal(limit)
	case *market:
	default:
		log.Fatalf("Must provide -limit, -market, and/or -stop flag(s).")
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

	var qty *decimal.Decimal
	if *qtyFlag > 0 {
		qty = Decimal(qtyFlag)
	}

	c, err := client.New()
	if err != nil {
		log.Fatalf("client.New: %v", err)
	}
	defer c.Close()

	req := alpaca.ReplaceOrderRequest{
		Qty:           qty,
		TimeInForce:   timeInForce,
		LimitPrice:    limitPrice,
		StopPrice:     stopPrice,
		ClientOrderID: uuid.New().String(),
	}
	order, err := c.ReplaceOrder(*id, req)
	if err != nil {
		log.Fatalf("ReplaceOrder: %v", err)
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
