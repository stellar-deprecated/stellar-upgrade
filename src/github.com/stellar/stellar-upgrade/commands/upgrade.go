package commands

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/rubblelabs/ripple/crypto"
	"github.com/spf13/cobra"
	"github.com/stellar/go-stellar-base"
	"github.com/stellar/stellar-upgrade/api"
)

type MessageData struct {
	NewAddress string `json:"newAddress"`
}

type Message struct {
	Data      string `json:"data"`
	PublicKey string `json:"publicKey"`
	Signature string `json:"signature"`
}

var upgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade stellars to XLMs",
	Long: "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Enter your %s account %s: ", Cyan("old network"), Cyan("secret seed"))
		var oldNetworkSeed string
		// TODO this should be no echo input
		fmt.Scanln(&oldNetworkSeed)
		oldNetworkSeedRaw, err := crypto.Base58Decode(oldNetworkSeed, oldNetworkAlphabet)
		if err != nil {
			fmt.Println("Your old network account secret seed is incorrect.")
			return
		}
		var rawSeed stellarbase.RawSeed
		copy(rawSeed[:], oldNetworkSeedRaw[1 : len(oldNetworkSeedRaw)-2]) // Payload

		publicKey, privateKey, err := stellarbase.GenerateKeyFromRawSeed(rawSeed)
		if err != nil {
			fmt.Println(err)
			return
		}

		keyData := publicKey.KeyData()
		keyDataBytes := keyData[:]
		publicKeyBase64 := base64.StdEncoding.EncodeToString(keyDataBytes)

		fmt.Printf("Enter your %s account %s: ", Cyan("new network"), Cyan("address"))
		var newNetworkAddress string
		fmt.Scanln(&newNetworkAddress)

		// TODO go-stellar-base does not export strkey functions
		// _, err = stellarbase.Decode(stellarbase.VersionByteAccountID, newNetworkAddress)
		// if err != nil {
		// 	fmt.Println("Your new network account address is incorrect.")
		// 	return
		// }

		// TODO confirmation step

		data := MessageData{
			NewAddress: newNetworkAddress,
		}

		dataJson, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		signature := privateKey.Sign(dataJson)
		signatureBytes := signature[:]
		signatureBase64 := base64.StdEncoding.EncodeToString(signatureBytes)

		message := Message{
			Data: string(dataJson),
			PublicKey: publicKeyBase64,
			Signature: signatureBase64,
		}

		requestJson, err := json.Marshal(message)
		if err != nil {
			fmt.Println(err)
			return
		}
		response, err := api.SendUpgradeRequest(requestJson)
		if err != nil {
			fmt.Println(err)
			return
		}
		
		if response.Status == "success" {
			fmt.Printf("%s Your XLM should arrive soon.\n", Green("Success!"))
		} else {
			fmt.Printf("%s"+response.Message+"\n", Red("Error: "))
		}
	},
}
