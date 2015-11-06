package commands

import (
	"fmt"

	"github.com/howeyc/gopass"
)

type CommandInput interface {
	GetOldNetworkSeedFromConsole() string
	GetNewNetworkAddressFromConsole() string
	GetConfirmationFromConsole(oldNetworkAddress, newNetworkAddress string) bool
}

type Input struct{}

func (Input) GetOldNetworkSeedFromConsole() string {
	fmt.Printf("Enter your %s account %s: [input will be hidden] ", Cyan("old network"), Cyan("secret seed"))
	oldNetworkSeedBytes := gopass.GetPasswd()
	return string(oldNetworkSeedBytes)
}

func (Input) GetNewNetworkAddressFromConsole() string {
	fmt.Printf("Enter the %s that your lumens will be sent to: ", Cyan("new network address"))
	var newNetworkAddress string
	fmt.Scanln(&newNetworkAddress)
	return newNetworkAddress
}

func (Input) GetConfirmationFromConsole(oldNetworkAddress, newNetworkAddress string) bool {
	fmt.Printf("Please confirm your addresses are correct:\n")
	//fmt.Printf("Old network address: %s\n", Cyan(oldNetworkAddress))
	fmt.Printf("New network address: %s\n", Cyan(newNetworkAddress))
	fmt.Printf("Correct? [y/N] ")

	var confirm string
	fmt.Scanln(&confirm)

	if confirm != "y" && confirm != "Y" {
		return false
	} else {
		return true
	}
}
