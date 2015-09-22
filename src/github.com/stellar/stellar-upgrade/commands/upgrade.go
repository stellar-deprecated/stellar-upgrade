package commands

import (
	"fmt"

	"github.com/rubblelabs/ripple/crypto"
	"github.com/spf13/cobra"
	"github.com/stellar/go-stellar-base"
	"github.com/stellar/stellar-upgrade/api"
)

var upgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the old network account to the new network",
	Run: func(cmd *cobra.Command, args []string) {
		updateCommand := UpdateCommand{
			Input: Input{},	
			ApiObject: api.Api{},
		}
		message := updateCommand.Run()
		if message != "" {
			fmt.Println(message)	
		}
	},
}

type UpdateCommand struct {
	Input CommandInput
	ApiObject api.NetworkApi
}

func (command UpdateCommand) Run() string {
	oldNetworkSeed := command.Input.GetOldNetworkSeedFromConsole()
	oldNetworkSeedRaw, err := crypto.Base58Decode(oldNetworkSeed, oldNetworkAlphabet)
	if err != nil {
		return "Your old network account secret seed is incorrect."
	}
	var rawSeed stellarbase.RawSeed
	copy(rawSeed[:], oldNetworkSeedRaw[1:len(oldNetworkSeedRaw)-2]) // Payload

	publicKey, privateKey, err := stellarbase.GenerateKeyFromRawSeed(rawSeed)
	if err != nil {
		return "Error generating signing keys from your secret seed"
	}

	oldNetworkAddress := publicKey.Address()
	newNetworkAddress := command.Input.GetNewNetworkAddressFromConsole()
	messageData := api.MessageData{NewAddress: newNetworkAddress}
	confirmed := command.Input.GetConfirmationFromConsole(oldNetworkAddress, messageData.NewAddress)
	if !confirmed {
		return "Exiting...";
	}

	response, err := command.ApiObject.SendUpgradeRequest(messageData, publicKey, privateKey)
	if err != nil {
		return "Error building or sending request to upgrade API."
	}

	if response.Status == "success" {
		return fmt.Sprintf("%s Your XLM should arrive soon.", Green("Success!"))
	} else {
		return fmt.Sprintf("%s"+response.Message, Red("Error: "))
	}
}
