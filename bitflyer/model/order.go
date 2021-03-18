package model

import "github.com/k-ueki/sfdbot/bitflyer"

type (
	SimpleOrderRequest struct {
		Code        string  `json:"product_code"`
		Type        string  `json:"child_order_type"`
		Side        string  `json:"side"`
		Price       float64 `json:"price"`
		Size        float64 `json:"size"`
		Expire      int64   `json:"minute_to_expire,omitempty"`
		TimeInForce string  `json:"time_in_force,omitempty"`
	}

	ShowOrderRequest struct {
		Code string `json:"product_code"`
		ID   string `json:"child_order_id"`
	}

	GetOrderResponse struct {
		ID          int64   `json:"id"`
		OrderID     string  `json:"child_order_id"`
		Code        string  `json:"product_code"`
		Side        string  `json:"side"`
		Type        string  `json:"child_order_type"`
		Price       float64 `json:"price"`
		Size        float64 `json:"size"`
		State       string  `json:"state"`
		ExpiredDate string  `json:"expired_date"`
		OrderDate   string  `json:"child_order_date"`
	}

	CancelOrderRequest struct {
		Code                    string `json:"product_code"`
		ParentOrderID           string `json:"parent_order_id,omitempty"`
		ParentOrderAcceptanceID string `json:"parent_order_acceptance_id,omitempty"`
	}
)

func NewCancelOrder(accepranceID string) *CancelOrderRequest {
	return &CancelOrderRequest{
		Code:                    bitflyer.CodeFXBTCJPY,
		ParentOrderAcceptanceID: accepranceID,
	}
}
