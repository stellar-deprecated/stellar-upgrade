package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const base = "http://localhost:3001"

type UpgradeResponse struct {
	Status  string
	Code    string
	Message string
}

type StatusResponse struct {
	OldAddress string
	Claimed    bool
	Upgraded   bool
}

func SendUpgradeRequest(requestJson []byte) (*UpgradeResponse, error) {
	url := base + "/upgrade/upgrade"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestJson))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var response UpgradeResponse
	err = json.Unmarshal(body, &response)
	return &response, err
}

func SendStatusRequest(address string) (*StatusResponse, error) {
	url := base + "/upgrade/balance"
	req, err := http.NewRequest("GET", url, nil)

	q := req.URL.Query()
	q.Add("address", address)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, errors.New("Address not found.")
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Server error.")
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var response StatusResponse
	err = json.Unmarshal(body, &response)
	return &response, err
}
