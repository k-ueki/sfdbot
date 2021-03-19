package application

import (
	"log"

	"github.com/k-ueki/sfdbot/bitflyer"
	"github.com/k-ueki/sfdbot/bitflyer/model"
)

func Start() error {
	api := bitflyer.NewAPIClient()
	preTicker := model.Ticker{}

	tickerCh := make(chan model.Ticker)
	go bitflyer.GetTickerStream(bitflyer.CodeBTCJPY, tickerCh)
	for ticker := range tickerCh {
		//fmt.Println(ticker)
		ltp := ticker.Ticker.Ltp

		// positionの有無
		positions, err := api.GetPosisionList()
		if err != nil {
			log.Fatal(err)
			break
		}

		log.Println("sfdPosition: ", *sfdPosition.HaveOrder, sfdPosition.AcceptanceID)

		if preTicker.Ticker.Ltp != ltp {
			if len(positions) == 0 {
				//PositionがなければCancel & IFD指値
				if sfdPosition.HaveOrder.Sell {
					if err := api.Cancel(model.CancelChildOrderRequest{
						Code:              bitflyer.CodeFXBTCJPY,
						OrderAcceptanceID: sfdPosition.AcceptanceID,
					}); err != nil {
						log.Fatal(err)
						break
					}
					log.Println("canceled previous order")
					sfdPosition.Reset()
				}

				orderAcceptanceID, err := api.SellLimit(int64(ltp*1.05) + 30)
				if err != nil {
					log.Fatal(err)
					break
				}
				if *orderAcceptanceID != "" {
					sfdPosition.SetSellOrder(*orderAcceptanceID)
				}

			} else {
				/*
					positionありの場合
						sfdPosition {
							HavePosition = {Buy: false, Sell: true}
							HaveOrder = nil
							AcceptanceID = "JRO00-ssss-ssss"
						}
				*/
				if sfdPosition.HaveOrder.Buy {
					if err := api.Cancel(model.CancelChildOrderRequest{
						Code:              bitflyer.CodeFXBTCJPY,
						OrderAcceptanceID: sfdPosition.AcceptanceID,
					}); err != nil {
						log.Fatal(err)
						break
					}
					log.Println("canceled previous order")
					sfdPosition.Reset()
				}

				orderAcceptanceID, err := api.BuyLimit(int64(ltp*1.05) - 15)
				if err != nil {
					log.Fatal(err)
					break
				}
				if *orderAcceptanceID != "" {
					sfdPosition.SetBuyOrder(*orderAcceptanceID)
				}
			}
			preTicker = ticker
		}
	}

	return nil
}
