package bitflyer

import (
	"encoding/json"
	"github.com/k-ueki/sfdbot/bitflyer/model"
	"net/http"
)

const(
	postSimpleOrderUrl = "/v1/me/sendchildorder"

	OrderTypeLimit = "LIMIT"
	OrderTypeMarket = "MARKET"

	OrderSideBuy = "BUY"
	OrderSideSell = "SELL"
)

func (api *APIClient)Execution(order model.SimpleOrderRequest)(*string,error){
	body,err:=json.Marshal(order)
	if err != nil {
		return nil,err
	}

	resp,err:=api.Request(http.MethodPost, postSimpleOrderUrl,nil,body)
	if err != nil {
		return nil,err
	}
	return getOrderID(resp),nil
}

func (api *APIClient) Buy(orderType string, size float64) (*string, error) {
	order := model.SimpleOrderRequest{
		Code:    CodeBTCJPYFX,
		Side:          OrderSideBuy,
		Type: orderType,
		Size:           size,
	}
	return api.Execution(order)
}

func (api *APIClient) Sell(productCode, orderType string, size float64) (*string, error) {
	order := model.SimpleOrderRequest{
		Code:    CodeBTCJPYFX,
		Side:           OrderSideSell,
		Type: orderType,
		Size:           size,
	}
	return api.Execution(order)
}

func getOrderID(b []byte) *string {
	resp := struct {
		ChildOrderAcceptanceID string `json:"child_order_acceptance_id"`
	}{}
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil
	}
	return &resp.ChildOrderAcceptanceID
}
