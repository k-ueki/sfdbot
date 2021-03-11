package bitflyer

import (
	"encoding/json"
	"fmt"
	"github.com/k-ueki/sfdbot/bitflyer/model"
	"net/http"
)

const(
	getCollateralHistoryUrl="/v1/me/getcollateralhistory"
)

func (api *APIClient) GetCollateralHistories() ([]*model.Collateral, error) {
	resp, err := api.Request(http.MethodGet, getCollateralHistoryUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	var list []*model.Collateral
	if err := json.Unmarshal(resp, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (api *APIClient) GetLatestCollateralHistory() (*model.Collateral, error) {
	resp, err := api.Request(http.MethodGet, genCollateralHistoryUrlWithCount(1), nil, nil)
	if err != nil {
		return nil, err
	}

	var list model.Collateral
	if err := json.Unmarshal(resp, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func genCollateralHistoryUrlWithCount(count int64) string {
	return fmt.Sprintf("%s?count=%d", getCollateralHistoryUrl, count)
}
