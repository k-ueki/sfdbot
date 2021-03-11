package bitflyer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/k-ueki/sfdbot/bitflyer/model"
	"github.com/k-ueki/sfdbot/util"
)

const (
	getTickerUrl = "https://lightning.bitflyer.com/api/trade/ticker/all?v=1"

	codeBTCJPY = "BTC_JPY"
)

func GetTicker(code string) (*model.Ticker, error) {
	val, err := util.HttpGet(getTickerUrl)
	if err != nil {
		return nil, err
	}

	tickers:=[]model.Ticker{}
	if err:=json.Unmarshal(val,&tickers);err!=nil{
		return nil,err
	}

	for _,ticker:=range tickers{
		if ticker.ProductCode==code{
			return &ticker,nil
		}
	}
	return nil, errors.New(fmt.Sprintf("cannot find %v",code))
}
