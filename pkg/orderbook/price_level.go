package orderbook

import "github.com/shopspring/decimal"

type PriceLevel struct {
	Price       decimal.Decimal
	TotalAmount decimal.Decimal
}
