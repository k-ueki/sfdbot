package bitflyer

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/k-ueki/sfdbot/bitflyer/model"
	"github.com/k-ueki/sfdbot/util"
)

const (
	getTickerUrl = "https://lightning.bitflyer.com/api/trade/ticker/all?v=1"

	CodeBTCJPY   = "BTC_JPY"
	CodeFXBTCJPY = "FX_BTC_JPY"
)

func GetTickerStream(code string, ch chan model.Ticker) {
	for {
		ticker, err := GetTicker(code)
		if err != nil {
			log.Fatal(err)
			return
		}
		ch <- *ticker
		time.Sleep(time.Millisecond * 800)
	}
}
func (api *APIClient) GetRealTimeTicker(code string, ch chan model.RealTimeTicker) {
	channel := fmt.Sprintf("lightning_ticker_%s", code)
	if err := api.WebsocketClient.WriteJSON(&model.JsonRPC2{
		Version: "2.0",
		Method:  "subscribe",
		Params:  &model.SubscribeParam{channel},
	}); err != nil {
		log.Fatal("subscribe: ", err)
		return
	}

	for {
		msg := new(model.JsonRPC2)
		if err := api.WebsocketClient.ReadJSON(&msg); err != nil {
			ws, _, err := websocket.DefaultDialer.Dial(websocketURL, nil)
			if err != nil {
				log.Fatal("cannot restore websocket")
			}
			api.WebsocketClient = ws
			continue
		}

		if msg.Method == "channelMessage" {
			switch v := msg.Params.(type) {
			case map[string]interface{}:
				for key, binary := range v {
					if key == "message" {
						var ticker model.RealTimeTicker
						if err := bindResponse(binary, &ticker); err != nil {
							break
						}
						ch <- ticker
					}
				}
			}
		}
	}
}

func GetTicker(code string) (*model.Ticker, error) {
	val, err := util.HttpGet(getTickerUrl)
	if err != nil {
		return nil, err
	}

	tickers := []model.Ticker{}
	if err := json.Unmarshal(val, &tickers); err != nil {
		return nil, err
	}

	for _, ticker := range tickers {
		if ticker.ProductCode == code {
			return &ticker, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("cannot find %v", code))
}

func bindResponse(response interface{}, v interface{}) error {
	ms, err := json.Marshal(response)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(ms, &v); err != nil {
		return err
	}
	return nil
}
