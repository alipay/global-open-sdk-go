package model

type StockInfo struct {
	ListedRegion string `json:"listedRegion,omitempty"`
	TickerSymbol string `json:"tickerSymbol,omitempty"`
}
