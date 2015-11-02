package main

import (
	"github.com/spf13/viper"
	"github.com/stellar/stellar-upgrade/commands"
)

func main() {
	viper.SetDefault("ApiRoot", "https://api.stellar.org")
	commands.Execute()
}
