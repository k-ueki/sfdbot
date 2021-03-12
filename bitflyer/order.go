package bitflyer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/k-ueki/sfdbot/bitflyer/model"
)

const (
	postSimpleOrderUrl = "/v1/me/sendchildorder"
	getOrderUrl        = "/v1/me/getchildorders"

	OrderTypeLimit  = "LIMIT"
	OrderTypeMarket = "MARKET"

	OrderSideBuy  = "BUY"
	OrderSideSell = "SELL"
)

func (api *APIClient) Execution(order model.SimpleOrderRequest) (*string, error) {
	body, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	resp, err := api.Request(http.MethodPost, postSimpleOrderUrl, nil, body)
	if err != nil {
		return nil, err
	}
	return getOrderID(resp), nil
}

func (api *APIClient) Buy(orderType string, size float64) (*string, error) {
	order := model.SimpleOrderRequest{
		Code: CodeBTCJPYFX,
		Side: OrderSideBuy,
		Type: orderType,
		Size: size,
	}
	return api.Execution(order)
}

func (api *APIClient) Sell(orderType string, size float64) (*string, error) {
	order := model.SimpleOrderRequest{
		Code: CodeBTCJPYFX,
		Side: OrderSideSell,
		Type: orderType,
		Size: size,
	}
	return api.Execution(order)
}

func (api *APIClient) GetOrderByID(id string) (*model.GetOrderResponse, error) {
	val, err := api.Request(http.MethodGet, genGetOrderURLWithID(id), nil, nil)
	if err != nil {
		return nil, err
	}

	resp := model.GetOrderResponse{}
	if err := json.Unmarshal(val, resp); err != nil {
		return nil, err
	}

	return &resp, nil
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

func genGetOrderURLWithID(id string) string {
	return fmt.Sprintf(`%s?child_order_acceptance_id=%s&product_code=%s`, getOrderUrl, id, CodeBTCJPYFX)
}
