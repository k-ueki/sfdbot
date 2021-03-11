package bitflyer

import (
	"github.com/k-ueki/sfdbot/bitflyer/model"
	"github.com/k-ueki/sfdbot/util"
)

const (
	getTickerUrl = "https://lightning.bitflyer.com/api/trade/ticker/all?v=1"
)

func GetTicker() (*model.Ticker, error) {
	tickers, err := util.HttpGet(getTickerUrl)
	if err != nil {
		return nil, err
	}

	resp := model.Ticker{}
	return &resp, nil
}
