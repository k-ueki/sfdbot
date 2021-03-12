package application

import (
	"log"

	"github.com/k-ueki/sfdbot/config"

	"github.com/k-ueki/sfdbot/bitflyer"
	"github.com/k-ueki/sfdbot/bitflyer/model"
)

func Start() error {
	api := bitflyer.NewAPIClient()
	c := config.Config

	for {
		//positions,err:=api.GetPosisionList()
		//if err != nil {
		//	return err
		//}
		//fmt.Println(*positions[0])
		//if len(positions)!=0{
		//	time.Sleep(time.Second)
		//	continue
		//}
		//positions,_ := api.GetPosisionList()

		tickerCh := make(chan model.Ticker)
		go bitflyer.GetTickerStream(bitflyer.CodeBTCJPY, tickerCh)
		for ticker := range tickerCh {
			position := sfdPosition.HavePosition
			if position == nil {
				if ticker.PriceDisparity >= 0.0503 {
					SendToSlack("::SELL::")
					_, err := api.Sell(bitflyer.OrderTypeMarket, c.TradeSize)
					if err != nil {
						log.Fatal(err)
						break
					}
					sfdPosition.SetSellPosition()
				}
			} else {
				if position.Sell && ticker.PriceDisparity < 0.0496 {
					SendToSlack("::BUY::")
					_, err := api.Buy(bitflyer.OrderTypeMarket, c.TradeSize)
					if err != nil {
						log.Fatal(err)
						break
					}

					col, err := api.GetLastCollateralHistory()
					if err != nil {
						return err
					}
					SendToSlack(col.String())

					sfdPosition.Reset()
				}
			}
		}
	}

	return nil
}
