package bitflyer

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/k-ueki/sfdbot/config"
)

const (
	baseURL      = "https://api.bitflyer.com"
	websocketURL = "wss://ws.lightstream.bitflyer.com/json-rpc"
)

type (
	APIClient struct {
		Key             string
		Secret          string
		Client          *http.Client
		WebsocketClient *websocket.Conn
	}
)

func NewAPIClient() *APIClient {
	c := config.Config
	ws, _, err := websocket.DefaultDialer.Dial(websocketURL, nil)
	if err != nil {
		log.Fatal("cannot connect websocket")
		return nil
	}
	return &APIClient{
		Key:             c.ApiKey,
		Secret:          c.ApiSecret,
		Client:          new(http.Client),
		WebsocketClient: ws,
	}
}

func (api *APIClient) Request(method, endpoint string, query map[string]string, body []byte) ([]byte, error) {
	resp, err := api.RequestSimpleResponse(method, endpoint, query, body)
	if err != nil {
		return nil, err
	}
	return convertResponseToByte(resp)
}

func (api *APIClient) RequestSimpleResponse(method, endpoint string, query map[string]string, body []byte) (*http.Response, error) {
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

	return api.Client.Do(req)
}

func convertResponseToByte(resp *http.Response) ([]byte, error) {
	return ioutil.ReadAll(resp.Body)
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
