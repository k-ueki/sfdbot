package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) ([]byte, error) {
	val, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer val.Body.Close()
	resp, err := ioutil.ReadAll(val.Body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func HttpPost(url, jsonContent string) error {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonContent)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
