package commands

import (
	"fmt"

	"github.com/rubblelabs/ripple/crypto"
	"github.com/spf13/cobra"
	"github.com/stellar/go-stellar-base/keypair"
	"github.com/stellar/stellar-upgrade/api"
)

var upgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the old network account to the new network",
	Run: func(cmd *cobra.Command, args []string) {
		updateCommand := UpdateCommand{
			Input:     Input{},
			ApiObject: api.Api{},
		}
		message := updateCommand.Run()
		if message != "" {
			fmt.Println(message)
		}
	},
}

type UpdateCommand struct {
	Input     CommandInput
	ApiObject api.NetworkApi
}

func (command UpdateCommand) Run() string {
	oldNetworkSeed := command.Input.GetOldNetworkSeedFromConsole()
	oldNetworkSeedRaw, err := crypto.Base58Decode(oldNetworkSeed, oldNetworkAlphabet)
	if err != nil {
		return "Your old network account secret seed is incorrect."
	}
	var rawSeed [32]byte
	copy(rawSeed[:], oldNetworkSeedRaw[1:len(oldNetworkSeedRaw)-2]) // Payload

	kp, err := keypair.FromRawSeed(rawSeed)
	if err != nil {
		return "Error generating signing keys from your secret seed"
	}

	fmt.Println("The stellar network uses a new format for public and secret keys.")
	fmt.Println("The new format for the seed you pasted above is:")
	fmt.Println("")
	fmt.Printf("Public Key: %s\n", kp.Address())
	fmt.Printf("Secret Key: %s\n", kp.(*keypair.Full).Seed())
	fmt.Println("")

	oldNetworkAddress := kp.Address()
	newNetworkAddress := command.Input.GetNewNetworkAddressFromConsole()
	messageData := api.MessageData{NewAddress: newNetworkAddress}
	confirmed := command.Input.GetConfirmationFromConsole(oldNetworkAddress, messageData.NewAddress)
	if !confirmed {
		return "Exiting..."
	}

	response, err := command.ApiObject.SendUpgradeRequest(messageData, kp)
	if err != nil {
		return "Error building or sending request to upgrade API."
	}

	if response.Status == "success" {
		return fmt.Sprintf("%s Your XLM should arrive soon.", Green("Success!"))
	} else {
		return fmt.Sprintf("%s"+response.Message, Red("Error: "))
	}
}
