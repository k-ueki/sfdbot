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

	JsonRPC2 struct {
		Version string      `json:"jsonrpc"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params"`
		Result  interface{} `json:"result"`
		Error   interface{} `json:"error"`
		ID      *int        `json:"id,omitempty"`
	}

	SubscribeParam struct {
		Channel string `json:"channel"`
	}

	RealTimeTicker struct {
		ProductCode     string  `json:"product_code"`
		Timestamp       string  `json:"timestamp"`
		TickID          int     `json:"tick_id"`
		BestBid         float64 `json:"best_bid"`
		BestAsk         float64 `json:"best_ask"`
		BestBidSize     float64 `json:"best_bid_size"`
		BestAskSize     float64 `json:"best_ask_size"`
		TotalBidDepth   float64 `json:"total_bid_depth"`
		TotalAskDepth   float64 `json:"total_ask_depth"`
		Ltp             float64 `json:"ltp"`
		Volume          float64 `json:"volume"`
		VolumeByProduct float64 `json:"volume_by_product"`
	}
)
