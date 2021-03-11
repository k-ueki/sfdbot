package application

import (
	"fmt"
	"github.com/k-ueki/sfdbot/bitflyer"
	"github.com/k-ueki/sfdbot/bitflyer/model"
)

func Start()error{
	api := bitflyer.NewAPIClient()
	positions,err:=api.GetPosisionList()
	if err != nil {
		return err
	}
	fmt.Println("positions: ",positions)

	tickerCh:=make(chan model.Ticker)
	go bitflyer.GetTickerStream(bitflyer.CodeBTCJPY,tickerCh)
	for ticker:=range tickerCh{
		fmt.Println(ticker)
	}

	return nil
}