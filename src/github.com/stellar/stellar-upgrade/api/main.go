package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
	"github.com/stellar/go-stellar-base/keypair"
	"github.com/stellar/go-stellar-base/strkey"
)

type MessageData struct {
	NewAddress string `json:"newAddress"`
}

type Message struct {
	Data      string `json:"data"`
	PublicKey string `json:"publicKey"`
	Signature string `json:"signature"`
}

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

type NetworkApi interface {
	SendUpgradeRequest(data MessageData, kp keypair.KP) (*UpgradeResponse, error)
	SendStatusRequest(address string) (*StatusResponse, error)
}

type Api struct{}

func (Api) SendUpgradeRequest(data MessageData, kp keypair.KP) (*UpgradeResponse, error) {
	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	signature, err := kp.Sign(dataJson)
	if err != nil {
		return nil, err
	}

	signatureBase64 := base64.StdEncoding.EncodeToString(signature[:])
	keyData := strkey.MustDecode(strkey.VersionByteAccountID, kp.Address())
	publicKeyBase64 := base64.StdEncoding.EncodeToString(keyData)

	message := Message{
		Data:      string(dataJson),
		PublicKey: publicKeyBase64,
		Signature: signatureBase64,
	}

	requestJson, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	url := viper.GetString("ApiRoot") + "/upgrade/upgrade"
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

func (Api) SendStatusRequest(address string) (*StatusResponse, error) {
	url := viper.GetString("ApiRoot") + "/upgrade/balance"
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
