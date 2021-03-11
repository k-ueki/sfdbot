package bitflyer

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/k-ueki/sfdbot/config"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const(
	baseURL      = "https://api.bitflyer.com"
)

type (
	APIClient struct{
		Key             string
		Secret          string
		Client *http.Client
	}
)

func NewAPIClient() *APIClient {
	c:=config.Config
	return &APIClient{
		Key:             c.ApiKey,
		Secret:          c.ApiSecret,
		Client: new(http.Client),
	}
}

func (api *APIClient) Request(method, endpoint string, query map[string]string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(method, baseURL+endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for key, val := range query {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()

	for key, val := range api.getHeader(method, endpoint, body) {
		req.Header.Add(key, val)
	}

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBytes, nil
}

func (api *APIClient) getHeader(method, endpoint string, body []byte) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	msg := timestamp + method + endpoint + string(body)
	sign := makeAccessSign(msg, api.Secret)
	return map[string]string{
		"ACCESS-KEY":       api.Key,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
	}
}

func makeAccessSign(msg, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))
}