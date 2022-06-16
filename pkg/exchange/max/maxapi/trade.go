package max

//go:generate -command GetRequest requestgen -method GET
//go:generate -command PostRequest requestgen -method POST

import (
	"github.com/c9s/bbgo/pkg/fixedpoint"
	"github.com/c9s/bbgo/pkg/types"
)

type MarkerInfo struct {
	Fee         string `json:"fee"`
	FeeCurrency string `json:"fee_currency"`
	OrderID     int    `json:"order_id"`
}

type TradeInfo struct {
	// Maker tells you the maker trade side
	Maker string      `json:"maker,omitempty"`
	Bid   *MarkerInfo `json:"bid,omitempty"`
	Ask   *MarkerInfo `json:"ask,omitempty"`
}

type Liquidity string

// Trade represents one returned trade on the max platform.
type Trade struct {
	ID          uint64                     `json:"id" db:"exchange_id"`
	WalletType  WalletType                 `json:"wallet_type,omitempty"`
	Price       fixedpoint.Value           `json:"price"`
	Volume      fixedpoint.Value           `json:"volume"`
	Funds       fixedpoint.Value           `json:"funds"`
	Market      string                     `json:"market"`
	MarketName  string                     `json:"market_name"`
	CreatedAt   types.MillisecondTimestamp `json:"created_at"`
	Side        string                     `json:"side"`
	OrderID     uint64                     `json:"order_id"`
	Fee         fixedpoint.Value           `json:"fee"` // float number as string
	FeeCurrency string                     `json:"fee_currency"`
	Liquidity   Liquidity                  `json:"liquidity"`
	Info        TradeInfo                  `json:"info,omitempty"`
}

func (t Trade) IsBuyer() bool {
	return t.Side == "bid" || t.Side == "buy"
}

func (t Trade) IsMaker() bool {
	return t.Info.Maker == t.Side
}

type TradeService struct {
	client *RestClient
}
