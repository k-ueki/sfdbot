package application

import (
	"fmt"
	"log"

	"github.com/k-ueki/sfdbot/bitflyer"
	"github.com/k-ueki/sfdbot/bitflyer/model"
)

func Start() error {
	api := bitflyer.NewAPIClient()
	preTicker := model.Ticker{}

	for {
		tickerCh := make(chan model.Ticker)
		go bitflyer.GetTickerStream(bitflyer.CodeBTCJPY, tickerCh)
		for ticker := range tickerCh {
			//st := time.Now().UnixNano()
			fmt.Println(ticker)
			ltp := ticker.Ticker.Ltp
			// Ltpが前のTickerと同じであれば無視
			if preTicker.Ticker.Ltp == ticker.Ticker.Ltp {
				break
			}

			// positionの有無
			positions, err := api.GetPosisionList()
			if err != nil {
				log.Fatal(err)
				break
			}
			if len(positions) != 0 {
				sfdPosition.Reset()
			}

			fmt.Println("sfdPosition: ", *sfdPosition.HaveOrder, sfdPosition.AcceptanceID)

			if len(positions) == 0 {
				//PositionがなければCancel & IFD指値
				if sfdPosition.HaveOrder.Sell {
					if err := api.Cancel(sfdPosition.AcceptanceID); err != nil {
						log.Fatal(err)
						break
					}
				}

				orderAcceptanceID, err := api.SellLimit(ltp * 1.05)
				if err != nil {
					log.Fatal(err)
					break
				}
				sfdPosition.SetSellOrder(*orderAcceptanceID)

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
					order := model.NewCancelChildOrder(sfdPosition.AcceptanceID)
					if err := api.Cancel(order); err != nil {
						break
					}
				}

				orderAcceptanceID, err := api.BuyLimit(ltp*1.05 - 10)
				if err != nil {
					log.Fatal(err)
					break
				}
				sfdPosition.SetBuyOrder(*orderAcceptanceID)
			}
			preTicker = ticker
			//fin := time.Now().UnixNano()
			//fmt.Println(float64(fin-st) / 1000000000)
		}
	}

	return nil
}
