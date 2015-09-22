package commands

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "stellar-upgrade",
	Short: "stellar-upgrade upgrades your old network account to the new network",
}

func Execute() {
	RootCmd.AddCommand(upgrade)
	RootCmd.AddCommand(status)
	RootCmd.Execute()
}
