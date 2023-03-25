package orderbook

import "github.com/shopspring/decimal"

type OrderSide uint32

const (
	SideBuy OrderSide = iota
	SideSell
)

type Order struct {
	Price  decimal.Decimal
	Amount decimal.Decimal
	Side   OrderSide
}
