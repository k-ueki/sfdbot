package bitflyer

import (
	"encoding/json"
	"github.com/k-ueki/sfdbot/bitflyer/model"
	"net/http"
)

const (
getPositionListUrl = "/v1/me/getpositions"
)

func (api *APIClient) GetPosisionList()([]*model.Position,error){
	resp, err := api.Request(http.MethodGet, getPositionListUrl+"?product_code=FX_BTC_JPY", nil, nil)
	if err != nil {
		return nil, err
	}

	var list []*model.Position
	if err := json.Unmarshal(resp, &list); err != nil {
		return nil, err
	}
	return list,nil
}
