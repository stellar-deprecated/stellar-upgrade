package commands

import (
	"fmt"

	"github.com/rubblelabs/ripple/crypto"
	"github.com/spf13/cobra"
	"github.com/stellar/stellar-upgrade/api"
)

var status = &cobra.Command{
	Use:   "status [address]",
	Short: "Displays your account upgrade status",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("[address] parameter is required.")
			return
		}
		address := args[0]
		_, err := crypto.Base58Decode(address, oldNetworkAlphabet)
		if err != nil {
			fmt.Println("Your old network account address is incorrect.")
			return
		}

		response, err := api.Api{}.SendStatusRequest(address)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Address: " + response.OldAddress)
		fmt.Print("Claimed: ")
		if response.Claimed {
			fmt.Printf("%s\n", Green("Yes"))
		} else {
			fmt.Println("No")
		}
		fmt.Print("Upgraded: ")
		if response.Upgraded {
			fmt.Printf("%s\n", Green("Yes"))
		} else {
			fmt.Println("No")
		}
	},
}
