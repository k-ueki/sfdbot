package application

import (
	"fmt"
	"github.com/k-ueki/sfdbot/bitflyer"
	"github.com/k-ueki/sfdbot/bitflyer/model"
)

func Start()error{
	//api := bitflyer.NewAPIClient()
	//c:=config.Config

	for{
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

		tickerCh:=make(chan model.Ticker)
		go bitflyer.GetTickerStream(bitflyer.CodeBTCJPY,tickerCh)
		for ticker:=range tickerCh{
			position:=sfdPosition.HavePosition
			if position==nil {
				if ticker.PriceDisparity >= 0.0502 {
					fmt.Println(":::::::::::::::::::::::::::SELL::::::::::::::::::::::::::::::::")
					//_,err:=api.Sell(bitflyer.OrderTypeMarket,c.TradeSize)
					//if err != nil {
					//	log.Fatal(err)
					//	continue
					//}
					sfdPosition.HavePosition = &HavePosition{
						Buy:  false,
						Sell: true,
					}
				}
			}else{
				if position.Sell && ticker.PriceDisparity<0.0499{
					fmt.Println("::::::::::::::::::::::::::::BUY::::::::::::::::::::::::::::::::::")
					//_,err:=api.Buy(bitflyer.OrderTypeMarket,c.TradeSize)
					//if err != nil {
					//	log.Fatal(err)
					//	continue
					//}
					//sfdPosition.Reset()
					sfdPosition.HavePosition=nil
				}
			}
		}

	}

	return nil
}