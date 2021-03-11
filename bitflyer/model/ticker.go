package model

type (
	Ticker struct {
		ProductCode    string     `json:"product_code"`
		Timestamp      string     `json:"timestamp"`
		ID             int64      `json:"ticker_id"`
		Ticker         TickerInfo `json:"ticker"`
		PriceDisparity float64    `json:"price_disparity"`
	}

	TickerInfo struct {
		Ltp float64 `json:"LTP"`
	}
)
