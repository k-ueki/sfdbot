package bitflyer

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/k-ueki/sfdbot/config"

	"github.com/k-ueki/sfdbot/bitflyer/model"
)

const (
	postSimpleOrderUrl = "/v1/me/sendchildorder"
	getOrderUrl        = "/v1/me/getchildorders"
	postCancelOrderUrl = "/v1/me/cancelparentorder"

	OrderTypeLimit  = "LIMIT"
	OrderTypeMarket = "MARKET"

	OrderSideBuy  = "BUY"
	OrderSideSell = "SELL"
)

var (
	simpleBuyOrder = model.SimpleOrderRequest{
		Code: CodeFXBTCJPY,
		Type: OrderTypeMarket,
		Side: OrderSideBuy,
		Size: config.Config.TradeSize,
	}

	simpleSellOrder = model.SimpleOrderRequest{
		Code: CodeFXBTCJPY,
		Type: OrderTypeMarket,
		Side: OrderSideSell,
		Size: config.Config.TradeSize,
	}
)

func (api *APIClient) Execution(method string, url string, order interface{}) (*string, error) {
	body, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	resp, err := api.Request(method, url, nil, body)
	if err != nil {
		return nil, err
	}
	return getOrderID(resp), nil
}

func (api *APIClient) Buy() (*string, error) {
	return api.Execution(http.MethodPost, postSimpleOrderUrl, simpleBuyOrder)
}

func (api *APIClient) Sell() (*string, error) {
	return api.Execution(http.MethodPost, postSimpleOrderUrl, simpleSellOrder)
}

func (api *APIClient) Cancel(orderID string) error {
	order := model.NewCancelOrder(orderID)
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	resp, err := api.RequestSimpleResponse(http.MethodPost, postCancelOrderUrl, nil, body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("failed to cancel: %s", orderID))
	}
	return nil
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
	return fmt.Sprintf(`%s?child_order_acceptance_id=%s&product_code=%s`, getOrderUrl, id, CodeFXBTCJPY)
}
